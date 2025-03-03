package comvipadpapiopenserviceuniongoodsv2servicesimilarrecommendwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsV2Service similarRecommendWithOauth 相似推荐商品列表-需要oauth授权
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
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 状态信息
	Data    Data   `json:"data"`    // 响应数据
}

type Data struct {
	Total     int         `json:"total"`     // 总记录数
	PageSize  int         `json:"pageSize"`  // 每页记录数
	Page      int         `json:"page"`      // 当前页码
	GoodsList []GoodsInfo `json:"goodsList"` // 商品列表
}

type GoodsInfo struct {
	GoodsId         string   `json:"goodsId"`         // 商品ID
	GoodsName       string   `json:"goodsName"`       // 商品名称
	GoodsDesc       string   `json:"goodsDesc"`       // 商品描述
	DestUrl         string   `json:"destUrl"`         // 商品落地页URL
	GoodsThumbUrl   string   `json:"goodsThumbUrl"`   // 商品缩略图URL
	GoodsMainPicUrl string   `json:"goodsMainPicUrl"` // 商品主图URL
	CategoryId      int      `json:"categoryId"`      // 商品类目ID
	CategoryName    string   `json:"categoryName"`    // 商品类目名称
	MarketPrice     float64  `json:"marketPrice"`     // 市场价，单位：元
	VipPrice        float64  `json:"vipPrice"`        // 唯品会价，单位：元
	CommissionRate  float64  `json:"commissionRate"`  // 佣金比例，单位：%
	Commission      float64  `json:"commission"`      // 佣金金额，单位：元
	Discount        float64  `json:"discount"`        // 折扣，单位：%
	GoodsScore      float64  `json:"goodsScore"`      // 商品评分
	SalesVolume     int      `json:"salesVolume"`     // 销量
	BrandId         int      `json:"brandId"`         // 品牌ID
	BrandName       string   `json:"brandName"`       // 品牌名称
	BrandStoreSn    string   `json:"brandStoreSn"`    // 品牌店铺号
	BrandStoreName  string   `json:"brandStoreName"`  // 品牌店铺名称
	StoreId         int      `json:"storeId"`         // 店铺ID
	StoreName       string   `json:"storeName"`       // 店铺名称
	Comments        int      `json:"comments"`        // 评论数
	GoodLabels      []string `json:"goodLabels"`      // 商品标签列表
	HasStock        bool     `json:"hasStock"`        // 是否有库存
}
