package --PACKNAME--

import (
	"encoding/json"
	"github.com/mimicode/tksdk/vipopensdk/response"
)

//--APIDESC--
type Response struct {
	response.TopResponse
	--RESULT2-- --RESULT2-- `json:"--RESULT--"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type --RESULT2-- struct {
}