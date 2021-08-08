package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.vegas.tlj.create( 淘宝客-推广者-淘礼金创建 )
//https://open.taobao.com/api.htm?docId=40173&docType=2
type TbkDgVegasTljCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgVegasTljCreateRequest) CheckParameters() {
	utils2.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils2.CheckNumber(tk.Parameters.Get("item_id"), "item_id")
	utils2.CheckNumber(tk.Parameters.Get("total_num"), "total_num")
	utils2.CheckNumber(tk.Parameters.Get("user_total_win_num_limit"), "user_total_win_num_limit")
	utils2.CheckNotNull(tk.Parameters.Get("security_switch"), "security_switch")
	utils2.CheckMaxLength(tk.Parameters.Get("name"), 10, "name")
	utils2.CheckMinFloatValue(tk.Parameters.Get("per_face"), 1, "per_face")
	utils2.CheckNotNull(tk.Parameters.Get("send_start_time"), "send_start_time")

}

//添加请求参数
func (tk *TbkDgVegasTljCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgVegasTljCreateRequest) GetApiName() string {
	return "taobao.tbk.dg.vegas.tlj.create"
}

//返回请求参数
func (tk *TbkDgVegasTljCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
