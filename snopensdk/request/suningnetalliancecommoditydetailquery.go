package request

//suning.netalliance.commoditydetail.query 推广商品详情信息接口
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.commoditydetail.query
type SuningNetallianceCommoditydetailQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceCommoditydetailQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceCommoditydetailQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceCommoditydetailQueryRequest) GetApiName() string {
	return "suning.netalliance.commoditydetail.query"
}

//方法名称
func (tk *SuningNetallianceCommoditydetailQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceCommoditydetailQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
