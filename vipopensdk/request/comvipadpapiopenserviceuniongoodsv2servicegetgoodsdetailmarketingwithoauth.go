package request

// com.vip.adp.api.open.service.UnionGoodsV2Service getGoodsDetailMarketingWithOauth 获取指定商品id的商品信息(支持精准计算渠道用户到手价)-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/getGoodsDetailMarketingWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest) GetMethod() string {
	return "getGoodsDetailMarketingWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetGoodsDetailMarketingWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
