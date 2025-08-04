package alibabacpsgenerateintroduceurlbyparam

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 批量生成推广链接服务(部分条件)响应
type Response struct {
	response.TopResponse
	// 结果对象，注意返回列表不保证所有offerId都有，只返回有效的商品，请根据返回objectId匹配
	Result ResultVO `json:"result"`
}

// ResultVO 结果对象
type ResultVO struct {
	// 错误码
	ErrorCode string `json:"errorCode"`
	// 错误信息
	ErrorMsg string `json:"errorMsg"`
	// 结果
	Result []ExtendIntroduceUrlVO `json:"result"`
	// 状态
	Success bool `json:"success"`
}

// ExtendIntroduceUrlVO 推广链接结果
type ExtendIntroduceUrlVO struct {
	// 推广链接
	ClickUrlVO ClickUrlVO `json:"clickUrlVO"`
}

// ClickUrlVO 推广链接
type ClickUrlVO struct {
	// 长链接
	LongClickUrl string `json:"longClickUrl"`
	// 对象id，商品生成长链接时为商品id
	ObjectId int64 `json:"objectId"`
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
	// 对于存在 errorMsg errorCode 的响应内容 需要判断这两个值是否为空
	if r.Result.ErrorMsg != "" || r.Result.ErrorCode != "" {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = fmt.Sprintf("errorMsg: %s, errorCode: %s", r.Result.ErrorMsg, r.Result.ErrorCode)
	}
}