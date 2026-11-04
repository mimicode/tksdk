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

// required_link_type 可选字段名（官方字段表：用于指定 link_info_dto 出参）
//
// 特别说明：超红链接默认不返回，仅 required_link_type 入参指定下列字段才返回；
// 非超红活动期间请求超红字段会返回普通推广链接。多个字段用英文逗号拼接传入。
// 官方文档：https://open.taobao.com/api.htm?docId=65412&docType=2&scopeId=28320
const (
	RequiredLinkTypeCouponSuperedLongURL   = "coupon_supered_long_url"   // 超红二合一长链接
	RequiredLinkTypeCouponSuperedLongTpwd  = "coupon_supered_long_tpwd"  // 超红二合一长口令
	RequiredLinkTypeCouponSuperedShortURL  = "coupon_supered_short_url"  // 超红二合一短链接
	RequiredLinkTypeCouponSuperedShortTpwd = "coupon_supered_short_tpwd" // 超红二合一短口令
	RequiredLinkTypeCPSSuperedLongURL      = "cps_supered_long_url"      // 超红cps长链接
	RequiredLinkTypeCPSSuperedLongTpwd     = "cps_supered_long_tpwd"     // 超红cps长口令
	RequiredLinkTypeCPSSuperedShortURL     = "cps_supered_short_url"     // 超红cps短链接
	RequiredLinkTypeCPSSuperedShortTpwd    = "cps_supered_short_tpwd"    // 超红cps短口令
)

// 设置 required_link_type 参数（可选）
// val 为需要返回的 link_info_dto 字段名，多个用英文逗号拼接；
// 例如：RequiredLinkTypeCouponSuperedShortURL + "," + RequiredLinkTypeCPSSuperedShortTpwd
func (tk *TbkScGeneralLinkConvertRequest) SetRequiredLinkType(val string) {
	tk.AddParameter("required_link_type", val)
}
