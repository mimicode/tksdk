package comvipadpapiopenserviceunionurlv2servicegetwxcodewithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionUrlV2Service
// getWxCodeWithOauth 2.0.0 生成微信小程序码-需要oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getWxCodeWithOauth
type Response struct {
	response.TopResponse
	Success WxCodeGenResponse `json:"success"`
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

// WxCodeGenResponse 响应结构
type WxCodeGenResponse struct {
	WxCodeInfo WxCodeInfo `json:"wxCodeInfo"` // 微信小程序码信息数据
}

// WxCodeInfo 微信小程序码信息结构
type WxCodeInfo struct {
	VipWxUrl        string `json:"vipWxUrl"`        // 唯品会微信小程序链接
	VipWxCode       string `json:"vipWxCode"`       // 唯品会微信小程序码
	ShortLink       string `json:"shortLink"`       // 微信小程序短链
	Remark          string `json:"remark"`          // 提示信息
	VipZfbUrl       string `json:"vipZfbUrl"`       // 唯品会支付宝小程序链接
	VipZfbSchemeUrl string `json:"vipZfbSchemeUrl"` // 唯品会支付宝小程序scheme链接
	VipZfbHttpsUrl  string `json:"vipZfbHttpsUrl"`  // 唯品会支付宝小程序https链接
}
