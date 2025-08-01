package i688opensdk

import (
	"context"
	"errors"
	"fmt"
	utils2 "github.com/mimicode/tksdk/utils"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	ApiGatewayUrl = "http://gw.open.1688.com/openapi"
	ApiSignMethod = "hmac-sha1"
	ApiFormat     = "json"
	ApiVersion    = "param2/1"
)

type DefaultRequest interface {
	AddParameter(string, string)
	GetParameters() url.Values
	GetApiName() string
	GetBusinessModule() string
	CheckParameters()
}

type DefaultResponse interface {
	//解析返回结果
	WrapResult(result string)
	IsError() bool
}

type TopClient struct {
	AppKey         string
	AppSecret      string
	ProxyUrl       string
	RequestTimeOut time.Duration
	HttpClient     *http.Client
	SysParameters  *url.Values //系统变量
}

func (u *TopClient) getTimeOut() time.Duration {
	if u.RequestTimeOut == 0 {
		return 30 * time.Second
	}
	return u.RequestTimeOut
}

func (u *TopClient) Init(appKey, appSecret, sessionkey string) {
	u.AppKey = appKey
	u.AppSecret = appSecret
	u.SysParameters = &url.Values{}
	//1688分配给应用的AppKey
	u.SysParameters.Add("_aop_appkey", appKey)
	//签名的摘要算法，可选值为：hmac-sha1
	u.SysParameters.Add("_aop_signature_method", ApiSignMethod)
	if sessionkey != "" {
		//用户授权令牌
		u.SysParameters.Add("access_token", sessionkey)
	}

	//时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8
	u.SysParameters.Add("_aop_timestamp", utils2.NowTime().Format("2006-01-02 15:04:05"))
	//响应格式。默认为xml格式，可选值：xml，json。
	u.SysParameters.Add("_aop_response_format", ApiFormat)
}

func (u *TopClient) CreateSign(businessModule, method string, params url.Values) string {
	// 1. 构造签名因子：urlPath
	urlPath := fmt.Sprintf("%s/%s/%s/%s", ApiVersion, businessModule, method, u.AppKey)

	// 2. 构造签名因子：拼装的参数
	// 将参数的key和value拼在一起，按照首字母排序
	paramKeys := make([]string, 0, len(params))
	for k := range params {
		if k != "_aop_signature" { // 签名参数不参与签名
			paramKeys = append(paramKeys, k)
		}
	}
	sort.Strings(paramKeys)

	paramStr := ""
	for _, k := range paramKeys {
		paramStr += k + params.Get(k)
	}

	// 3. 合并两个签名因子
	signStr := urlPath + paramStr

	// 4. 对合并后的签名因子执行hmac_sha1算法
	sign := strings.ToUpper(utils2.HmacSha1(signStr, u.AppSecret))
	return sign
}

func (u *TopClient) CreateStrParam(businessModule, method string, params url.Values) string {
	//合并参数
	newParams := url.Values{}
	for k, v := range *u.SysParameters {
		for _, vv := range v {
			newParams.Add(k, vv)
		}
	}

	for k, v := range params {
		for _, vv := range v {
			newParams.Add(k, vv)
		}
	}

	// 添加签名
	sign := u.CreateSign(businessModule, method, newParams)
	newParams.Add("_aop_signature", sign)

	return newParams.Encode()
}

// 发送POST请求
func (u *TopClient) PostRequest(businessModule, method string, uri string) (string, error) {
	if u.HttpClient == nil {
		dc := &net.Dialer{Timeout: 5 * time.Second}
		if len(u.ProxyUrl) > 0 {
			u.HttpClient = &http.Client{
				Transport: &http.Transport{
					Proxy: func(r *http.Request) (*url.URL, error) {
						return url.Parse(u.ProxyUrl)
					},
					DisableKeepAlives: true,
					DialContext:       dc.DialContext,
				},
				Timeout: u.getTimeOut(),
			}
		} else {
			u.HttpClient = &http.Client{
				Transport: &http.Transport{
					DisableKeepAlives: true,
					DialContext:       dc.DialContext,
				},
				Timeout: u.getTimeOut(),
			}
		}
	}

	// 构建完整的API URL
	apiUrl := fmt.Sprintf("%s/%s/%s/%s/%s?%s", ApiGatewayUrl, ApiVersion, businessModule, method, u.AppKey, uri)

	request, err := http.NewRequest(http.MethodPost, apiUrl, nil)
	if err != nil {
		return "", nil
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	request.WithContext(ctx)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := u.HttpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	//响应状态是不是 200
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("Response statusCode:%d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func (u *TopClient) Execute(businessModule, method string, params url.Values) (string, error) {
	//拼装请求参数
	uri := u.CreateStrParam(businessModule, method, params)

	return u.PostRequest(businessModule, method, uri)
}

func (u *TopClient) Exec(request DefaultRequest, response DefaultResponse) error {
	//检测参数
	request.CheckParameters()
	//API接口名称
	method := request.GetApiName()
	if method == "" {
		panic("API name missing")
	}
	//业务模块
	businessModule := request.GetBusinessModule()
	if businessModule == "" {
		panic("Business module missing")
	}

	//请求参数
	params := request.GetParameters()
	result, err := u.Execute(businessModule, method, params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil
}