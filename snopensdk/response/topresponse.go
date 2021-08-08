package response

//错误等通用信息
type TopResponse struct {
	ErrorResponse ErrorInfo `json:"error_response"` //错误信息
	Body          string    `json:"body"`           //原始信息
}

func (t *TopResponse) IsError() bool {
	return t.ErrorResponse.ErrorCode != "" || t.ErrorResponse.ErrorMsg != ""
}

type ErrorInfo struct {
	ErrorCode string `json:"error_code"` //主错误
	ErrorMsg  string `json:"error_msg"`  //主错误提示
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}
