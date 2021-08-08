package request

//suning.netalliance.recommendcommodity.query 小编推荐接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.recommendcommodity.query
type SuningNetallianceRecommendcommodityQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceRecommendcommodityQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceRecommendcommodityQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceRecommendcommodityQueryRequest) GetApiName() string {
	return "suning.netalliance.recommendcommodity.query"
}

//方法名称
func (tk *SuningNetallianceRecommendcommodityQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceRecommendcommodityQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
