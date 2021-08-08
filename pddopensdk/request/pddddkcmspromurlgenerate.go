package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.cms.prom.url.generate生成商城-频道推广链接
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.cms.prom.url.generate
type PddDdkCmsPromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkCmsPromUrlGenerateRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("p_id_list"), "p_id_list")

}

//添加请求参数
func (tk *PddDdkCmsPromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkCmsPromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.cms.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkCmsPromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
