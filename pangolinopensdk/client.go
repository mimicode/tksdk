package pangolinopensdk

import (
	"context"
	"encoding/json"
	"errors"
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
	// 穿山甲电商联盟接口域名
	ApiGatewayUrl = "https://ecom.pangolin-sdk-toutiao.com"
	// API协议版本
	ApiVersion = "1"
	// 签名算法类型
	SignTypeMD5 = "MD5"
)

// DefaultRequest 请求接口
type DefaultRequest interface {
	AddParameter(key string, val interface{})
	GetParameters() map[string]interface{}
	GetApiName() string
	GetVersion() string
	CheckParameters()
}

// DefaultResponse 响应接口
type DefaultResponse interface {
	WrapResult(result string)
	IsError() bool
}

// TopClient 客户端
type TopClient struct {
	AppKey       string // 在穿山甲媒体平台上注册的应用id
	UserID       string // 主账号user_id
	RoleID       string // 子账号role_id，超管的role_id等于其user_id
	SecureKey    string // 安全密钥
	HttpClient   *http.Client
	ProxyUrl     string
	RequestTimeOut time.Duration
}

// Init 初始化客户端
func (u *TopClient) Init(appKey, userID, roleID, secureKey string) {
	u.AppKey = appKey
	u.UserID = userID
	u.RoleID = roleID
	u.SecureKey = secureKey
}

// getTimeOut 获取超时时间
func (u *TopClient) getTimeOut() time.Duration {
	if u.RequestTimeOut == 0 {
		return 30 * time.Second
	}
	return u.RequestTimeOut
}

// CreateSign 生成签名
// 签名方法：MD5(app_id + role_id + timestamp + secure_key + data)
func (u *TopClient) CreateSign(data string) string {
	timestamp := time.Now().Unix()
	signStr := fmt.Sprintf("%s%s%d%s%s", u.AppKey, u.RoleID, timestamp, u.SecureKey, data)
	return utils.Md5(signStr)
}

// PostRequest 发送POST请求
func (u *TopClient) PostRequest(uri string, params map[string]interface{}) (string, error) {
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

	// 序列化data字段
	dataBytes, _ := json.Marshal(params["data"])
	dataStr := string(dataBytes)

	// 构建请求体
	requestBody := map[string]interface{}{
		"app_id":   u.AppKey,
		"timestamp": time.Now().Unix(),
		"version":  ApiVersion,
		"sign":     u.CreateSign(dataStr),
		"sign_type": SignTypeMD5,
		"req_id":   utils.GetUUID(),
		"user_id":  u.UserID,
		"role_id":  u.RoleID,
		"data":     dataStr,
	}

	bodyBytes, _ := json.Marshal(requestBody)
	request, err := http.NewRequest(http.MethodPost, ApiGatewayUrl+uri, strings.NewReader(string(bodyBytes)))
	if err != nil {
		return "", err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	request.WithContext(ctx)

	request.Header.Add("Content-Type", "application/json")

	response, err := u.HttpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("Response statusCode:%d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// Execute 执行请求
func (u *TopClient) Execute(uri string, request DefaultRequest, response DefaultResponse) error {
	// 检测参数
	request.CheckParameters()

	// 请求参数
	params := request.GetParameters()
	result, err := u.PostRequest(uri, params)
	if err != nil {
		return err
	}
	response.WrapResult(result)
	return nil
}
