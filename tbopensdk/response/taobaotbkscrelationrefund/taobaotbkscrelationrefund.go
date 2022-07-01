package taobaotbkscrelationrefund

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.relation.refund( 淘宝客-服务商-维权退款订单查询 )
type Response struct {
	response.TopResponse
	TbkScRelationRefundResponse TbkScRelationRefundResponse `json:"tbk_sc_relation_refund_response"`
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

type TbkScRelationRefundResponse struct {
	ResultMsg    string `json:"result_msg"`
	ResultCode   int64  `json:"result_code"`
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int64  `json:"biz_error_code"`
	Data         Data   `json:"data"`
}

type Data struct {
	PageNo     int64    `json:"page_no"`
	PageSize   int64    `json:"page_size"`
	TotalCount int64    `json:"total_count"`
	Results    []Result `json:"results"`
}

type Result struct {
	TbTradeParentId             int64       `json:"tb_trade_parent_id"`
	SpecialId                   int64       `json:"special_id"`
	RelationId                  int64       `json:"relation_id"`
	Tk3rdPubId                  int64       `json:"tk_3_rd_pub_id"`
	TkPubId                     int64       `json:"tk_pub_id"`
	TkSubsidyFeeRefund3rdPub    string      `json:"tk_subsidy_fee_refund_3_rd_pub"`
	TkCommissionFeeRefund3rdPub string      `json:"tk_commission_fee_refund_3_rd_pub"`
	TkSubsidyFeeRefundPub       string      `json:"tk_subsidy_fee_refund_pub"`
	TkCommissionFeeRefundPub    string      `json:"tk_commission_fee_refund_pub"`
	TkRefundSuitTime            string      `json:"tk_refund_suit_time"`
	TkRefundTime                string      `json:"tk_refund_time"`
	EarningTime                 string      `json:"earning_time"`
	TbTradeCreateTime           string      `json:"tb_trade_create_time"`
	RefundStatus                int         `json:"refund_status"` //维权创建(淘客结算回执) 4,维权成功(淘客结算回执) 2,维权失败(淘客结算回执) 3,发生多次维权，待处理 11,从淘客处补扣（钱已结给淘客） 等待扣款 12,从淘客处补扣（钱已结给淘客） 扣款成功 13,从卖家处补扣（钱已结给卖家） 等待扣款 14,从卖家处补扣（钱已结给卖家） 扣款成功 15
	TbAuctionTitle              string      `json:"tb_auction_title"`
	TbTradeId                   int64       `json:"tb_trade_id"`
	RefundFee                   interface{} `json:"refund_fee"`
	TbTradeFinishPrice          interface{} `json:"tb_trade_finish_price"`
	TkPubShowReturnFee          interface{} `json:"tk_pub_show_return_fee"`
	Tk3rdPubShowReturnFee       interface{} `json:"tk_3_rd_pub_show_return_fee"`
	RefundType                  int         `json:"refund_type"`
	AlscPid                     string      `json:"alsc_pid"`
	AlscId                      string      `json:"alsc_id"`
	ModifiedTime                string      `json:"modified_time"`
}
