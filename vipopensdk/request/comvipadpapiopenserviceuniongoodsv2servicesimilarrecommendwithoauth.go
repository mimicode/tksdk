package request

// ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest com.vip.adp.api.open.service.UnionGoodsV2Service
// similarRecommendWithOauth 2.0.0 相似推荐商品列表-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=similarRecommendWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest) GetMethod() string {
	return "similarRecommendWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceSimilarRecommendWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
