package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.item.coupon.get( 好券清单API【导购】 )
//http://open.taobao.com/api.htm?docId=29821&docType=2&scopeId=11655
type TbkDgItemCouponGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgItemCouponGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkDgItemCouponGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgItemCouponGetRequest) GetApiName() string {
	return "taobao.tbk.dg.item.coupon.get"
}

//返回请求参数
func (tk *TbkDgItemCouponGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
