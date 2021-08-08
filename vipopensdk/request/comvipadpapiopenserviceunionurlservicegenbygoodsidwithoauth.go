package request

//com.vip.adp.api.open.service.UnionUrlService 根据商品id生成联盟链接-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=genByGoodsIdWithOauth
type ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest) GetMethod() string {
	return "genByGoodsIdWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
