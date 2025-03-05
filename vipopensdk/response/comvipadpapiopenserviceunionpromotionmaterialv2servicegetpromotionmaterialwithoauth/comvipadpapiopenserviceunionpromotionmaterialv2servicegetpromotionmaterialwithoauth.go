package comvipadpapiopenserviceunionpromotionmaterialv2servicegetpromotionmaterialwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionPromotionMaterialV2Service
// getPromotionMaterialWithOauth 2.0.0 获取推广素材列表-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0/getPromotionMaterialWithOauth
type Response struct {
	response.TopResponse
	Success PromotionMaterialResponse `json:"success"`
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

// PromotionMaterialResponse 响应结构
type PromotionMaterialResponse struct {
	Total     int                 `json:"total"`     // 总记录数
	PageSize  int                 `json:"pageSize"`  // 每页记录数
	PageNo    int                 `json:"page"`      // 当前页码
	Materials []PromotionMaterial `json:"materials"` // 推广素材列表
}

// PromotionMaterial 素材信息结构
type PromotionMaterial struct {
	MaterialId       int64           `json:"materialId"`       // 素材ID
	ItemId           string          `json:"itemId"`           // 标的ID
	ItemType         string          `json:"itemType"`         // 标的类型：GOODS-单商品
	PromotionText    string          `json:"promotionText"`    // 推广文案
	PromotionImages  []string        `json:"promotionImages"`  // 推广图片:输出有序
	BasicInfo        TargetBasicInfo `json:"basicInfo"`        // 推广标的基础信息
	CreateTime       int64           `json:"createTime"`       // 创建时间: 时间戳，单位毫秒
	PromoteStartTime int64           `json:"promoteStartTime"` // 推广起始时间: 时间戳，单位毫秒
	PromoteEndTime   int64           `json:"promoteEndTime"`   // 推广结束时间: 时间戳，单位毫秒
	Module           int64           `json:"module"`           // 所属模块
	ModuleName       string          `json:"moduleName"`       // 所属模块名称
	Weight           int             `json:"weight"`           // 推广权重，用于确定推广该素材的优先级，权重值越大，优先级越高
	AdCode           string          `json:"adCode"`           // 用来标识获取推广物料的来源
}

// TargetBasicInfo 推广标的基础信息
type TargetBasicInfo struct {
	Name            string `json:"name"`            // 标的名称
	VipPrice        string `json:"vipPrice"`        // 唯品价
	MarketPrice     string `json:"marketPrice"`     // 市场价
	Commission      string `json:"commission"`      // 佣金
	CommissionRatio string `json:"commissionRatio"` // 佣金比例
}
