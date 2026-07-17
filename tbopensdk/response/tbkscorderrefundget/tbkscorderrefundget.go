package tbkscorderrefundget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.order.refund.get( 淘宝客-推广-全量售后退款订单查询 )
// https://developer.alibaba.com/docs/api.htm?apiId=65381
type Response struct {
	response.TopResponse
	TbkScOrderRefundGetResponse TbkScOrderRefundGetResponse `json:"tbk_sc_order_refund_get_response"`
}

// 解析输出结果
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

type TbkScOrderRefundGetResponse struct {
	Data         Data   `json:"data"`           //返回数据
	ResultCode   int64  `json:"result_code"`    //接口返回值信息，跟rpc架构保持一致
	BizErrorDesc string `json:"biz_error_desc"` //业务错误信息
	BizErrorCode int64  `json:"biz_error_code"` //业务错误码 101,102,103
	ResultMsg    string `json:"result_msg"`     //返回信息
}

type Data struct {
	Result        []PublisherRefundOrderDTO `json:"result"`         //真正的业务数据结构，售后退款订单明细列表
	PrePage       int64                     `json:"pre_page"`       //上一页
	NextPage      int64                     `json:"next_page"`      //下一页
	PageNo        int64                     `json:"page_no"`        //页码
	PageSize      int64                     `json:"page_size"`      //每页订单量
	HasNext       bool                      `json:"has_next"`       //是否有下一页
	PositionIndex string                    `json:"position_index"` //位点字段，由调用方原样传递
	HasPre        bool                      `json:"has_pre"`        //是否有上一页
}

type PublisherRefundOrderDTO struct {
	RefundSuitId               string `json:"refund_suit_id"`                 //维权编号，当前订单发生维权退款的编号（非淘宝订单编号），如订单发生多次维权则会产生多个维权编号
	TbTradeParentId            string `json:"tb_trade_parent_id"`             //淘宝父订单编号（买家在淘宝后台显示的订单编号）
	TbTradeId                  string `json:"tb_trade_id"`                    //淘宝子订单编号
	TbTradeCreateTime          string `json:"tb_trade_create_time"`           //订单创建时间
	EarningTime                string `json:"earning_time"`                   //订单结算时间
	TkRefundTime               string `json:"tk_refund_time"`                 //维权创建时间
	TkRefundSuitTime           string `json:"tk_refund_suit_time"`            //维权完成时间
	ModifiedTime               string `json:"modified_time"`                  //订单更新时间
	ItemTitle                  string `json:"item_title"`                     //商品标题
	TkOrderRole                string `json:"tk_order_role"`                  //推广者角色（二方、三方）
	RefundStatus               int64  `json:"refund_status"`                  //维权状态：4-维权创建，2-维权成功，3-维权失败，11-发生多次维权待处理，12/13-从淘客处补扣等待/成功，14/15-从卖家处补扣等待/成功
	TbTradeFinishPrice         string `json:"tb_trade_finish_price"`          //结算金额（订单确认收货后的成交金额）
	RefundFee                  string `json:"refund_fee"`                     //维权金额（买家申请维权退款的金额）
	PubShareFee                string `json:"pub_share_fee"`                  //结算预估收入=结算金额*提成，以订单确认收货后的成交金额为基数预估的收入
	TkCommissionFeeRefund      string `json:"tk_commission_fee_refund"`       //应退还佣金（不含技术服务费和渠道专项服务费）
	TkSubsidyFeeRefund         string `json:"tk_subsidy_fee_refund"`          //应退还补贴（不含技术服务费和渠道专项服务费）
	TkCommissionAlimmRefundFee string `json:"tk_commission_alimm_refund_fee"` //应退还佣金对应的技术服务费
	TkSubsidyAlimmRefundFee    string `json:"tk_subsidy_alimm_refund_fee"`    //应退还补贴对应的技术服务费
	TkCommissionAgentRefundFee string `json:"tk_commission_agent_refund_fee"` //应退还佣金对应的渠道专项服务费
	TkSubsidyAgentRefundFee    string `json:"tk_subsidy_agent_refund_fee"`    //应退还补贴对应的渠道专项服务费
	ShowReturnFee              string `json:"show_return_fee"`                //应退还预估收入：订单发生维权退款应退还的预估收入（佣金+补贴，含技术服务费和渠道专项服务费）
	RelationId                 int64  `json:"relation_id"`                    //渠道关系id
	SpecialId                  int64  `json:"special_id"`                     //会员关系id
	RightsId                   string `json:"rights_id"`                      //淘礼金ID，每次创建的淘礼金有唯一的识别ID，可在订单中查询
}
