package alibabacpslistofferpagequery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 获取联盟offer列表响应
type Response struct {
	response.TopResponse
	Result   []OpenOfferDTO `json:"result"`
	TotalRow int            `json:"totalRow"`
}

// OpenOfferDTO 联盟offer信息
type OpenOfferDTO struct {
	LoginId              string   `json:"loginId"`              // login id
	CompanyName          string   `json:"companyName"`          // 公司名称
	OfferId              int64    `json:"offerId"`              // offer id
	SellerId             int64    `json:"sellerId"`             // 卖家ID
	Title                string   `json:"title"`                // offer名称
	Url                  string   `json:"url"`                  // offer链接
	ImgUrl               string   `json:"imgUrl"`               // offer主图
	Ratio                float64  `json:"ratio"`                // 佣金比例
	Type                 int      `json:"type"`                 // 推广类型
	Price                float64  `json:"price"`                // 商品价格
	Unit                 string   `json:"unit"`                 // 单位
	SaleQuantity         int      `json:"saleQuantity"`         // 销量(月)
	CategoryId           string   `json:"categoryId"`           // 类目树id
	TkCnt                int      `json:"tkCnt"`                // 月推广量
	TkCommission         float64  `json:"tkCommission"`         // 月支出佣金
	QuantityBegin        int      `json:"quantityBegin"`        // 起批量
	SlsjFlag             bool     `json:"slsjFlag"`             // 实力商家标志
	SdrzFlag             bool     `json:"sdrzFlag"`             // 深度认证标志
	JkhyFlag             bool     `json:"jkhyFlag"`             // 进口货源标志
	YjdfFlag             bool     `json:"yjdfFlag"`             // 一件代发标志
	TpServiceYear        int      `json:"tpServiceYear"`        // 诚信通年限
	YhqFlag              bool     `json:"yhqFlag"`              // 优惠券标志
	YhqDiscountFee       float64  `json:"yhqDiscountFee"`       // 优惠券面额，单位为元
	YhqRemainingCount    int      `json:"yhqRemainingCount"`    // 优惠券余量
	YhqExt               string   `json:"yhqExt"`               // 优惠券其他信息
	OldBuyerRatio        float64  `json:"oldBuyerRatio"`        // 老买家佣金比例
	ImageInfos           []string `json:"imageInfos"`           // offer主图列表
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
	
	// 检查响应中是否包含错误信息
	// 注意：1688 API的错误信息通常在具体的响应字段中，这里可以根据实际情况调整
}