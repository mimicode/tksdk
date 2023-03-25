package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbOrderRefundRequest alibaba.alsc.union.kb.order.refund( openapi订单售中退 )
// https://open.taobao.com/api.htm?docId=59885&docType=2&scopeId=24987
type AlibabaAlscUnionKbOrderRefundRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbOrderRefundRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("order_refund_dto"), "order_refund_dto")
	var OrderRefundDto = tk.parseSubParameters(tk.Parameters.Get("order_refund_dto"))
	utils.CheckNotNull(tk.getMapVal(OrderRefundDto, "voucher_list"), "voucher_list")
	var OrderRefundDtoVoucherList = tk.parseSubParameters(tk.getMapVal(OrderRefundDto, "voucher_list"))
	utils.CheckNotNull(tk.getMapVal(OrderRefundDtoVoucherList, "item_id"), "item_id")
	utils.CheckNotNull(tk.getMapVal(OrderRefundDtoVoucherList, "voucher_id"), "voucher_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbOrderRefundRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbOrderRefundRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.order.refund"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbOrderRefundRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbOrderRefundRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbOrderRefundRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
