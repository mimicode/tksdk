package request

// com.vip.adp.api.open.service.UnionUrlService vipLinkCheck 进行cps链接的解析
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/vipLinkCheck
type ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest) GetMethod() string {
	return "vipLinkCheck"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceVipLinkCheckRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
