package request

// ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest com.vip.adp.api.open.service.UnionOrderV2Service
// virtualOrderListWithOauth 2.0.0 获取虚拟订单列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=virtualOrderListWithOauth
type ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionOrderV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest) GetMethod() string {
	return "virtualOrderListWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceVirtualOrderListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
