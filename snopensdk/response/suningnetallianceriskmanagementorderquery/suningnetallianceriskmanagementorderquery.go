package suningnetallianceriskmanagementorderquery

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.riskmanagementorder.query 风控订单查询
type Response struct {
	response2.TopResponse
	SnResponseContent SnResponseContent `json:"sn_responseContent"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = "-1"
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	} else {
		if t.SnResponseContent.SnError != nil {
			t.ErrorResponse.ErrorCode = t.SnResponseContent.SnError.ErrorCode
			t.ErrorResponse.ErrorMsg = t.SnResponseContent.SnError.ErrorMsg
		}
	}
}

type SnResponseContent struct {
	SnBody  SnBody   `json:"sn_body"`
	SnError *SnError `json:"sn_error"`
}
type SnError struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type SnBody struct {
	QueryRiskmanagementorder []QueryRiskmanagementorder `json:"queryRiskmanagementorder"`
}

/*
	orderId	String	2222	订单号
	orderItemId	String	666	订单行号
	payTime	String	2019-10-23 10:33:12	支付时间
	custNum	String	7018222104	推广人会员编码
*/
type QueryRiskmanagementorder struct {
	OrderID     string `json:"orderId"`
	PayTime     string `json:"payTime"`
	OrderItemID string `json:"orderItemId"`
	CustNum     string `json:"custNum"`
}

type SnHead struct {
	TotalSize     string `json:"totalSize"`
	PageTotal     string `json:"pageTotal"`
	PageNo        string `json:"pageNo"`
	ReturnMessage string `json:"returnMessage"`
}
