package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.oauth.pid.mediaid.bind批量绑定推广位的媒体id
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.pid.mediaid.bind&permissionId=7
type PddDdkOauthPidMediaidBindRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthPidMediaidBindRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("media_id"), "media_id")
	utils.CheckNotNull(tk.Parameters.Get("pid_list"), "pid_list")

}

//添加请求参数
func (tk *PddDdkOauthPidMediaidBindRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthPidMediaidBindRequest) GetApiName() string {
	return "pdd.ddk.oauth.pid.mediaid.bind"
}

//返回请求参数
func (tk *PddDdkOauthPidMediaidBindRequest) GetParameters() url.Values {
	return *tk.Parameters
}
