package request

// com.vip.adp.api.open.service.UnionOrderService virtualOrderList 获取订单信息列表
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderService-1.3.0/virtualOrderList
type ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionOrderService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest) GetMethod() string {
	return "virtualOrderList"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionOrderServiceVirtualOrderListRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
