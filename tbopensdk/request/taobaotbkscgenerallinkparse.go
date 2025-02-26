package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.general.link.parse( 淘宝客-服务商-万能解析 )
//https://open.taobao.com/api.htm?docId=70535&docType=2&scopeId=31093
type TbkScGeneralLinkParseRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScGeneralLinkParseRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkScGeneralLinkParseRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScGeneralLinkParseRequest) GetApiName() string {
	return "taobao.tbk.sc.general.link.parse"
}

//返回请求参数
func (tk *TbkScGeneralLinkParseRequest) GetParameters() url.Values {
	return *tk.Parameters
}
