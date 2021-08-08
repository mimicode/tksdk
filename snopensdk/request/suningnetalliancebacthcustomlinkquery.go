package request

//suning.netalliance.bacthcustomlink.query 批量转链
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.bacthcustomlink.query
type SuningNetallianceBacthcustomlinkQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceBacthcustomlinkQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceBacthcustomlinkQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceBacthcustomlinkQueryRequest) GetApiName() string {
	return "suning.netalliance.bacthcustomlink.query"
}

//方法名称
func (tk *SuningNetallianceBacthcustomlinkQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceBacthcustomlinkQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
