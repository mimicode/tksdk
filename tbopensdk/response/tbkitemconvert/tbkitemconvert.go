package tbkitemconvert

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.item.convert( 淘宝客商品链接转换 )
type Response struct {
	response.TopResponse
	TbkItemConvertResult Result `json:"tbk_item_convert_response"`
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
	Results   Results `json:"results"`
	RequestID string  `json:"request_id"`
}

type Results struct {
	NTbkItem []NTbkItem `json:"n_tbk_item"`
}

type NTbkItem struct {
	ClickURL string `json:"click_url"`
	NumIid   int64  `json:"num_iid"`
}
