package response

// TopResponse 通用响应结构
type TopResponse struct {
	Code    int64  `json:"code"`    // 错误码，0表示成功
	Desc    string `json:"desc"`   // 错误描述
	RawData string `json:"-"`      // 原始响应数据
}

// IsError 判断是否有错误
func (t *TopResponse) IsError() bool {
	return t.Code != 0
}
