package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbOrderPayRequest alibaba.alsc.union.kb.order.pay( openapi预下单订单支付 )
// https://open.taobao.com/api.htm?docId=61523&docType=2&scopeId=24987
type AlibabaAlscUnionKbOrderPayRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbOrderPayRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("order_pay_dto"), "order_pay_dto")
	var OrderPayDto = tk.parseSubParameters(tk.Parameters.Get("order_pay_dto"))
	utils.CheckNotNull(tk.getMapVal(OrderPayDto, "outer_order_id"), "outer_order_id")
	utils.CheckNotNull(tk.getMapVal(OrderPayDto, "biz_order_id"), "biz_order_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbOrderPayRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbOrderPayRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.order.pay"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbOrderPayRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbOrderPayRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbOrderPayRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
