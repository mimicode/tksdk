package request

//com.vip.adp.api.open.service.UnionGoodsService 获取联盟在推商品列表V2-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=goodsListV2WithOauth
type ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest) GetMethod() string {
	return "goodsListV2WithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
