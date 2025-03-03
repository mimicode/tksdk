package request

// com.vip.adp.api.open.service.UnionGoodsService getGoodsListByBrandAndJxCodeWithOauth 根据组货码与品牌获取商品列表详情
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getGoodsListByBrandAndJxCodeWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest) GetMethod() string {
	return "getGoodsListByBrandAndJxCodeWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetGoodsListByBrandAndJxCodeWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
