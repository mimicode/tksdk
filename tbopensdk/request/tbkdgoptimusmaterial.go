package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.dg.optimus.material( 淘宝客物料下行-导购 )
//http://open.taobao.com/api.htm?docId=33947&docType=2&scopeId=11655
type TbkDgOptimusMaterialRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgOptimusMaterialRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

}

//添加请求参数
func (tk *TbkDgOptimusMaterialRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkDgOptimusMaterialRequest) GetApiName() string {
	return "taobao.tbk.dg.optimus.material"
}

//返回请求参数
func (tk *TbkDgOptimusMaterialRequest) GetParameters() url.Values {
	return *tk.Parameters
}
