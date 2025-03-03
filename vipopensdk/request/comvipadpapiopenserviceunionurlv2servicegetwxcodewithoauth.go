package request

// com.vip.adp.api.open.service.UnionUrlV2Service getWxCodeWithOauth 生成微信小程序码,需要Oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getWxCodeWithOauth
type ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest) GetMethod() string {
	return "getWxCodeWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetWxCodeWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
