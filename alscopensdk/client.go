package alscopensdk

import (
	"context"
	"errors"
	"fmt"
	utils2 "github.com/mimicode/tksdk/utils"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ApiGatewayUrl = "http://gw.api.taobao.com/router/rest"
	ApiSignMethod = "md5"
	ApiFormat     = "json"
	ApiVersion    = "2.0"
)

type DefaultRequest interface {
	AddParameter(string, string)
	GetParameters() url.Values
	GetApiName() string
	CheckParameters()
}

type DefaultResponse interface {
	//解析返回结果
	WrapResult(result string)
	IsError() bool
}

type TopClient struct {
	Appkey         string
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
	u.Appkey = appKey
	u.AppSecret = appSecret
	u.SysParameters = &url.Values{}
	//TOP分配给应用的AppKey
	u.SysParameters.Add("app_key", appKey)
	//签名的摘要算法，可选值为：hmac，md5
	u.SysParameters.Add("sign_method", ApiSignMethod)
	if sessionkey != "" {
		//用户登录授权成功后，TOP颁发给应用的授权信息,当此API的标签上注明：“需要授权”，
		// 则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
		u.SysParameters.Add("session", sessionkey)
	}

	//时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟
	u.SysParameters.Add("timestamp", utils2.NowTime().Format("2006-01-02 15:04:05"))
	//响应格式。默认为xml格式，可选值：xml，json。
	u.SysParameters.Add("format", ApiFormat)
	//API协议版本，可选值：2.0
	u.SysParameters.Add("v", ApiVersion)

	//是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false
	//u.SysParameters.Add("simplify","false")
	//合作伙伴身份标识
	//u.SysParameters.Add("partner_id","")
	//被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效
	//u.SysParameters.Add("target_app_key","")
}

func (u *TopClient) CreateSign(params url.Values) {

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
	//排序
	newParamsKey := utils2.SortParamters(newParams)
	//拼装签名字符串
	signStr := u.AppSecret
	for _, k := range newParamsKey {
		signStr += k + newParams.Get(k)
	}
	signStr += u.AppSecret
	sign := strings.ToUpper(utils2.Md5(signStr))
	//API输入参数签名结果
	u.SysParameters.Set("sign", sign)
}

func (u *TopClient) CreateStrParam(params url.Values) string {

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

	return newParams.Encode()
}

// 发送POST请求
func (u *TopClient) PostRequest(uri string) (string, error) {
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
	request, err := http.NewRequest(http.MethodPost, ApiGatewayUrl, strings.NewReader(uri))
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

func (u *TopClient) Execute(params url.Values) (string, error) {
	//签名
	u.CreateSign(params)
	//拼装请求参数
	uri := u.CreateStrParam(params)

	return u.PostRequest(uri)
}
func (u *TopClient) Exec(request DefaultRequest, response DefaultResponse) error {
	//检测参数
	request.CheckParameters()
	//API接口名称
	method := request.GetApiName()
	if method == "" {
		panic("API name missing")
	}
	u.SysParameters.Set("method", method)

	//请求参数
	params := request.GetParameters()
	result, err := u.Execute(params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil

}
