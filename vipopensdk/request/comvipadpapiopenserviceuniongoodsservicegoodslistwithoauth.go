package request

//com.vip.adp.api.open.service.UnionGoodsService 获取联盟在推商品列表-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=goodsListWithOauth
type ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest) GetMethod() string {
	return "goodsListWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
