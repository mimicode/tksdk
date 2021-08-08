package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

// PddDdkGoodsZsUnitUrlGenRequest pdd.ddk.goods.zs.unit.url.gen多多进宝转链接口
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.zs.unit.url.gen
type PddDdkGoodsZsUnitUrlGenRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsZsUnitUrlGenRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils2.CheckNotNull(tk.Parameters.Get("custom_parameters"), "custom_parameters")
	utils2.CheckNotNull(tk.Parameters.Get("source_url"), "source_url")
}

//添加请求参数
func (tk *PddDdkGoodsZsUnitUrlGenRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkGoodsZsUnitUrlGenRequest) GetApiName() string {
	return "pdd.ddk.goods.zs.unit.url.gen"
}

//返回请求参数
func (tk *PddDdkGoodsZsUnitUrlGenRequest) GetParameters() url.Values {
	return *tk.Parameters
}
