package request

import (
	"net/url"
)

//pdd.ddk.cashgift.data.query查询多多礼金效果数据
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.cashgift.data.query
type PddDdkCashgiftDataQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkCashgiftDataQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkCashgiftDataQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkCashgiftDataQueryRequest) GetApiName() string {
	return "pdd.ddk.cashgift.data.query"
}

//返回请求参数
func (tk *PddDdkCashgiftDataQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
