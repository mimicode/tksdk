package pddddklotteryurlgen

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.lottery.url.gen多多客生成转盘抽免单url
type Response struct {
	response2.TopResponse
	LotteryURLResponse LotteryURLResponse `json:"lottery_url_response"`
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

type LotteryURLResponse struct {
	Total     int64     `json:"total"`
	URLList   []URLList `json:"url_list"`
	RequestID string    `json:"request_id"`
}

type URLList struct {
	WeAppPagePath                 interface{}   `json:"we_app_page_path"`
	MultiSchemaURL                interface{}   `json:"multi_schema_url"`
	MobileURL                     string        `json:"mobile_url"`
	Sign                          string        `json:"sign"`
	MultiQqAppPagePath            interface{}   `json:"multi_qq_app_page_path"`
	MultiWeAppWebViewShortURL     interface{}   `json:"multi_we_app_web_view_short_url"`
	MultiGroupMobileURL           interface{}   `json:"multi_group_mobile_url"`
	ShortURL                      interface{}   `json:"short_url"`
	QqAppInfo                     interface{}   `json:"qq_app_info"`
	SingleURLList                 SingleURLList `json:"single_url_list"`
	MultiGroupURL                 interface{}   `json:"multi_group_url"`
	MobileSuperShortURL           interface{}   `json:"mobile_super_short_url"`
	SchemaURL                     interface{}   `json:"schema_url"`
	WeAppInfo                     interface{}   `json:"we_app_info"`
	WeiboAppWebViewURL            interface{}   `json:"weibo_app_web_view_url"`
	WeAppWebViewURL               string        `json:"we_app_web_view_url"`
	MultiWeiboAppWebViewURL       interface{}   `json:"multi_weibo_app_web_view_url"`
	MultiWeAppWebViewURL          interface{}   `json:"multi_we_app_web_view_url"`
	MultiURLList                  interface{}   `json:"multi_url_list"`
	MultiGroupMobileShortURL      interface{}   `json:"multi_group_mobile_short_url"`
	WeiboAppWebViewShortURL       interface{}   `json:"weibo_app_web_view_short_url"`
	URL                           string        `json:"url"`
	MultiGroupMobileSuperShortURL interface{}   `json:"multi_group_mobile_super_short_url"`
	MultiWeAppPagePath            interface{}   `json:"multi_we_app_page_path"`
	MultiGroupShortURL            interface{}   `json:"multi_group_short_url"`
	MultiWeiboAppWebViewShortURL  interface{}   `json:"multi_weibo_app_web_view_short_url"`
	MobileShortURL                interface{}   `json:"mobile_short_url"`
	QqAppPagePath                 interface{}   `json:"qq_app_page_path"`
	WeAppWebViewShortURL          interface{}   `json:"we_app_web_view_short_url"`
}

type SingleURLList struct {
	WeAppPagePath        interface{} `json:"we_app_page_path"`
	MobileSuperShortURL  interface{} `json:"mobile_super_short_url"`
	MobileURL            string      `json:"mobile_url"`
	SchemaURL            interface{} `json:"schema_url"`
	MobileShortURL       interface{} `json:"mobile_short_url"`
	WeAppWebViewURL      string      `json:"we_app_web_view_url"`
	QqAppPagePath        interface{} `json:"qq_app_page_path"`
	URL                  string      `json:"url"`
	ShortURL             interface{} `json:"short_url"`
	WeAppWebViewShortURL interface{} `json:"we_app_web_view_short_url"`
}
