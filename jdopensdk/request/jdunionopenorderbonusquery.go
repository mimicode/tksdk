package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenOrderBonusQueryRequest jd.union.open.order.bonus.query 奖励订单查询接口
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.bonus.query
type JdUnionOpenOrderBonusQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenOrderBonusQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenOrderBonusQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenOrderBonusQueryRequest) GetApiName() string {
	return "jd.union.open.order.bonus.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenOrderBonusQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
