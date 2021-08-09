package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.wireless.share.tpwd.query( 查询解析淘口令 )
//http://open.taobao.com/api.htm?docId=32461&docType=2&scopeId=11998
type WirelessShareTpwdQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *WirelessShareTpwdQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("password_content"), "password_content")
}

//添加请求参数
func (tk *WirelessShareTpwdQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *WirelessShareTpwdQueryRequest) GetApiName() string {
	return "taobao.wireless.share.tpwd.query"
}

//返回请求参数
func (tk *WirelessShareTpwdQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
