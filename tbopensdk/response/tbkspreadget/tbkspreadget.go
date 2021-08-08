package tbkspreadget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.spread.get( 物料传播方式获取 )
type Response struct {
	response.TopResponse
	TbkSpreadGetResult Result `json:"tbk_spread_get_response"`
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
	RequestID    string              `json:"request_id"`
}

type Results struct {
	TbkSpread []TbkSpread `json:"tbk_spread"`
}

type TbkSpread struct {
	Content string `json:"content"`
	ErrMsg  string `json:"err_msg"`
}
