package pddddkgoodsrecommendget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.goods.recommend.get多多进宝商品推荐API
type Response struct {
	response2.TopResponse
	GoodsBasicDetailResponse GoodsBasicDetailResponse `json:"goods_basic_detail_response"`
}

//解析输出结果
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
	List     []List `json:"list"`
	ListID   string `json:"list_id"`
	SearchID string `json:"search_id"`
	Total    int64  `json:"total"`
}

type List struct {
	CatID                string  `json:"cat_id"`
	CatIDS               []int64 `json:"cat_ids"`
	CategoryID           string  `json:"category_id"`
	CategoryName         string  `json:"category_name"`
	CouponDiscount       int64   `json:"coupon_discount"`
	CouponEndTime        int64   `json:"coupon_end_time"`
	CouponMinOrderAmount int64   `json:"coupon_min_order_amount"`
	CouponPrice          int64   `json:"coupon_price"`
	CouponRemainQuantity int64   `json:"coupon_remain_quantity"`
	CouponStartTime      int64   `json:"coupon_start_time"`
	CouponTotalQuantity  int64   `json:"coupon_total_quantity"`
	CreateAt             int64   `json:"create_at"`
	DescTxt              string  `json:"desc_txt"`
	GoodsDesc            string  `json:"goods_desc"`
	GoodsGalleryUrls     string  `json:"goods_gallery_urls"`
	GoodsID              int64   `json:"goods_id"`
	GoodsImageURL        string  `json:"goods_image_url"`
	GoodsName            string  `json:"goods_name"`
	GoodsRate            int64   `json:"goods_rate"`
	GoodsSign            string  `json:"goods_sign"`
	GoodsThumbnailURL    string  `json:"goods_thumbnail_url"`
	GoodsType            int64   `json:"goods_type"`
	HasCoupon            bool    `json:"has_coupon"`
	LgstTxt              string  `json:"lgst_txt"`
	MallID               int64   `json:"mall_id"`
	MallName             string  `json:"mall_name"`
	MarketFee            int64   `json:"market_fee"`
	MerchantType         string  `json:"merchant_type"`
	MinGroupPrice        int64   `json:"min_group_price"`
	MinNormalPrice       int64   `json:"min_normal_price"`
	OptID                string  `json:"opt_id"`
	OptIDS               []int64 `json:"opt_ids"`
	OptName              string  `json:"opt_name"`
	PredictPromotionRate int64   `json:"predict_promotion_rate"`
	PromotionRate        int64   `json:"promotion_rate"`
	QrCodeImageURL       string  `json:"qr_code_image_url"`
	SalesTip             string  `json:"sales_tip"`
	SearchID             string  `json:"search_id"`
	ServTxt              string  `json:"serv_txt"`
	ShareDesc            string  `json:"share_desc"`
	ShareRate            int64   `json:"share_rate"`
}
