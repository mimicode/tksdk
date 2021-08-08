package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.all.order.list.increment.get查询所有授权的多多客订单
//https://open.pinduoduo.com/application/document/api?permissionId=7
type PddDdkAllOrderListIncrementGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkAllOrderListIncrementGetRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("end_update_time"), "end_update_time")
	utils.CheckNumber(tk.Parameters.Get("start_update_time"), "start_update_time")

}

//添加请求参数
func (tk *PddDdkAllOrderListIncrementGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkAllOrderListIncrementGetRequest) GetApiName() string {
	return "pdd.ddk.all.order.list.increment.get"
}

//返回请求参数
func (tk *PddDdkAllOrderListIncrementGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
