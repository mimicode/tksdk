package --APINAME--

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response --APIDESC--
type Response struct {
	response.TopResponse
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
