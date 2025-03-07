package comvipadpapiopenserviceunionuserv2serviceuserverifywithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionUserV2Service
// userVerifyWithOauth 2.0.0 CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/userVerifyWithOauth
type Response struct {
	response.TopResponse
	Success UserVerifyResponse `json:"result"`
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

// 响应结构
type UserVerifyResponse struct {
	Result  int    `json:"result"`  // 校验结果, 0：否, 1：是 -2：查询异常,请重试
	ExtData string `json:"extData"` // 补充信息
}
