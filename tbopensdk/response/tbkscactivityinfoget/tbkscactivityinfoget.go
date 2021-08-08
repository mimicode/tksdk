package tbkscactivityinfoget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.activity.info.get( 淘宝客-服务商-官方活动信息获取 新接口 )
type Response struct {
	response.TopResponse
	TbkScActivityInfoGetResponseResult ResponseResult `json:"tbk_sc_activity_info_get_response"`
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

type ResponseResult struct {
	Data      ResponseResultData `json:"data"`
	RequestID string             `json:"request_id"`
}

type ResponseResultData struct {
	WxQrcodeURL       string `json:"wx_qrcode_url"`
	ClickURL          string `json:"click_url"`
	ShortClickURL     string `json:"short_click_url"`
	TerminalType      string `json:"terminal_type"`
	MaterialOSSURL    string `json:"material_oss_url"`
	PageName          string `json:"page_name"`
	PageStartTime     string `json:"page_start_time"`
	PageEndTime       string `json:"page_end_time"`
	WxMiniprogramPath string `json:"wx_miniprogram_path"`
}
