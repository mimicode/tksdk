package request

// ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest com.vip.adp.api.open.service.UnionUserV2Service
// unbindOpenIdWithOauth 2.0.0 解绑已授权的openId接口服务(需要工具商权限)
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/unbindOpenIdWithOauth
type ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest) GetMethod() string {
	return "unbindOpenIdWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUnbindOpenIdWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
