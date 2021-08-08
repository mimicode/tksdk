package request

import (
	"net/url"
)

//taobao.ju.items.search( 聚划算商品搜索接口 )
//http://open.taobao.com/api.htm?docId=28762&docType=2&scopeId=11655
type JuItemsSearchRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JuItemsSearchRequest) CheckParameters() {

}

//添加请求参数
func (tk *JuItemsSearchRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *JuItemsSearchRequest) GetApiName() string {
	return "taobao.ju.items.search"
}

//返回请求参数
func (tk *JuItemsSearchRequest) GetParameters() url.Values {
	return *tk.Parameters
}
