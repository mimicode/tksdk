package tbkdgvegastljcreate

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.vegas.tlj.create( 淘宝客-推广者-淘礼金创建 )
type Response struct {
	response.TopResponse
	CreateResponse CreateResponse `json:"tbk_dg_vegas_tlj_create_response"`
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
	} else {
		if t.CreateResponse.Result.MsgCode != "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Msg = fmt.Sprintf("MsgCode:%s MsgInfo:%s", t.CreateResponse.Result.MsgCode, t.CreateResponse.Result.MsgInfo)
			t.ErrorResponse.SubCode = t.CreateResponse.Result.MsgCode
			t.ErrorResponse.SubMsg = t.CreateResponse.Result.MsgInfo
			t.ErrorResponse.RequestID = t.CreateResponse.RequestID
		} else if t.CreateResponse.Result.Model == nil || t.CreateResponse.Result.Model.SendURL == "" {
			t.ErrorResponse.Code = -3
			t.ErrorResponse.Msg = "用户无权创建淘礼金"
			t.ErrorResponse.SubCode = "用户无权创建淘礼金"
			t.ErrorResponse.SubMsg = "用户无权创建淘礼金"
			t.ErrorResponse.RequestID = t.CreateResponse.RequestID
		}
	}

}

type CreateResponse struct {
	Result    Result `json:"result"`
	RequestID string `json:"request_id"`
}

type Result struct {
	MsgCode string `json:"msg_code"`
	MsgInfo string `json:"msg_info"`
	Success bool   `json:"success"`
	Model   *Model `json:"model"`
}

/*
	rights_id	String	asfasdfasd	淘礼金Id
	send_url	String	https://www.taobao.com	淘礼金领取Url
	vegas_code	String	asfasdfasd	投放code【百川商品详情页业务专用】
	available_fee	String	20.23	创建完成后资金账户可用资金，单位元，保留2位小数
	item_today_num_left	Number	10	媒体针对此商品今日剩余可领取淘礼金数量
*/

type Model struct {
	RightsID         string `json:"rights_id"`
	SendURL          string `json:"send_url"`
	VegasCode        string `json:"vegas_code"`
	AvailableFee     string `json:"available_fee"`
	ItemTodayNumLeft int64  `json:"item_today_num_left"`
}
