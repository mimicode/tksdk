package request

// com.vip.adp.api.open.service.UnionUrlService getChannelUrlByType 渠道根据落地页类型生链接口
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/getChannelUrlByType
type ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest) CheckParameters() {

}

// 添加请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest) GetMethod() string {
	return "getChannelUrlByType"
}

// 方法名称
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest) GetVersion() string {
	return "1.3.0"
}

// 返回请求参数
func (tk *ComVipAdpApiOpenServiceUnionUrlServiceGetChannelUrlByTypeRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
