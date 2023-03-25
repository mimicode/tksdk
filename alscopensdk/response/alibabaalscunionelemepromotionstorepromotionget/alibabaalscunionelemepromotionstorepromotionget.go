package alibabaalscunionelemepromotionstorepromotionget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.eleme.promotion.storepromotion.get( 本地联盟饿了么单店推广单店铺查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_eleme_promotion_storepromotion_get_response"`
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
	BizType                string        `json:"biz_type"`
	CommissionRate         string        `json:"commission_rate"`
	IndistinctMonthlySales string        `json:"indistinct_monthly_sales"`
	Link                   Link          `json:"link"`
	OverlayCoupon          OverlayCoupon `json:"overlay_coupon"`
	ShopLogo               string        `json:"shop_logo"`
	Title                  string        `json:"title"`
}

type Link struct {
	MiniQrcode string `json:"mini_qrcode"`
	Picture    string `json:"picture"`
	WxAppid    string `json:"wx_appid"`
	WxPath     string `json:"wx_path"`
}

type OverlayCoupon struct {
	Amount      string  `json:"amount"`
	EndTime     string  `json:"end_time"`
	StartTime   string  `json:"start_time"`
	TemplateID  float64 `json:"template_id"`
	ValidPeriod int64   `json:"valid_period"`
}
