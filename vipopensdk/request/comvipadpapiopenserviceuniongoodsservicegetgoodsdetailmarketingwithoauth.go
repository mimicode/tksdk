package request

// com.vip.adp.api.open.service.UnionGoodsService getGoodsDetailMarketingWithOauth 获取指定商品id的商品信息(支持精准计算渠道用户到手价)-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getGoodsDetailMarketingWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest) GetMethod() string {
	return "getGoodsDetailMarketingWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsDetailMarketingWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
