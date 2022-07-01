package jdunionopenpositioncreate

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.position.create 创建推广位【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_position_create_responce"`
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
		//解析createResult
		if t.Responce.CreateResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty createResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.CreateResultStr), &t.Responce.CreateResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.CreateResult.Code
				t.ErrorResponse.Message = t.Responce.CreateResult.Message
				t.ErrorResponse.RequestID = t.Responce.CreateResult.RequestID
			}
		}
	}
	t.Responce.CreateResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code            string `json:"code"`
	CreateResultStr string `json:"createResult"`
	CreateResult    CreateResult
}

// CreateResult 具体内容
type CreateResult struct {
	Code      int64  `json:"code"`
	Data      Data   `json:"data"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type Data struct {
	ResultList map[string]int64 `json:"resultList"`
	SiteID     int64            `json:"siteId"`
	Type       int64            `json:"type"`
	UnionID    int64            `json:"unionId"`
}
