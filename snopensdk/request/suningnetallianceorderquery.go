package request

//suning.netalliance.order.query 网盟订单信息批量查询
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.order.query
type SuningNetallianceOrderQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceOrderQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceOrderQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceOrderQueryRequest) GetApiName() string {
	return "suning.netalliance.order.query"
}

//方法名称
func (tk *SuningNetallianceOrderQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceOrderQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
