package request

// com.vip.adp.api.open.service.UnionGoodsService getGoodsListMarketingWithOauth 获取指定商品id集合的商品信息(支持精准计算渠道用户到手价)-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getGoodsListMarketingWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest) GetMethod() string {
	return "getGoodsListMarketingWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListMarketingWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
