package taobaotbkorderrefundget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.order.refund.get( 淘宝客-推广者-全量售后退款订单查询 )
type Response struct {
	response.TopResponse
	TbkOrderRefundGetResponse TbkOrderRefundGetResponse `json:"tbk_order_refund_get_response"`
}

//解析输出结果
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

type TbkOrderRefundGetResponse struct {
	Result     Result `json:"result"`
	ErrorCode  string `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
	SubCode    string `json:"sub_code"`
	SubMsg     string `json:"sub_msg"`
	RequestId  string `json:"request_id"`
}

type Result struct {
	PageNo       int64                     `json:"page_no"`
	PageSize     int64                     `json:"page_size"`
	HasNext      bool                      `json:"has_next"`
	HasPre       bool                      `json:"has_pre"`
	PositionIndex string                    `json:"position_index"`
	Results      []PublisherRefundOrderDTO `json:"result"`
}

type PublisherRefundOrderDTO struct {
	RefundSuitId                    string `json:"refund_suit_id"`
	TbTradeParentId                 string `json:"tb_trade_parent_id"`
	TbTradeId                       string `json:"tb_trade_id"`
	TbTradeCreateTime               string `json:"tb_trade_create_time"`
	EarningTime                     string `json:"earning_time"`
	TkRefundTime                    string `json:"tk_refund_time"`
	TkRefundSuitTime                string `json:"tk_refund_suit_time"`
	ModifiedTime                    string `json:"modified_time"`
	ItemTitle                       string `json:"item_title"`
	TkOrderRole                     string `json:"tk_order_role"`
	RefundStatus                    int64  `json:"refund_status"`
	TbTradeFinishPrice              string `json:"tb_trade_finish_price"`
	RefundFee                       string `json:"refund_fee"`
	PubShareFee                     string `json:"pub_share_fee"`
	TkCommissionFeeRefund           string `json:"tk_commission_fee_refund"`
	TkSubsidyFeeRefund              string `json:"tk_subsidy_fee_refund"`
	TkCommissionAlimmRefundFee      string `json:"tk_commission_alimm_refund_fee"`
	TkSubsidyAlimmRefundFee         string `json:"tk_subsidy_alimm_refund_fee"`
	TkCommissionAgentRefundFee      string `json:"tk_commission_agent_refund_fee"`
	TkSubsidyAgentRefundFee         string `json:"tk_subsidy_agent_refund_fee"`
	ShowReturnFee                   string `json:"show_return_fee"`
	RelationId                      int64  `json:"relation_id"`
	SpecialId                       int64  `json:"special_id"`
}
