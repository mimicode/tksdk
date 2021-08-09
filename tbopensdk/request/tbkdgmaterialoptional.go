package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.material.optional( 通用物料搜索API（导购） )
//http://open.taobao.com/api.htm?docId=35896&docType=2&scopeId=11655
type TbkDgMaterialOptionalRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgMaterialOptionalRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkDgMaterialOptionalRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgMaterialOptionalRequest) GetApiName() string {
	return "taobao.tbk.dg.material.optional"
}

//返回请求参数
func (tk *TbkDgMaterialOptionalRequest) GetParameters() url.Values {
	return *tk.Parameters
}
