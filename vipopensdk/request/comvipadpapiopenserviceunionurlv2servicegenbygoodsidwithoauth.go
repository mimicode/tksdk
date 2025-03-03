package request

// com.vip.adp.api.open.service.UnionUrlV2Service genByGoodsIdWithOauth 根据商品id生成联盟链接-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/genByGoodsIdWithOauth
type ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest) GetMethod() string {
	return "genByGoodsIdWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByGoodsIdWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
