package request

// ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest com.vip.adp.api.open.service.UnionUrlV2Service
// vipLinkCheckWithOuth 2.0.0 进行cps链接的解析,需要申请渠道包权限进行Oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/vipLinkCheckWithOuth
type ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlV2Service"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest) GetMethod() string {
	return "vipLinkCheckWithOuth"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest) GetVersion() string {
	return "2.0.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlV2ServiceVipLinkCheckWithOuthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
