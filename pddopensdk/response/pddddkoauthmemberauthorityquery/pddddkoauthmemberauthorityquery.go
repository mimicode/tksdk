package pddddkoauthmemberauthorityquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.member.authority.query查询是否绑定备案
type Response struct {
	response.TopResponse
	AuthorityQueryResponse AuthorityQueryResponse `json:"authority_query_response"`
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

type AuthorityQueryResponse struct {
	Bind      int64  `json:"bind"` //	1-已绑定；0-未绑定
	RequestID string `json:"request_id"`
}
