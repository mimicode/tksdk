package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.rp.prom.url.generate生成营销工具推广链接
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.rp.prom.url.generate
type PddDdkRpPromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkRpPromUrlGenerateRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("p_id_list"), "p_id_list")

}

//添加请求参数
func (tk *PddDdkRpPromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkRpPromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.rp.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkRpPromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
