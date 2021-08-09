package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.optimus.material( 淘宝客擎天柱通用物料API - 社交 )
//http://open.taobao.com/api.htm?docId=37884&docType=2&scopeId=11655
type TbkScOptimusMaterialRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScOptimusMaterialRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")

}

//添加请求参数
func (tk *TbkScOptimusMaterialRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScOptimusMaterialRequest) GetApiName() string {
	return "taobao.tbk.sc.optimus.material"
}

//返回请求参数
func (tk *TbkScOptimusMaterialRequest) GetParameters() url.Values {
	return *tk.Parameters
}
