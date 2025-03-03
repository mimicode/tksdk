package vipapisoauthoauthservicerefreshtoken

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// vipapis.oauth.OauthService 刷新令牌服务
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Success struct {
	AccessToken        string `json:"access_token"`         // 访问令牌
	CreateAt           string `json:"create_at"`            // 创建时间
	ExpiresIn          int64  `json:"expires_in"`           // 有效时长(秒)
	ExpiresTime        string `json:"expires_time"`         // 过期时间
	IsExpired          bool   `json:"is_expired"`           // 是否过期
	RefreshToken       string `json:"refresh_token"`        // 刷新令牌值
	RefreshExpiresTime string `json:"refresh_expires_time"` // 刷新令牌过期时间
	Status             string `json:"status"`               // 状态
}
