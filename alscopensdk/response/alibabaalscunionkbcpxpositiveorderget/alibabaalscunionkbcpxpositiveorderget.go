package alibabaalscunionkbcpxpositiveorderget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kbcpx.positive.order.get( 本地生活媒体推广订单明细报表查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kbcpx_positive_order_get_response"`
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
	OrderDetailReportDTO []OrderDetailReportDTO `json:"order_detail_report_d_t_o"`
}

type OrderDetailReportDTO struct {
	Title                  string `json:"title"`
	PicURL                 string `json:"pic_url"`
	ShopName               string `json:"shop_name"`
	PayAmount              string `json:"pay_amount"`
	SettleAmount           string `json:"settle_amount"`
	TraceTime              string `json:"trace_time"`
	TkCreateTime           string `json:"tk_create_time"`
	PayTime                string `json:"pay_time"`
	ReceiveTime            string `json:"receive_time"`
	SettleTime             string `json:"settle_time"`
	Income                 string `json:"income"`
	Settle                 string `json:"settle"`
	ItemID                 string `json:"item_id"`
	ProductNum             int64  `json:"product_num"`
	UnitPrice              string `json:"unit_price"`
	CategoryName           string `json:"category_name"`
	BizOrderID             int64  `json:"biz_order_id"`
	ParentOrderID          int64  `json:"parent_order_id"`
	MainItemID             string `json:"main_item_id"`
	MainItemTitle          string `json:"main_item_title"`
	OrderState             int64  `json:"order_state"`
	OrderItemStatusName    string `json:"order_item_status_name"`
	SettleState            int64  `json:"settle_state"`
	FullSettleAmount       string `json:"full_settle_amount"`
	CommissionRate         string `json:"commission_rate"`
	CommissionFee          string `json:"commission_fee"`
	SubsidyRate            string `json:"subsidy_rate"`
	SubsidyFee             string `json:"subsidy_fee"`
	IncomeRate             string `json:"income_rate"`
	StratifyRate           string `json:"stratify_rate"`
	DeductRate             string `json:"deduct_rate"`
	PlatformCommissionRate string `json:"platform_commission_rate"`
	PlatformCommissionFee  string `json:"platform_commission_fee"`
	ChannelRate            string `json:"channel_rate"`
	ChannelFee             string `json:"channel_fee"`
	MediaID                string `json:"media_id"`
	MediaName              string `json:"media_name"`
	AdZoneID               string `json:"ad_zone_id"`
	AdZoneName             string `json:"ad_zone_name"`
	ActivityFee            string `json:"activity_fee"`
	ActivityServiceFee     string `json:"activity_service_fee"`
	ActivityServiceRate    string `json:"activity_service_rate"`
	GmtModified            string `json:"gmt_modified"`
	Tag                    string `json:"tag"`
	Sid                    string `json:"sid"`
	PlatformType           int64  `json:"platform_type"`
	ActivityID             int64  `json:"activity_id"`
	UsedStoreID            string `json:"used_store_id"`
	PID                    string `json:"pid"`
	RelationOrderID        int64  `json:"relation_order_id"`
	FlowType               int64  `json:"flow_type"`
	OrderItemStatus        int64  `json:"order_item_status"`
	ActivityInfoRemarkList string `json:"activity_info_remark_list"`
	ChannelRightID         string `json:"channel_right_id"`
}
