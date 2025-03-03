package comvipadpapiopenserviceuniongoodsv2servicequerywithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsV2Service queryWithOauth 根据关键词查询商品列表-需要oauth授权
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
	Total         int64           `json:"total"`         // 总记录数
	GoodsInfoList []GoodsInfoList `json:"goodsInfoList"` // 商品信息列表
	PageSize      int64           `json:"pageSize"`      // 每页条数
	Page          int64           `json:"page"`          // 页码
}

type GoodsInfoList struct {
	CategoryName    string  `json:"categoryName"`    // 商品分类名称
	CommissionRate  float64 `json:"commissionRate"`  // 佣金比例
	GoodsId         string  `json:"goodsId"`         // 商品ID
	GoodsName       string  `json:"goodsName"`       // 商品名称
	GoodsThumbUrl   string  `json:"goodsThumbUrl"`   // 商品缩略图
	GoodsType       int     `json:"goodsType"`       // 商品类型：0-普通商品，1-品牌特卖商品
	MarketPrice     float64 `json:"marketPrice"`     // 市场价
	SalePrice       float64 `json:"salePrice"`       // 销售价
	VipPrice        float64 `json:"vipPrice"`        // 会员价
	Commission      float64 `json:"commission"`      // 佣金金额
	Discount        float64 `json:"discount"`        // 折扣
	Cat1stName      string  `json:"cat1stName"`      // 一级分类名称
	Cat2ndName      string  `json:"cat2ndName"`      // 二级分类名称
	BrandName       string  `json:"brandName"`       // 品牌名称
	BrandStoreSn    string  `json:"brandStoreSn"`    // 品牌店铺SN
	BrandLogoUrl    string  `json:"brandLogoUrl"`    // 品牌logo图片
	SchemeEndTime   int64   `json:"schemeEndTime"`   // 推广计划结束时间
	SchemeStartTime int64   `json:"schemeStartTime"` // 推广计划开始时间
	SellTimeFrom    int64   `json:"sellTimeFrom"`    // 开始售卖时间
	SellTimeTo      int64   `json:"sellTimeTo"`      // 结束售卖时间
	Weight          int     `json:"weight"`          // 权重
	StoreInfo       Store   `json:"storeInfo"`       // 店铺信息
}

type Store struct {
	StoreId   int64  `json:"storeId"`   // 店铺ID
	StoreName string `json:"storeName"` // 店铺名称
	StoreUrl  string `json:"storeUrl"`  // 店铺URL
}
