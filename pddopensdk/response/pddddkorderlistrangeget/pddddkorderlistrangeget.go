package pddddkorderlistrangeget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.order.list.range.get用时间段查询推广订单接口
type Response struct {
	response.TopResponse
	OrderListGetResponse OrderListGetResponse `json:"order_list_get_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type OrderListGetResponse struct {
	LastOrderID string      `json:"last_order_id"`
	OrderList   []OrderList `json:"order_list"`
}

type OrderList struct {
	ActivityTags               []int64 `json:"activity_tags"`
	AuthDuoID                  int64   `json:"auth_duo_id"`
	BandanRiskConsult          int64   `json:"bandan_risk_consult"`
	BatchNo                    string  `json:"batch_no"`
	CashGiftID                 int64   `json:"cash_gift_id"`
	CatIDS                     []int64 `json:"cat_ids"`
	CPANew                     int64   `json:"cpa_new"`
	CustomParameters           string  `json:"custom_parameters"`
	FailReason                 string  `json:"fail_reason"`
	GoodsCategoryName          string  `json:"goods_category_name"`
	GoodsID                    int64   `json:"goods_id"`
	GoodsName                  string  `json:"goods_name"`
	GoodsPrice                 int64   `json:"goods_price"`
	GoodsQuantity              int64   `json:"goods_quantity"`
	GoodsSign                  string  `json:"goods_sign"`
	GoodsThumbnailURL          string  `json:"goods_thumbnail_url"`
	GroupID                    int64   `json:"group_id"`
	IsDirect                   int64   `json:"is_direct"`
	MallID                     int64   `json:"mall_id"`
	MallName                   string  `json:"mall_name"`
	NoSubsidyReason            string  `json:"no_subsidy_reason"`
	OrderAmount                int64   `json:"order_amount"`
	OrderCreateTime            int64   `json:"order_create_time"`
	OrderGroupSuccessTime      int64   `json:"order_group_success_time"`
	OrderModifyAt              int64   `json:"order_modify_at"`
	OrderPayTime               int64   `json:"order_pay_time"`
	OrderReceiveTime           int64   `json:"order_receive_time"`
	OrderSettleTime            int64   `json:"order_settle_time"`
	OrderSn                    string  `json:"order_sn"`
	OrderStatus                int64   `json:"order_status"`
	OrderStatusDesc            string  `json:"order_status_desc"`
	OrderVerifyTime            int64   `json:"order_verify_time"`
	PID                        string  `json:"p_id"`
	PlatformDiscount           int64   `json:"platform_discount"`
	PriceCompareStatus         int64   `json:"price_compare_status"`
	PromotionAmount            int64   `json:"promotion_amount"`
	PromotionRate              int64   `json:"promotion_rate"`
	RedPacketType              int64   `json:"red_packet_type"`
	SepDuoID                   int64   `json:"sep_duo_id"`
	SepMarketFee               int64   `json:"sep_market_fee"`
	SepParameters              string  `json:"sep_parameters"`
	SepPID                     string  `json:"sep_pid"`
	SepRate                    int64   `json:"sep_rate"`
	ShareAmount                int64   `json:"share_amount"`
	ShareRate                  int64   `json:"share_rate"`
	SubsidyAmount              int64   `json:"subsidy_amount"`
	SubsidyDuoAmountLevel      int64   `json:"subsidy_duo_amount_level"`
	SubsidyDuoAmountTenMillion int64   `json:"subsidy_duo_amount_ten_million"`
	SubsidyOrderRemark         string  `json:"subsidy_order_remark"`
	SubsidyType                int64   `json:"subsidy_type"`
	Type                       int64   `json:"type"`
	ZsDuoID                    int64   `json:"zs_duo_id"`
}
