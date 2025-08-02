package alibabacpsgetcpsrecommendofferlist

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取个性化推荐的cps商品(含推广链接)响应
type Response struct {
	response.TopResponse
	Result *PageResultDTO `json:"result"`
}

// PageResultDTO 分页结果
type PageResultDTO struct {
	Success   bool                      `json:"success"`
	ErrorCode string                    `json:"errorCode"`
	ErrorMsg  string                    `json:"errorMsg"`
	Result    []*CpsRecommendOfferDTO   `json:"result"`
}

// CpsRecommendOfferDTO CPS推荐商品信息
type CpsRecommendOfferDTO struct {
	OfferId      int64   `json:"offerId"`      // 商品id
	Title        string  `json:"title"`        // 标题
	SaleCount    int64   `json:"saleCount"`    // 月销量
	Price        float64 `json:"price"`        // 价格
	LongClickUrl string  `json:"longClickUrl"` // 推广长链接
	ImgUrl       string  `json:"imgUrl"`       // 首图
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
	// 对于存在 errorMsg errorCode 的响应内容 需要判断这两个值是否为空
	if r.Result != nil && (r.Result.ErrorMsg != "" || r.Result.ErrorCode != "") {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = fmt.Sprintf("errorMsg: %s, errorCode: %s", r.Result.ErrorMsg, r.Result.ErrorCode)
	}
}