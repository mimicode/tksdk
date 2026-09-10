package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenPromotionToolsIntelligenceQueryRequest jd.union.open.promotion.tools.intelligence.query 工具商线报推广【申请】：获取京东商品段子、优惠信息、店铺及类目活动爆料信息，接口出参直接是转链后链接
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.tools.intelligence.query
type JdUnionOpenPromotionToolsIntelligenceQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenPromotionToolsIntelligenceQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenPromotionToolsIntelligenceQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenPromotionToolsIntelligenceQueryRequest) GetApiName() string {
	return "jd.union.open.promotion.tools.intelligence.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenPromotionToolsIntelligenceQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
