package request

// com.vip.adp.api.open.service.UnionGoodsService getByGoodsIdsV2 获取指定商品id集合的商品信息(新)
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getByGoodsIdsV2
type ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request) GetMethod() string {
	return "getByGoodsIdsV2"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2Request) GetParameters() map[string]interface{} {
	return tk.Parameters
}
