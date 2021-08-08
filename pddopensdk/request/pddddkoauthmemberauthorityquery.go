package request

import (
	"net/url"
)

//pdd.ddk.oauth.member.authority.query查询是否绑定备案
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.member.authority.query&permissionId=7
type PddDdkOauthMemberAuthorityQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthMemberAuthorityQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkOauthMemberAuthorityQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthMemberAuthorityQueryRequest) GetApiName() string {
	return "pdd.ddk.oauth.member.authority.query"
}

//返回请求参数
func (tk *PddDdkOauthMemberAuthorityQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
