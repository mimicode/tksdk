package request

// com.vip.adp.api.open.service.UnionGoodsV2Service queryCompetitorMatchingProductWithOauth 获取外网商品匹配结果
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/queryCompetitorMatchingProductWithOauth
type ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest) GetMethod() string {
	return "queryCompetitorMatchingProductWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsV2ServiceQueryCompetitorMatchingProductWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
