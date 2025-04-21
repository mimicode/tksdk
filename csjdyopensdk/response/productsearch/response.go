package productsearch

import (
	"encoding/json"
	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

type Response struct {
	response.TopResponse
	Data ResponseData `json:"data"`
}

func (r *Response) WrapResult(result string) {
	r.Body = result
	err := json.Unmarshal([]byte(result), r)
	//解析错误
	if err != nil {
		r.Code = -1
		r.Desc = err.Error()
	}
}

type ResponseData struct {
	Total    int64     `json:"total"`
	Products []Product `json:"products"`
}

type Product struct {
	ProductId    int64  `json:"product_id"`     // 商品id
	Title        string `json:"title"`          // 商品名称
	IsKolProduct bool   `json:"is_kol_product"` // 是否有达人特殊佣金
	Price        int64  `json:"price"`          // 商品价格，单位分
	InStock      bool   `json:"in_stock"`       // 有无库存
	FirstCid     int64  `json:"first_cid"`      // 商品一级类目
	SecondCid    int64  `json:"second_cid"`     // 商品二级类目
	ThirdCid     int64  `json:"third_cid"`      // 商品三级类目
	Sales        int    `json:"sales"`          // 商品历史销量
	Cover        string `json:"cover"`          // 商品主图
	DetailUrl    string `json:"detail_url"`     // 商品链接
	ShopId       int64  `json:"shop_id"`        // 商铺id
	ShopName     string `json:"shop_name"`      // 商铺名称
	CouponPrice  int64  `json:"coupon_price"`   // 券后价格，单位分（0或者没传则为原价）
	CosRatio     int    `json:"cos_ratio"`      // 分佣比例，百分比乘以100，比如1%返回1*100=100
	CosFee       int    `json:"cos_fee"`        // 佣金金额，单位分
	ActivityId   int    `json:"activity_id"`    // 获取活动商品。1: 返回超值购商品；0:返回全量商品
	Ext          string `json:"ext"`            // 一个加密字段，需要在转链接口当中回传
	PostFree     bool   `json:"post_free"`      // 是否包邮
	LimitMinSale bool   `json:"limit_min_sale"` // 是否多件起购
}
