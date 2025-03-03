package request

// com.vip.adp.api.open.service.UnionPromotionMaterialService 获取推广素材列表——需要授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialService-1.3.0/getPromotionMaterialWithOauth
type ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPromotionMaterialService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest) GetMethod() string {
	return "getPromotionMaterialWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
