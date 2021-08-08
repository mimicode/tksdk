package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.item.info.get( 淘宝客商品详情（简版） )
//http://open.taobao.com/api.htm?docId=24518&docType=2
type TbkItemInfoGetRequest struct {
	Parameters *url.Values //请求参数
}

//参数检测
func (tk *TbkItemInfoGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("num_iids"), "num_iids")
	utils2.CheckMaxListSize(tk.Parameters.Get("num_iids"), 40, "num_iids")
}

//添加请求参数
func (tk *TbkItemInfoGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemInfoGetRequest) GetApiName() string {
	return "taobao.tbk.item.info.get"
}

//返回请求参数
func (tk *TbkItemInfoGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
