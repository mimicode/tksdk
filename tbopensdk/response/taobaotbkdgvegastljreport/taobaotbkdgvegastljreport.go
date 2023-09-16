package taobaotbkdgvegastljreport

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.vegas.tlj.report( 淘宝客-推广者-淘礼金效果数据 )
type Response struct {
	response.TopResponse
	TbkDgVegasTljReportResponse TbkDgVegasTljReportResponse `json:"tbk_dg_vegas_tlj_report_response"`
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

type TbkDgVegasTljReportResponse struct {
	Model         Model  `json:"model"`
	MsgInfo       string `json:"msg_info"`
	MsgCode       string `json:"msg_code"`
	ResultSuccess bool   `json:"result_success"`
}

type Model struct {
	Extra Extra `json:"extra"`
}

type Extra struct {
	GetRate               string `json:"get_rate"`
	UseRate               string `json:"use_rate"`
	AlipayNum             int64  `json:"alipay_num"`
	AlipayAmt             string `json:"alipay_amt"`
	PrePubShareFeeForDisp string `json:"pre_pub_share_fee_for_disp"`
	CMSettleAmt           string `json:"cm_settle_amt"`
	WinPV                 int64  `json:"win_pv"`
	WinSumAmt             string `json:"win_sum_amt"`
	RemainingNum          int64  `json:"remaining_num"`
	RemainingAmt          string `json:"remaining_amt"`
	UseNum                int64  `json:"use_num"`
	UseSumAmt             string `json:"use_sum_amt"`
	RefundNum             int64  `json:"refund_num"`
	RefundSumAmt          string `json:"refund_sum_amt"`
}
