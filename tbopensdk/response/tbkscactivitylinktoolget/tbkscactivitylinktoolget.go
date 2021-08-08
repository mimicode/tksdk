package tbkscactivitylinktoolget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.activitylink.toolget( 淘宝联盟官方活动推广API-工具 )
type Response struct {
	response.TopResponse
	TbkScActivitylinkToolgetResult Result `json:"tbk_sc_activitylink_toolget_response"`
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

type Result struct {
	BizErrorCode int64  `json:"biz_error_code"`
	Data         string `json:"data"`
	ResultCode   int64  `json:"result_code"`
	RequestID    string `json:"request_id"`
}
