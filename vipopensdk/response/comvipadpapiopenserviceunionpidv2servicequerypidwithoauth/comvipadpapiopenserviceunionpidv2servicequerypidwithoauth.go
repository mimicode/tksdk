package comvipadpapiopenserviceunionpidv2servicequerypidwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionPidV2Service
// queryPidWithOauth 2.0.0 查询PID-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPidV2Service-2.0.0&methodName=queryPidWithOauth
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

// Success CpsUnionPidQueryResponse结构
type Success struct {
	PidInfoList []PidInfo `json:"pidInfoList"` // 推广位Pid信息列表
	Total       int       `json:"total"`       // 总条数
}

// PidInfo 推广位信息结构
type PidInfo struct {
	Pid           string `json:"pid"`           // 推广位ID
	PidName       string `json:"pidName"`       // 推广位名称
	CreateTime    int64  `json:"createTime"`    // 该推广位创建的时间
	Ucode         string `json:"ucode"`         // 所属ucode
	B2cUserId     string `json:"b2cUserId"`     // 所属b2c_user_id
	Status        int    `json:"status"`        // pid状态，1表示可用，2表示禁用，3表示未在联盟创建
	HasValidOrder bool   `json:"hasValidOrder"` // 是否有有效订单
	MediaId       int64  `json:"mediaId"`       // 所属的媒体id
}
