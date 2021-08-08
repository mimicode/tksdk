package response

//错误等通用信息
type TopResponse struct {
	ErrorResponse ErrorInfo `json:"error_response"` //错误信息
	Body          string    `json:"body"`           //原始信息
}

func (t *TopResponse) IsError() bool {
	return t.ErrorResponse.ErrorCode != 0 || t.ErrorResponse.SubCode != ""
}

type ErrorInfo struct {
	ErrorCode int64  `json:"error_code"` //主错误
	ErrorMsg  string `json:"error_msg"`  //主错误提示
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}

//type Welcome struct {
//	ErrorResponse ErrorResponse `json:"error_response"`
//}
//
//type ErrorResponse struct {
//	ErrorMsg  string `json:"error_msg"`
//	SubMsg    string `json:"sub_msg"`
//	SubCode   string `json:"sub_code"`
//	ErrorCode int64  `json:"error_code"`
//	RequestID string `json:"request_id"`
//}
//{"error_response":{"error_msg":"参数错误:p_id_name_list","sub_msg":"参数错误:p_id_name_list","sub_code":"10000","error_code":10000,"request_id":"16278930378660952"}}
