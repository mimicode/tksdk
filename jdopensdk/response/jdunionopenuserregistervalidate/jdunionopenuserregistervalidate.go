package jdunionopenuserregistervalidate

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.user.register.validate 实时RTA接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_user_register_validate_responce"`
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
		//解析validateResult
		if t.Responce.ValidateResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty validateResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.ValidateResultStr), &t.Responce.ValidateResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.ValidateResult.Code
				t.ErrorResponse.Message = t.Responce.ValidateResult.Message
			}
		}
	}
	t.Responce.ValidateResultStr = ""
}

// Responce 响应结果
type Responce struct {
	ValidateResultStr string `json:"validateResult"`
	ValidateResult    ValidateResult
}

// ValidateResult 返回结果
type ValidateResult struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// Data 数据明细
type Data struct {
	UserResp UserResp `json:"userResp"`
}

// UserResp 用户状态结果
type UserResp struct {
	JdUser int64 `json:"jdUser"` //1：实时不满足要求，2：实时满足要求
}
