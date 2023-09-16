package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.vegas.tlj.report( 淘宝客-推广者-淘礼金效果数据 )
//https://open.taobao.com/api.htm?docId=58736&docType=2
type TbkDgVegasTljReportRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgVegasTljReportRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("rights_id"), "rights_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkDgVegasTljReportRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgVegasTljReportRequest) GetApiName() string {
	return "taobao.tbk.dg.vegas.tlj.report"
}

//返回请求参数
func (tk *TbkDgVegasTljReportRequest) GetParameters() url.Values {
	return *tk.Parameters
}
