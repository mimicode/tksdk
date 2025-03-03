package request

// com.vip.adp.api.open.service.UnionPidV2Service 创建推广位PID-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidV2Service-2.0.0/genPidWithOauth
type ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPidV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest) GetMethod() string {
	return "genPidWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceGenPidWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
