package request

// com.vip.adp.api.open.service.UnionGoodsService getCategorys 类目查询接口
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getCategorys
type ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest) GetMethod() string {
	return "getCategorys"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetCategorysRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
