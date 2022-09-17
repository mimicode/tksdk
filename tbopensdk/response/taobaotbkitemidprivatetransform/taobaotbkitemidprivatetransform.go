package taobaotbkitemidprivatetransform

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.itemid.private.transform( 淘宝客-推广者-商品id转化（二方）（专有） )
type Response struct {
	response.TopResponse
	TbkItemidPrivateTransformResponse TbkItemidPrivateTransformResponse `json:"tbk_itemid_private_transform_response"`
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

type TbkItemidPrivateTransformResponse struct {
	Results Results `json:"results"`
}

type Results struct {
	ItemIDTransformDTO []ItemIDTransformDTO `json:"item_id_transform_d_t_o"`
}

type ItemIDTransformDTO struct {
	OriginalItemID string `json:"original_item_id"`
	ItemID         string `json:"item_id"`
}
