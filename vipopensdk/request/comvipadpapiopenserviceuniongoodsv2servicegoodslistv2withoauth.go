package request

// ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest com.vip.adp.api.open.service.UnionGoodsV2Service
// goodsListV2WithOauth 2.0.0 获取联盟在推商品列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=goodsListV2WithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest) GetMethod() string {
	return "goodsListV2WithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceGoodsListV2WithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
