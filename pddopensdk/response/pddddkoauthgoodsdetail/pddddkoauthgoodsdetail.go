package pddddkoauthgoodsdetail

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.goods.detail多多进宝商品详情查询
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
	CategoryName                string   `json:"category_name"`
	CouponRemainQuantity        int64    `json:"coupon_remain_quantity"`
	PromotionRate               int64    `json:"promotion_rate"`
	ServiceTags                 []int64  `json:"service_tags"`
	MallID                      int64    `json:"mall_id"`
	MallName                    string   `json:"mall_name"`
	MallCouponEndTime           int64    `json:"mall_coupon_end_time"`
	LgstTxt                     string   `json:"lgst_txt"`
	GoodsName                   string   `json:"goods_name"`
	GoodsGalleryUrls            []string `json:"goods_gallery_urls"`
	GoodsID                     int64    `json:"goods_id"`
	BrandName                   string   `json:"brand_name"`
	PredictPromotionRate        int64    `json:"predict_promotion_rate"`
	GoodsDesc                   string   `json:"goods_desc"`
	OptName                     string   `json:"opt_name"`
	ShareRate                   int64    `json:"share_rate"`
	OptIDS                      []int64  `json:"opt_ids"`
	GoodsImageURL               string   `json:"goods_image_url"`
	MallImgURL                  string   `json:"mall_img_url"`
	HasMallCoupon               bool     `json:"has_mall_coupon"`
	UnifiedTags                 []string `json:"unified_tags"`
	VideoUrls                   []string `json:"video_urls"`
	CouponStartTime             int64    `json:"coupon_start_time"`
	MinGroupPrice               int64    `json:"min_group_price"`
	CouponDiscount              int64    `json:"coupon_discount"`
	CouponEndTime               int64    `json:"coupon_end_time"`
	ZsDuoID                     int64    `json:"zs_duo_id"`
	MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity"`
	PlanType                    int64    `json:"plan_type"`
	ExtraCouponAmount           int64    `json:"extra_coupon_amount"`
	CatIDS                      []int64  `json:"cat_ids"`
	CouponMinOrderAmount        int64    `json:"coupon_min_order_amount"`
	CategoryID                  int64    `json:"category_id"`
	MallCouponDiscountPct       int64    `json:"mall_coupon_discount_pct"`
	CouponTotalQuantity         int64    `json:"coupon_total_quantity"`
	MallCouponMinOrderAmount    int64    `json:"mall_coupon_min_order_amount"`
	MerchantType                int64    `json:"merchant_type"`
	SalesTip                    string   `json:"sales_tip"`
	OnlySceneAuth               bool     `json:"only_scene_auth"`
	DescTxt                     string   `json:"desc_txt"`
	GoodsThumbnailURL           string   `json:"goods_thumbnail_url"`
	OptID                       int64    `json:"opt_id"`
	ActivityTags                []int64  `json:"activity_tags"`
	HasCoupon                   bool     `json:"has_coupon"`
	MinNormalPrice              int64    `json:"min_normal_price"`
	MallCouponStartTime         int64    `json:"mall_coupon_start_time"`
	ServTxt                     string   `json:"serv_txt"`
	MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity"`
	MallCouponMaxDiscountAmount int64    `json:"mall_coupon_max_discount_amount"`
	MallCPS                     int64    `json:"mall_cps"`
	GoodsSign                   string   `json:"goods_sign"`
}
