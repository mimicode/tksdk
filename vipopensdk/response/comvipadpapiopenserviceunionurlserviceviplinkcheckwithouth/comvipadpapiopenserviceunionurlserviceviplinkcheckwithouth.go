package comvipadpapiopenserviceunionurlserviceviplinkcheckwithouth

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/vipopensdk/response"
)

//com.vip.adp.api.open.service.UnionUrlService 检测一段文本中是否有唯品会链接-需要oauth授权
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
	SuccessMap map[string]SuccessMap `json:"successMap"`
	FailMap    map[string]SuccessMap `json:"failMap"`
}

type SuccessMap struct {
	LinkType string `json:"linkType"`
	LandUrl  string `json:"landUrl"`
	GoodsId  string `json:"goodsId"`
	BrandId  string `json:"brandId"`
}

/*
linkType	VipLinkTypeEnum	否			链接类型
landUrl	String	否			落地页url
goodsId	String	否			商品id
brandId	String	否			专场id

*/
