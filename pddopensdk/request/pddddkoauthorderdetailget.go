package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.order.detail.get获取订单详情
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.order.detail.get&permissionId=7
type PddDdkOauthOrderDetailGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthOrderDetailGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("order_sn"), "order_sn")

}

//添加请求参数
func (tk *PddDdkOauthOrderDetailGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthOrderDetailGetRequest) GetApiName() string {
	return "pdd.ddk.oauth.order.detail.get"
}

//返回请求参数
func (tk *PddDdkOauthOrderDetailGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
