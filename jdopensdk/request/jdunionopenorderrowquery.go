package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenOrderRowQueryRequest jd.union.open.order.row.query 查询推广订单及佣金信息
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.row.query
type JdUnionOpenOrderRowQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenOrderRowQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenOrderRowQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenOrderRowQueryRequest) GetApiName() string {
	return "jd.union.open.order.row.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenOrderRowQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
