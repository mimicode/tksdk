package request

import (
	"net/url"
)

//pdd.ddk.oauth.goods.prom.url.generate生成多多进宝推广链接
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.prom.url.generate&permissionId=7
type PddDdkOauthGoodsPromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsPromUrlGenerateRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkOauthGoodsPromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsPromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkOauthGoodsPromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
