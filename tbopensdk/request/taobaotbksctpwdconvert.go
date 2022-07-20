package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.tpwd.convert( 淘宝客-服务商-淘口令解析&转链 )
//https://open.taobao.com/api.htm?docId=43873&docType=2&scopeId=16401
type TbkScTpwdConvertRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScTpwdConvertRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNotNull(tk.Parameters.Get("password_content"), "password_content")

}

//添加请求参数
func (tk *TbkScTpwdConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScTpwdConvertRequest) GetApiName() string {
	return "taobao.tbk.sc.tpwd.convert"
}

//返回请求参数
func (tk *TbkScTpwdConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}
