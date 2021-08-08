package request

//suning.netalliance.couponinfo.query 查询券领用情况
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.couponinfo.query
type SuningNetallianceCouponinfoQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceCouponinfoQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceCouponinfoQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceCouponinfoQueryRequest) GetApiName() string {
	return "suning.netalliance.couponinfo.query"
}

//方法名称
func (tk *SuningNetallianceCouponinfoQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceCouponinfoQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
