package request

import (
	"net/url"
)

//taobao.shopcats.list.get( 获取前台展示的店铺类目 )
//http://open.taobao.com/api.htm?docId=64&docType=2&scopeId=386
type ShopcatsListGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *ShopcatsListGetRequest) CheckParameters() {
}

//添加请求参数
func (tk *ShopcatsListGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *ShopcatsListGetRequest) GetApiName() string {
	return "taobao.shopcats.list.get"
}

//返回请求参数
func (tk *ShopcatsListGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
