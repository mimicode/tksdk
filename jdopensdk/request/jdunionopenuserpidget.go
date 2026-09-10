package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenUserPidGetRequest jd.union.open.user.pid.get 获取PID【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.user.pid.get
type JdUnionOpenUserPidGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenUserPidGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenUserPidGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenUserPidGetRequest) GetApiName() string {
	return "jd.union.open.user.pid.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenUserPidGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
