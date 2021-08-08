package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.rp.prom.url.generate生成营销工具推广链接
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.rp.prom.url.generate&permissionId=7
type PddDdkOauthRpPromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthRpPromUrlGenerateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("p_id_list"), "p_id_list")

}

//添加请求参数
func (tk *PddDdkOauthRpPromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthRpPromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.oauth.rp.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkOauthRpPromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
