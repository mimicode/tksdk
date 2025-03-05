package comvipadpapiopenserviceunionurlv2serviceviplinkcheckwithouth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionUrlV2Service
// vipLinkCheckWithOuth 2.0.0 进行cps链接的解析,需要申请渠道包权限进行Oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/vipLinkCheckWithOuth
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
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
	SuccessMap map[string]interface{} `json:"successMap"` // 解析成功的结果 map结构，key为输入的待解析的链接，value为解析结果
	FailMap    map[string]interface{} `json:"failMap"`    // 解析失败的结构 map结构，key为输入的待解析的链接，value为解析结果
}
