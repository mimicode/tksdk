package alibabacpstradebilllist

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 联盟账单列表查询响应
type Response struct {
	response.TopResponse
	Result   []OpenTradeBillDTO `json:"result"`
	TotalRow int                `json:"totalRow"`
}

// OpenTradeBillDTO CPS交易账单数据传输对象
type OpenTradeBillDTO struct {
	// 主订单号
	BizId int64 `json:"bizId"`
	// 子订单号
	BizSubId int64 `json:"bizSubId"`
	// 商品ID
	FeedId int64 `json:"feedId"`
	// 商品名称
	Name string `json:"name"`
	// 成交金额
	TradeAmount float64 `json:"tradeAmount"`
	// 成交商品数
	TradeNumber float64 `json:"tradeNumber"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 确定收货时间
	ConfirmTime string `json:"confirmTime"`
	// 订单状态
	OrderState int `json:"orderState"`
	// 结算时间
	SettleTime string `json:"settleTime"`
	// 结算状态
	SettleState int `json:"settleState"`
	// 佣金比例
	Ratio float64 `json:"ratio"`
	// 佣金金额
	Commission float64 `json:"commission"`
	// 子推客佣金
	SubCommission float64 `json:"subCommission"`
	// 推广类型 0：商品，1：店铺，2：活动，9：三方
	Type int `json:"type"`
	// 媒体
	MediaId int64 `json:"mediaId"`
	// 推广位
	MediaZoneId int64 `json:"mediaZoneId"`
	// 推客id
	TkId int64 `json:"tkId"`
	// 子推客id
	SubTkId int64 `json:"subTkId"`
	// 主推客的分佣比例
	TkRatio float64 `json:"tkRatio"`
	// 子推客分佣比例
	SubTkRatio float64 `json:"subTkRatio"`
	// 平台分成比例
	PlatformRatio float64 `json:"platformRatio"`
	// 买家id
	BuyerId int64 `json:"buyerId"`
	// 卖家id
	SellerId int64 `json:"sellerId"`
	// 交易类型，对应交易线的类型
	TradeType string `json:"tradeType"`
	// 补贴金额比例
	SubsidizationRatio float64 `json:"subsidizationRatio"`
	// 推客补贴金额
	SubsidizationAmount float64 `json:"subsidizationAmount"`
	// 子推客补贴金额
	SubSubsidizationAmount float64 `json:"subSubsidizationAmount"`
	// ext字段
	Ext string `json:"ext"`
	// 维权状态
	RightsState int `json:"rightsState"`
	// 平台 1:1688
	Platform int `json:"platform"`
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
