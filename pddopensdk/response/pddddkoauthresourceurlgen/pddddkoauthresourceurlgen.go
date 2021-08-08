package pddddkoauthresourceurlgen

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.resource.url.gen拼多多主站频道推广接口
type Response struct {
	response.TopResponse
	ResourceURLResponse ResourceURLResponse `json:"resource_url_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type ResourceURLResponse struct {
	SingleURLList SingleURLList `json:"single_url_list"`
	Sign          string        `json:"sign"`
	WeAppInfo     WeAppInfo     `json:"we_app_info"`
	RequestID     string        `json:"request_id"`
}

type SingleURLList struct {
	WeAppPagePath        string `json:"we_app_page_path"`
	MobileURL            string `json:"mobile_url"`
	MobileShortURL       string `json:"mobile_short_url"`
	WeAppWebViewURL      string `json:"we_app_web_view_url"`
	URL                  string `json:"url"`
	ShortURL             string `json:"short_url"`
	WeAppWebViewShortURL string `json:"we_app_web_view_short_url"`
	SchemaUrl            string `json:"schema_url"` //手动填充 该值拼多多不返回
}

type WeAppInfo struct {
	WeAppIconURL      string `json:"we_app_icon_url"`
	UserName          string `json:"user_name"`
	PagePath          string `json:"page_path"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	AppID             string `json:"app_id"`
	Desc              string `json:"desc"`
}
