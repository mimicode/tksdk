package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbStoreItemQueryRequest alibaba.alsc.union.kb.store.item.query( 本地联盟口碑门店商品列表 )
// https://open.taobao.com/api.htm?docId=63268&docType=2&scopeId=24408
type AlibabaAlscUnionKbStoreItemQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbStoreItemQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("store_id"), "store_id")
	utils.CheckNotNull(tk.Parameters.Get("biz_type"), "biz_type")
	utils.CheckNotNull(tk.Parameters.Get("pid"), "pid")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbStoreItemQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbStoreItemQueryRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.store.item.query"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbStoreItemQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbStoreItemQueryRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbStoreItemQueryRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
