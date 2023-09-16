package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.vegas.tlj.stop( 淘宝客-推广者-淘礼金暂停发放 )
//https://open.taobao.com/api.htm?docId=58713&docType=2
type TbkDgVegasTljStopRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgVegasTljStopRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("rights_id"), "rights_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkDgVegasTljStopRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgVegasTljStopRequest) GetApiName() string {
	return "taobao.tbk.dg.vegas.tlj.stop"
}

//返回请求参数
func (tk *TbkDgVegasTljStopRequest) GetParameters() url.Values {
	return *tk.Parameters
}
