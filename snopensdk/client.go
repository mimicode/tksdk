package snopensdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
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
	ApiGatewayUrl = "https://open.suning.com/api/http/sopRequest"
	ApiFormat     = "json"
)

type DefaultRequest interface {
	AddParameter(string, interface{})
	GetParameters() map[string]interface{}
	GetApiName() string
	GetVersion() string
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
	RequestTimeOut time.Duration
	HttpClient     *http.Client
	ProxyUrl       string
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
	u.SysParameters.Add("appKey", appKey)
	if sessionkey != "" {
		u.SysParameters.Add("access_token", sessionkey)
	}
	u.SysParameters.Add("appRequestTime", utils2.NowTime().Format("2006-01-02 15:04:05"))
	u.SysParameters.Add("format", ApiFormat)
}

func (u *TopClient) CreateSign(params map[string]interface{}) {
	//排序
	//拼装签名字符串
	signStr := u.AppSecret
	signStr += u.SysParameters.Get("appMethod")
	signStr += u.SysParameters.Get("appRequestTime")
	signStr += u.SysParameters.Get("appKey")
	signStr += u.SysParameters.Get("versionNo")
	marshal, _ := json.Marshal(params)
	signStr += base64.StdEncoding.EncodeToString(marshal)
	sign := utils2.Md5(signStr)
	//API输入参数签名结果
	u.SysParameters.Set("signInfo", sign)
}

func (u *TopClient) CreateStrParam() string {
	return u.SysParameters.Encode()
}

//发送POST请求
func (u *TopClient) PostRequest(uri, body string) (string, error) {
	if u.HttpClient == nil {
		if len(u.ProxyUrl) > 0 {
			u.HttpClient = &http.Client{
				Transport: &http.Transport{
					Proxy: func(r *http.Request) (*url.URL, error) {
						return url.Parse(u.ProxyUrl)
					},
					DisableKeepAlives: true,
					DialContext:       (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
				},
				Timeout: u.getTimeOut(),
			}
		} else {
			u.HttpClient = &http.Client{
				Transport: &http.Transport{
					DisableKeepAlives: true,
					DialContext:       (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
				},
				Timeout: u.getTimeOut(),
			}
		}
	}
	request, err := http.NewRequest(http.MethodPost, ApiGatewayUrl+"/"+uri, strings.NewReader(body))
	if err != nil {
		return "", nil
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	request.WithContext(ctx)

	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("AppMethod", u.SysParameters.Get("appMethod"))
	request.Header.Add("AppRequestTime", u.SysParameters.Get("appRequestTime"))
	request.Header.Add("Format", ApiFormat)
	request.Header.Add("signInfo", u.SysParameters.Get("signInfo"))
	request.Header.Add("AppKey", u.Appkey)
	request.Header.Add("VersionNo", u.SysParameters.Get("versionNo"))
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

func (u *TopClient) Execute(params map[string]interface{}) (string, error) {
	//签名
	u.CreateSign(params)
	marshal, _ := json.Marshal(params)
	return u.PostRequest(u.SysParameters.Get("appMethod"), string(marshal))
}
func (u *TopClient) Exec(request DefaultRequest, response DefaultResponse) error {
	//检测参数
	request.CheckParameters()
	//API接口名称
	appMethod := request.GetApiName()
	if appMethod == "" {
		panic("API appMethod missing")
	}
	u.SysParameters.Set("appMethod", appMethod)
	version := request.GetVersion()
	u.SysParameters.Set("versionNo", version)
	//请求参数
	params := request.GetParameters()
	result, err := u.Execute(params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil

}
