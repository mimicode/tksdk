package taobaotbkrelationrefund

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.relation.refund( 淘宝客-推广者-售后退款订单查询 )
type Response struct {
	response.TopResponse
	TbkRelationRefundResponse TbkRelationRefundResponse `json:"tbk_relation_refund_response"`
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

type TbkRelationRefundResponse struct {
	ResultMsg    string     `json:"result_msg"`
	Data         PageResult `json:"data"`
	ErrorCode    string     `json:"error_code"`
	ErrorMsg     string     `json:"error_msg"`
	SubCode      string     `json:"sub_code"`
	SubMsg       string     `json:"sub_msg"`
	RequestId    string     `json:"request_id"`
}

type PageResult struct {
	PageNo     string   `json:"page_no"`
	PageSize   string   `json:"page_size"`
	TotalCount string   `json:"total_count"`
	Results    []Result `json:"results"`
}

type Result struct {
	TbTradeParentId             int64  `json:"tb_trade_parent_id"`
	SpecialId                   int64  `json:"special_id"`
	RelationId                  int64  `json:"relation_id"`
	Tk3rdPubId                  int64  `json:"tk_3_rd_pub_id"`
	TkPubId                     int64  `json:"tk_pub_id"`
	TkSubsidyFeeRefund3rdPub    string `json:"tk_subsidy_fee_refund_3_rd_pub"`
	TkCommissionFeeRefund3rdPub string `json:"tk_commission_fee_refund_3_rd_pub"`
	TkSubsidyFeeRefundPub       string `json:"tk_subsidy_fee_refund_pub"`
	TkCommissionFeeRefundPub    string `json:"tk_commission_fee_refund_pub"`
	TkRefundSuitTime            string `json:"tk_refund_suit_time"`
	TkRefundTime                string `json:"tk_refund_time"`
	EarningTime                 string `json:"earning_time"`
	TbTradeCreateTime           string `json:"tb_trade_create_time"`
	RefundStatus                int64  `json:"refund_status"`
	TbAuctionTitle              string `json:"tb_auction_title"`
	TbTradeId                   int64  `json:"tb_trade_id"`
	RefundFee                   string `json:"refund_fee"`
	TbTradeFinishPrice          string `json:"tb_trade_finish_price"`
	TkPubShowReturnFee          string `json:"tk_pub_show_return_fee"`
	Tk3rdPubShowReturnFee       string `json:"tk_3_rd_pub_show_return_fee"`
	RefundType                  int64  `json:"refund_type"`
	AlscPid                     string `json:"alsc_pid"`
	AlscId                      string `json:"alsc_id"`
	ModifiedTime                string `json:"modified_time"`
}
