package request

// com.vip.adp.api.open.service.UnionPromotionMaterialService getPromotionMaterialDetailWithOauth 获取推广素材详情——需要授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialService-1.3.0/getPromotionMaterialDetailWithOauth
type ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPromotionMaterialService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest) GetMethod() string {
	return "getPromotionMaterialDetailWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialServiceGetPromotionMaterialDetailWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
