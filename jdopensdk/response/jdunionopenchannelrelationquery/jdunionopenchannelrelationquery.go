package jdunionopenchannelrelationquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.channel.relation.query 渠道关系查询接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_channel_relation_query_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析queryResult
		if t.Responce.QueryResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty queryResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.QueryResultStr), &t.Responce.QueryResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.QueryResult.Code
				t.ErrorResponse.Message = t.Responce.QueryResult.Message
			}
		}
	}
	t.Responce.QueryResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"queryResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code       int64                      `json:"code"`
	Data       []ChannelRelationQueryResp `json:"data"`
	Message    string                     `json:"message"`
	TotalCount int64                      `json:"totalCount"` //总数
}

// ChannelRelationQueryResp 渠道关系明细
type ChannelRelationQueryResp struct {
	ChannelID   int64  `json:"channelId"`   //渠道关系ID
	CreateTime  string `json:"createTime"`  //备案时间，格式：yyyy-MM-dd HH:mm:ss
	Note        string `json:"note"`        //渠道备注名(渠道侧备注)：仅支持传入中文、字母、数字、下划线或中划线，最多15个字符
	ChannelNote string `json:"channelNote"` //合作方备注名(合作方给渠道的备注且唯一)：仅支持传入中文、字母、数字、下划线或中划线，最多15个字符
	ID          int64  `json:"id"`          //索引Id，可传入参数startIndex来查询下一页
}
