package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenGoodsBigfieldQueryRequest jd.union.open.goods.bigfield.query 商品详情查询接口,大字段信息
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.bigfield.query
type JdUnionOpenGoodsBigfieldQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsBigfieldQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsBigfieldQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsBigfieldQueryRequest) GetApiName() string {
	return "jd.union.open.goods.bigfield.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsBigfieldQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
