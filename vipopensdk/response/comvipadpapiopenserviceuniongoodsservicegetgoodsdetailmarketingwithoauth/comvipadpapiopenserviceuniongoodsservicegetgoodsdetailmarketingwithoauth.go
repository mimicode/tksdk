package comvipadpapiopenserviceuniongoodsservicegetgoodsdetailmarketingwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService getGoodsDetailMarketingWithOauth 获取指定商品id的商品信息(支持精准计算渠道用户到手价)-需要oauth授权
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
	GoodsInfo GoodsInfo `json:"goodsInfo"` // 商品信息
}

type GoodsInfo struct {
	MarketPrice      string     `json:"marketPrice"`      // 市场价
	CommissionRate   string     `json:"commissionRate"`   // 佣金比例
	GoodsID          string     `json:"goodsId"`          // 商品ID
	Discount         string     `json:"discount"`         // 折扣
	CategoryName     string     `json:"categoryName"`     // 分类名称
	HaiTAO           int64      `json:"haiTao"`           // 是否海淘
	Cat2NdName       string     `json:"cat2ndName"`       // 二级分类名称
	Cat1StName       string     `json:"cat1stName"`       // 一级分类名称
	VipPrice         string     `json:"vipPrice"`         // VIP价格
	Commission       string     `json:"commission"`       // 佣金金额
	Cat1StID         int64      `json:"cat1stId"`         // 一级分类ID
	GoodsName        string     `json:"goodsName"`        // 商品名称
	BrandName        string     `json:"brandName"`        // 品牌名称
	BrandLogoFull    string     `json:"brandLogoFull"`    // 品牌完整logo
	BrandStoreSn     string     `json:"brandStoreSn"`     // 品牌店铺编号
	SellTimeFrom     int64      `json:"sellTimeFrom"`     // 开始售卖时间
	SchemeStartTime  int64      `json:"schemeStartTime"`  // 方案开始时间
	SchemeEndTime    int64      `json:"schemeEndTime"`    // 方案结束时间
	SourceType       int64      `json:"sourceType"`       // 来源类型
	SellTimeTo       int64      `json:"sellTimeTo"`       // 结束售卖时间
	BrandID          int64      `json:"brandId"`          // 品牌ID
	GoodsThumbURL    string     `json:"goodsThumbUrl"`    // 商品缩略图URL
	Cat2NdID         int64      `json:"cat2ndId"`         // 二级分类ID
	StoreInfo        StoreInfo  `json:"storeInfo"`        // 店铺信息
	GoodsMainPicture string     `json:"goodsMainPicture"` // 商品主图
	DestURL          string     `json:"destUrl"`          // 目标URL
	CategoryID       int64      `json:"categoryId"`       // 分类ID
	Status           int64      `json:"status"`           // 状态
	CouponInfo       CouponInfo `json:"couponInfo"`       // 优惠券信息
}

type StoreInfo struct {
	StoreName string `json:"storeName"` // 店铺名称
	StoreID   string `json:"storeId"`   // 店铺ID
}

type CouponInfo struct {
	CouponName string `json:"couponName"` // 优惠券名称
	Buy        string `json:"buy"`        // 使用门槛
	CouponNo   string `json:"couponNo"`   // 优惠券编号
	Fav        string `json:"fav"`        // 优惠金额
}
