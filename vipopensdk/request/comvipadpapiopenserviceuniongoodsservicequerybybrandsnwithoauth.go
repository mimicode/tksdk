package request

// com.vip.adp.api.open.service.UnionGoodsService queryByBrandSnWithOauth 根据品牌SN查询商品列表-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/queryByBrandSnWithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest) GetMethod() string {
	return "queryByBrandSnWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceQueryByBrandSnWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
