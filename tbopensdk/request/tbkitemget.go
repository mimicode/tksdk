package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.item.get( 淘宝客商品查询 )
//http://open.taobao.com/api.htm?docId=24515&docType=2&scopeId=11655
type TbkItemGetRequest struct {
	Parameters *url.Values //请求参数
}

//参数检测
func (tk *TbkItemGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("fields"), "fields")
}

//添加请求参数
func (tk *TbkItemGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemGetRequest) GetApiName() string {
	return "taobao.tbk.item.get"
}

//返回请求参数
func (tk *TbkItemGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
