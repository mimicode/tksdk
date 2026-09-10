package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenPromotionIntelligenceQueryRequest jd.union.open.promotion.intelligence.query 线报推广【申请】：京东商品段子、优惠信息、店铺及类目活动爆料信息，接口出参直接是转链后链接
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.intelligence.query
type JdUnionOpenPromotionIntelligenceQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenPromotionIntelligenceQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenPromotionIntelligenceQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenPromotionIntelligenceQueryRequest) GetApiName() string {
	return "jd.union.open.promotion.intelligence.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenPromotionIntelligenceQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
