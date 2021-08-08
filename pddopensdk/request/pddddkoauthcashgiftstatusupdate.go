package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.cashgift.status.update多多礼金状态更新
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cashgift.status.update&permissionId=7
type PddDdkOauthCashgiftStatusUpdateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthCashgiftStatusUpdateRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("cash_gift_id"), "cash_gift_id")
	utils.CheckNumber(tk.Parameters.Get("update_type"), "update_type")

}

//添加请求参数
func (tk *PddDdkOauthCashgiftStatusUpdateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthCashgiftStatusUpdateRequest) GetApiName() string {
	return "pdd.ddk.oauth.cashgift.status.update"
}

//返回请求参数
func (tk *PddDdkOauthCashgiftStatusUpdateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
