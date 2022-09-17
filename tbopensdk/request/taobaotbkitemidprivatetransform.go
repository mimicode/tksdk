package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.itemid.private.transform( 淘宝客-推广者-商品id转化（二方）（专有） )
//https://open.taobao.com/api.htm?docId=64171&docType=2&scopeId=27337
type TbkItemidPrivateTransformRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkItemidPrivateTransformRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("item_ids"), "item_ids")

}

//添加请求参数
func (tk *TbkItemidPrivateTransformRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemidPrivateTransformRequest) GetApiName() string {
	return "taobao.tbk.itemid.private.transform"
}

//返回请求参数
func (tk *TbkItemidPrivateTransformRequest) GetParameters() url.Values {
	return *tk.Parameters
}
