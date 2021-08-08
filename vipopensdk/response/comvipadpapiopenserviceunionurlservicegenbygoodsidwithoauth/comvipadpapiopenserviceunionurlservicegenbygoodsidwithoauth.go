package comvipadpapiopenserviceunionurlservicegenbygoodsidwithoauth

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/vipopensdk/response"
)

//com.vip.adp.api.open.service.UnionUrlService 根据商品id生成联盟链接-需要oauth授权
type Response struct {
	response2.TopResponse
	Result Result `json:"result"`
}

//解析输出结果
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
	VipWxURL       string `json:"vipWxUrl"`
	DeeplinkURL    string `json:"deeplinkUrl"`
	LongURL        string `json:"longUrl"`
	Source         string `json:"source"`
	ULURL          string `json:"ulUrl"`
	URL            string `json:"url"`
	TraFrom        string `json:"traFrom"`
	NoEvokeLongURL string `json:"noEvokeLongUrl"`
}
