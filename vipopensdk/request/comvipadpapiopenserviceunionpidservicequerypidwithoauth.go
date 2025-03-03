package request

// com.vip.adp.api.open.service.UnionPidService queryPidWithOauth 查询推广位PID-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidService-1.0.0/queryPidWithOauth
type ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionPidService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest) GetMethod() string {
	return "queryPidWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest) GetVersion() string {
	return "1.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionPidServiceQueryPidWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
