package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenStatisticsPromotionQueryRequest jd.union.open.statistics.promotion.query 推广效果数据查询接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.statistics.promotion.query
type JdUnionOpenStatisticsPromotionQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenStatisticsPromotionQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenStatisticsPromotionQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenStatisticsPromotionQueryRequest) GetApiName() string {
	return "jd.union.open.statistics.promotion.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenStatisticsPromotionQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
