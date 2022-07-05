package tbkscorderdetailsget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.order.details.get( 淘宝客【服务商】所有订单查询 )
type Response struct {
	response.TopResponse
	TbkScOrderDetailsGetResponse Result `json:"tbk_sc_order_details_get_response"`
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

type Result struct {
	Data      Data   `json:"data"`
	RequestID string `json:"request_id"`
}

type Data struct {
	HasNext       bool    `json:"has_next"`
	HasPre        bool    `json:"has_pre"`
	PageNo        int64   `json:"page_no"`
	PageSize      int64   `json:"page_size"`
	PositionIndex string  `json:"position_index"`
	Results       Results `json:"results"`
}

type Results struct {
	PublisherOrderDto []PublisherOrderDto `json:"publisher_order_dto"`
}

type PublisherOrderDto struct {
	AdzoneID                           int64              `json:"adzone_id"`
	AdzoneName                         string             `json:"adzone_name"`
	AlimamaRate                        string             `json:"alimama_rate"`
	AlimamaShareFee                    string             `json:"alimama_share_fee"`
	AlipayTotalPrice                   string             `json:"alipay_total_price"`
	AppKey                             *string            `json:"app_key,omitempty"`
	ClickTime                          string             `json:"click_time"`
	DepositPrice                       string             `json:"deposit_price"`
	FlowSource                         string             `json:"flow_source"`
	IncomeRate                         string             `json:"income_rate"`
	ItemCategoryName                   string             `json:"item_category_name"`
	ItemID                             interface{}        `json:"item_id,omitempty"`
	ItemImg                            string             `json:"item_img"`
	ItemLink                           *string            `json:"item_link,omitempty"`
	ItemNum                            int64              `json:"item_num"`
	ItemPrice                          *string            `json:"item_price,omitempty"`
	ItemTitle                          string             `json:"item_title"`
	OrderType                          string             `json:"order_type"`
	PubID                              int64              `json:"pub_id"`
	PubShareFee                        string             `json:"pub_share_fee"`
	PubSharePreFee                     string             `json:"pub_share_pre_fee"`
	PubShareRate                       string             `json:"pub_share_rate"`
	RefundTag                          int64              `json:"refund_tag"`
	SellerNick                         string             `json:"seller_nick"`
	SellerShopTitle                    string             `json:"seller_shop_title"`
	ServiceFeeDtoList                  *ServiceFeeDtoList `json:"service_fee_dto_list,omitempty"`
	SiteID                             int64              `json:"site_id"`
	SiteName                           string             `json:"site_name"`
	SubsidyFee                         string             `json:"subsidy_fee"`
	SubsidyRate                        string             `json:"subsidy_rate"`
	SubsidyType                        string             `json:"subsidy_type"`
	TBDepositTime                      string             `json:"tb_deposit_time"`
	TBPaidTime                         string             `json:"tb_paid_time"`
	TerminalType                       string             `json:"terminal_type"`
	TkCommissionFeeForMediaPlatform    string             `json:"tk_commission_fee_for_media_platform"`
	TkCommissionPreFeeForMediaPlatform string             `json:"tk_commission_pre_fee_for_media_platform"`
	TkCommissionRateForMediaPlatform   string             `json:"tk_commission_rate_for_media_platform"`
	TkCreateTime                       string             `json:"tk_create_time"`
	TkDepositTime                      string             `json:"tk_deposit_time"`
	TkOrderRole                        int64              `json:"tk_order_role"`
	TkPaidTime                         string             `json:"tk_paid_time"`
	TkStatus                           int64              `json:"tk_status"`
	TkTotalRate                        string             `json:"tk_total_rate"`
	TotalCommissionFee                 string             `json:"total_commission_fee"`
	TotalCommissionRate                string             `json:"total_commission_rate"`
	TradeID                            string             `json:"trade_id"`
	TradeParentID                      string             `json:"trade_parent_id"`
	PayPrice                           *string            `json:"pay_price,omitempty"`
	TkEarningTime                      *string            `json:"tk_earning_time,omitempty"`
	RelationID                         int64              `json:"relation_id"`
	SpecialId                          int64              `json:"special_id"`
}

type ServiceFeeDtoList struct {
	ServiceFeeDto []ServiceFeeDto `json:"service_fee_dto"`
}

type ServiceFeeDto struct {
	SharePreFee       string `json:"share_pre_fee"`
	ShareFee          string `json:"share_fee"`
	ShareRelativeRate string `json:"share_relative_rate"`
	TkShareRoleType   int64  `json:"tk_share_role_type"`
}
