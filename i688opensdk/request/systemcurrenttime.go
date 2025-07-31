package request

import (
	"net/url"
)

// alibaba.system.currentTime( 获取系统当前时间 )
// https://open.1688.com/api/system/currentTime.html
type SystemCurrentTimeRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *SystemCurrentTimeRequest) CheckParameters() {
	// 无需参数检查
}

// 添加请求参数
func (tk *SystemCurrentTimeRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称
func (tk *SystemCurrentTimeRequest) GetApiName() string {
	return "system/currentTime"
}

// 返回请求参数
func (tk *SystemCurrentTimeRequest) GetParameters() url.Values {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	return *tk.Parameters
}
