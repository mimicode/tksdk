package request

// com.vip.adp.api.open.service.UnionUrlV2Service genByVIPUrlWithOauth 根据唯品会链接生成联盟链接-需要申请渠道工具商权限包权限
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/genByVIPUrlWithOauth
type ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest) GetMethod() string {
	return "genByVIPUrlWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGenByVIPUrlWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
