package tbkjutqgget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.ju.tqg.get( 淘抢购api )
type Response struct {
	response.TopResponse
	TbkJuTqgGetResult Result `json:"tbk_ju_tqg_get_response"`
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
	Results      Results `json:"results"`
	TotalResults int64   `json:"total_results"`
	RequestID    string             `json:"request_id"`
}

type Results struct {
	Results []ResultResults `json:"results"`
}

type ResultResults struct {
	CategoryName string `json:"category_name"`
	ClickURL     string `json:"click_url"`
	EndTime      string `json:"end_time"`
	NumIid       int64  `json:"num_iid"`
	PicURL       string `json:"pic_url"`
	ReservePrice string `json:"reserve_price"`
	SoldNum      int64  `json:"sold_num"`
	StartTime    string `json:"start_time"`
	Title        string `json:"title"`
	TotalAmount  int64  `json:"total_amount"`
	ZkFinalPrice string `json:"zk_final_price"`
}
