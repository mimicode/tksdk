package pddddkoauthgoodspidquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.goods.pid.query多多客已生成推广位信息查询
type Response struct {
	response.TopResponse
	PIDQueryResponse PIDQueryResponse `json:"p_id_query_response"`
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

type PIDQueryResponse struct {
	PIDList    []PIDList `json:"p_id_list"`
	TotalCount int64     `json:"total_count"`
	RequestID  string    `json:"request_id"`
}

type PIDList struct {
	CreateTime int64  `json:"create_time"`
	MediaID    int64  `json:"media_id"`
	PIDName    string `json:"pid_name"`
	PID        string `json:"p_id"`
}
