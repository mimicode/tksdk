package alibabaalscunionkbitemquery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.item.query( 本地联盟口碑商品列表 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_item_query_response"`
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
	ErrorMessage string `json:"error_message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	SessionID  string  `json:"session_id"`
	PageNumber int64   `json:"page_number"`
	PageSize   int64   `json:"page_size"`
	Total      int64   `json:"total"`
	Records    Records `json:"records"`
}

type Records struct {
	KBItemPromotionDTO []KBItemPromotionDTO `json:"kb_item_promotion_d_t_o"`
}

type KBItemPromotionDTO struct {
	ItemID              string `json:"item_id"`
	Title               string `json:"title"`
	MainPicture         string `json:"main_picture"`
	OriginalPriceCent   int64  `json:"original_price_cent"`
	ActivityPriceCent   int64  `json:"activity_price_cent"`
	PriceWithCouponCent int64  `json:"price_with_coupon_cent"`
	CouponPriceCent     int64  `json:"coupon_price_cent"`
	NinetySales         int64  `json:"ninety_sales"`
	TotalSales          int64  `json:"total_sales"`
	Link                Link   `json:"link"`
}

type Link struct {
	AlipaySchemeURL string `json:"alipay_scheme_url"`
	AlipayH5URL     string `json:"alipay_h5_url"`
}
