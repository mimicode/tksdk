package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.order.list.increment.get&permissionId=2
type PddDdkOrderListIncrementGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOrderListIncrementGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("start_update_time"), "start_update_time")
	utils2.CheckNotNull(tk.Parameters.Get("end_update_time"), "end_update_time")

}

//添加请求参数
func (tk *PddDdkOrderListIncrementGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOrderListIncrementGetRequest) GetApiName() string {
	return "pdd.ddk.order.list.increment.get"
}

//返回请求参数
func (tk *PddDdkOrderListIncrementGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
