package request

//suning.netalliance.extensionlink.get 商品和券二合一接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.extensionlink.get
type SuningNetallianceExtensionlinkGetRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceExtensionlinkGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceExtensionlinkGetRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceExtensionlinkGetRequest) GetApiName() string {
	return "suning.netalliance.extensionlink.get"
}

//方法名称
func (tk *SuningNetallianceExtensionlinkGetRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceExtensionlinkGetRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
