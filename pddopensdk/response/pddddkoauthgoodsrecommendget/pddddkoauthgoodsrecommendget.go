package pddddkoauthgoodsrecommendget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.goods.recommend.get运营频道商品查询API
type Response struct {
	response.TopResponse
	GoodsBasicDetailResponse GoodsBasicDetailResponse `json:"goods_basic_detail_response"`
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

type GoodsBasicDetailResponse struct {
	Total     int64  `json:"total"`
	List      []List `json:"list"`
	RequestID string `json:"request_id"`
	SearchID  string `json:"search_id"`
}

type List struct {
	CategoryName               string    `json:"category_name"`
	CouponRemainQuantity       int64     `json:"coupon_remain_quantity"`
	PromotionRate              int64     `json:"promotion_rate"`
	CatIDS                     []int64   `json:"cat_ids"`
	CouponMinOrderAmount       int64     `json:"coupon_min_order_amount"`
	CategoryID                 string    `json:"category_id"`
	MallID                     int64     `json:"mall_id"`
	MallName                   string    `json:"mall_name"`
	CouponTotalQuantity        int64     `json:"coupon_total_quantity"`
	MerchantType               string    `json:"merchant_type"`
	LgstTxt                    string    `json:"lgst_txt"`
	GoodsName                  string    `json:"goods_name"`
	SalesTip                   string    `json:"sales_tip"`
	GoodsID                    int64     `json:"goods_id"`
	PredictPromotionRate       int64     `json:"predict_promotion_rate"`
	DescTxt                    string    `json:"desc_txt"`
	GoodsDesc                  string    `json:"goods_desc"`
	OptName                    string    `json:"opt_name"`
	RealtimeSalesTip           string    `json:"realtime_sales_tip"`
	ShareRate                  int64     `json:"share_rate"`
	GoodsThumbnailURL          string    `json:"goods_thumbnail_url"`
	OptIDS                     []int64   `json:"opt_ids"`
	OptID                      string    `json:"opt_id"`
	SearchID                   string    `json:"search_id"`
	GoodsImageURL              string    `json:"goods_image_url"`
	ActivityTags               []int64   `json:"activity_tags"`
	HasCoupon                  bool      `json:"has_coupon"`
	MinNormalPrice             int64     `json:"min_normal_price"`
	ServTxt                    string    `json:"serv_txt"`
	UnifiedTags                []*string `json:"unified_tags"`
	CouponStartTime            int64     `json:"coupon_start_time"`
	MinGroupPrice              int64     `json:"min_group_price"`
	CouponDiscount             int64     `json:"coupon_discount"`
	GoodsSign                  string    `json:"goods_sign"`
	CouponEndTime              int64     `json:"coupon_end_time"`
	SubsidyAmount              *int64    `json:"subsidy_amount,omitempty"`
	SubsidyDuoAmountTenMillion *int64    `json:"subsidy_duo_amount_ten_million,omitempty"`
	ActivityPromotionRate      *int64    `json:"activity_promotion_rate,omitempty"`
}
