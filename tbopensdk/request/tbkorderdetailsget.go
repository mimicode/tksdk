package request

import (
	utils "github.com/mimicode/tksdk/utils"
	"net/url"
)

// taobao.tbk.order.details.get( 淘宝客-推广者-所有订单查询 )
// https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
type TbkOrderDetailsGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkOrderDetailsGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("end_time"), "end_time")
	utils.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")

}

// 添加请求参数
func (tk *TbkOrderDetailsGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称
func (tk *TbkOrderDetailsGetRequest) GetApiName() string {
	return "taobao.tbk.order.details.get"
}

// 返回请求参数
func (tk *TbkOrderDetailsGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
