package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.cashgift.create创建多多礼金
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.cashgift.create
type PddDdkCashgiftCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkCashgiftCreateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("acquire_end_time"), "acquire_end_time")
	utils.CheckNotNull(tk.Parameters.Get("acquire_start_time"), "acquire_start_time")

}

//添加请求参数
func (tk *PddDdkCashgiftCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkCashgiftCreateRequest) GetApiName() string {
	return "pdd.ddk.cashgift.create"
}

//返回请求参数
func (tk *PddDdkCashgiftCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
