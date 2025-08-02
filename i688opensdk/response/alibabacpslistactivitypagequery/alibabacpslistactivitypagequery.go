package alibabacpslistactivitypagequery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取联盟活动列表响应
type Response struct {
	response.TopResponse
	AlibabaCpsListActivityPageQueryResponse AlibabaCpsListActivityPageQueryResponse `json:"result"`
	TotalRow                                int                                     `json:"totalRow"` // 总记录数
}

// AlibabaCpsListActivityPageQueryResponse 响应结构
type AlibabaCpsListActivityPageQueryResponse []OpenUnionActivityDTO

// OpenUnionActivityDTO 联盟活动信息
type OpenUnionActivityDTO struct {
	ID         int64   `json:"id"`         // 活动id
	Type       int     `json:"type"`       // 活动类型，0普通类型
	Title      string  `json:"title"`      // 活动标题
	BannerUrl  string  `json:"bannerUrl"`  // 活动banner链接
	LinkUrl    string  `json:"linkUrl"`    // 活动落地页
	StartDate  string  `json:"startDate"`  // 开始日期
	EndDate    string  `json:"endDate"`    // 结束日期
	Ratio      float64 `json:"ratio"`      // 平均佣金比例
	ProductCnt int     `json:"productCnt"` // 活动实体数量
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