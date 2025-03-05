package comvipadpapiopenserviceunionuserv2serviceunbindopenidwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionUserV2Service
// unbindOpenIdWithOauth 2.0.0 解绑已授权的openId接口服务(需要工具商权限)
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/unbindOpenIdWithOauth
type Response struct {
	response.TopResponse
	Success UnBindOpenIdResponse `json:"success"`
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
type UnBindOpenIdResponse struct {
	Code    int    `json:"code"`    // 返回码：1-成功，其它为失败，失败信息可参考message字段
	Message string `json:"message"` // 返回说明信息
}
