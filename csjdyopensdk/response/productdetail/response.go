package productdetail

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 商品详情响应
type Response struct {
	response.TopResponse
	Data ResponseData `json:"data"`
}

func (r *Response) WrapResult(result string) {
	r.Body = result
	err := json.Unmarshal([]byte(result), r)
	if err != nil {
		r.Code = -1
		r.Desc = err.Error()
	}
}

// ResponseData 响应数据
type ResponseData struct {
	Total    int64     `json:"total"`    // 商品总数
	Products []Product `json:"products"` // 商品列表
}

// Product 商品详情
type Product struct {
	ProductId           int64           `json:"product_id"`           // 商品id
	Title               string          `json:"title"`                // 商品名称
	IsKolProduct        bool            `json:"is_kol_product"`       // 是否有达人特殊佣金
	Price               int64           `json:"price"`               // 商品价格（单位：分）
	CosRatio            int             `json:"cos_ratio"`           // 分佣比例
	CosFee              int             `json:"cos_fee"`             // 佣金金额（单位：分）
	FirstCid            int             `json:"first_cid"`           // 商品一级类目
	SecondCid           int             `json:"second_cid"`          // 商品二级类目
	ThirdCid            int             `json:"third_cid"`           // 商品三级类目
	InStock             bool            `json:"in_stock"`            // 是否有库存
	Sales               int             `json:"sales"`               // 商品历史销量
	Cover               string          `json:"cover"`               // 商品主图
	Imgs                []string        `json:"imgs"`                 // 商品轮播图
	DetailUrl           string          `json:"detail_url"`           // 商品链接
	ShopId              int64           `json:"shop_id"`             // 商铺id
	ShopName            string          `json:"shop_name"`            // 商铺名称
	CommentScore        float64         `json:"comment_score"`       // 商品评分
	CommentNum          int             `json:"comment_num"`         // 商品评价数目
	OrderNum            int             `json:"order_num"`           // 近30天商品总订单量
	ViewNum             int             `json:"view_num"`            // 近30天商品总浏览量
	KolNum              int             `json:"kol_num"`             // 近30天推广总达人数
	DailyStatistics     []DailyStatistic `json:"daily_statistics"`   // 近30天推广明细
	LogisticsInfo       string          `json:"logistics_info"`      // 商品物流说明
	HasSxt              bool            `json:"has_sxt"`             // 是否有短视频随心推资质
	ShopTotalScore      ShopTotalScore `json:"shop_total_score"`    // 店铺评分
	CouponPrice         int64           `json:"coupon_price"`        // 券后价格
	BrandInfo           BrandInfo       `json:"brand_info"`          // 品牌信息
	Tags                Tags            `json:"tags"`                // 标签信息
	AvailableCoupons    []AvailableCoupon `json:"available_coupons"` // 优惠券列表
	CommissionType      int             `json:"commission_type"`     // 佣金类型
	InsActivityParam    string          `json:"ins_activity_param"`  // 团长参数
	PublicPlanCosRatio  int             `json:"public_plan_cos_ratio"` // 公共计划佣金比例
	PublicPlanDetailUrl string          `json:"public_plan_detail_url"` // 公共计划商品链接
}

// DailyStatistic 近30天推广明细
type DailyStatistic struct {
	Date     string `json:"date"`      // 日期
	OrderNum int    `json:"order_num"` // 当日商品订单量
	ViewNum  int    `json:"view_num"`  // 当日商品浏览量
	KolNum   int    `json:"kol_num"`   // 当日推广达人数
}

// ShopTotalScore 店铺评分
type ShopTotalScore struct {
	ShopScore      float64 `json:"shop_score"`       // 商家体验分
	ProductScore   float64 `json:"product_score"`    // 商品体验分
	LogisticsScore float64 `json:"logistics_score"`  // 物流体验分
	ServiceScore   float64 `json:"service_score"`    // 商家服务分
}

// BrandInfo 品牌信息
type BrandInfo struct {
	BrandId     int64  `json:"brand_id"`      // 品牌id
	BrandNameCN string `json:"brand_name_cn"`  // 品牌中文名
	BrandNameEN string `json:"brand_name_en"`  // 品牌英文名
}

// Tags 标签信息
type Tags struct {
	HasShopBrandTag bool `json:"has_shop_brand_tag"` // 是否有品牌旗舰店标签
	Success        bool `json:"success"`             // 获取数据是否成功
}

// AvailableCoupon 优惠券信息
type AvailableCoupon struct {
	CouponType     int    `json:"coupon_type"`      // 券类型：0平台券，1平台券，2店铺券，3主播券
	TypeDesc       string `json:"type_desc"`        // 优惠券类型描述
	DiscountDesc   string `json:"discount_desc"`    // 优惠内容
	ApplyStartTime string `json:"apply_start_time"` // 领取开始时间
	ApplyEndTime   string `json:"apply_end_time"`   // 领取结束时间
	ValidityType   int    `json:"validity_type"`    // 有效期类型：1固定有效期，2浮动有效期
	UseStartTime   string `json:"use_start_time"`   // 使用开始时间
	UseEndTime     string `json:"use_end_time"`     // 使用结束时间
	ValidPeriod    int    `json:"valid_period"`     // 浮动有效期（单位：秒）
}
