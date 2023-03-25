package alibabaalscunionkborderquery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.order.query( openapi订单查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_order_query_response"`
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

type Data struct {
	BizOrderID  string      `json:"biz_order_id"`
	VoucherList VoucherList `json:"voucher_list"`
	OrderStatus string      `json:"order_status"`
}

type VoucherList struct {
	OrderVoucherDetailDto []OrderVoucherDetailDto `json:"order_voucher_detail_dto"`
}

type OrderVoucherDetailDto struct {
	ItemID        string `json:"item_id"`
	Title         string `json:"title"`
	VoucherID     string `json:"voucher_id"`
	TicketCode    string `json:"ticket_code"`
	VoucherStatus string `json:"voucher_status"`
	TotalAmount   int64  `json:"total_amount"`
	UsedAmount    int64  `json:"used_amount"`
	RefundAmount  int64  `json:"refund_amount"`
	EffectTime    string `json:"effect_time"`
	ExpireTime    string `json:"expire_time"`
	EXTInfo       string `json:"ext_info"`
	RefundType    string `json:"refund_type"`
	TicketURL     string `json:"ticket_url"`
}
