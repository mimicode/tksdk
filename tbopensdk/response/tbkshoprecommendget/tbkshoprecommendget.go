package tbkshoprecommendget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.shop.recommend.get( 淘宝客店铺关联推荐查询 )
type Response struct {
	response.TopResponse
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
