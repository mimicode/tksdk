package request

// com.vip.adp.api.open.service.UnionGoodsService getBrandSnListWithOauth 精选品牌列表查询-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getBrandSnListWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest) GetMethod() string {
	return "getBrandSnListWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandSnListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
