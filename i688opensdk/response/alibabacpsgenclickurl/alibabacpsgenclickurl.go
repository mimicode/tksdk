package alibabacpsgenclickurl

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 批量生成联盟推广url点击信息响应
type Response struct {
	response.TopResponse
	// 响应结果数组
	Result []OpenClickUrlDTO `json:"result"`
}

// OpenClickUrlDTO 推广链接信息结构体
type OpenClickUrlDTO struct {
	// 推广实体id
	ObjectId int64 `json:"objectId"`
	// 长链接(图搜结果页) 示例值: https://s.click.1688.com/t?e=7175776D9BFAC498E1BAA1BCF153E1BD1572C8CCC074EF8B5ECDBE0208A7735AA661347D71EB87DF39C0515439B0C52E83E002417C4617065ECDBE0208A7735A3C953691953C1042124C58689192A3A2A8A465BAA97F8D61DD4852BE0CE5F0EB668CABA5A1D53443B9B8B717A40D8FE3
	LongClickUrl string `json:"longClickUrl"`
	// 短链接 示例值: https://c.tb.cn/G1.Z4wTRDa
	ShortClickUrl string `json:"shortClickUrl"`
	// 支付宝口令 示例值: 吱口令：【芯壹测试企业购sku-sku报价】￥ipZoSS32hcav￥，复制这条信息，在【手机阿里】或【支付宝】查看。
	AlipayUrl string `json:"alipayUrl"`
	// 搜索码 示例值: ￥1688精选181￥，复制这条信息，在【手机阿里】或【支付宝】查看。
	SearchCode string `json:"searchCode"`
	// pK字段（实际返回中存在）
	PK int64 `json:"pK"`
}

// WrapResult 包装结果
func (r *Response) WrapResult(result string) {
	// 保存原始结果到Body
	r.Body = result
	// 尝试解析具体的响应结构
	if err := json.Unmarshal([]byte(result), r); err != nil {
		// 如果解析失败，设置错误信息
		r.ErrorResponse.Code = 500
		r.ErrorResponse.Msg = "Failed to parse response: " + err.Error()
		return
	}
}