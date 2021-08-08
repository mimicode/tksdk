package response

// TopResponse 错误等通用信息
type TopResponse struct {
	ErrorResponse ErrorInfo `json:"error_response"` //错误信息
	Body          string    `json:"body"`           //原始信息
}

func (t *TopResponse) IsError() bool {
	return t.ErrorResponse.Code != 200
}

type ErrorInfo struct {
	Code      int64    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}
