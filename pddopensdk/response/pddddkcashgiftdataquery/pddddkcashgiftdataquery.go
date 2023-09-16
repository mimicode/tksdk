package pddddkcashgiftdataquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.cashgift.data.query查询多多礼金效果数据
type Response struct {
	response.TopResponse
	CashgiftDataResponse CashgiftDataResponse `json:"cashgift_data_response"`
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

type CashgiftDataResponse struct {
	AvailableBalance int64  `json:"available_balance"`
	List             []List `json:"list"`
	Total            int64  `json:"total"`
}

type List struct {
	Amount             int64           `json:"amount"`
	CashGiftID         int64           `json:"cash_gift_id"`
	CashGiftName       string          `json:"cash_gift_name"`
	CouponAmount       int64           `json:"couponAmount"`
	FetchAmount        int64           `json:"fetch_amount"`
	FetchQuantity      int64           `json:"fetch_quantity"`
	GoodsInfoList      []GoodsInfoList `json:"goods_info_list"`
	OrderCouponAmount  int64           `json:"order_coupon_amount"`
	OrderGmv           int64           `json:"order_gmv"`
	OrderQuantity      int64           `json:"order_quantity"`
	PrePromotionAmount int64           `json:"pre_promotion_amount"`
	Quantity           int64           `json:"quantity"`
	RefundAmount       int64           `json:"refund_amount"`
	RefundQuantity     int64           `json:"refund_quantity"`
	Status             int64           `json:"status"`
}

type GoodsInfoList struct {
	CouponDiscount int64  `json:"coupon_discount"`
	GoodsName      string `json:"goods_name"`
	GoodsPrice     int64  `json:"goods_price"`
	GoodsSign      string `json:"goods_sign"`
	Rate           int64  `json:"rate"`
}
