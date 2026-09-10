package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenChannelRelationGetRequest jd.union.open.channel.relation.get 渠道关系ID生成接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.channel.relation.get
type JdUnionOpenChannelRelationGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenChannelRelationGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenChannelRelationGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenChannelRelationGetRequest) GetApiName() string {
	return "jd.union.open.channel.relation.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenChannelRelationGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
