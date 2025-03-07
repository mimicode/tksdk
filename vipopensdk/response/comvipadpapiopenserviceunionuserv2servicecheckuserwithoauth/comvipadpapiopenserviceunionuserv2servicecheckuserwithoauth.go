package comvipadpapiopenserviceunionuserv2servicecheckuserwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionUserV2Service
// checkUserWithOauth 2.0.0 根据openId判断用户是否是渠道用户信息, 注意:如果scene传值为1，是渠道新用户判断，而非唯品会新用户判断
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/checkUserWithOauth
type Response struct {
	response.TopResponse
	Success CheckUserResponse `json:"result"`
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
type CheckUserResponse struct {
	Result int `json:"result"` // 校验结果, 0：否, 1：是
}
