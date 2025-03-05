package comvipadpapiopenserviceunionurlservicegenbyvipurlwithoauth

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUrlService 根据唯品会链接生成联盟链接-需要oauth授权
type Response struct {
	response2.TopResponse
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
	URLInfoList []URLInfoList `json:"urlInfoList"`
}

type URLInfoList struct {
	NoEvokeURL     string `json:"noEvokeUrl"`
	VipQuickAppURL string `json:"vipQuickAppUrl"`
	DeeplinkURL    string `json:"deeplinkUrl"`
	LongURL        string `json:"longUrl"`
	Source         string `json:"source"`
	ULURL          string `json:"ulUrl"`
	URL            string `json:"url"`
	TraFrom        string `json:"traFrom"`
	NoEvokeLongURL string `json:"noEvokeLongUrl"`
	VipWxUrl       string `json:"vipWxUrl"`  //小程序 仅根据商品id获取时 返回
	VipWxCode      string `json:"vipWxCode"` //小程序码 需要高级权限
}
