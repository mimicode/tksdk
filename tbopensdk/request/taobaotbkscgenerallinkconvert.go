package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.general.link.convert( 淘宝客-服务商-万能转链 )
//https://open.taobao.com/api.htm?docId=65412&docType=2&scopeId=28320
type TbkScGeneralLinkConvertRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScGeneralLinkConvertRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkScGeneralLinkConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScGeneralLinkConvertRequest) GetApiName() string {
	return "taobao.tbk.sc.general.link.convert"
}

//返回请求参数
func (tk *TbkScGeneralLinkConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}

// 设置 required_link_type 参数（可选）
// 用于指定link_info_dto字段出参，多个字段用英文逗号拼接；默认不返回超红链接，
// 仅入参指定coupon_supered_*/cps_supered_*等字段才返回对应链接
func (tk *TbkScGeneralLinkConvertRequest) SetRequiredLinkType(val string) {
	tk.AddParameter("required_link_type", val)
}
