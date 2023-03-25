package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbOrderCreateRequest alibaba.alsc.union.kb.order.create( openapi订单创建 )
// https://open.taobao.com/api.htm?docId=59883&docType=2&scopeId=24987
type AlibabaAlscUnionKbOrderCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbOrderCreateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("order_dto"), "order_dto")
	var subFields = make(map[string]string)
	subFields = tk.parseSubParameters(tk.Parameters.Get("order_dto"))
	utils.CheckNotNull(tk.getMapVal(subFields, "outer_order_id"), "outer_order_id")
	utils.CheckNotNull(tk.getMapVal(subFields, "item_id"), "item_id")
	utils.CheckNotNull(tk.getMapVal(subFields, "title"), "title")
	utils.CheckNumber(tk.getMapVal(subFields, "quantity"), "quantity")
	utils.CheckNumber(tk.getMapVal(subFields, "pay_order_fee"), "pay_order_fee")
	utils.CheckNumber(tk.getMapVal(subFields, "order_fee"), "order_fee")
	utils.CheckNumber(tk.getMapVal(subFields, "sell_price"), "sell_price")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbOrderCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbOrderCreateRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.order.create"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbOrderCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbOrderCreateRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbOrderCreateRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
