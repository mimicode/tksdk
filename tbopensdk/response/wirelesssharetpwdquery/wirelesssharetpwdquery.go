package wirelesssharetpwdquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.wireless.share.tpwd.query( 查询解析淘口令 )
type Response struct {
	response.TopResponse
	WirelessShareTpwdQueryResult Result `json:"wireless_share_tpwd_query_response"`
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
	Content     string `json:"content"`
	NativeURL   string `json:"native_url"`
	PicURL      string `json:"pic_url"`
	Suc         bool   `json:"suc"`
	ThumbPicURL string `json:"thumb_pic_url"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	RequestID   string `json:"request_id"`
}
