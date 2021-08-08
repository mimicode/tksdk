package pddddkallorderlistincrementget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.all.order.list.increment.get查询所有授权的多多客订单
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
	TotalCount int64       `json:"total_count"`
	OrderList  []OrderList `json:"order_list"`
	RequestID  string      `json:"request_id"`
}

type OrderList struct {
	SepMarketFee          int64   `json:"sep_market_fee"`
	GoodsPrice            int64   `json:"goods_price"`
	SepDuoID              int64   `json:"sep_duo_id"`
	PromotionRate         int64   `json:"promotion_rate"`
	Type                  int64   `json:"type"`
	SubsidyDuoAmountLevel int64   `json:"subsidy_duo_amount_level"`
	CatIDS                []int64 `json:"cat_ids"`
	OrderStatus           int64   `json:"order_status"`
	OrderCreateTime       int64   `json:"order_create_time"`
	IsDirect              int64   `json:"is_direct"`
	OrderGroupSuccessTime int64   `json:"order_group_success_time"`
	MallID                int64   `json:"mall_id"`
	OrderAmount           int64   `json:"order_amount"`
	PriceCompareStatus    int64   `json:"price_compare_status"`
	OrderModifyAt         int64   `json:"order_modify_at"`
	AuthDuoID             int64   `json:"auth_duo_id"`
	CPANew                int64   `json:"cpa_new"`
	GoodsName             string  `json:"goods_name"`
	GoodsQuantity         int64   `json:"goods_quantity"`
	GoodsID               int64   `json:"goods_id"`
	SepRate               int64   `json:"sep_rate"`
	SubsidyType           int64   `json:"subsidy_type"`
	CustomParameters      string  `json:"custom_parameters"`
	GoodsThumbnailURL     string  `json:"goods_thumbnail_url"`
	ShareRate             int64   `json:"share_rate"`
	PromotionAmount       int64   `json:"promotion_amount"`
	OrderPayTime          int64   `json:"order_pay_time"`
	ActivityTags          []int64 `json:"activity_tags"`
	GroupID               float64 `json:"group_id"`
	SceneAtMarketFee      int64   `json:"scene_at_market_fee"`
	OrderStatusDesc       string  `json:"order_status_desc"`
	ShareAmount           int64   `json:"share_amount"`
	GoodsSign             string  `json:"goods_sign"`
	OrderSn               string  `json:"order_sn"`
	PID                   string  `json:"p_id"`
	ZsDuoID               int64   `json:"zs_duo_id"`
}
