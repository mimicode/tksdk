package comvipadpapiopenserviceunionpromotionmaterialservicegetpromotionmaterialdetailwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionPromotionMaterialService getPromotionMaterialDetailWithOauth 获取推广素材详情——需要授权
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
	MaterialId   string   `json:"materialId"`   // 素材ID
	Title        string   `json:"title"`        // 素材标题
	Content      string   `json:"content"`      // 素材内容
	Images       []string `json:"images"`       // 素材图片列表
	MaterialType int      `json:"materialType"` // 素材类型
	CreateTime   string   `json:"createTime"`   // 创建时间
	UpdateTime   string   `json:"updateTime"`   // 更新时间
	Status       int      `json:"status"`       // 素材状态
	GoodsIds     []string `json:"goodsIds"`     // 关联商品ID列表
	GoodsList    []Goods  `json:"goodsList"`    // 关联商品信息列表
}

type Goods struct {
	GoodsId        string  `json:"goodsId"`        // 商品ID
	GoodsName      string  `json:"goodsName"`      // 商品名称
	GoodsImage     string  `json:"goodsImage"`     // 商品主图
	MarketPrice    float64 `json:"marketPrice"`    // 市场价
	VipPrice       float64 `json:"vipPrice"`       // 唯品会价
	Commission     float64 `json:"commission"`     // 佣金
	CommissionRate float64 `json:"commissionRate"` // 佣金比例
	DetailUrl      string  `json:"detailUrl"`      // 商品详情页链接
	PromoteUrl     string  `json:"promoteUrl"`     // 推广链接
}
