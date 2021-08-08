package tbkshopconvert

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.shop.convert( 淘宝客店铺链接转换 )
type Response struct {
	response.TopResponse
	TbkShopConvertResult Result `json:"tbk_shop_convert_response"`
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
	NTbkShop []NTbkShop `json:"n_tbk_shop"`
}

type NTbkShop struct {
	ClickURL string `json:"click_url"`
	UserID   int64  `json:"user_id"`
}
