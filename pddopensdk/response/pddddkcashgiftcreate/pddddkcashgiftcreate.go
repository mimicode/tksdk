package pddddkcashgiftcreate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.cashgift.create创建多多礼金
type Response struct {
	response.TopResponse
	CreateCashgiftResponse CreateCashgiftResponse `json:"create_cashgift_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type CreateCashgiftResponse struct {
	CashGiftID int64 `json:"cash_gift_id"`
	Success    bool  `json:"success"`
}
