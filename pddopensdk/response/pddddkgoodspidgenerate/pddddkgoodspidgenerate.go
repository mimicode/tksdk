package pddddkgoodspidgenerate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.goods.pid.generate（创建多多进宝推广位)
type Response struct {
	response.TopResponse
	PIDGenerateResponse PIDGenerateResponse `json:"p_id_generate_response"`
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

type PIDGenerateResponse struct {
	PIDList   []PIDList `json:"p_id_list"`
	RequestID string    `json:"request_id"`
}

type PIDList struct {
	CreateTime int64  `json:"create_time"`
	PIDName    string `json:"pid_name"`
	PID        string `json:"p_id"`
}
