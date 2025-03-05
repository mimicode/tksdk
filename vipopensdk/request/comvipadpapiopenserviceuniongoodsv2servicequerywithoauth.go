package request

// ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest com.vip.adp.api.open.service.UnionGoodsV2Service
// queryWithOauth 2.0.0 根据关键词查询商品列表-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=queryWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest) GetMethod() string {
	return "queryWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
