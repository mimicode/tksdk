package request

//suning.netalliance.selectrecommendcommodity.query 商品精选接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.selectrecommendcommodity.query
type SuningNetallianceSelectrecommendcommodityQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceSelectrecommendcommodityQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceSelectrecommendcommodityQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceSelectrecommendcommodityQueryRequest) GetApiName() string {
	return "suning.netalliance.selectrecommendcommodity.query"
}

//方法名称
func (tk *SuningNetallianceSelectrecommendcommodityQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceSelectrecommendcommodityQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
