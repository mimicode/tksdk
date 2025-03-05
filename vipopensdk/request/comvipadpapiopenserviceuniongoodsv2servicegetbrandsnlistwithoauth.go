package request

// ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest com.vip.adp.api.open.service.UnionGoodsV2Service
// getBrandSnListWithOauth 2.0.0 精选品牌列表查询-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getBrandSnListWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest) GetMethod() string {
	return "getBrandSnListWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGetBrandSnListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
