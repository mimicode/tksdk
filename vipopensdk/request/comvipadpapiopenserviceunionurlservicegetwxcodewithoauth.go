package request

// com.vip.adp.api.open.service.UnionUrlService getWxCodeWithOauth 生成微信小程序码,需要Oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/getWxCodeWithOauth
type ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest) GetMethod() string {
	return "getWxCodeWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetWxCodeWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
