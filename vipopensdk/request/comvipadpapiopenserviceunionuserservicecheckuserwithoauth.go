package request

// com.vip.adp.api.open.service.UnionUserService checkUserWithOauth 根据openId判断用户是否是渠道用户信息
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserService-1.0.0/checkUserWithOauth
type ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest) GetMethod() string {
	return "checkUserWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
