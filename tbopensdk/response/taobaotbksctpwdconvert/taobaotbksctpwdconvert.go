package taobaotbksctpwdconvert

import (
	"encoding/json"

	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.tpwd.convert( 淘宝客-服务商-淘口令解析&转链 )
type Response struct {
	response.TopResponse
	TbkScTpwdConvertResponse TbkScTpwdConvertResponse `json:"tbk_sc_tpwd_convert_response"`
}

// 解析输出结果
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

type TbkScTpwdConvertResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	NumIid     string `json:"num_iid"`
	ClickUrl   string `json:"click_url"`
	SellerId   string `json:"seller_id"`
	OriginUrl  string `json:"origin_url"`
	OriginPid  string `json:"origin_pid"`
	BizSceneID string `json:"biz_scene_id"`
}
