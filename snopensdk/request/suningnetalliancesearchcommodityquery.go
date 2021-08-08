package request

//suning.netalliance.searchcommodity.query 关键字商品查询接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.searchcommodity.query
type SuningNetallianceSearchcommodityQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceSearchcommodityQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceSearchcommodityQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceSearchcommodityQueryRequest) GetApiName() string {
	return "suning.netalliance.searchcommodity.query"
}

//方法名称
func (tk *SuningNetallianceSearchcommodityQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceSearchcommodityQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
