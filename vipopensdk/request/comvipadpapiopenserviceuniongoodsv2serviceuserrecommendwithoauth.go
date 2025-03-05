package request

// ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest com.vip.adp.api.open.service.UnionGoodsV2Service
// userRecommendWithOauth 2.0.0 猜你喜欢商品列表-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=userRecommendWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest) GetMethod() string {
	return "userRecommendWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceUserRecommendWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
