package alibabacpslistmediainfo

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取媒体推广位信息响应
type Response struct {
	response.TopResponse
	AlibabaCpsListMediaInfoResponse AlibabaCpsListMediaInfoResponse `json:"alibaba_cps_listMediaInfo_response"`
}

// AlibabaCpsListMediaInfoResponse 响应结构
type AlibabaCpsListMediaInfoResponse struct {
	Result []OpenUnionMediaDTO `json:"result"`
	RequestID string `json:"requestId"`
}

// OpenUnionMediaDTO 媒体推广位信息
type OpenUnionMediaDTO struct {
	MediaId     int64  `json:"mediaId"`     // 媒体ID
	MediaName   string `json:"mediaName"`   // 媒体名称
	MediaType   string `json:"mediaType"`   // 媒体类型
	MediaStatus string `json:"mediaStatus"` // 媒体状态
	CreateTime  string `json:"createTime"`  // 创建时间
	UpdateTime  string `json:"updateTime"`  // 更新时间
	Pid         string `json:"pid"`         // 推广位ID
	SiteId      int64  `json:"siteId"`      // 站点ID
	AdzoneId    int64  `json:"adzoneId"`    // 广告位ID
}

func (r *Response) WrapResult(body []byte) error {
	return json.Unmarshal(body, r)
}