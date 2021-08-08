package pddddkresourceurlgen

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.resource.url.gen（生成多多进宝频道推广）
type Response struct {
	response2.TopResponse
	ResourceURLResponse ResponseResourceURLResponse `json:"resource_url_response"`
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

type ResponseResourceURLResponse struct {
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
	SchemaUrl            string `json:"schema_url"`
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
