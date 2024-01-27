package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.material.optional.upgrade( 淘宝客-服务商-物料搜索升级版 )
//https://open.taobao.com/api.htm?docId=64758&docType=2&scopeId=13991
type TbkScMaterialOptionalUpgradeRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScMaterialOptionalUpgradeRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkScMaterialOptionalUpgradeRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScMaterialOptionalUpgradeRequest) GetApiName() string {
	return "taobao.tbk.sc.material.optional.upgrade"
}

//返回请求参数
func (tk *TbkScMaterialOptionalUpgradeRequest) GetParameters() url.Values {
	return *tk.Parameters
}
