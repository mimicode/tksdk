package pddddkcmspromurlgenerate

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.cms.prom.url.generate生成商城-频道推广链接
type Response struct {
	response2.TopResponse
	CMSPromotionURLGenerateResponse CMSPromotionURLGenerateResponse `json:"cms_promotion_url_generate_response"`
}

//解析输出结果
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
	SingleURLList        SingleURLList `json:"single_url_list"`
	MobileURL            string        `json:"mobile_url"`
	Sign                 string        `json:"sign"`
	WeAppInfo            WeAppInfo     `json:"we_app_info"`
	MobileShortURL       string        `json:"mobile_short_url"`
	WeAppWebViewURL      string        `json:"we_app_web_view_url"`
	URL                  string        `json:"url"`
	ShortURL             string        `json:"short_url"`
	WeAppWebViewShortURL string        `json:"we_app_web_view_short_url"`
	SchemaURL            string        `json:"schema_url"`
}

type SingleURLList struct {
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
