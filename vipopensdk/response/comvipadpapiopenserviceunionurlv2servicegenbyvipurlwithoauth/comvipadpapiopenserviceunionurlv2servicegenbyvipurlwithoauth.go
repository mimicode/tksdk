package comvipadpapiopenserviceunionurlv2servicegenbyvipurlwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUrlV2Service genByVIPUrlWithOauth 根据唯品会链接生成联盟链接-需要申请渠道工具商权限包权限
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
	Code int    `json:"code"` // 状态码，1-成功，0-失败
	Msg  string `json:"msg"`  // 状态信息
	Data Data   `json:"data"` // 数据
}

type Data struct {
	UrlInfoList []UrlInfo `json:"urlInfoList"` // 链接数据
}

type UrlInfo struct {
	Source          string `json:"source"`          // 链接生成的数据源：如果根据商品id生成链接，该值商品id，如果根据链接生成链接，该值为唯品会链接
	Url             string `json:"url"`             // CPS短链接
	LongUrl         string `json:"longUrl"`         // CPS长连接
	UlUrl           string `json:"ulUrl"`           // CPS通用连接
	DeeplinkUrl     string `json:"deeplinkUrl"`     // CPS Deeplink链接
	TraFrom         string `json:"traFrom"`         // 小程序CPS参数：通用小程序跟单参数
	NoEvokeUrl      string `json:"noEvokeUrl"`      // CPS短链接：不唤起快应用
	NoEvokeLongUrl  string `json:"noEvokeLongUrl"`  // CPS长链接：不唤起快应用
	VipWxUrl        string `json:"vipWxUrl"`        // 唯品会微信小程序链接
	VipWxCode       string `json:"vipWxCode"`       // 唯品会小程序码：需获取小程序码高级权限
	VipQuickAppUrl  string `json:"vipQuickAppUrl"`  // 唯品会快应用链接
	VipZfbUrl       string `json:"vipZfbUrl"`       // 唯品会支付宝小程序链接
	VipZfbSchemeUrl string `json:"vipZfbSchemeUrl"` // 唯品会支付宝小程序scheme链接
	VipZfbHttpsUrl  string `json:"vipZfbHttpsUrl"`  // 唯品会支付宝小程序https链接
	OnlyCommand     string `json:"onlyCommand"`     // 唯品会唯口令
}
