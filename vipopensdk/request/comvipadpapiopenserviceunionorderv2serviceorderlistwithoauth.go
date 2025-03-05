package request

// ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest com.vip.adp.api.open.service.UnionOrderV2Service
// orderListWithOauth 2.0.0 获取订单列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=orderListWithOauth
type ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionOrderV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest) GetMethod() string {
	return "orderListWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceOrderListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
