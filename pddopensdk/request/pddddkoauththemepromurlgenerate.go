package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.theme.prom.url.generate（多多进宝主题推广链接生成
//https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.oauth.theme.prom.url.generate
type PddDdkOauthThemePromUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthThemePromUrlGenerateRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils2.CheckNotNull(tk.Parameters.Get("theme_id_list"), "theme_id_list")

}

//添加请求参数
func (tk *PddDdkOauthThemePromUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthThemePromUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.oauth.theme.prom.url.generate"
}

//返回请求参数
func (tk *PddDdkOauthThemePromUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
