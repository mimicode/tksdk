package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbStoreQueryRequest alibaba.alsc.union.kb.store.query( 本地联盟口碑商户列表 )
// https://open.taobao.com/api.htm?docId=63267&docType=2&scopeId=24408
type AlibabaAlscUnionKbStoreQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbStoreQueryRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page_number"), "page_number")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNotNull(tk.Parameters.Get("biz_type"), "biz_type")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbStoreQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbStoreQueryRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.store.query"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbStoreQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbStoreQueryRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbStoreQueryRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
