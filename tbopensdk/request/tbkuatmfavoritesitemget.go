package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.uatm.favorites.item.get( 获取淘宝联盟选品库的宝贝信息 )
//http://open.taobao.com/api.htm?docId=26619&docType=2&scopeId=11655
type TbkUatmFavoritesItemGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkUatmFavoritesItemGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNotNull(tk.Parameters.Get("favorites_id"), "favorites_id")
	utils.CheckNumber(tk.Parameters.Get("favorites_id"), "favorites_id")

}

//添加请求参数
func (tk *TbkUatmFavoritesItemGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkUatmFavoritesItemGetRequest) GetApiName() string {
	return "taobao.tbk.uatm.favorites.item.get"
}

//返回请求参数
func (tk *TbkUatmFavoritesItemGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
