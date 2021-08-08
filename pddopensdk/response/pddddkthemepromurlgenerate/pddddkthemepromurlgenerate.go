package pddddkthemepromurlgenerate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.theme.prom.url.generate（多多进宝主题推广链接生成
type Response struct {
	response.TopResponse
	ThemePromotionURLGenerateResponse ResponseThemePromotionURLGenerateResponse `json:"theme_promotion_url_generate_response"`
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

type ResponseThemePromotionURLGenerateResponse struct {
	URLList   []ResponseURLList `json:"url_list"`
	RequestID string            `json:"request_id"`
}

type ResponseURLList struct {
	MultiGroupMobileShortURL string    `json:"multi_group_mobile_short_url"`
	MultiGroupURL            string    `json:"multi_group_url"`
	MobileURL                string    `json:"mobile_url"`
	MultiGroupShortURL       string    `json:"multi_group_short_url"`
	SchemaURL                string    `json:"schema_url"`
	WeAppInfo                WeAppInfo `json:"we_app_info"`
	MobileShortURL           string    `json:"mobile_short_url"`
	MultiGroupMobileURL      string    `json:"multi_group_mobile_url"`
	URL                      string    `json:"url"`
	ShortURL                 string    `json:"short_url"`
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
