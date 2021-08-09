package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.newuser.tborder.get( 淘宝客新用户订单API--导购 )
//http://open.taobao.com/api.htm?docId=33892&docType=2&scopeId=11655
type TbkDgNewuserOrderGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgNewuserOrderGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("activity_id"), "activity_id")

}

//添加请求参数
func (tk *TbkDgNewuserOrderGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgNewuserOrderGetRequest) GetApiName() string {
	return "taobao.tbk.dg.newuser.tborder.get"
}

//返回请求参数
func (tk *TbkDgNewuserOrderGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
