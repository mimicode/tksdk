package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.goods.zs.unit.url.gen多多进宝转链接口
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.zs.unit.url.gen&permissionId=7
type PddDdkOauthGoodsZsUnitUrlGenRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsZsUnitUrlGenRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils.CheckNotNull(tk.Parameters.Get("source_url"), "source_url")

}

//添加请求参数
func (tk *PddDdkOauthGoodsZsUnitUrlGenRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsZsUnitUrlGenRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.zs.unit.url.gen"
}

//返回请求参数
func (tk *PddDdkOauthGoodsZsUnitUrlGenRequest) GetParameters() url.Values {
	return *tk.Parameters
}
