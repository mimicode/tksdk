package alibabaalscunionkbcpxpunishorderget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kbcpx.punish.order.get( 本地生活媒体推广用户反作弊订单数据报表 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kbcpx_punish_order_get_response"`
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
	TotalCount    int64  `json:"total_count"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Result struct {
	PunishOrderDetailReportDTO []PunishOrderDetailReportDTO `json:"punish_order_detail_report_d_t_o"`
}

type PunishOrderDetailReportDTO struct {
	Title                 string `json:"title"`
	PicURL                string `json:"pic_url"`
	ShopName              string `json:"shop_name"`
	PayAmount             string `json:"pay_amount"`
	SettleTime            string `json:"settle_time"`
	ExplainStartTime      string `json:"explain_start_time"`
	ExplainEndTime        string `json:"explain_end_time"`
	SerialNo              string `json:"serial_no"`
	MainItemID            string `json:"main_item_id"`
	MainItemTitle         string `json:"main_item_title"`
	ExplainState          string `json:"explain_state"`
	ReturnCommissionState int64  `json:"return_commission_state"`
	Settle                string `json:"settle"`
	GmtModified           string `json:"gmt_modified"`
	Sid                   string `json:"sid"`
	PlatformType          int64  `json:"platform_type"`
	ActivityID            int64  `json:"activity_id"`
	BizOrderID            string `json:"biz_order_id"`
}
