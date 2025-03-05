package request

// ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest com.vip.adp.api.open.service.UnionGoodsV2Service
// queryByBrandSnWithOauth 2.0.0 根据品牌SN查询商品列表-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=queryByBrandSnWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest) GetMethod() string {
	return "queryByBrandSnWithOauth"
}

//方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryByBrandSnWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
