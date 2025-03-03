package comvipadpapiopenserviceunionurlv2serviceviplinkcheckwithouth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionUrlV2Service vipLinkCheckWithOuth 进行cps链接的解析,需要申请渠道包权限进行Oauth授权
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

// Success 响应结果
type Success struct {
	Code int    `json:"code"` // 响应码
	Msg  string `json:"msg"`  // 响应信息
	Data Data   `json:"data"` // 响应数据
}

// Data 响应数据
type Data struct {
	IsVipLink bool   `json:"isVipLink"` // 是否是唯品会链接
	GoodsId   string `json:"goodsId"`   // 商品ID
	BrandId   string `json:"brandId"`   // 品牌ID
	ChannelId string `json:"channelId"` // 渠道ID
	Remark    string `json:"remark"`    // 备注信息
}
