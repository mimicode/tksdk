package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.publisher.info.get( 淘宝客信息查询 - 社交 )
//http://open.taobao.com/api.htm?docId=37989&docType=2&scopeId=14474
type TbkScPublisherInfoGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScPublisherInfoGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("info_type"), "info_type")
	utils.CheckNumber(tk.Parameters.Get("info_type"), "info_type")
	utils.CheckNotNull(tk.Parameters.Get("relation_app"), "relation_app")
}

//添加请求参数
func (tk *TbkScPublisherInfoGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScPublisherInfoGetRequest) GetApiName() string {
	return "taobao.tbk.sc.publisher.info.get"
}

//返回请求参数
func (tk *TbkScPublisherInfoGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
