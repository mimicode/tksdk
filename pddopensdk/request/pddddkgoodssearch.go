package request

import (
	"net/url"
)

//pdd.ddk.goods.search（多多进宝商品查询）
//https://open.pinduoduo.com/#/apidocument/port?id=27
type PddDdkGoodsSearchRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsSearchRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkGoodsSearchRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkGoodsSearchRequest) GetApiName() string {
	return "pdd.ddk.goods.search"
}

//返回请求参数
func (tk *PddDdkGoodsSearchRequest) GetParameters() url.Values {
	return *tk.Parameters
}
