package comvipadpapiopenserviceunionurlservicegetwxcodewithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUrlService getWxCodeWithOauth 生成微信小程序码,需要Oauth授权
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Result struct {
	WxCodeInfo WxCodeInfo `json:"wxCodeInfo"` // 微信小程序码信息数据
}

type WxCodeInfo struct {
	VipWxUrl        string `json:"vipWxUrl"`        // 唯品会微信小程序链接
	VipWxCode       string `json:"vipWxCode"`       // 唯品会微信小程序码
	ShortLink       string `json:"shortLink"`       // 微信小程序短链
	Remark          string `json:"remark"`          // 提示信息
	VipZfbUrl       string `json:"vipZfbUrl"`       // 唯品会支付宝小程序链接
	VipZfbSchemeUrl string `json:"vipZfbSchemeUrl"` // 唯品会支付宝小程序scheme链接
	VipZfbHttpsUrl  string `json:"vipZfbHttpsUrl"`  // 唯品会支付宝小程序https链接
}
