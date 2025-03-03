package request

// com.vip.adp.api.open.service.UnionOrderV2Service refundOrderListWithOauth 获取维权订单列表，订单归属人为授权用户
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0/refundOrderListWithOauth
type ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionOrderV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest) GetMethod() string {
	return "refundOrderListWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderV2ServiceRefundOrderListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
