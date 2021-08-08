package request

//suning.netalliance.hoistinglink.query 获取吊起链接接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.hoistinglink.query
type SuningNetallianceHoistinglinkQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceHoistinglinkQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceHoistinglinkQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceHoistinglinkQueryRequest) GetApiName() string {
	return "suning.netalliance.hoistinglink.query"
}

//方法名称
func (tk *SuningNetallianceHoistinglinkQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceHoistinglinkQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
