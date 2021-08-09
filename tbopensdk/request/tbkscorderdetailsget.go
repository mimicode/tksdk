package request

import (
	utils "github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.order.details.get( 淘宝客【服务商】所有订单查询 )
//https://open.taobao.com/api.htm?docId=43755&docType=2&scopeId=16322
type TbkScOrderDetailsGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScOrderDetailsGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("end_time"), "end_time")
	utils.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")

}

//添加请求参数
func (tk *TbkScOrderDetailsGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScOrderDetailsGetRequest) GetApiName() string {
	return "taobao.tbk.sc.order.details.get"
}

//返回请求参数
func (tk *TbkScOrderDetailsGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
