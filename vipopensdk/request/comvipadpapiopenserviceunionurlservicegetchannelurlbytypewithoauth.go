package request

// com.vip.adp.api.open.service.UnionUrlService getChannelUrlByTypeWithOauth 渠道根据落地页类型生链接口,需要Oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/getChannelUrlByTypeWithOauth
type ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest) GetMethod() string {
	return "getChannelUrlByTypeWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
