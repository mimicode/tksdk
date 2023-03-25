package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbItemStoreDetailGetRequest alibaba.alsc.union.kb.item.store.detail.get( 本地联盟口碑门店详情 )
// https://open.taobao.com/api.htm?docId=62009&docType=2&scopeId=24408
type AlibabaAlscUnionKbItemStoreDetailGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbItemStoreDetailGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")
	var subFields = make(map[string]string)
	subFields = tk.parseSubParameters(tk.Parameters.Get("query_request"))
	utils.CheckNotNull(tk.getMapVal(subFields, "store_id"), "store_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbItemStoreDetailGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbItemStoreDetailGetRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.item.store.detail.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbItemStoreDetailGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbItemStoreDetailGetRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbItemStoreDetailGetRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
