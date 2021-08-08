package request

//com.vip.adp.api.open.service.UnionOrderService 维权订单总收益需要扣除该接口返回-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=refundOrderListWithOauth
type ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionOrderService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest) GetMethod() string {
	return "refundOrderListWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
