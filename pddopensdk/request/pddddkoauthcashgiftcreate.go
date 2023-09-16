package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// pdd.ddk.oauth.cashgift.create创建多多礼金
// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cashgift.create
type PddDdkOauthCashgiftCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthCashgiftCreateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("acquire_end_time"), "acquire_end_time")
	utils.CheckNotNull(tk.Parameters.Get("acquire_start_time"), "acquire_start_time")

}

// 添加请求参数
func (tk *PddDdkOauthCashgiftCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称
func (tk *PddDdkOauthCashgiftCreateRequest) GetApiName() string {
	return "pdd.ddk.oauth.cashgift.create"
}

// 返回请求参数
func (tk *PddDdkOauthCashgiftCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
