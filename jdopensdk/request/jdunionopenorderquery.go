package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenOrderQueryRequest jd.union.open.order.query 订单查询接口
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.query
type JdUnionOpenOrderQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenOrderQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenOrderQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenOrderQueryRequest) GetApiName() string {
	return "jd.union.open.order.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenOrderQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
