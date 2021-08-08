package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.cashgift.create创建多多礼金
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cashgift.create&permissionId=7
type PddDdkOauthCashgiftCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthCashgiftCreateRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("acquire_end_time"), "acquire_end_time")
	utils.CheckNumber(tk.Parameters.Get("acquire_start_time"), "acquire_start_time")
	utils.CheckNumber(tk.Parameters.Get("coupon_amount"), "coupon_amount")
	utils.CheckNumber(tk.Parameters.Get("quantity"), "quantity")
	utils.CheckNumber(tk.Parameters.Get("validity_time_type"), "validity_time_type")
	utils.CheckNotNull(tk.Parameters.Get("source_url"), "source_url")

}

//添加请求参数
func (tk *PddDdkOauthCashgiftCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthCashgiftCreateRequest) GetApiName() string {
	return "pdd.ddk.oauth.cashgift.create"
}

//返回请求参数
func (tk *PddDdkOauthCashgiftCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
