package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.item.coupon.get( 单品加券检索api )
//http://open.taobao.com/api.htm?docId=28110&docType=2&scopeId=12332
type TbkItemCouponGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkItemCouponGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")
}

//添加请求参数
func (tk *TbkItemCouponGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemCouponGetRequest) GetApiName() string {
	return "taobao.tbk.item.coupon.get"
}

//返回请求参数
func (tk *TbkItemCouponGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
