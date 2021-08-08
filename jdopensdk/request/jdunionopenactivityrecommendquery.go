package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenActivityRecommendQueryRequest jd.union.open.activity.recommend.query 提供联盟官方活动查询
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.activity.recommend.query
type JdUnionOpenActivityRecommendQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenActivityRecommendQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenActivityRecommendQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenActivityRecommendQueryRequest) GetApiName() string {
	return "jd.union.open.activity.recommend.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenActivityRecommendQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
