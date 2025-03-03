package request

// com.vip.adp.api.open.service.UnionUrlV2Service 渠道根据落地页类型生链接口,需要申请渠道工具商权限包权限
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getChannelUrlByTypeWithOauth
type ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest) GetMethod() string {
	return "getChannelUrlByTypeWithOauth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceGetChannelUrlByTypeWithOauthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
