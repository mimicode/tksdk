package request

// ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest com.vip.adp.api.open.service.UnionPromotionMaterialV2Service
// getPromotionMaterialDetailWithOauth 2.0.0 获取推广素材详情——需要授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0&methodName=getPromotionMaterialDetailWithOauth
type ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest) GetMethod() string {
	return "getPromotionMaterialDetailWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPromotionMaterialV2ServiceGetPromotionMaterialDetailWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
