package alibabaalscunionelemepromotionofficialactivityget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.eleme.promotion.officialactivity.get( 本地联盟饿了么推广官方活动查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_eleme_promotion_officialactivity_get_response"`
}

// WrapResult 解析输出结果
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

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int    `json:"biz_error_code"`
	RequestID    string `json:"request_id"`
	Data         Data   `json:"data"`
	ResultCode   int64  `json:"result_code"`
	Message      string `json:"message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Link        Link   `json:"link"`
}

type Link struct {
	WxAppid       string `json:"wx_appid"`
	WxPath        string `json:"wx_path"`
	Picture       string `json:"picture"`
	AlipayMiniURL string `json:"alipay_mini_url"`
	H5URL         string `json:"h5_url"`
	TBQrCode      string `json:"tb_qr_code"`
	MiniQrcode    string `json:"mini_qrcode"`
	TBMiniQrcode  string `json:"tb_mini_qrcode"`
	EleSchemeURL  string `json:"ele_scheme_url"`
	H5ShortLink   string `json:"h5_short_link"`
}
