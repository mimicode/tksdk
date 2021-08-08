package pddddkgoodsdetail

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response  pdd.ddk.goods.detail（多多进宝商品详情查询）
type Response struct {
	response.TopResponse
	GoodsDetailResponse GoodsDetailResponse `json:"goods_detail_response"`
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

type GoodsDetailResponse struct {
	GoodsDetails []GoodsDetail `json:"goods_details"`
	RequestID    string        `json:"request_id"`
}

type GoodsDetail struct {
	CategoryName                string      `json:"category_name"`
	CouponRemainQuantity        int64       `json:"coupon_remain_quantity"`
	ServPct                     float64     `json:"serv_pct"`
	PromotionRate               int64       `json:"promotion_rate"`
	PredictPromotionRate        int64       `json:"predict_promotion_rate"`
	MallID                      int64       `json:"mall_id"`
	MallName                    string      `json:"mall_name"`
	DescPct                     float64     `json:"desc_pct"`
	MallCouponEndTime           interface{} `json:"mall_coupon_end_time"`
	GoodsName                   string      `json:"goods_name"`
	GoodsEvalCount              int64       `json:"goods_eval_count"`
	GoodsID                     int64       `json:"goods_id"`
	GoodsGalleryUrls            []string    `json:"goods_gallery_urls"`
	GoodsDesc                   string      `json:"goods_desc"`
	OptName                     string      `json:"opt_name"`
	OptIDS                      []int64     `json:"opt_ids"`
	GoodsImageURL               string      `json:"goods_image_url"`
	HasMallCoupon               bool        `json:"has_mall_coupon"`
	MinGroupPrice               int64       `json:"min_group_price"`
	CouponStartTime             int64       `json:"coupon_start_time"`
	CouponDiscount              int64       `json:"coupon_discount"`
	CouponEndTime               int64       `json:"coupon_end_time"`
	AvgDesc                     int64       `json:"avg_desc"`
	MallCouponRemainQuantity    interface{} `json:"mall_coupon_remain_quantity"`
	PlanType                    int64       `json:"plan_type"`
	AvgServ                     int64       `json:"avg_serv"`
	AvgLgst                     int64       `json:"avg_lgst"`
	SoldQuantity                int64       `json:"sold_quantity"`
	CatIDS                      []int64     `json:"cat_ids"`
	CouponMinOrderAmount        int64       `json:"coupon_min_order_amount"`
	LgstPct                     float64     `json:"lgst_pct"`
	CategoryID                  int64       `json:"category_id"`
	MallCouponDiscountPct       interface{} `json:"mall_coupon_discount_pct"`
	GoodsEvalScore              *float64    `json:"goods_eval_score"`
	CatID                       interface{} `json:"cat_id"`
	CouponTotalQuantity         int64       `json:"coupon_total_quantity"`
	MallCouponMinOrderAmount    interface{} `json:"mall_coupon_min_order_amount"`
	MerchantType                int64       `json:"merchant_type"`
	PlanTypeAll                 int64       `json:"plan_type_all"`
	MallCouponID                interface{} `json:"mall_coupon_id"`
	GoodsThumbnailURL           string      `json:"goods_thumbnail_url"`
	OptID                       int64       `json:"opt_id"`
	MinNormalPrice              int64       `json:"min_normal_price"`
	HasCoupon                   bool        `json:"has_coupon"`
	MallCouponStartTime         interface{} `json:"mall_coupon_start_time"`
	MallRate                    int64       `json:"mall_rate"`
	MallCouponTotalQuantity     interface{} `json:"mall_coupon_total_quantity"`
	CreateAt                    int64       `json:"create_at"`
	MallCouponMaxDiscountAmount interface{} `json:"mall_coupon_max_discount_amount"`
	MallCPS                     int64       `json:"mall_cps"`

	LgstTxt string `json:"lgst_txt"`

	SalesTip      string `json:"sales_tip"`
	OnlySceneAuth bool   `json:"only_scene_auth"`

	DescTxt string `json:"desc_txt"`

	ShareRate int64 `json:"share_rate"`

	ServTxt   string        `json:"serv_txt"`
	VideoUrls []interface{} `json:"video_urls"`

	GoodsSign string `json:"goods_sign"`

	ZsDuoID int64 `json:"zs_duo_id"`
}
