package request

//com.vip.adp.api.open.service.UnionOrderService 获取订单信息列表-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=vipLinkCheckWithOuth
type ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionOrderService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest) GetMethod() string {
	return "orderListWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
