package csjdyopensdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mimicode/tksdk/utils"
)

const (
	ApiGatewayUrl = "https://ecom.pangolin-sdk-toutiao.com"
	ApiSignMethod = "MD5"
	ApiVersion    = "1"
)

type DefaultRequest interface {
	AddParameter(string, interface{})
	GetParameters() map[string]interface{} // data参数，此时传入的是map
	GetApiName() string                    // 接口名称 其实是接口的path
	CheckParameters() error
}

type DefaultResponse interface {
	WrapResult(result string)
	IsError() bool
}

type TopClient struct {
	Appkey         string
	AppSecret      string
	ProxyUrl       string
	RequestTimeOut time.Duration
	HttpClient     *http.Client
	SysParameters  map[string]interface{} //系统变量
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
	u.SysParameters = map[string]interface{}{}
	//TOP分配给应用的AppKey
	u.SysParameters["app_id"] = appKey
	//API协议版本，当前支持"1"。若不传，默认为"1"
	u.SysParameters["version"] = ApiVersion
	//签名算法类型。目前只支持MD5，若不传，默认为MD5
	u.SysParameters["sign_type"] = ApiSignMethod
}

func (u *TopClient) createSign(params map[string]interface{}) string {
	return utils.Md5(fmt.Sprintf("app_id=%v&data=%v&req_id=%v&timestamp=%v%v", params["app_id"], params["data"], params["req_id"], params["timestamp"], u.AppSecret))
}

// CreateRequestParams
//
// data 参数
func (u *TopClient) CreateRequestParams(params map[string]interface{}) map[string]interface{} {
	// 合并参数
	newParams := map[string]interface{}{}
	// 固定参数
	for k, v := range u.SysParameters {
		newParams[k] = v
	}

	//秒时间戳，和服务器时间相差超过 10分钟会报错。和服务器时间相差超过
	newParams["timestamp"] = utils.NowTime().Unix()
	//唯一id，方便问题排查
	newParams["req_id"] = utils.GetUUID()

	jsonData, _ := json.Marshal(params)
	newParams["data"] = string(jsonData)
	//  签名结果
	newParams["sign"] = u.createSign(newParams)
	return newParams
}

// PostRequest 发送POST请求
func (u *TopClient) PostRequest(uri string, jsonData []byte) (string, error) {
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
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", ApiGatewayUrl, uri), strings.NewReader(string(jsonData)))
	if err != nil {
		return "", nil
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := u.HttpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	//响应状态是不是 200
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response statusCode:%d", response.StatusCode)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), err

}

func (u *TopClient) Execute(method string, params map[string]interface{}) (string, error) {
	//签名
	data := u.CreateRequestParams(params)
	jsonData, _ := json.Marshal(data)
	return u.PostRequest(method, jsonData)
}
func (u *TopClient) Exec(request DefaultRequest, response DefaultResponse) error {
	//检测参数
	if err := request.CheckParameters(); err != nil {
		return err
	}

	//API接口名称
	method := request.GetApiName()
	if method == "" {
		return fmt.Errorf("API name missing")
	}

	//请求参数
	params := request.GetParameters()
	result, err := u.Execute(method, params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil

}
