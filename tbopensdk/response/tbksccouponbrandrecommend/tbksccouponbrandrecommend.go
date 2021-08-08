package tbksccouponbrandrecommend

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.coupon.brand.recommend( 品牌券API【社交】 )
type Response struct {
	response.TopResponse
	TbkScCouponBrandRecommendResult Result `json:"tbk_sc_coupon_brand_recommend_response"`
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
	Results   Results `json:"results"`
	RequestID string  `json:"request_id"`
}

type Results struct {
	TbkCoupon []TbkCoupon `json:"tbk_coupon"`
}

type TbkCoupon struct {
	Category          int64       `json:"category"`
	CommissionRate    string      `json:"commission_rate"`
	CouponEndTime     string      `json:"coupon_end_time"`
	CouponInfo        string      `json:"coupon_info"`
	CouponRemainCount int64       `json:"coupon_remain_count"`
	CouponStartTime   string      `json:"coupon_start_time"`
	CouponTotalCount  int64       `json:"coupon_total_count"`
	ItemDescription   string      `json:"item_description"`
	ItemURL           string      `json:"item_url"`
	Nick              string      `json:"nick"`
	NumIid            int64       `json:"num_iid"`
	PictURL           string      `json:"pict_url"`
	SellerID          int64       `json:"seller_id"`
	ShopTitle         string      `json:"shop_title"`
	SmallImages       SmallImages `json:"small_images"`
	Title             string      `json:"title"`
	UserType          int64       `json:"user_type"`
	Volume            int64       `json:"volume"`
	ZkFinalPrice      string      `json:"zk_final_price"`
}

type SmallImages struct {
	String []string `json:"string"`
}
