package request

// com.vip.adp.api.open.service.UnionGoodsService goodsListV2 获取联盟在推商品列表-V2版本
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/goodsListV2
type ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request) GetMethod() string {
	return "goodsListV2"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2Request) GetParameters() map[string]interface{} {
	return tk.Parameters
}
