package pddddkoauthorderdetailget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.order.detail.get获取订单详情
type Response struct {
	response.TopResponse
	OrderDetailResponse OrderDetailResponse `json:"order_detail_response"`
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

type OrderDetailResponse struct {
	SepMarketFee          int64   `json:"sep_market_fee"`
	GoodsPrice            int64   `json:"goods_price"`
	SepDuoID              int64   `json:"sep_duo_id"`
	PID                   string  `json:"pid"`
	PromotionRate         int64   `json:"promotion_rate"`
	CPSSign               string  `json:"cps_sign"`
	Type                  int64   `json:"type"`
	SubsidyDuoAmountLevel int64   `json:"subsidy_duo_amount_level"`
	OrderStatus           int64   `json:"order_status"`
	CatIDS                []int64 `json:"cat_ids"`
	OrderCreateTime       int64   `json:"order_create_time"`
	IsDirect              int64   `json:"is_direct"`
	MallID                int64   `json:"mall_id"`
	OrderAmount           int64   `json:"order_amount"`
	PriceCompareStatus    int64   `json:"price_compare_status"`
	OrderModifyAt         int64   `json:"order_modify_at"`
	AuthDuoID             int64   `json:"auth_duo_id"`
	CPANew                int64   `json:"cpa_new"`
	GoodsName             string  `json:"goods_name"`
	BatchNo               string  `json:"batch_no"`
	URLLastGenerateTime   int64   `json:"url_last_generate_time"`
	GoodsQuantity         int64   `json:"goods_quantity"`
	GoodsID               int64   `json:"goods_id"`
	SepParameters         string  `json:"sep_parameters"`
	SepRate               int64   `json:"sep_rate"`
	SubsidyType           int64   `json:"subsidy_type"`
	ShareRate             int64   `json:"share_rate"`
	CustomParameters      string  `json:"custom_parameters"`
	GoodsThumbnailURL     string  `json:"goods_thumbnail_url"`
	PromotionAmount       int64   `json:"promotion_amount"`
	OrderPayTime          int64   `json:"order_pay_time"`
	ActivityTags          []int64 `json:"activity_tags"`
	GroupID               float64 `json:"group_id"`
	SepPID                string  `json:"sep_pid"`
	ReturnStatus          int64   `json:"return_status"`
	OrderStatusDesc       string  `json:"order_status_desc"`
	ShareAmount           int64   `json:"share_amount"`
	RequestID             string  `json:"request_id"`
	GoodsSign             string  `json:"goods_sign"`
	OrderSn               string  `json:"order_sn"`
	ZsDuoID               int64   `json:"zs_duo_id"`
}
