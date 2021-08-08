package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.ju.tqg.get( 淘抢购api )
//http://open.taobao.com/api.htm?docId=27543&docType=2&scopeId=11655
type TbkJuTqgGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkJuTqgGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils2.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")
	utils2.CheckNotNull(tk.Parameters.Get("end_time"), "end_time")
	utils2.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils2.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkJuTqgGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkJuTqgGetRequest) GetApiName() string {
	return "taobao.tbk.ju.tqg.get"
}

//返回请求参数
func (tk *TbkJuTqgGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
