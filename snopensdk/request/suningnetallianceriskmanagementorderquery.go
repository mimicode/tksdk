package request

//suning.netalliance.riskmanagementorder.query 风控订单查询
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.riskmanagementorder.query
type SuningNetallianceRiskmanagementorderQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceRiskmanagementorderQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceRiskmanagementorderQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceRiskmanagementorderQueryRequest) GetApiName() string {
	return "suning.netalliance.riskmanagementorder.query"
}

//方法名称
func (tk *SuningNetallianceRiskmanagementorderQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceRiskmanagementorderQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
