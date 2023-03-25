package alibabaalscunionelemepromotionstorepromotionquery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.eleme.promotion.storepromotion.query( 本地联盟饿了么单店推广店铺列表 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_eleme_promotion_storepromotion_query_response"`
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
	Message      string `json:"message"`
	ResultCode   int64  `json:"result_code"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	PageSize  int64   `json:"page_size"`
	Records   Records `json:"records"`
	SessionID string  `json:"session_id"`
}

type Records struct {
	StorePromotionDto []StorePromotionDto `json:"store_promotion_dto"`
}

type StorePromotionDto struct {
	BizType                string `json:"biz_type"`
	Commission             int64  `json:"commission"`
	CommissionRate         string `json:"commission_rate"`
	IndistinctMonthlySales string `json:"indistinct_monthly_sales"`
	Link                   Link   `json:"link"`
	ShopID                 string `json:"shop_id"`
	ShopLogo               string `json:"shop_logo"`
	Title                  string `json:"title"`
}

type Link struct {
	WxAppid string `json:"wx_appid"`
	WxPath  string `json:"wx_path"`
}
