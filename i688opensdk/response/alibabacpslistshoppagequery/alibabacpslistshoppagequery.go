package alibabacpslistshoppagequery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取联盟商家列表响应
type Response struct {
	response.TopResponse
	Result   []OpenUnionShopDTO `json:"result"`
	TotalRow int                `json:"totalRow"`
}

// OpenUnionShopDTO 联盟商家信息
type OpenUnionShopDTO struct {
	SellerId    int64   `json:"sellerId"`    // 卖家id
	LoginId     string  `json:"loginId"`     // 商家loginId
	CompanyName string  `json:"companyName"` // 店铺名称
	TradeGrade  float64 `json:"tradeGrade"`  // 交易勋章
	Ratio       float64 `json:"ratio"`       // 平均佣金比例
	ProductCnt  int     `json:"productCnt"`  // 商品数量
	TkCnt       int     `json:"tkCnt"`       // 30天推广量
	LinkUrl     string  `json:"linkUrl"`     // 店铺首页
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