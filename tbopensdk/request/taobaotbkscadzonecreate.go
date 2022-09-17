package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.adzone.create( 淘宝客-服务商-创建推广者位 )
//https://open.taobao.com/api.htm?docId=34751&docType=2&scopeId=13878
type TbkScAdzoneCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScAdzoneCreateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_name"), "adzone_name")

}

//添加请求参数
func (tk *TbkScAdzoneCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScAdzoneCreateRequest) GetApiName() string {
	return "taobao.tbk.sc.adzone.create"
}

//返回请求参数
func (tk *TbkScAdzoneCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
