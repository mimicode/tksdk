package comvipadpapiopenserviceunionurlv2servicegetchannelurlbytypewithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUrlV2Service 渠道根据落地页类型生链接口,需要申请渠道工具商权限包权限
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
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

type Success struct {
	Code int    `json:"code"` // 返回码：0-成功，其它失败
	Msg  string `json:"msg"`  // 返回信息
	Data Data   `json:"data"` // 返回数据体
}

type Data struct {
	ShortUrl        string `json:"shortUrl"`        // 页面短链
	ShortLink       string `json:"shortLink"`       // 页面小程序短链
	Remark          string `json:"remark"`          // 提示信息
	VipWxUrl        string `json:"vipWxUrl"`        // 唯品会微信小程序链接
	VipZfbUrl       string `json:"vipZfbUrl"`       // 唯品会支付宝小程序链接
	VipZfbSchemeUrl string `json:"vipZfbSchemeUrl"` // 唯品会支付宝小程序scheme链接
	VipZfbHttpsUrl  string `json:"vipZfbHttpsUrl"`  // 唯品会支付宝小程序https链接
	DeeplinkUrl     string `json:"deeplinkUrl"`     // CPS Deeplink链接
}
