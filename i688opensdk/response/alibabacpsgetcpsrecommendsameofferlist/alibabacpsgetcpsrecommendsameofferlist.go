package alibabacpsgetcpsrecommendsameofferlist

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取同款推荐商品响应
type Response struct {
	response.TopResponse
	AlibabaCpsGetCpsRecommendSameOfferListResponse *PageResultDTO `json:"result"`
}

type AlibabaCpsGetCpsRecommendSameOfferListResponse struct {
	Response
}

// PageResultDTO 分页结果对象
type PageResultDTO struct {
	ErrorCode string                    `json:"errorCode"` // 错误码
	ErrorMsg  string                    `json:"errorMsg"`  // 错误信息
	Result    []CpsRecommendOfferDTO    `json:"result"`    // 结果列表
	Success   bool                      `json:"success"`   // 成功状态
}

// CpsRecommendOfferDTO 推荐商品信息
type CpsRecommendOfferDTO struct {
	DetailUrl    string  `json:"detailUrl"`    // 详情页链接(非cps品使用该链接，不分佣)
	ImgUrl       string  `json:"imgUrl"`       // 首图
	LongClickUrl string  `json:"longClickUrl"` // 推广长链接(非cps品不返回，不分佣)
	OfferId      int64   `json:"offerId"`      // 商品id
	Price        float64 `json:"price"`        // 价格
	SaleCount    int64   `json:"saleCount"`    // 月销量
	Title        string  `json:"title"`        // 标题
}

func (r *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), r)
	//保存原始信息
	r.Body = result
	//解析错误
	if unmarshal != nil {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = unmarshal.Error()
	}
}