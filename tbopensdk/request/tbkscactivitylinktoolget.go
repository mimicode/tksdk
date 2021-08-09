package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.activitylink.toolget( 淘宝联盟官方活动推广API-工具 )
//http://open.taobao.com/api.htm?docId=41921&docType=2&scopeId=15675
type TbkScActivitylinkToolgetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScActivitylinkToolgetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")

	utils.CheckNotNull(tk.Parameters.Get("promotion_scene_id"), "promotion_scene_id")
}

//添加请求参数
func (tk *TbkScActivitylinkToolgetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScActivitylinkToolgetRequest) GetApiName() string {
	return "taobao.tbk.sc.activitylink.toolget"
}

//返回请求参数
func (tk *TbkScActivitylinkToolgetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
