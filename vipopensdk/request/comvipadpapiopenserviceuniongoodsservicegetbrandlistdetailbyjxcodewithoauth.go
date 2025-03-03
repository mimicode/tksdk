package request

// com.vip.adp.api.open.service.UnionGoodsService getBrandListDetailByJxCodeWithOauth 根据组货码获取品牌详情
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getBrandListDetailByJxCodeWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest) GetMethod() string {
	return "getBrandListDetailByJxCodeWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandListDetailByJxCodeWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
