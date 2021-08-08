package request

import (
	"net/url"
)

//pdd.ddk.oauth.goods.search多多进宝商品查询
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.search&permissionId=7
type PddDdkOauthGoodsSearchRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsSearchRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkOauthGoodsSearchRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsSearchRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.search"
}

//返回请求参数
func (tk *PddDdkOauthGoodsSearchRequest) GetParameters() url.Values {
	return *tk.Parameters
}
