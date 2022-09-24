package request

//com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id集合的商品信息(新)-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=getByGoodsIdsV2WithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest) GetMethod() string {
	return "getByGoodsIdsV2WithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
