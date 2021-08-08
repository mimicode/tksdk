package response

//错误等通用信息
type TopResponse struct {
	Body          string `json:"body"`          //原始信息
	ReturnCode    string `json:"returnCode"`    //主错误
	ReturnMessage string `json:"returnMessage"` //主错误提示
}

func (t *TopResponse) IsError() bool {
	return t.ReturnCode != "0" || t.ReturnMessage != ""
}
