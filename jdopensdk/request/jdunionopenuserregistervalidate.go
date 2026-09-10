package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenUserRegisterValidateRequest jd.union.open.user.register.validate 实时RTA接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.user.register.validate
type JdUnionOpenUserRegisterValidateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenUserRegisterValidateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenUserRegisterValidateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenUserRegisterValidateRequest) GetApiName() string {
	return "jd.union.open.user.register.validate"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenUserRegisterValidateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
