package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest alibaba.alsc.union.kb.bbt.item.store.relation.query( 本地联盟爆爆团商品门店关系 )
// https://open.taobao.com/api.htm?docId=59976&docType=2&scopeId=24987
type AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")
	var QueryRequest = tk.parseSubParameters(tk.Parameters.Get("query_request"))
	utils.CheckNotNull(tk.getMapVal(QueryRequest, "item_id"), "item_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.bbt.item.store.relation.query"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
