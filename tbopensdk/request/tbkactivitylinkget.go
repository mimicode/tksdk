package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.activitylink.get( 淘宝联盟官方活动推广API-媒体 )
//http://open.taobao.com/api.htm?docId=41918&docType=2&scopeId=11655
type TbkActivitylinkGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkActivitylinkGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNotNull(tk.Parameters.Get("promotion_scene_id"), "promotion_scene_id")
	utils.CheckNumber(tk.Parameters.Get("promotion_scene_id"), "promotion_scene_id")

}

//添加请求参数
func (tk *TbkActivitylinkGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkActivitylinkGetRequest) GetApiName() string {
	return "taobao.tbk.activitylink.get"
}

//返回请求参数
func (tk *TbkActivitylinkGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
