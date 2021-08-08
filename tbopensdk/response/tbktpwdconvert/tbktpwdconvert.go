package tbktpwdconvert

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.tpwd.convert( 淘口令转链 )
type Response struct {
	response.TopResponse
	TbkTpwdConvertResult Result `json:"tbk_tpwd_convert_response"`
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
	NumIid    string `json:"num_iid"`
	ClickURL  string `json:"click_url"`
	SellerID  string `json:"seller_id"`
	OriginURL string `json:"origin_url"`
	OriginPID string `json:"origin_pid"`
}
