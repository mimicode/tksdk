package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.cms.prom.url.generate生成商城推广链接接口
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cms.prom.url.generate&permissionId=7
type PddDdkOauthCmsPromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthCmsPromUrlGenerateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("p_id_list"), "p_id_list")

}

//添加请求参数
func (tk *PddDdkOauthCmsPromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthCmsPromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.oauth.cms.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkOauthCmsPromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
