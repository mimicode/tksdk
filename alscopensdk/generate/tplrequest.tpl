package request

import (
    "encoding/json"
    --UTIL--
	"net/url"
)

// --APINAME--Request --APIDESC--
//--APIURL--
type --APINAME--Request struct {
	Parameters *url.Values //请求参数
}

func (tk *--APINAME--Request) CheckParameters() {
	--APIPARAMCHECK--
}

// AddParameter 添加请求参数
func (tk *--APINAME--Request) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *--APINAME--Request) GetApiName() string {
	return "--APIORGNAME--"
}

// GetParameters 返回请求参数
func (tk *--APINAME--Request) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *--APINAME--Request) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *--APINAME--Request) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}