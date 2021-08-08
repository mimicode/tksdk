package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.member.authority.query 查询是否绑定备案
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.member.authority.query
type PddDdkMemberAuthorityQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkMemberAuthorityQueryRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils2.CheckNotNull(tk.Parameters.Get("custom_parameters"), "custom_parameters")

}

//添加请求参数
func (tk *PddDdkMemberAuthorityQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkMemberAuthorityQueryRequest) GetApiName() string {
	return "pdd.ddk.member.authority.query"
}

//返回请求参数
func (tk *PddDdkMemberAuthorityQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
