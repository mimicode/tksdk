package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.tpwd.convert( 淘口令转链 )
//http://open.taobao.com/api.htm?docId=32932&docType=2&scopeId=11653
type TbkTpwdConvertRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkTpwdConvertRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils.CheckNotNull(tk.Parameters.Get("password_content"), "password_content")
}

//添加请求参数
func (tk *TbkTpwdConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkTpwdConvertRequest) GetApiName() string {
	return "taobao.tbk.tpwd.convert"
}

//返回请求参数
func (tk *TbkTpwdConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}
