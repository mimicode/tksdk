package response

// 错误等通用信息
type TopResponse struct {
	ErrorResponse ErrorInfo `json:"error_response"` //错误信息
	Body          string    `json:"body"`           //原始信息
}

func (t *TopResponse) IsError() bool {
	return t.ErrorResponse.Code != 0 || t.ErrorResponse.SubCode != ""
}

type ErrorInfo struct {
	Code      int64  `json:"code"` //主错误
	Msg       string `json:"msg"`  //主错误提示
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}
