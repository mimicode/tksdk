package response

// TopResponse 错误等通用信息
type TopResponse struct {
	ErrorInfo
	Body string `json:"body"` //原始信息
}

func (t *TopResponse) IsError() bool {
	return t.Code != 0
}

type ErrorInfo struct {
	Code int64  `json:"code"`
	Desc string `json:"desc,omitempty"` // 当code不为0时 有值
}
