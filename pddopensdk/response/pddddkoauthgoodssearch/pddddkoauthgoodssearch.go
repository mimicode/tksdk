package pddddkoauthgoodssearch

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.goods.search多多进宝商品查询
type Response struct {
	response.TopResponse
	GoodsSearchResponse GoodsSearchResponse `json:"goods_search_response"`
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

type GoodsSearchResponse struct {
	GoodsList  []GoodsList `json:"goods_list"`
	ListID     string      `json:"list_id"`
	TotalCount int64       `json:"total_count"`
	RequestID  string      `json:"request_id"`
	SearchID   string      `json:"search_id"`
}

type GoodsList struct {
	CategoryName                string    `json:"category_name"`
	CouponRemainQuantity        int64     `json:"coupon_remain_quantity"`
	PromotionRate               int64     `json:"promotion_rate"`
	ServiceTags                 []int64   `json:"service_tags"`
	MallID                      int64     `json:"mall_id"`
	MallName                    string    `json:"mall_name"`
	MallCouponEndTime           int64     `json:"mall_coupon_end_time"`
	LgstTxt                     string    `json:"lgst_txt"`
	GoodsName                   string    `json:"goods_name"`
	GoodsID                     int64     `json:"goods_id"`
	BrandName                   string    `json:"brand_name"`
	PredictPromotionRate        int64     `json:"predict_promotion_rate"`
	GoodsDesc                   string    `json:"goods_desc"`
	OptName                     string    `json:"opt_name"`
	ShareRate                   int64     `json:"share_rate"`
	OptIDS                      []int64   `json:"opt_ids"`
	GoodsImageURL               string    `json:"goods_image_url"`
	HasMallCoupon               bool      `json:"has_mall_coupon"`
	UnifiedTags                 []*string `json:"unified_tags"`
	CouponStartTime             int64     `json:"coupon_start_time"`
	MinGroupPrice               int64     `json:"min_group_price"`
	CouponDiscount              int64     `json:"coupon_discount"`
	CouponEndTime               int64     `json:"coupon_end_time"`
	ZsDuoID                     int64     `json:"zs_duo_id"`
	MallCouponRemainQuantity    int64     `json:"mall_coupon_remain_quantity"`
	PlanType                    int64     `json:"plan_type"`
	CatIDS                      []int64   `json:"cat_ids"`
	CouponMinOrderAmount        int64     `json:"coupon_min_order_amount"`
	CategoryID                  int64     `json:"category_id"`
	MallCouponDiscountPct       int64     `json:"mall_coupon_discount_pct"`
	CouponTotalQuantity         int64     `json:"coupon_total_quantity"`
	MallCouponMinOrderAmount    int64     `json:"mall_coupon_min_order_amount"`
	MerchantType                int64     `json:"merchant_type"`
	SalesTip                    string    `json:"sales_tip"`
	OnlySceneAuth               bool      `json:"only_scene_auth"`
	DescTxt                     string    `json:"desc_txt"`
	MallCouponID                int64     `json:"mall_coupon_id"`
	GoodsThumbnailURL           string    `json:"goods_thumbnail_url"`
	OptID                       int64     `json:"opt_id"`
	SearchID                    string    `json:"search_id"`
	ActivityTags                []int64   `json:"activity_tags"`
	HasCoupon                   bool      `json:"has_coupon"`
	MinNormalPrice              int64     `json:"min_normal_price"`
	MallCouponStartTime         int64     `json:"mall_coupon_start_time"`
	ServTxt                     string    `json:"serv_txt"`
	MallCouponTotalQuantity     int64     `json:"mall_coupon_total_quantity"`
	MallCouponMaxDiscountAmount int64     `json:"mall_coupon_max_discount_amount"`
	MallCPS                     int64     `json:"mall_cps"`
	GoodsSign                   string    `json:"goods_sign"`
	ActivityType                *int64    `json:"activity_type,omitempty"`
	ExtraCouponAmount           *int64    `json:"extra_coupon_amount,omitempty"`
	CltCpnEndTime               *int64    `json:"clt_cpn_end_time,omitempty"`
	CltCpnMinAmt                *int64    `json:"clt_cpn_min_amt,omitempty"`
	CltCpnRemainQuantity        *int64    `json:"clt_cpn_remain_quantity,omitempty"`
	CltCpnBatchSn               *string   `json:"clt_cpn_batch_sn,omitempty"`
	CltCpnDiscount              *int64    `json:"clt_cpn_discount,omitempty"`
	CltCpnQuantity              *int64    `json:"clt_cpn_quantity,omitempty"`
	CltCpnStartTime             *int64    `json:"clt_cpn_start_time,omitempty"`
}
