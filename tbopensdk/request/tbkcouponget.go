package request

import (
	"net/url"
)

//taobao.tbk.coupon.get( 阿里妈妈推广券信息查询 )
//http://open.taobao.com/api.htm?docId=31106&docType=2&scopeId=11655
type TbkCouponGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkCouponGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *TbkCouponGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkCouponGetRequest) GetApiName() string {
	return "taobao.tbk.coupon.get"
}

//返回请求参数
func (tk *TbkCouponGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
