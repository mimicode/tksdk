package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.sellercats.list.get( 获取前台展示的店铺内卖家自定义商品类目 )
//http://open.taobao.com/api.htm?docId=65&docType=2&scopeId=386
type SellercatsListGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *SellercatsListGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("nick"), "nick")
}

//添加请求参数
func (tk *SellercatsListGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *SellercatsListGetRequest) GetApiName() string {
	return "taobao.sellercats.list.get"
}

//返回请求参数
func (tk *SellercatsListGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
