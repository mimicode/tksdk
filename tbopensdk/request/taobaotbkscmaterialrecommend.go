package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.material.recommend( 淘宝客-服务商-物料精选升级版 )
//https://open.taobao.com/api.htm?docId=62191&docType=2&scopeId=16287
type TbkScMaterialRecommendRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScMaterialRecommendRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNotNull(tk.Parameters.Get("material_id"), "material_id")

}

//添加请求参数
func (tk *TbkScMaterialRecommendRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScMaterialRecommendRequest) GetApiName() string {
	return "taobao.tbk.sc.material.recommend"
}

//返回请求参数
func (tk *TbkScMaterialRecommendRequest) GetParameters() url.Values {
	return *tk.Parameters
}
