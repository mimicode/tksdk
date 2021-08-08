package request

//suning.netalliance.searchcommoditynew.query 关键字商品查询接口(新)
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.searchcommoditynew.query
type SuningNetallianceSearchcommoditynewQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceSearchcommoditynewQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceSearchcommoditynewQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceSearchcommoditynewQueryRequest) GetApiName() string {
	return "suning.netalliance.searchcommoditynew.query"
}

//方法名称
func (tk *SuningNetallianceSearchcommoditynewQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceSearchcommoditynewQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
