package comvipadpapiopenserviceuniongoodsservicegetbygoodsidsv2withoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id集合的商品信息(新)-需要oauth授权
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Success struct {
	Code int    `json:"code"` // 状态码，1-成功，0-失败
	Msg  string `json:"msg"`  // 状态信息
	Data Data   `json:"data"` // 数据
}

type Data struct {
	GoodsInfoList []GoodsInfoList `json:"goodsInfoList"` // 商品信息列表
}

type GoodsInfoList struct {
	GoodsIDSWithSameSpu    []string               `json:"goodsIdsWithSameSpu"`    // 同SPU商品ID列表
	MarketPrice            string                 `json:"marketPrice"`            // 市场价
	CommissionRate         string                 `json:"commissionRate"`         // 佣金比例
	GoodsID                string                 `json:"goodsId"`                // 商品ID
	Discount               string                 `json:"discount"`               // 折扣
	CouponPriceType        int64                  `json:"couponPriceType"`        // 优惠券价格类型
	GoodsCarouselPictures  []string               `json:"goodsCarouselPictures"`  // 商品轮播图
	GoodsDetailPictures    []string               `json:"goodsDetailPictures"`    // 商品详情图
	CategoryName           string                 `json:"categoryName"`           // 分类名称
	HaiTAO                 int64                  `json:"haiTao"`                 // 是否海淘
	PrepayInfo             PrepayInfo             `json:"prepayInfo"`             // 预付信息
	Cat2NdName             string                 `json:"cat2ndName"`             // 二级分类名称
	IsSubsidyActivityGoods bool                   `json:"isSubsidyActivityGoods"` // 是否补贴活动商品
	Cat1StName             string                 `json:"cat1stName"`             // 一级分类名称
	VipPrice               string                 `json:"vipPrice"`               // VIP价格
	Commission             string                 `json:"commission"`             // 佣金金额
	Sn                     string                 `json:"sn"`                     // 商品编号
	Cat1StID               int64                  `json:"cat1stId"`               // 一级分类ID
	GoodsName              string                 `json:"goodsName"`              // 商品名称
	StoreServiceCapability StoreServiceCapability `json:"storeServiceCapability"` // 店铺服务能力
	BrandName              string                 `json:"brandName"`              // 品牌名称
	BrandLogoFull          string                 `json:"brandLogoFull"`          // 品牌完整logo
	BrandStoreSn           string                 `json:"brandStoreSn"`           // 品牌店铺编号
	SkuInfos               []SkuInfo              `json:"skuInfos"`               // SKU信息列表
	CPSInfo                map[string]string      `json:"cpsInfo"`                // CPS信息
	SellTimeFrom           int64                  `json:"sellTimeFrom"`           // 开始售卖时间
	SchemeStartTime        int64                  `json:"schemeStartTime"`        // 方案开始时间
	CommentsInfo           CommentsInfo           `json:"commentsInfo"`           // 评论信息
	SchemeEndTime          int64                  `json:"schemeEndTime"`          // 方案结束时间
	SourceType             int64                  `json:"sourceType"`             // 来源类型
	SellTimeTo             int64                  `json:"sellTimeTo"`             // 结束售卖时间
	BrandID                int64                  `json:"brandId"`                // 品牌ID
	GoodsThumbURL          string                 `json:"goodsThumbUrl"`          // 商品缩略图URL
	Cat2NdID               int64                  `json:"cat2ndId"`               // 二级分类ID
	SpuID                  string                 `json:"spuId"`                  // SPU ID
	StoreInfo              StoreInfo              `json:"storeInfo"`              // 店铺信息
	EstimatePrice          string                 `json:"estimatePrice"`          // 预估价格
	GoodsMainPicture       string                 `json:"goodsMainPicture"`       // 商品主图
	DestURL                string                 `json:"destUrl"`                // 目标URL
	CategoryID             int64                  `json:"categoryId"`             // 分类ID
	Status                 int64                  `json:"status"`                 // 状态
	CouponInfo             CouponInfo             `json:"couponInfo"`             // 优惠券信息
}

type CommentsInfo struct {
	Comments          int64  `json:"comments"`
	GoodCommentsShare string `json:"goodCommentsShare"`
}

type PrepayInfo struct {
	IsPrepay int64 `json:"isPrepay"`
}

type SkuInfo struct {
	SizeID          string            `json:"sizeId"`
	MarketPrice     string            `json:"marketPrice"`
	SaleStockStatus int64             `json:"saleStockStatus"`
	VipPrice        string            `json:"vipPrice"`
	SaleProps       map[string]string `json:"saleProps"`
}

type StoreInfo struct {
	StoreName string `json:"storeName"`
	StoreID   string `json:"storeId"`
}

type StoreServiceCapability struct {
	StoreScore    string `json:"storeScore"`
	StoreRankRate string `json:"storeRankRate"`
}

type CouponInfo struct {
	UseEndTime        int64  `json:"useEndTime"`
	TotalAmount       int64  `json:"totalAmount"`
	CouponName        string `json:"couponName"`
	ActivateEndTime   int64  `json:"activateEndTime"`
	Buy               string `json:"buy"`
	UseBeginTime      int64  `json:"useBeginTime"`
	CouponNo          string `json:"couponNo"`
	Fav               string `json:"fav"`
	ActivateBeginTime int64  `json:"activateBeginTime"`
	ActivedAmount     int64  `json:"activedAmount"`
}
