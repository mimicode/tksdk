package alibabacpsparseurlortoken

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 解析推广物料信息响应
type Response struct {
	response.TopResponse
	Result ParseResult `json:"result"`
}

// ParseResult 解析推广物料信息响应结果
type ParseResult struct {
	Data        CpsMaterialVO `json:"data"`        // 数据
	ErrorCode   string        `json:"errorCode"`   // 错误码，示例值: 1
	ErrorInfo   string        `json:"errorInfo"`   // 错误信息，示例值: 系统错误
	Success     bool          `json:"success"`     // 是否成功，示例值: true
}

// CpsMaterialVO CPS推广物料信息
type CpsMaterialVO struct {
	MaterialId    string `json:"materialId"`    // 物料ID，示例值: 1
	ParseType     string `json:"parseType"`     // 解析类型，示例值: SHORT_URL，LONG_URL，ZHI_TOKEN
	PromotionType string `json:"promotionType"` // 推广类型，示例值: OFFER，ACTIVITY，SHOP
}

// WrapResult 包装结果
func (r *Response) WrapResult(result string) {
	// 保存原始结果到Body
	r.Body = result
	// 尝试解析具体的响应结构
	if err := json.Unmarshal([]byte(result), r); err != nil {
		// 如果解析失败，设置错误信息
		r.ErrorResponse.Code = 500
		r.ErrorResponse.Msg = "Failed to parse response: " + err.Error()
		return
	}
	// 检查业务层面的错误
	if r.Result.ErrorCode != "" && r.Result.ErrorCode != "0" {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = fmt.Sprintf("business error - errorCode: %s, errorInfo: %s", r.Result.ErrorCode, r.Result.ErrorInfo)
	}
}