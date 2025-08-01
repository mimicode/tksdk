package alibabacpslistmediainfo

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取媒体推广位信息响应
type Response struct {
	response.TopResponse
	AlibabaCpsListMediaInfoResponse AlibabaCpsListMediaInfoResponse `json:"result"`
}

// AlibabaCpsListMediaInfoResponse 响应结构
type AlibabaCpsListMediaInfoResponse []OpenUnionMediaDTO

// OpenUnionMediaDTO 媒体推广位信息
type OpenUnionMediaDTO struct {
	TkID           int64  `json:"tkId"`           // 推客id
	MediaID        int64  `json:"mediaId"`        // 媒体id
	MediaTitle     string `json:"mediaTitle"`     // 媒体名称
	MediaType      int    `json:"mediaType"`      // 媒体类型
	AuditState     int    `json:"auditState"`     // 审核状态
	AuditReason    string `json:"auditReason"`    // 审核原因
	MediaZoneID    int64  `json:"mediaZoneId"`    // 媒体推广位id
	MediaZoneTitle string `json:"mediaZoneTitle"` // 推广位名称
}

func (r *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), r)
	//保存原始信息
	r.Body = result
	//解析错误
	if unmarshal != nil {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = unmarshal.Error()
	}
}
