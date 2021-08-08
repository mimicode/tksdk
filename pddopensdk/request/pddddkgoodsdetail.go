package request

import (
	"net/url"
)

//pdd.ddk.goods.detail（多多进宝商品详情查询）
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.detail
type PddDdkGoodsDetailRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsDetailRequest) CheckParameters() {
}

//添加请求参数
func (tk *PddDdkGoodsDetailRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkGoodsDetailRequest) GetApiName() string {
	return "pdd.ddk.goods.detail"
}

//返回请求参数
func (tk *PddDdkGoodsDetailRequest) GetParameters() url.Values {
	return *tk.Parameters
}
