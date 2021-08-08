package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.resource.url.gen拼多多主站频道推广接口
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.resource.url.gen&permissionId=7
type PddDdkOauthResourceUrlGenRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthResourceUrlGenRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("pid"), "pid")

}

//添加请求参数
func (tk *PddDdkOauthResourceUrlGenRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthResourceUrlGenRequest) GetApiName() string {
	return "pdd.ddk.oauth.resource.url.gen"
}

//返回请求参数
func (tk *PddDdkOauthResourceUrlGenRequest) GetParameters() url.Values {
	return *tk.Parameters
}
