package request

// ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest com.vip.adp.api.open.service.UnionGoodsService 根据关键词查询商品列表-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=queryWithOauth
type ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest) GetMethod() string {
	return "queryWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
