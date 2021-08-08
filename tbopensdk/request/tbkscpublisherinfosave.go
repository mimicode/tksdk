package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.sc.publisher.info.save( 淘宝客渠道信息备案 - 社交 )
//http://open.taobao.com/api.htm?docId=37988&docType=2&scopeId=14474
type TbkScPublisherInfoSaveRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScPublisherInfoSaveRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("inviter_code"), "inviter_code")
	utils2.CheckNotNull(tk.Parameters.Get("info_type"), "info_type")
	utils2.CheckNumber(tk.Parameters.Get("info_type"), "info_type")
}

//添加请求参数
func (tk *TbkScPublisherInfoSaveRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScPublisherInfoSaveRequest) GetApiName() string {
	return "taobao.tbk.sc.publisher.info.save"
}

//返回请求参数
func (tk *TbkScPublisherInfoSaveRequest) GetParameters() url.Values {
	return *tk.Parameters
}
