package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenStatisticsRedpacketAgentQueryRequest jd.union.open.statistics.redpacket.agent.query 工具商查询帮助其他推客转的链接的红包发放数据
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.statistics.redpacket.agent.query
type JdUnionOpenStatisticsRedpacketAgentQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenStatisticsRedpacketAgentQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenStatisticsRedpacketAgentQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenStatisticsRedpacketAgentQueryRequest) GetApiName() string {
	return "jd.union.open.statistics.redpacket.agent.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenStatisticsRedpacketAgentQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
