package tbkscinvitecodeget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.invitecode.get( 淘宝客邀请码生成-社交 )
type Response struct {
	response.TopResponse
	TbkScInvitecodeGetResult Result `json:"tbk_sc_invitecode_get_response"`
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
	InviterCode string `json:"inviter_code"`
}
