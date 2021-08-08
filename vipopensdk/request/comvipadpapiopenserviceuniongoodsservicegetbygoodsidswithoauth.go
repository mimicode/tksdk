package request

//com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id结合的商品信息-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=getByGoodsIdsWithOauth
type ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest) GetMethod() string {
	return "getByGoodsIdsWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
