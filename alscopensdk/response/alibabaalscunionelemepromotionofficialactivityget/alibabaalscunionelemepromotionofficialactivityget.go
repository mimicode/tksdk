package alibabaalscunionelemepromotionofficialactivityget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.eleme.promotion.officialactivity.get( 本地联盟饿了么推广官方活动查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_eleme_promotion_officialactivity_get_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int    `json:"biz_error_code"`
	RequestID    string `json:"request_id"`
	Data         Data   `json:"data"`
	ResultCode   int64  `json:"result_code"`
	Message      string `json:"message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

// Data 活动数据
type Data struct {
	ID          string `json:"id"`          // 活动ID
	Title       string `json:"title"`       // 标题
	Description string `json:"description"` // 描述
	Picture     string `json:"picture"`     // 活动创意图片
	StartTime   int64  `json:"start_time"`  // 起始时间（秒）
	EndTime     int64  `json:"end_time"`    // 结束时间（秒）
	Link        Link   `json:"link"`        // 推广链接
}

// Link 推广链接
type Link struct {
	WxAppid         string                `json:"wx_appid"`         // 小程序appId
	WxPath          string                `json:"wx_path"`          // 微信小程序path链接
	Picture         string                `json:"picture"`          // 推广图片地址
	AlipayMiniURL   string                `json:"alipay_mini_url"`  // 支付宝小程序推广链接
	H5URL           string                `json:"h5_url"`           // h5推广地址
	TBQrCode        string                `json:"tb_qr_code"`       // 淘宝二维码图片地址
	MiniQrcode      string                `json:"mini_qrcode"`      // 微信独立二维码
	TBMiniQrcode    string                `json:"tb_mini_qrcode"`   // 淘宝独立二维码
	EleSchemeURL    string                `json:"ele_scheme_url"`   // 饿了么唤端链接
	H5ShortLink     string                `json:"h5_short_link"`    // h5推广地址短链
	H5MiniQrcode    string                `json:"h5_mini_qrcode"`   // H5二维码图片地址
	ElemeWord       string                `json:"eleme_word"`       // 饿了么口令（有效期30天，建议到期前重新获取）
	TbSchemeURL     string                `json:"tb_scheme_url"`    // 淘宝唤端链接
	TaobaoWord      string                `json:"taobao_word"`      // 淘口令（有效期30天，建议到期前重新获取）
	FullTaobaoWord  string                `json:"full_taobao_word"` // 带文案淘口令（有效期30天，建议到期前重新获取）
	H5Promotion     TopH5PromotionDto     `json:"h5_promotion"`     // 推广链接
	TaobaoPromotion TopTaobaoPromotionDto `json:"taobao_promotion"` // 淘宝链接
	WxPromotion     TopWxPromotionDto     `json:"wx_promotion"`     // 微信链接
	AlipayPromotion TopAlipayPromotionDto `json:"alipay_promotion"` // 支付宝链接
	AppPromotion    TopAppPromotionDto    `json:"app_promotion"`    // app链接
}

// TopH5PromotionDto 推广链接
type TopH5PromotionDto struct {
	H5URL     string `json:"h5_url"`     // 长连接
	ShortLink string `json:"short_link"` // 短连接
	H5QrCode  string `json:"h5_qr_code"` // 推广二维码
	H5ImgURL  string `json:"h5_img_url"` // 推广海报
	TjH5URL   string `json:"tj_h5_url"`  // 推荐有奖拉新链接
}

// TopTaobaoPromotionDto 淘宝链接
type TopTaobaoPromotionDto struct {
	H5URL              string `json:"h5_url"`               // 推广链接
	H5ShortURL         string `json:"h5_short_url"`         // 短连接
	TbQrCode           string `json:"tb_qr_code"`           // 二维码
	ImgURL             string `json:"img_url"`              // 海报
	AppID              string `json:"app_id"`               // appid
	AppPath            string `json:"app_path"`             // apppath
	SchemeURL          string `json:"scheme_url"`           // schemeurl
	TbWatchword        string `json:"tb_watchword"`         // 淘口令
	TbWatchwordSuggest string `json:"tb_watchword_suggest"` // 淘口令完整版
}

// TopWxPromotionDto 微信链接
type TopWxPromotionDto struct {
	WxAppID         string `json:"wx_app_id"`          // appid
	WxPath          string `json:"wx_path"`            // apppath
	WxQrCode        string `json:"wx_qr_code"`         // 二维码
	ImgURL          string `json:"img_url"`            // 海报
	ReduceWxAppID   string `json:"reduce_wx_app_id"`   // 立减的appid
	ReduceWxPath    string `json:"reduce_wx_path"`     // 立减的apppath
	ReduceWxQrCode  string `json:"reduce_wx_qr_code"`  // 立减的二维码
	ReduceImgURL    string `json:"reduce_img_url"`     // 立减的海报
	H5ShortLink     string `json:"h5_short_link"`      // h5短链
	MarketWxAppID   string `json:"market_wx_app_id"`   // 媒体出资appid
	MarketWxAppPath string `json:"market_wx_app_path"` // 媒体出资apppath
}

// TopAlipayPromotionDto 支付宝链接
type TopAlipayPromotionDto struct {
	AppID                  string `json:"app_id"`                   // appId
	AppPath                string `json:"app_path"`                 // appPath
	AlipaySchemeURL        string `json:"alipay_scheme_url"`        // scheme唤端地址
	H5URL                  string `json:"h5_url"`                   // h5地址
	AlipayWatchword        string `json:"alipay_watchword"`         // 支付宝口令
	AlipayWatchwordSuggest string `json:"alipay_watchword_suggest"` // 支付宝完整口令
	ImgURL                 string `json:"img_url"`                  // 海报地址
	AlipayQrCode           string `json:"alipay_qr_code"`           // 二维码地址
	ShortLink              string `json:"short_link"`               // h5短连接
	H5SchemeURL            string `json:"h5_scheme_url"`            // http换端地址
}

// TopAppPromotionDto app链接
type TopAppPromotionDto struct {
	DeepLink      string `json:"deep_link"`       // deeplink地址
	ElemeWord     string `json:"eleme_word"`      // 饿口令
	WordValidDate string `json:"word_valid_date"` // 饿口令有效期
}
