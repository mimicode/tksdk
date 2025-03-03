package request

// com.vip.adp.api.open.service.UnionUserV2Service checkUser 根据openId判断用户是否是渠道新用户
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/checkUser
type ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest) GetMethod() string {
	return "checkUser"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
