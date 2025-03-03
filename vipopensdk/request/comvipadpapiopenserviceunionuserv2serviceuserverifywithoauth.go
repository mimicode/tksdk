package request

// com.vip.adp.api.open.service.UnionUserV2Service userVerifyWithOauth CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/userVerifyWithOauth
type ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest) GetMethod() string {
	return "userVerifyWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceUserVerifyWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
