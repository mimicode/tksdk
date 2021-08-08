package request

//suning.netalliance.inverstmentcommodity.query 高佣专区商品查询接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.inverstmentcommodity.query
type SuningNetallianceInverstmentcommodityQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceInverstmentcommodityQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceInverstmentcommodityQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceInverstmentcommodityQueryRequest) GetApiName() string {
	return "suning.netalliance.inverstmentcommodity.query"
}

//方法名称
func (tk *SuningNetallianceInverstmentcommodityQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceInverstmentcommodityQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
