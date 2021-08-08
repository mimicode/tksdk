package request

//suning.netalliance.appletextensionlink.get 商品和券二合一（小程序）
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.appletextensionlink.get
type SuningNetallianceAppletextensionlinkGetRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceAppletextensionlinkGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceAppletextensionlinkGetRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceAppletextensionlinkGetRequest) GetApiName() string {
	return "suning.netalliance.appletextensionlink.get"
}

//方法名称
func (tk *SuningNetallianceAppletextensionlinkGetRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceAppletextensionlinkGetRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
