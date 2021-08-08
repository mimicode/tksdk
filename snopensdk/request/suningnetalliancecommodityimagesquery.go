package request

//suning.netalliance.commodityimages.query 商品图文详情查询
//https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.commodityimages.query
type SuningNetallianceCommodityimagesQueryRequest struct {
	Parameters map[string]interface{} //请求参数
}

func (tk *SuningNetallianceCommodityimagesQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *SuningNetallianceCommodityimagesQueryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

//返回接口名称
func (tk *SuningNetallianceCommodityimagesQueryRequest) GetApiName() string {
	return "suning.netalliance.commodityimages.query"
}

//方法名称
func (tk *SuningNetallianceCommodityimagesQueryRequest) GetVersion() string {
	return "v1.2"
}

//返回请求参数
func (tk *SuningNetallianceCommodityimagesQueryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}
