package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenChannelRelationQueryRequest jd.union.open.channel.relation.query 渠道关系查询接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.channel.relation.query
type JdUnionOpenChannelRelationQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenChannelRelationQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenChannelRelationQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenChannelRelationQueryRequest) GetApiName() string {
	return "jd.union.open.channel.relation.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenChannelRelationQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
