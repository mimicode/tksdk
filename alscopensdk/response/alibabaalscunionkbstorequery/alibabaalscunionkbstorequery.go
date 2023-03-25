package alibabaalscunionkbstorequery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.store.query( 本地联盟口碑商户列表 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_store_query_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int    `json:"biz_error_code"`
	RequestID    string `json:"request_id"`

	Data         Data   `json:"data"`
	ResultCode   int64  `json:"result_code"`
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	SessionID  string  `json:"session_id"`
	PageNumber int64   `json:"page_number"`
	PageSize   int64   `json:"page_size"`
	Total      int64   `json:"total"`
	Records    Records `json:"records"`
}

type Records struct {
	Integer []Integer `json:"integer"`
}

type Integer struct {
	StoreID    string     `json:"store_id"`
	Name       string     `json:"name"`
	Cover      string     `json:"cover"`
	Categories Categories `json:"categories"`
	Location   Location   `json:"location"`
	Distance   int64      `json:"distance"`
}

type Categories struct {
	Category []Category `json:"category"`
}

type Category struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Address   string `json:"address"`
}
