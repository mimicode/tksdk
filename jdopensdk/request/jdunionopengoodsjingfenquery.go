package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenGoodsJingfenQueryRequest jd.union.open.goods.jingfen.query 京粉精选商品查询接口
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.jingfen.query
type JdUnionOpenGoodsJingfenQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsJingfenQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsJingfenQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsJingfenQueryRequest) GetApiName() string {
	return "jd.union.open.goods.jingfen.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsJingfenQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
