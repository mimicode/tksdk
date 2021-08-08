package pddddkoauthcmspromurlgenerate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.cms.prom.url.generate生成商城推广链接接口
type Response struct {
	response.TopResponse
	CMSPromotionURLGenerateResponse CMSPromotionURLGenerateResponse `json:"cms_promotion_url_generate_response"`
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

type CMSPromotionURLGenerateResponse struct {
	Total     int64     `json:"total"`
	URLList   []URLList `json:"url_list"`
	RequestID string    `json:"request_id"`
}

type URLList struct {
	MultiWeAppWebViewURL      string            `json:"multi_we_app_web_view_url"`
	MultiGroupMobileShortURL  string            `json:"multi_group_mobile_short_url"`
	MultiURLList              MultiURLListClass `json:"multi_url_list"`
	MobileURL                 string            `json:"mobile_url"`
	Sign                      string            `json:"sign"`
	MultiWeAppWebViewShortURL string            `json:"multi_we_app_web_view_short_url"`
	MultiGroupMobileURL       string            `json:"multi_group_mobile_url"`
	URL                       string            `json:"url"`
	ShortURL                  string            `json:"short_url"`
	MultiGroupURL             string            `json:"multi_group_url"`
	SingleURLList             MultiURLListClass `json:"single_url_list"`
	MultiGroupShortURL        string            `json:"multi_group_short_url"`
	WeAppInfo                 WeAppInfo         `json:"we_app_info"`
	MobileShortURL            string            `json:"mobile_short_url"`
	WeAppWebViewURL           string            `json:"we_app_web_view_url"`
	WeAppWebViewShortURL      string            `json:"we_app_web_view_short_url"`
	SchemaURL                 string            `json:"schema_url"`
}

type MultiURLListClass struct {
	MobileURL            string `json:"mobile_url"`
	SchemaURL            string `json:"schema_url"`
	MobileShortURL       string `json:"mobile_short_url"`
	WeAppWebViewURL      string `json:"we_app_web_view_url"`
	URL                  string `json:"url"`
	ShortURL             string `json:"short_url"`
	WeAppWebViewShortURL string `json:"we_app_web_view_short_url"`
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
