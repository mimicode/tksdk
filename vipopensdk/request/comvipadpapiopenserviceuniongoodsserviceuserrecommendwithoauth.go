package request

//com.vip.adp.api.open.service.UnionGoodsService 猜你喜欢商品列表-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=userRecommendWithOauth
type ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest) GetMethod() string {
	return "userRecommendWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
