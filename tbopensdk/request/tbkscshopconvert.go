package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.shop.convert( 淘宝客-服务商-店铺链接转换 )
//https://open.taobao.com/api.htm?docId=43878&docType=2&scopeId=16403
type TbkScShopConvertRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScShopConvertRequest) CheckParameters() {
	utils2.CheckNumber(tk.Parameters.Get("site_id"), "site_id")
	utils2.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils2.CheckNotNull(tk.Parameters.Get("user_ids"), "user_ids")
	utils2.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkScShopConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScShopConvertRequest) GetApiName() string {
	return "taobao.tbk.sc.shop.convert"
}

//返回请求参数
func (tk *TbkScShopConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}
