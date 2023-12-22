package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

//pdd.ddk.tmc.activity.list千万神券
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.tmc.activity.list
type PddDdkTmcActivityListRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkTmcActivityListRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page_num"), "page_num")
utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
utils.CheckNotNull(tk.Parameters.Get("start_time_lower"), "start_time_lower")
utils.CheckNotNull(tk.Parameters.Get("start_time_upper"), "start_time_upper")

}

//添加请求参数
func (tk *PddDdkTmcActivityListRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkTmcActivityListRequest) GetApiName() string {
	return "pdd.ddk.tmc.activity.list"
}

//返回请求参数
func (tk *PddDdkTmcActivityListRequest) GetParameters() url.Values {
	return *tk.Parameters
}
