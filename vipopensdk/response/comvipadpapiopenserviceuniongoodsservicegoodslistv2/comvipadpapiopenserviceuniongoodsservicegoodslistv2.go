package comvipadpapiopenserviceuniongoodsservicegoodslistv2

import (
	"encoding/json"
	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService goodsListV2 获取联盟在推商品列表-V2版本
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Result struct {
}
