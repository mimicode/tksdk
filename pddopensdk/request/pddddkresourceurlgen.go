package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.resource.url.gen（生成多多进宝频道推广）
//https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.resource.url.gen
type PddDdkResourceUrlGenRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkResourceUrlGenRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")

}

//添加请求参数
func (tk *PddDdkResourceUrlGenRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkResourceUrlGenRequest) GetApiName() string {
	return "pdd.ddk.resource.url.gen"
}

//返回请求参数
func (tk *PddDdkResourceUrlGenRequest) GetParameters() url.Values {
	return *tk.Parameters
}
