package vipopensdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mimicode/tksdk/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	ApiGatewayUrl = "https://gw.vipapis.com"
	ApiFormat     = "JSON"
)

type DefaultRequest interface {
	AddParameter(string, interface{})
	GetParameters() map[string]interface{}
	GetApiName() string
	GetMethod() string
	GetVersion() string
	CheckParameters()
}

type DefaultResponse interface {
	//解析返回结果
	WrapResult(result string)
	IsError() bool
}

type TopClient struct {
	Appkey        string
	AppSecret     string
	ProxyUrl      string
	SysParameters *url.Values //系统变量
}

func (u *TopClient) Init(appKey, appSecret, sessionkey string) {
	u.Appkey = appKey
	u.AppSecret = appSecret
	u.SysParameters = &url.Values{}
	u.SysParameters.Add("appKey", appKey)
	if sessionkey != "" {
		u.SysParameters.Add("accessToken", sessionkey)
	}
	u.SysParameters.Add("timestamp", strconv.FormatInt(utils.NowTime().Unix(), 10))
	u.SysParameters.Add("format", ApiFormat)
	u.SysParameters.Add("language", "zh")
}

func (u *TopClient) CreateSign(params map[string]interface{}) {

	//排序
	newParamsKey := utils.SortParamters(*u.SysParameters)
	//拼装签名字符串
	signStr := ""
	for _, k := range newParamsKey {
		signStr += k + u.SysParameters.Get(k)
	}
 	marshal, _ := json.Marshal(params)
	signStr += string(marshal)
	sign := strings.ToUpper(utils.Hmac(u.AppSecret, signStr))
	//API输入参数签名结果
	//D724D8CCE929A4118B5ABB7D54634875
	u.SysParameters.Set("sign", sign)
}

func (u *TopClient) CreateStrParam() string {

	return u.SysParameters.Encode()
}

// PostRequest 发送POST请求
func (u *TopClient) PostRequest(uri, body string) (string, error) {
	var client *http.Client
	if len(u.ProxyUrl) > 0 {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: func(r *http.Request) (*url.URL, error) {
					return url.Parse(u.ProxyUrl)
				},
			},
			Timeout: 30 * time.Second,
		}
	} else {
		client = http.DefaultClient
	}
	request, err := http.NewRequest(http.MethodPost, ApiGatewayUrl+"?"+uri, strings.NewReader(body))
	if err != nil {
		return "", nil
	}
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	response, err := client.Do(request)
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
	//拼装请求参数
	uri := u.CreateStrParam()
	marshal, _ := json.Marshal(params)
	return u.PostRequest(uri, string(marshal))
}
func (u *TopClient) Exec(request DefaultRequest, response DefaultResponse) error {
	//检测参数
	request.CheckParameters()
	//API接口名称
	service := request.GetApiName()
	if service == "" {
		panic("API service missing")
	}
	u.SysParameters.Set("service", service)
	method := request.GetMethod()
	if method == "" {
		panic("API method missing")
	}
	version := request.GetVersion()
	if method == "" {
		panic("API version missing")
	}
	u.SysParameters.Set("method", method)
	u.SysParameters.Set("version", version)
	//请求参数
	params := request.GetParameters()
	result, err := u.Execute(params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil

}
