package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenGoodsRecommendQueryRequest jd.union.open.goods.recommend.query 相似品推荐接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.recommend.query
type JdUnionOpenGoodsRecommendQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsRecommendQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsRecommendQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsRecommendQueryRequest) GetApiName() string {
	return "jd.union.open.goods.recommend.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsRecommendQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
