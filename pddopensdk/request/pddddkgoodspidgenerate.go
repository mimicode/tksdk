package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.goods.pid.generate（创建多多进宝推广位)
//https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.goods.pid.generate
type PddDdkGoodsPidGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsPidGenerateRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("number"), "number")
	utils2.CheckNumber(tk.Parameters.Get("number"), "number")

}

//添加请求参数
func (tk *PddDdkGoodsPidGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkGoodsPidGenerateRequest) GetApiName() string {
	return "pdd.ddk.goods.pid.generate"
}

//返回请求参数
func (tk *PddDdkGoodsPidGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
