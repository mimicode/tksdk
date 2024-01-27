package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.item.info.upgrade.get( 淘宝客-公用-淘宝客商品详情查询升级版（简易版） )
//https://open.taobao.com/api.htm?docId=64763&docType=2&scopeId=16189
type TbkItemInfoUpgradeGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkItemInfoUpgradeGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("item_id"), "item_id")

}

//添加请求参数
func (tk *TbkItemInfoUpgradeGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemInfoUpgradeGetRequest) GetApiName() string {
	return "taobao.tbk.item.info.upgrade.get"
}

//返回请求参数
func (tk *TbkItemInfoUpgradeGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
