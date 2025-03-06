package comvipadpapiopenserviceunionurlv2servicegetchannelurlbytypewithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionUrlV2Service
// getChannelUrlByTypeWithOauth 2.0.0 渠道根据落地页类型生链接口
// 该接口需要申请渠道工具商权限包权限和oauth授权
// https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getChannelUrlByTypeWithOauth
type Response struct {
	response.TopResponse
	Success GetChannelUrlByTypeResponse `json:"result"` // 响应结果
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

// GetChannelUrlByTypeResponse 响应结构
type GetChannelUrlByTypeResponse struct {
	Code int                     `json:"code"` // 返回码：0-成功，其它失败
	Msg  string                  `json:"msg"`  // 返回信息
	Data GetChannelUrlByTypeData `json:"data"` // 返回数据体
}

// GetChannelUrlByTypeData 返回数据结构
type GetChannelUrlByTypeData struct {
	ShortUrl        string `json:"shortUrl"`        // 页面短链
	ShortLink       string `json:"shortLink"`       // 页面小程序短链
	Remark          string `json:"remark"`          // 提示信息
	VipWxUrl        string `json:"vipWxUrl"`        // 唯品会微信小程序链接，用于实现未安装唯品会app的用户跳转唯品会微信小程序，和从小程序跳转唯品会微信小程序
	VipZfbUrl       string `json:"vipZfbUrl"`       // 唯品会支付宝小程序链接，用于实现未安装唯品会app的用户跳转唯品会支付宝小程序，和从小程序跳转唯品会支付宝小程序
	VipZfbSchemeUrl string `json:"vipZfbSchemeUrl"` // 唯品会支付宝小程序scheme链接，用于实现未安装唯品会app的用户跳转唯品会支付宝小程序，和从小程序跳转唯品会支付宝小程序
	VipZfbHttpsUrl  string `json:"vipZfbHttpsUrl"`  // 唯品会支付宝小程序https链接，用于实现未安装唯品会app的用户跳转唯品会支付宝小程序，和从小程序跳转唯品会支付宝小程序
	DeeplinkUrl     string `json:"deeplinkUrl"`     // CPS Deeplink链接
}
