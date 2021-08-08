package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenGoodsPromotiongoodsinfoQueryRequest jd.union.open.goods.promotiongoodsinfo.query 通过SKUID查询推广商品
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.promotiongoodsinfo.query
type JdUnionOpenGoodsPromotiongoodsinfoQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsPromotiongoodsinfoQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsPromotiongoodsinfoQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsPromotiongoodsinfoQueryRequest) GetApiName() string {
	return "jd.union.open.goods.promotiongoodsinfo.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsPromotiongoodsinfoQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
