package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbBbtItemQueryRequest alibaba.alsc.union.kb.bbt.item.query( 本地联盟爆爆团商品列表 )
// https://open.taobao.com/api.htm?docId=59941&docType=2&scopeId=24987
type AlibabaAlscUnionKbBbtItemQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbBbtItemQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbBbtItemQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbBbtItemQueryRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.bbt.item.query"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbBbtItemQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbBbtItemQueryRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbBbtItemQueryRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
