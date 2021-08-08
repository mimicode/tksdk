package pddddkgoodszsuniturlgen

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.goods.zs.unit.url.gen多多进宝转链接口
type Response struct {
	response2.TopResponse
	GoodsZsUnitGenerateResponse GoodsZsUnitGenerateResponse `json:"goods_zs_unit_generate_response"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type GoodsZsUnitGenerateResponse struct {
	MobileShortURL           string `json:"mobile_short_url"`
	MobileURL                string `json:"mobile_url"`
	MultiGroupMobileShortURL string `json:"multi_group_mobile_short_url"`
	MultiGroupMobileURL      string `json:"multi_group_mobile_url"`
	MultiGroupShortURL       string `json:"multi_group_short_url"`
	MultiGroupURL            string `json:"multi_group_url"`
	ShortURL                 string `json:"short_url"`
	URL                      string `json:"url"`
}
