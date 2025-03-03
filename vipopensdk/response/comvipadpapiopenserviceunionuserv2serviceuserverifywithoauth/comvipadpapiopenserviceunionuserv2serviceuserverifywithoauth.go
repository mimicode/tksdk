package comvipadpapiopenserviceunionuserv2serviceuserverifywithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUserV2Service userVerifyWithOauth CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等
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
	Result  int    `json:"result"`  // 校验结果, 0：否, 1：是 -2：查询异常,请重试
	ExtData string `json:"extData"` // 补充信息
}
