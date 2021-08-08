package suningnetalliancecustompromotionurlquery

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.custompromotionurl.query 获取自定义推广链接接口
type Response struct {
	response2.TopResponse
	SnResponseContent SnResponseContent `json:"sn_responseContent"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = "-1"
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	} else {
		if t.SnResponseContent.SnError != nil {
			t.ErrorResponse.ErrorCode = t.SnResponseContent.SnError.ErrorCode
			t.ErrorResponse.ErrorMsg = t.SnResponseContent.SnError.ErrorMsg
		}
	}
}

type SnResponseContent struct {
	SnBody  SnBody   `json:"sn_body"`
	SnError *SnError `json:"sn_error"`
}
type SnError struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type SnBody struct {
	QueryCustompromotionurl QueryCustompromotionurl `json:"queryCustompromotionurl"`
}

type QueryCustompromotionurl struct {
	ExtendURL string `json:"extendUrl"`
	ShortURL  string `json:"shortUrl"`
}
