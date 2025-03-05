package request

// ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest com.vip.adp.api.open.service.UnionPidV2Service
// queryPidWithOauth 2.0.0 查询PID-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPidV2Service-2.0.0&methodName=queryPidWithOauth
type ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPidV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest) GetMethod() string {
	return "queryPidWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPidV2ServiceQueryPidWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
