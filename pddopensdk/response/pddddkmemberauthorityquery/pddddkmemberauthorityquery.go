package pddddkmemberauthorityquery

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.member.authority.query 查询是否绑定备案
type Response struct {
	response2.TopResponse
	AuthorityQueryResponse AuthorityQueryResponse `json:"authority_query_response"`
}

//解析输出结果
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

type AuthorityQueryResponse struct {
	Bind      int64  `json:"bind"`
	RequestID string `json:"request_id"`
}
