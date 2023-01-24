package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.order.list.range.get用时间段查询推广订单接口
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.order.list.range.get
type PddDdkOrderListRangeGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOrderListRangeGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")
	utils.CheckNotNull(tk.Parameters.Get("end_time"), "end_time")

}

//添加请求参数
func (tk *PddDdkOrderListRangeGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOrderListRangeGetRequest) GetApiName() string {
	return "pdd.ddk.order.list.range.get"
}

//返回请求参数
func (tk *PddDdkOrderListRangeGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
