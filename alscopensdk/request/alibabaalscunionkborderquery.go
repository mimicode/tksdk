package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbOrderQueryRequest alibaba.alsc.union.kb.order.query( openapi订单查询 )
// https://open.taobao.com/api.htm?docId=59884&docType=2&scopeId=24987
type AlibabaAlscUnionKbOrderQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbOrderQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("order_query_dto"), "order_query_dto")
	var subFields = make(map[string]string)
	subFields = tk.parseSubParameters(tk.Parameters.Get("order_query_dto"))
	utils.CheckNotNull(tk.getMapVal(subFields, "biz_order_id"), "biz_order_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbOrderQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbOrderQueryRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.order.query"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbOrderQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbOrderQueryRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbOrderQueryRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
