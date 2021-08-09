package request

import (
	utils "github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.activity.info.get( 淘宝客-服务商-官方活动信息获取 新接口 )
//https://open.taobao.com/api.htm?docId=48417&docType=2&scopeId=18353
type TbkScActivityInfoGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScActivityInfoGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("activity_material_id"), "activity_material_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")

}

//添加请求参数
func (tk *TbkScActivityInfoGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScActivityInfoGetRequest) GetApiName() string {
	return "taobao.tbk.sc.activity.info.get"
}

//返回请求参数
func (tk *TbkScActivityInfoGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
