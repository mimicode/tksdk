package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbBbtItemDetailGetRequest alibaba.alsc.union.kb.bbt.item.detail.get( 本地联盟爆爆团商品详情 )
// https://open.taobao.com/api.htm?docId=59975&docType=2&scopeId=24987
type AlibabaAlscUnionKbBbtItemDetailGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbBbtItemDetailGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")
	var QueryRequest = tk.parseSubParameters(tk.Parameters.Get("query_request"))
	utils.CheckNotNull(tk.getMapVal(QueryRequest, "item_id"), "item_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbBbtItemDetailGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbBbtItemDetailGetRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.bbt.item.detail.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbBbtItemDetailGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbBbtItemDetailGetRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbBbtItemDetailGetRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
