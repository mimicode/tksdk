package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.invitecode.get( 淘宝客邀请码生成-社交 )
//http://open.taobao.com/api.htm?docId=38046&docType=2&scopeId=14474
type TbkScInvitecodeGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScInvitecodeGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("relation_app"), "relation_app")
	utils.CheckNotNull(tk.Parameters.Get("code_type"), "code_type")
	utils.CheckNumber(tk.Parameters.Get("code_type"), "code_type")
}

//添加请求参数
func (tk *TbkScInvitecodeGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScInvitecodeGetRequest) GetApiName() string {
	return "taobao.tbk.sc.invitecode.get"
}

//返回请求参数
func (tk *TbkScInvitecodeGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
