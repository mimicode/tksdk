package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.goods.pid.generate多多进宝推广位创建接口
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.pid.generate&permissionId=7
type PddDdkOauthGoodsPidGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsPidGenerateRequest) CheckParameters() {
	utils.CheckMinValue(tk.Parameters.Get("number"), 1, "number")
	utils.CheckMaxValue(tk.Parameters.Get("number"), 100, "number")
	utils.CheckNumber(tk.Parameters.Get("number"), "number")

}

//添加请求参数
func (tk *PddDdkOauthGoodsPidGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsPidGenerateRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.pid.generate"
}

//返回请求参数
func (tk *PddDdkOauthGoodsPidGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
