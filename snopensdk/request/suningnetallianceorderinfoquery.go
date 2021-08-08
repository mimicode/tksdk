package request

//suning.netalliance.orderinfo.query 订单批量查询接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.riskmanagementorder.query
type SuningNetallianceOrderinfoQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceOrderinfoQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceOrderinfoQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceOrderinfoQueryRequest) GetApiName() string {
	return "suning.netalliance.orderinfo.query"
}

//方法名称
func (tk *SuningNetallianceOrderinfoQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceOrderinfoQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
