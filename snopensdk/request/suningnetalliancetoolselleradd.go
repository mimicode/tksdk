package request

//suning.netalliance.toolseller.add 生成工具商PID接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.toolseller.add
type SuningNetallianceToolsellerAddRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceToolsellerAddRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceToolsellerAddRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceToolsellerAddRequest) GetApiName() string {
	return "suning.netalliance.toolseller.add"
}

//方法名称
func (tk *SuningNetallianceToolsellerAddRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceToolsellerAddRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
