package comvipadpapiopenserviceunionpromotionmaterialv2servicegetpromotionmaterialdetailwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionPromotionMaterialV2Service
// getPromotionMaterialDetailWithOauth 2.0.0 获取推广素材详情-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0/getPromotionMaterialDetailWithOauth
type Response struct {
	response.TopResponse
	Success PromotionMaterialDetailResponse `json:"result"`
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

// PromotionMaterialDetailResponse 响应结构
type PromotionMaterialDetailResponse struct {
	MaterialInfos []MaterialInfo `json:"materialInfos"` // 推广素材详情列表
}

// MaterialInfo 素材信息结构
type MaterialInfo struct {
	MaterialType  string `json:"materialType"`  // 素材类型，对应请求materialType用逗号分割后的值，例如PICTURE
	MaterialValue string `json:"materialValue"` // 素材内容，如果素材类型是PICTURE，素材内容则为图片链接
	Weight        int    `json:"weight"`        // 推广权重，用于确定推广该素材的优先级，权重值越大，优先级越高
}
