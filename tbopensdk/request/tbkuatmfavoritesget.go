package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.uatm.favorites.get( 获取淘宝联盟选品库列表 )
//http://open.taobao.com/api.htm?docId=26620&docType=2&scopeId=11655
type TbkUatmFavoritesGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkUatmFavoritesGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("fields"), "fields")

}

//添加请求参数
func (tk *TbkUatmFavoritesGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkUatmFavoritesGetRequest) GetApiName() string {
	return "taobao.tbk.uatm.favorites.get"
}

//返回请求参数
func (tk *TbkUatmFavoritesGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
