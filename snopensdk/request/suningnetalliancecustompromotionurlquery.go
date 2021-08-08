package request

//suning.netalliance.custompromotionurl.query 获取自定义推广链接接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.custompromotionurl.query
type SuningNetallianceCustompromotionurlQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceCustompromotionurlQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceCustompromotionurlQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceCustompromotionurlQueryRequest) GetApiName() string {
	return "suning.netalliance.custompromotionurl.query"
}

//方法名称
func (tk *SuningNetallianceCustompromotionurlQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceCustompromotionurlQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
