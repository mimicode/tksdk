package comvipadpapiopenserviceunionurlv2servicecheckuserwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUrlV2Service 根据openId判断用户是否是渠道用户信息, 注意:如果scene传值为1，是渠道新用户判断，而非唯品会新用户判断
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
	Code int    `json:"code"` // 状态码，1-成功，0-失败
	Msg  string `json:"msg"`  // 状态信息
	Data Data   `json:"data"` // 数据
}

type Data struct {
	IsNewUser bool `json:"isNewUser"` // 是否是新用户，true-是新用户，false-不是新用户
}
