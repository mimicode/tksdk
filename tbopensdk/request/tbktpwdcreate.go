package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.tpwd.create( 淘宝客淘口令 )
//http://open.taobao.com/api.htm?docId=31127&docType=2&scopeId=11655
type TbkTpwdCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkTpwdCreateRequest) CheckParameters() {
	//utils2.CheckNotNull(tk.Parameters.Get("text"), "text")
	utils2.CheckNotNull(tk.Parameters.Get("url"), "url")

}

//添加请求参数
func (tk *TbkTpwdCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkTpwdCreateRequest) GetApiName() string {
	return "taobao.tbk.tpwd.create"
}

//返回请求参数
func (tk *TbkTpwdCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
