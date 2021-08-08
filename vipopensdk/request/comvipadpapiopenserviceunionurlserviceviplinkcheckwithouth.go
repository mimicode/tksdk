package request

//com.vip.adp.api.open.service.UnionUrlService 检测一段文本中是否有唯品会链接-需要oauth授权
//http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=vipLinkCheckWithOuth
type ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest) CheckParameters() {

}

//添加请求参数
func (tk *ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest) GetApiName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest) GetMethod() string {
	return "vipLinkCheckWithOuth"
}

//方法名称
func (tk *ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest) GetVersion() string {
	return "1.0.0"
}

//返回请求参数
func (tk *ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
