package request

// com.vip.adp.api.open.service.UnionUserService checkUser 根据openId判断用户是否是渠道新用户
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserService-1.0.0/checkUser
type ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest) GetMethod() string {
	return "checkUser"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest) GetVersion() string {
	return "1.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserServiceCheckUserRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
