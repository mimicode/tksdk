package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenChannelInvitecodeGetRequest jd.union.open.channel.invitecode.get 邀请码获取接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.channel.invitecode.get
type JdUnionOpenChannelInvitecodeGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenChannelInvitecodeGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenChannelInvitecodeGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenChannelInvitecodeGetRequest) GetApiName() string {
	return "jd.union.open.channel.invitecode.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenChannelInvitecodeGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
