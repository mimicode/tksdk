package request

import (
	"net/url"
)

//pdd.ddk.goods.recommend.get多多进宝商品推荐API
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.recommend.get
type PddDdkGoodsRecommendGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsRecommendGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkGoodsRecommendGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkGoodsRecommendGetRequest) GetApiName() string {
	return "pdd.ddk.goods.recommend.get"
}

//返回请求参数
func (tk *PddDdkGoodsRecommendGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
