package shopgetbytitle

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.shop.getbytitle( 根据店铺名称获取店铺信息 )
type Response struct {
	response.TopResponse
	ShopGetbytitleResult Result `json:"shop_getbytitle_response"`
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
	IsError    bool   `json:"is_error"`
	ResultShop string `json:"result_shop"`
	RequestID  string `json:"request_id"`
}
