package pddopensdk

import (
	"errors"
	"fmt"
	utils2 "github.com/mimicode/tksdk/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	ApiGatewayUrl = "https://gw-api.pinduoduo.com/api/router"
	ApiFormat     = "JSON"
	ApiVersion    = "V1"
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
	Appkey        string
	AppSecret     string
	ProxyUrl      string
	SysParameters *url.Values //系统变量
}

func (u *TopClient) Init(appKey, appSecret, sessionkey string) {
	u.Appkey = appKey
	u.AppSecret = appSecret
	u.SysParameters = &url.Values{}
	u.SysParameters.Add("client_id", appKey)
	if sessionkey != "" {
		u.SysParameters.Add("access_token", sessionkey)
	}
	u.SysParameters.Add("timestamp", strconv.FormatInt(utils2.NowTime().Unix(), 10))
	u.SysParameters.Add("data_type", ApiFormat)
	u.SysParameters.Add("version", ApiVersion)
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

//发送POST请求
func (u *TopClient) PostRequest(uri string) (string, error) {
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
	request, err := http.NewRequest(http.MethodPost, ApiGatewayUrl, strings.NewReader(uri))
	if err != nil {
		return "", nil
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
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
	u.SysParameters.Set("type", method)

	//请求参数
	params := request.GetParameters()
	result, err := u.Execute(params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil

}
