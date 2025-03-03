package comvipadpapiopenserviceunionpidv2servicequerypidwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionPidV2Service queryPidWithOauth 查询推广位PID-需要oauth授权
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
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 状态信息
	Data    Data   `json:"data"`    // 响应数据
}

type Data struct {
	Page     int       `json:"page"`     // 当前页码
	PageSize int       `json:"pageSize"` // 每页数量
	Total    int       `json:"total"`    // 总记录数
	PidList  []PidInfo `json:"pidList"`  // PID列表
}

type PidInfo struct {
	PidId       string `json:"pidId"`       // PID ID
	PidName     string `json:"pidName"`     // PID名称
	CreateTime  string `json:"createTime"`  // 创建时间
	UpdateTime  string `json:"updateTime"`  // 更新时间
	Status      int    `json:"status"`      // PID状态
	PidTypeId   int    `json:"pidTypeId"`   // PID类型ID
	PidTypeName string `json:"pidTypeName"` // PID类型名称
}
