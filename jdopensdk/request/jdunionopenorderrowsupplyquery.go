package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenOrderRowSupplyQueryRequest jd.union.open.order.row.supply.query 供开订单行查询接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.row.supply.query
type JdUnionOpenOrderRowSupplyQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenOrderRowSupplyQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenOrderRowSupplyQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenOrderRowSupplyQueryRequest) GetApiName() string {
	return "jd.union.open.order.row.supply.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenOrderRowSupplyQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
