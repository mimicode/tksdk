package pddddkoauthgoodspidgenerate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.oauth.goods.pid.generate多多进宝推广位创建接口
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
	PIDList        []PIDList `json:"p_id_list"`
	RemainPIDCount int64     `json:"remain_pid_count"`
	RequestID      string    `json:"request_id"`
}

type PIDList struct {
	CreateTime int64  `json:"create_time"`
	MediaID    int64  `json:"media_id"`
	PIDName    string `json:"pid_name"`
	PID        string `json:"p_id"`
}
