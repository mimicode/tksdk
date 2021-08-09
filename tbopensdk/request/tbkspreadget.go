package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.spread.get( 物料传播方式获取 )
//http://open.taobao.com/api.htm?docId=27832&docType=2&scopeId=12340
type TbkSpreadGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkSpreadGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("requests"), "requests")
}

//添加请求参数
func (tk *TbkSpreadGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkSpreadGetRequest) GetApiName() string {
	return "taobao.tbk.spread.get"
}

//返回请求参数
func (tk *TbkSpreadGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
