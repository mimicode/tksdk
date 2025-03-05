package request

// ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest com.vip.adp.api.open.service.UnionUserV2Service
// checkUserWithOauth 2.0.0 根据openId判断用户是否是渠道用户信息, 注意:如果scene传值为1，是渠道新用户判断，而非唯品会新用户判断
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/checkUserWithOauth
type ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUserV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest) GetMethod() string {
	return "checkUserWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUserV2ServiceCheckUserWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
