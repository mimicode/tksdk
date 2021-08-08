package request

//suning.netalliance.morerecommend.get 关联商品推荐接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.morerecommend.get
type SuningNetallianceMorerecommendGetRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceMorerecommendGetRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceMorerecommendGetRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceMorerecommendGetRequest) GetApiName() string {
	return "suning.netalliance.morerecommend.get"
}

//方法名称
func (tk *SuningNetallianceMorerecommendGetRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceMorerecommendGetRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
