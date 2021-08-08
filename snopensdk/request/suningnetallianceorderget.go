package request

//suning.netalliance.order.get 网盟订单信息单独查询
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.order.get
type SuningNetallianceOrderGetRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceOrderGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceOrderGetRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceOrderGetRequest) GetApiName() string {
	return "suning.netalliance.order.get"
}

//方法名称
func (tk *SuningNetallianceOrderGetRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceOrderGetRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
