package systemcurrenttime

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// alibaba.system.currentTime( 获取系统当前时间 )
type Response struct {
	response.TopResponse
	CurrentTimeResponse CurrentTimeResponse `json:"alibaba_system_currentTime_response"`
}

type CurrentTimeResponse struct {
	CurrentTime string `json:"currentTime"` // 当前时间，格式：yyyy-MM-dd HH:mm:ss
	RequestID   string `json:"request_id"`  // 请求ID
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
