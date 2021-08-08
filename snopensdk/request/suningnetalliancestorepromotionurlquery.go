package request

//suning.netalliance.storepromotionurl.query 获取商品或店铺推广链接接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.storepromotionurl.query
type SuningNetallianceStorepromotionurlQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceStorepromotionurlQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceStorepromotionurlQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceStorepromotionurlQueryRequest) GetApiName() string {
	return "suning.netalliance.storepromotionurl.query"
}

//方法名称
func (tk *SuningNetallianceStorepromotionurlQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceStorepromotionurlQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
