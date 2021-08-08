package request

import (
	"net/url"
)

//pdd.ddk.oauth.goods.recommend.get运营频道商品查询API
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.recommend.get&permissionId=7
type PddDdkOauthGoodsRecommendGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsRecommendGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkOauthGoodsRecommendGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsRecommendGetRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.recommend.get"
}

//返回请求参数
func (tk *PddDdkOauthGoodsRecommendGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
