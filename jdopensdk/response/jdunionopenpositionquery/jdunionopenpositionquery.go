package jdunionopenpositionquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.position.query 查询推广位【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_position_query_responce"`
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
				t.ErrorResponse.RequestID = t.Responce.QueryResult.RequestID
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
	Code      int64  `json:"code"`
	Data      Data   `json:"data"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type Data struct {
	PageNo   int64    `json:"pageNo"`
	PageSize int64    `json:"pageSize"`
	Result   []Result `json:"result"`
	Total    int64    `json:"total"`
}

type Result struct {
	ID        int64  `json:"id"`
	PID       string `json:"pid"`
	SiteID    int64  `json:"siteId"`
	SpaceName string `json:"spaceName"`
	Type      int64  `json:"type"`
}
