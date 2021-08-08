package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenOrderAgentQueryRequest jd.union.open.order.agent.query 工具商订单行查询接口
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.agent.query
type JdUnionOpenOrderAgentQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenOrderAgentQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenOrderAgentQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenOrderAgentQueryRequest) GetApiName() string {
	return "jd.union.open.order.agent.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenOrderAgentQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
