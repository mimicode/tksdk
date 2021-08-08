package request

//com.vip.adp.api.open.service.UnionGoodsService 相似推荐商品列表-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=similarRecommendWithOauth
type ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest) GetMethod() string {
	return "similarRecommendWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
