package request

//suning.netalliance.ordersettlenew.get 订单结算信息接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.ordersettle.get
type SuningNetallianceOrdersettlenewGetRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceOrdersettlenewGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceOrdersettlenewGetRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceOrdersettlenewGetRequest) GetApiName() string {
	return "suning.netalliance.ordersettlenew.get"
}

//方法名称
func (tk *SuningNetallianceOrdersettlenewGetRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceOrdersettlenewGetRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
