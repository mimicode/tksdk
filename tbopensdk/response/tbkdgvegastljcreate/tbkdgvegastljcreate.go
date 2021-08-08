package tbkdgvegastljcreate

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/tbopensdk/response"
	tbkdgvegastljcreateresult2 "github.com/mimicode/tksdk/tbopensdk/response/tbkdgvegastljcreateresult"
)

//taobao.tbk.dg.vegas.tlj.create( 淘宝客-推广者-淘礼金创建 )
type Response struct {
	response.TopResponse
	TbkDgVegasTljCreateResponse tbkdgvegastljcreateresult2.TbkDgVegasTljCreateResponse `json:"tbk_dg_vegas_tlj_create_response"`
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
	} else {
		if t.TbkDgVegasTljCreateResponse.Result.MsgCode != "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Msg = fmt.Sprintf("MsgCode:%s MsgInfo:%s", t.TbkDgVegasTljCreateResponse.Result.MsgCode, t.TbkDgVegasTljCreateResponse.Result.MsgInfo)
			t.ErrorResponse.SubCode = t.TbkDgVegasTljCreateResponse.Result.MsgCode
			t.ErrorResponse.SubMsg = t.TbkDgVegasTljCreateResponse.Result.MsgInfo
			t.ErrorResponse.RequestID = t.TbkDgVegasTljCreateResponse.RequestID
		} else if t.TbkDgVegasTljCreateResponse.Result.Model == nil || t.TbkDgVegasTljCreateResponse.Result.Model.SendURL == "" {
			t.ErrorResponse.Code = -3
			t.ErrorResponse.Msg = "用户无权创建淘礼金"
			t.ErrorResponse.SubCode = "用户无权创建淘礼金"
			t.ErrorResponse.SubMsg = "用户无权创建淘礼金"
			t.ErrorResponse.RequestID = t.TbkDgVegasTljCreateResponse.RequestID
		}
	}

}
