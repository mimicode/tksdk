package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenStatisticsRedpacketQueryRequest jd.union.open.statistics.redpacket.query 京享红包效果数据
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.statistics.redpacket.query
type JdUnionOpenStatisticsRedpacketQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenStatisticsRedpacketQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenStatisticsRedpacketQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenStatisticsRedpacketQueryRequest) GetApiName() string {
	return "jd.union.open.statistics.redpacket.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenStatisticsRedpacketQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
