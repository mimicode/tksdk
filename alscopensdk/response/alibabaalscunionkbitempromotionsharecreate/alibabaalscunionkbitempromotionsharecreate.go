package alibabaalscunionkbitempromotionsharecreate

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.item.promotion.share.create( 本地生活媒体创建商品推广链接 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_item_promotion_share_create_response"`
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
	BizErrorDesc  string `json:"biz_error_desc"`
	BizErrorCode  int    `json:"biz_error_code"`
	RequestID     string `json:"request_id"`
	Result        Result `json:"result"`
	ResultSuccess bool   `json:"result_success"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Result struct {
	ImgURL                 string `json:"img_url"`
	WxAppid                string `json:"wx_appid"`
	WxPath                 string `json:"wx_path"`
	MiniQrCode             string `json:"mini_qr_code"`
	AlipayImgURL           string `json:"alipay_img_url"`
	AlipayWatchword        string `json:"alipay_watchword"`
	AlipayWatchwordSuggest string `json:"alipay_watchword_suggest"`
	AlipayMiniQrCode       string `json:"alipay_mini_qr_code"`
	AlipaySchemeURL        string `json:"alipay_scheme_url"`
	H5URL                  string `json:"h5_url"`
}
