package pddddkoauthgoodspromurlgenerate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.goods.prom.url.generate生成多多进宝推广链接
type Response struct {
	response.TopResponse
	GoodsPromotionURLGenerateResponse GoodsPromotionURLGenerateResponse `json:"goods_promotion_url_generate_response"`
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

type GoodsPromotionURLGenerateResponse struct {
	GoodsPromotionURLList []GoodsPromotionURLList `json:"goods_promotion_url_list"`
	RequestID             string                  `json:"request_id"`
}

type GoodsPromotionURLList struct {
	QqAppInfo            AppInfo `json:"qq_app_info"`
	MobileURL            string  `json:"mobile_url"`
	SchemaURL            string  `json:"schema_url"`
	WeAppInfo            AppInfo `json:"we_app_info"`
	MobileShortURL       string  `json:"mobile_short_url"`
	WeAppWebViewURL      string  `json:"we_app_web_view_url"`
	URL                  string  `json:"url"`
	ShortURL             string  `json:"short_url"`
	WeAppWebViewShortURL string  `json:"we_app_web_view_short_url"`
	FFXOpenAppUrl        string  `json:"ffx_open_app_url"`
}

type AppInfo struct {
	UserName          string  `json:"user_name"`
	PagePath          string  `json:"page_path"`
	QqAppIconURL      *string `json:"qq_app_icon_url,omitempty"`
	SourceDisplayName string  `json:"source_display_name"`
	Title             string  `json:"title"`
	AppID             string  `json:"app_id"`
	Desc              string  `json:"desc"`
	WeAppIconURL      *string `json:"we_app_icon_url,omitempty"`
}
