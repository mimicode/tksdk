package request

import (
	"net/url"
)

//pdd.ddk.oauth.goods.detail多多进宝商品详情查询
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.detail&permissionId=7
type PddDdkOauthGoodsDetailRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsDetailRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkOauthGoodsDetailRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsDetailRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.detail"
}

//返回请求参数
func (tk *PddDdkOauthGoodsDetailRequest) GetParameters() url.Values {
	return *tk.Parameters
}
