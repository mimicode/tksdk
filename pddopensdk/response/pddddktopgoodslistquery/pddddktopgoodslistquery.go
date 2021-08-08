package pddddktopgoodslistquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.top.goods.list.query（多多客获取爆款排行商品接口）
type Response struct {
	response.TopResponse
	TopGoodsListGetResponse TopGoodsListGetResponseResult `json:"top_goods_list_get_response"`
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

type TopGoodsListGetResponseResult struct {
	Total int64                               `json:"total"`
	List  []TopGoodsListGetResponseResultList `json:"list"`
	//RequestID string                              `json:"request_id"`
	//SearchID  string                              `json:"search_id"`
}

type TopGoodsListGetResponseResultList struct {
	CouponRemainQuantity int64 `json:"coupon_remain_quantity"`
	PromotionRate        int64 `json:"promotion_rate"`
	//CatIDS               []int64 `json:"cat_ids"`
	CouponMinOrderAmount int64  `json:"coupon_min_order_amount"`
	MallName             string `json:"mall_name"`
	CouponTotalQuantity  int64  `json:"coupon_total_quantity"`
	MerchantType         int64  `json:"merchant_type"`
	LgstTxt              string `json:"lgst_txt"`
	GoodsName            string `json:"goods_name"`
	SalesTip             string `json:"sales_tip"`
	GoodsID              int64  `json:"goods_id"`
	DescTxt              string `json:"desc_txt"`
	GoodsDesc            string `json:"goods_desc"`
	OptName              string `json:"opt_name"`
	GoodsThumbnailURL    string `json:"goods_thumbnail_url"`
	OptID                int64  `json:"opt_id"`
	//OptIDS               []int64 `json:"opt_ids"`
	//SearchID        string `json:"search_id"`
	GoodsImageURL   string `json:"goods_image_url"`
	MinNormalPrice  int64  `json:"min_normal_price"`
	HasCoupon       bool   `json:"has_coupon"`
	ServTxt         string `json:"serv_txt"`
	MinGroupPrice   int64  `json:"min_group_price"`
	CouponStartTime int64  `json:"coupon_start_time"`
	MallCPS         int64  `json:"mall_cps"`
	CouponDiscount  int64  `json:"coupon_discount"`
	CouponEndTime   int64  `json:"coupon_end_time"`
}
