package request

// com.vip.adp.api.open.service.UnionGoodsService getBrandDetailWithOauth 通过品牌SN查询品牌详情——需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getBrandDetailWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest) GetMethod() string {
	return "getBrandDetailWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetBrandDetailWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
