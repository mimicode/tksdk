package comvipadpapiopenserviceuniongoodsv2servicegetbrandsnlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionGoodsV2Service
// getBrandSnListWithOauth 2.0.0 精选品牌列表查询-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getBrandSnListWithOauth
type Response struct {
	response.TopResponse
	Success Success `json:"result"`
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

type Success struct {
	BrandSnList []string            `json:"brandSnList"` // 品牌Sn码列表
	Total       int                 `json:"total"`       // 总记录数
	BrandSnData map[string][]string `json:"brandSnData"` // 对应传入的品牌类型返回的品牌sn
}
