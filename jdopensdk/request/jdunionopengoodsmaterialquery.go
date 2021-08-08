package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenGoodsMaterialQueryRequest jd.union.open.goods.material.query 猜你喜欢商品推荐
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.material.query
type JdUnionOpenGoodsMaterialQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsMaterialQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsMaterialQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsMaterialQueryRequest) GetApiName() string {
	return "jd.union.open.goods.material.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsMaterialQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
