package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.shop.convert( 淘宝客店铺链接转换 )
//http://open.taobao.com/api.htm?docId=24523&docType=2&scopeId=11653
type TbkShopConvertRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkShopConvertRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils2.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils2.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils2.CheckNotNull(tk.Parameters.Get("user_ids"), "user_ids")
}

//添加请求参数
func (tk *TbkShopConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkShopConvertRequest) GetApiName() string {
	return "taobao.tbk.shop.convert"
}

//返回请求参数
func (tk *TbkShopConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}
