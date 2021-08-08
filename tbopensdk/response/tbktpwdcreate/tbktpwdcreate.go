package tbktpwdcreate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.tpwd.create( 淘宝客淘口令 )
type Response struct {
	response.TopResponse
	TbkTpwdCreateResult Result `json:"tbk_tpwd_create_response"`
}

//解析输出结果
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

type Result struct {
	Data      Data   `json:"data"`
	RequestID string `json:"request_id"`
}

type Data struct {
	Model          string `json:"model"`
	PasswordSimple string `json:"password_simple"`
}
