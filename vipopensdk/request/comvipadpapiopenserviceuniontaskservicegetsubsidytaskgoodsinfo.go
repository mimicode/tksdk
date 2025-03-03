package request

// com.vip.adp.api.open.service.UnionTaskService getSubsidyTaskGoodsInfo 补贴活动商品列表
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionTaskService-1.0.0/getSubsidyTaskGoodsInfo
type ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionTaskService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest) GetMethod() string {
	return "getSubsidyTaskGoodsInfo"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest) GetVersion() string {
	return "1.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionTaskServiceGetSubsidyTaskGoodsInfoRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
