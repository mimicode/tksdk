package comvipadpapiopenserviceuniongoodsservicegetbranddetailwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService getBrandDetailWithOauth 通过品牌SN查询品牌详情——需要oauth授权
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
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

type Result struct {
	BrandInfo BrandInfo `json:"brandInfo"` // 品牌信息
}

type BrandInfo struct {
	BrandId         int     `json:"brandId"`         // 品牌ID
	BrandName       string  `json:"brandName"`       // 品牌名称
	BrandLogo       string  `json:"brandLogo"`       // 品牌logo
	BrandDesc       string  `json:"brandDesc"`       // 品牌描述
	BrandStoreSn    string  `json:"brandStoreSn"`    // 品牌店铺号
	BrandStoreName  string  `json:"brandStoreName"`  // 品牌店铺名称
	BrandScore      float64 `json:"brandScore"`      // 品牌评分
	CommissionRate  float64 `json:"commissionRate"`  // 佣金比例，单位：%
	AvgDiscount     float64 `json:"avgDiscount"`     // 平均折扣，单位：%
	TotalGoodsCount int     `json:"totalGoodsCount"` // 商品总数
	SalesVolume     int     `json:"salesVolume"`     // 销量
}
