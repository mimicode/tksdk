package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenGoodsRankQueryRequest jd.union.open.goods.rank.query 联盟实时热销榜商品接口，支持榜单Id和排序类型查询榜单商品列表
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.rank.query
type JdUnionOpenGoodsRankQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsRankQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsRankQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsRankQueryRequest) GetApiName() string {
	return "jd.union.open.goods.rank.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsRankQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
