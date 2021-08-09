package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.newuser.tborder.sum( 拉新活动汇总API--导购 )
//http://open.taobao.com/api.htm?docId=36836&docType=2&scopeId=11655
type TbkDgNewuserOrderSumRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgNewuserOrderSumRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("activity_id"), "activity_id")
	utils.CheckNotNull(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNumber(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNotNull(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")

}

//添加请求参数
func (tk *TbkDgNewuserOrderSumRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgNewuserOrderSumRequest) GetApiName() string {
	return "taobao.tbk.dg.newuser.tborder.sum"
}

//返回请求参数
func (tk *TbkDgNewuserOrderSumRequest) GetParameters() url.Values {
	return *tk.Parameters
}
