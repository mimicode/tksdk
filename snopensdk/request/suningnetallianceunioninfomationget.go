package request

//suning.netalliance.unioninfomation.get 单个查询联盟商品信息
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.unioninfomation.get
type SuningNetallianceUnioninfomationGetRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceUnioninfomationGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceUnioninfomationGetRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceUnioninfomationGetRequest) GetApiName() string {
	return "suning.netalliance.unioninfomation.get"
}

//方法名称
func (tk *SuningNetallianceUnioninfomationGetRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceUnioninfomationGetRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
