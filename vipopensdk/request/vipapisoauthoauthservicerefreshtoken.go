package request

// vipapis.oauth.OauthService 刷新令牌服务
// https://vop.vip.com/home#/api/method/detail/vipapis.oauth.OauthService-1.0.0/refreshToken
type VipapisOauthOauthServiceRefreshTokenRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *VipapisOauthOauthServiceRefreshTokenRequest) CheckParameters() {

}

// 添加请求参数
func (tk *VipapisOauthOauthServiceRefreshTokenRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *VipapisOauthOauthServiceRefreshTokenRequest) GetApiName() string {
	return "vipapis.oauth.OauthService"
}

// 方法名称
func (tk *VipapisOauthOauthServiceRefreshTokenRequest) GetMethod() string {
	return "refreshToken"
}

// 方法名称
func (tk *VipapisOauthOauthServiceRefreshTokenRequest) GetVersion() string {
	return "1.0.0"
}

// 返回请求参数
func (tk *VipapisOauthOauthServiceRefreshTokenRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
