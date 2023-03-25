package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbBbtItemStoreDetailGetRequest alibaba.alsc.union.kb.bbt.item.store.detail.get( 本地联盟爆爆团门店详情 )
// https://open.taobao.com/api.htm?docId=59977&docType=2&scopeId=24987
type AlibabaAlscUnionKbBbtItemStoreDetailGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")
	var QueryRequest = tk.parseSubParameters(tk.Parameters.Get("query_request"))
	utils.CheckNotNull(tk.getMapVal(QueryRequest, "store_id"), "store_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.bbt.item.store.detail.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
