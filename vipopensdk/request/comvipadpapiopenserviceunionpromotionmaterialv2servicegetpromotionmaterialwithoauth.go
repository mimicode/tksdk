package request

// ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest com.vip.adp.api.open.service.UnionPromotionMaterialV2Service
// getPromotionMaterialWithOauth 2.0.0 获取推广素材列表——需要授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0&methodName=getPromotionMaterialWithOauth
type ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest) GetMethod() string {
	return "getPromotionMaterialWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
