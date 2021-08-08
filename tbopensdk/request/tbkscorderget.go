package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.sc.tborder.get( 淘宝客订单查询 - 社交 )
//http://open.taobao.com/api.htm?docId=38078&docType=2&scopeId=14814
type TbkScOrderGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScOrderGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils2.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")
}

//添加请求参数
func (tk *TbkScOrderGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScOrderGetRequest) GetApiName() string {
	return "taobao.tbk.sc.tborder.get"
}

//返回请求参数
func (tk *TbkScOrderGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
