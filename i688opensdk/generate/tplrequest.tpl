package request

import (
	"net/url"
	--UTIL--
)

//--APIDESC--
//--APIURL--
type --APINAME--Request struct {
	Parameters *url.Values //请求参数
}

func (tk *--APINAME--Request) CheckParameters() {
	--APIPARAMCHECK--
}

//添加请求参数
func (tk *--APINAME--Request) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *--APINAME--Request) GetApiName() string {
	return "--APIORGNAME--"
}

//返回请求参数
func (tk *--APINAME--Request) GetParameters() url.Values {
	return *tk.Parameters
}