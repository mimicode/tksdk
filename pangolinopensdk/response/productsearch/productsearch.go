package productsearch

import (
	"encoding/json"

	"github.com/mimicode/tksdk/pangolinopensdk/response"
)

// Response 商品列表响应
type Response struct {
	response.TopResponse
	Data Data `json:"data"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	t.RawData = result
	if err := json.Unmarshal([]byte(result), t); err != nil {
		t.Code = -1
		t.Desc = err.Error()
	}
}

// Data 响应数据
type Data struct {
	Total    int64      `json:"total"`    // 满足搜索条件的商品总数（最多2000）
	Products []Product  `json:"products"` // 商品列表
}

// Product 商品信息
type Product struct {
	ProductID           int64  `json:"product_id"`           // 商品id
	Title               string `json:"title"`                 // 商品名称
	IsKolProduct        bool   `json:"is_kol_product"`        // 是否有达人特殊佣金
	Price               int64  `json:"price"`                 // 商品价格（单位：分）
	InStock             bool   `json:"in_stock"`             // 是否有库存
	FirstCid            int64  `json:"first_cid"`            // 商品一级类目
	SecondCid           int64  `json:"second_cid"`           // 商品二级类目
	ThirdCid            int64  `json:"third_cid"`            // 商品三级类目
	Sales               int    `json:"sales"`                 // 商品历史销量
	Cover               string `json:"cover"`                // 商品主图
	DetailURL           string `json:"detail_url"`            // 商品链接
	ShopID              int64  `json:"shop_id"`              // 商铺id
	ShopName            string `json:"shop_name"`             // 商铺名称
	CouponPrice         int64  `json:"coupon_price"`         // 券后价格（单位：分）
	CosRatio            int    `json:"cos_ratio"`            // 分佣比例（百分比乘以100）
	CosFee              int    `json:"cos_fee"`              // 佣金金额（单位：分）
	ActivityID          int    `json:"activity_id"`           // 活动商品标识：1超值购，0全量
	Ext                 string `json:"ext"`                   // 加密字段（转链时需要回传）
	PostFree            bool   `json:"post_free"`             // 是否包邮
	LimitMinSale        bool   `json:"limit_min_sale"`        // 是否多件起购
	PublicPlanCosRatio  int    `json:"public_plan_cos_ratio"` // 公共计划佣金比例
	PublicPlanDetailURL string `json:"public_plan_detail_url"` // 公共计划商品链接
}
