package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.material.optional( 通用物料搜索API )
//http://open.taobao.com/api.htm?docId=35263&docType=2&scopeId=13991
type TbkScMaterialOptionalRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScMaterialOptionalRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")

	utils.CheckMaxValue(tk.Parameters.Get("page_size"), 100, "page_size")

}

//添加请求参数
func (tk *TbkScMaterialOptionalRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScMaterialOptionalRequest) GetApiName() string {
	return "taobao.tbk.sc.material.optional"
}

//返回请求参数
func (tk *TbkScMaterialOptionalRequest) GetParameters() url.Values {
	return *tk.Parameters
}
