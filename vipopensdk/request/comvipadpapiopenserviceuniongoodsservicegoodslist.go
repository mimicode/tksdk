package request

// com.vip.adp.api.open.service.UnionGoodsService goodsList 获取联盟在推商品列表
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/goodsList
type ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest) GetMethod() string {
	return "goodsList"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
