package pddddkgoodssearch

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.goods.search（多多进宝商品查询）
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
	TotalCount int64       `json:"total_count"`
	RequestID  string      `json:"request_id"`
	ListID     string      `json:"list_id"`
	SearchID   string      `json:"search_id"`
}

type GoodsList struct {
	AvgDesc int64 `json:"avg_desc"`

	AvgServ int64    `json:"avg_serv"`
	AvgLgst int64    `json:"avg_lgst"`
	ServPct *float64 `json:"serv_pct"`

	SoldQuantity int64 `json:"sold_quantity"`

	LgstPct *float64 `json:"lgst_pct"`

	GoodsEvalScore *float64    `json:"goods_eval_score"`
	CatID          interface{} `json:"cat_id"`

	DescPct *float64 `json:"desc_pct"`

	GoodsEvalCount int64 `json:"goods_eval_count"`

	GoodsGalleryUrls interface{} `json:"goods_gallery_urls"`

	MallRate int64 `json:"mall_rate"`

	CreateAt int64 `json:"create_at"`

	MallCouponRemainQuantity    int64   `json:"mall_coupon_remain_quantity"`
	PlanType                    int64   `json:"plan_type"`
	CategoryName                string  `json:"category_name"`
	CouponRemainQuantity        int64   `json:"coupon_remain_quantity"`
	PromotionRate               int64   `json:"promotion_rate"`
	CatIDS                      []int64 `json:"cat_ids"`
	CouponMinOrderAmount        int64   `json:"coupon_min_order_amount"`
	CategoryID                  int64   `json:"category_id"`
	MallCouponDiscountPct       int64   `json:"mall_coupon_discount_pct"`
	MallID                      int64   `json:"mall_id"`
	MallName                    string  `json:"mall_name"`
	CouponTotalQuantity         int64   `json:"coupon_total_quantity"`
	MallCouponEndTime           int64   `json:"mall_coupon_end_time"`
	MallCouponMinOrderAmount    int64   `json:"mall_coupon_min_order_amount"`
	MerchantType                int64   `json:"merchant_type"`
	LgstTxt                     string  `json:"lgst_txt"`
	GoodsName                   string  `json:"goods_name"`
	SalesTip                    string  `json:"sales_tip"`
	OnlySceneAuth               bool    `json:"only_scene_auth"`
	GoodsID                     int64   `json:"goods_id"`
	PredictPromotionRate        int64   `json:"predict_promotion_rate"`
	DescTxt                     string  `json:"desc_txt"`
	MallCouponID                int64   `json:"mall_coupon_id"`
	GoodsDesc                   string  `json:"goods_desc"`
	OptName                     string  `json:"opt_name"`
	ShareRate                   int64   `json:"share_rate"`
	GoodsThumbnailURL           string  `json:"goods_thumbnail_url"`
	OptIDS                      []int64 `json:"opt_ids"`
	OptID                       int64   `json:"opt_id"`
	SearchID                    string  `json:"search_id"`
	GoodsImageURL               string  `json:"goods_image_url"`
	ActivityTags                []int64 `json:"activity_tags"`
	HasMallCoupon               bool    `json:"has_mall_coupon"`
	HasCoupon                   bool    `json:"has_coupon"`
	MinNormalPrice              int64   `json:"min_normal_price"`
	MallCouponStartTime         int64   `json:"mall_coupon_start_time"`
	ServTxt                     string  `json:"serv_txt"`
	MallCouponTotalQuantity     int64   `json:"mall_coupon_total_quantity"`
	MallCouponMaxDiscountAmount int64   `json:"mall_coupon_max_discount_amount"`
	CouponStartTime             int64   `json:"coupon_start_time"`
	MinGroupPrice               int64   `json:"min_group_price"`
	MallCPS                     int64   `json:"mall_cps"`
	CouponDiscount              int64   `json:"coupon_discount"`
	GoodsSign                   string  `json:"goods_sign"`
	CouponEndTime               int64   `json:"coupon_end_time"`
	ZsDuoID                     int64   `json:"zs_duo_id"`
	CltCpnEndTime               *int64  `json:"clt_cpn_end_time,omitempty"`
	CltCpnMinAmt                *int64  `json:"clt_cpn_min_amt,omitempty"`
	CltCpnRemainQuantity        *int64  `json:"clt_cpn_remain_quantity,omitempty"`
	CltCpnBatchSn               *string `json:"clt_cpn_batch_sn,omitempty"`
	CltCpnDiscount              *int64  `json:"clt_cpn_discount,omitempty"`
	CltCpnQuantity              *int64  `json:"clt_cpn_quantity,omitempty"`
	CltCpnStartTime             *int64  `json:"clt_cpn_start_time,omitempty"`
	ActivityType                *int64  `json:"activity_type,omitempty"`
}
