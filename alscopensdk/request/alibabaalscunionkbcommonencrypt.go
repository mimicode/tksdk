package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbCommonEncryptRequest alibaba.alsc.union.kb.common.encrypt( 加密方法 )
// https://open.taobao.com/api.htm?docId=61584&docType=2&scopeId=24987
type AlibabaAlscUnionKbCommonEncryptRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbCommonEncryptRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("encrypt_model"), "encrypt_model")
	var EncryptModel = tk.parseSubParameters(tk.Parameters.Get("encrypt_model"))
	utils.CheckNotNull(tk.getMapVal(EncryptModel, "text"), "text")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbCommonEncryptRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbCommonEncryptRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.common.encrypt"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbCommonEncryptRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbCommonEncryptRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbCommonEncryptRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
