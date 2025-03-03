package request

// com.vip.adp.api.open.service.UnionUserService userVerifyWithOauth CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserService-1.0.0/userVerifyWithOauth
type ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest) GetMethod() string {
	return "userVerifyWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserServiceUserVerifyWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
