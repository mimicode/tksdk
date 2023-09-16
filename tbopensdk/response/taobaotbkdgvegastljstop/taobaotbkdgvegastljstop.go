package taobaotbkdgvegastljstop

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.vegas.tlj.stop( 淘宝客-推广者-淘礼金暂停发放 )
type Response struct {
	response.TopResponse
	TbkDgVegasTljStopResponse TbkDgVegasTljStopResponse `json:"tbk_dg_vegas_tlj_stop_response"`
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

type TbkDgVegasTljStopResponse struct {
	Model         Model  `json:"model"`
	MsgInfo       string `json:"msg_info"`
	MsgCode       string `json:"msg_code"`
	ResultSuccess bool   `json:"result_success"`
}

type Model struct {
	UpdateSuccess bool `json:"update_success"`
}
