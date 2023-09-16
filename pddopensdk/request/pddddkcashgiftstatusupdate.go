package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.cashgift.status.update多多礼金状态更新
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.cashgift.status.update
type PddDdkCashgiftStatusUpdateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkCashgiftStatusUpdateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("cash_gift_id"), "cash_gift_id")
	utils.CheckNotNull(tk.Parameters.Get("update_type"), "update_type")

}

//添加请求参数
func (tk *PddDdkCashgiftStatusUpdateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkCashgiftStatusUpdateRequest) GetApiName() string {
	return "pdd.ddk.cashgift.status.update"
}

//返回请求参数
func (tk *PddDdkCashgiftStatusUpdateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
