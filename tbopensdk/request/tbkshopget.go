package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.shop.get( 淘宝客店铺查询 )
//http://open.taobao.com/api.htm?docId=24521&docType=2&scopeId=11655
type TbkShopGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkShopGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils.CheckNotNull(tk.Parameters.Get("q"), "q")

}

//添加请求参数
func (tk *TbkShopGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkShopGetRequest) GetApiName() string {
	return "taobao.tbk.shop.get"
}

//返回请求参数
func (tk *TbkShopGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
