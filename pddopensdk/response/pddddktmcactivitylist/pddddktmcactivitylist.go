package pddddktmcactivitylist

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.tmc.activity.list千万神券
type Response struct {
	response.TopResponse
	TmcAtyListResponse TmcAtyListResponse `json:"tmc_aty_list_response"`
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

type TmcAtyListResponse struct {
	Total          int64          `json:"total"`
	QueryEndTime   string         `json:"query_end_time"`
	QueryStartTime string         `json:"query_start_time"`
	TmcAtyVoList   []TmcAtyVoList `json:"tmc_aty_vo_list"`
	RequestID      string         `json:"request_id"`
}

type TmcAtyVoList struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Name      string `json:"name"`
	ID        int64  `json:"id"`
	Type      int64  `json:"type"`
}
