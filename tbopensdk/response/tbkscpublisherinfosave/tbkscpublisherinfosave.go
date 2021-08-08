package tbkscpublisherinfosave

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.publisher.info.save( 淘宝客渠道信息备案 - 社交 )
type Response struct {
	response.TopResponse
	TbkScPublisherInfoSaveResult ResponseClass `json:"tbk_sc_publisher_info_save_response"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type ResponseClass struct {
	Data      ResponseClassData `json:"data"`
	RequestID string            `json:"request_id"`
}

type ResponseClassData struct {
	AccountName string `json:"account_name"`
	Desc        string `json:"desc"`
	RelationID  int64  `json:"relation_id"`
}
