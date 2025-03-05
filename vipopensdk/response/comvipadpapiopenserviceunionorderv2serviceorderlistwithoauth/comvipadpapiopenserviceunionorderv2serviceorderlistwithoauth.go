package comvipadpapiopenserviceunionorderv2serviceorderlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionOrderV2Service
// orderListWithOauth 2.0.0 获取订单列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=orderListWithOauth
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

// Success 成功响应结构
type Success struct {
	OrderInfoList []OrderInfo `json:"orderInfoList"` // 业绩查询响应结果
	Total         int         `json:"total"`         // 业绩总条数
	Page          int         `json:"page"`          // 当前页码
	PageSize      int         `json:"pageSize"`      // 页面大小
}

// OrderInfo 订单信息
type OrderInfo struct {
	OrderSn                          string            `json:"orderSn"`                          // 订单号
	Status                           int16             `json:"status"`                           // 订单状态:0-不合格，1-待定，2-已完结
	NewCustomer                      int16             `json:"newCustomer"`                      // 新老客：0-待定，1-新客，2-其他 5-沉睡用户
	ChannelTag                       string            `json:"channelTag"`                       // channelTag=chanTag=pid，推广位标识
	OrderTime                        int64             `json:"orderTime"`                        // 下单时间 时间戳 单位毫秒
	SignTime                         int64             `json:"signTime"`                         // 签收时间 时间戳 单位毫秒
	SettledTime                      int64             `json:"settledTime"`                      // 结算时间 时间戳 单位毫秒
	DetailList                       []OrderDetailInfo `json:"detailList"`                       // 商品明细
	LastUpdateTime                   int64             `json:"lastUpdateTime"`                   // 订单上次更新时间 时间戳 单位毫秒
	Settled                          int16             `json:"settled"`                          // 订单结算状态 0-未结算,1-已结算
	SelfBuy                          int               `json:"selfBuy"`                          // 是否自推自买 0-否，1-是
	OrderSubStatusName               string            `json:"orderSubStatusName"`               // 订单子状态
	Commission                       string            `json:"commission"`                       // 商品总佣金:单位元
	AfterSaleChangeCommission        string            `json:"afterSaleChangeCommission"`        // 售后订单佣金变动
	AfterSaleChangeGoodsCount        int               `json:"afterSaleChangeGoodsCount"`        // 售后订单总商品数量变动
	CommissionEnterTime              int64             `json:"commissionEnterTime"`              // 入账时间，时间戳，单位毫秒
	OrderSource                      string            `json:"orderSource"`                      // 订单来源
	Pid                              string            `json:"pid"`                              // 推广PID
	IsPrepay                         int               `json:"isPrepay"`                         // 是否预付订单:0-否，1-是
	StatParam                        string            `json:"statParam"`                        // 自定义统计参数
	IsSplit                          int               `json:"isSplit"`                          // 订单拆单标识: 0-否，1-是
	ParentSn                         string            `json:"parentSn"`                         // 订单母单号
	OrderTrackReason                 int               `json:"orderTrackReason"`                 // 订单归因方式
	AppKey                           string            `json:"appKey"`                           // 开发者调用的appKey
	TotalCost                        string            `json:"totalCost"`                        // 订单支付金额:单位元
	OpenId                           string            `json:"openId"`                           // 渠道用户唯一标识
	AdCode                           string            `json:"adCode"`                           // 标识获取推广物料的来源
	FdsBlackOrder                    int               `json:"fdsBlackOrder"`                    // 是否异常订单,1：是,0：否
	OrderQuality                     string            `json:"orderQuality"`                     // 订单质量：待定-高-中-低
	BaseCommissionSettleRate         string            `json:"baseCommissionSettleRate"`         // 基础佣金结算比例
	DeductionCommission              string            `json:"DeductionCommission"`              // 扣减佣金
	EstimateCommissionAfterDeduction string            `json:"estimateCommissionAfterDeduction"` // 扣减后预估佣金
	OrderType                        int16             `json:"orderType"`                        // 订单类型：0-待定，1-新客订单，2-其他订单 5-沉睡用户订单
}

// OrderDetailInfo 订单商品明细
type OrderDetailInfo struct {
	GoodsId                          string                `json:"goodsId"`                          // 商品id
	GoodsName                        string                `json:"goodsName"`                        // 商品名称
	GoodsThumb                       string                `json:"goodsThumb"`                       // 商品缩略图
	GoodsCount                       int                   `json:"goodsCount"`                       // 商品数量
	CommissionTotalCost              string                `json:"commissionTotalCost"`              // 商品计佣金额
	CommissionRate                   string                `json:"commissionRate"`                   // 商品佣金比例
	Commission                       string                `json:"commission"`                       // 商品佣金金额
	CommCode                         string                `json:"commCode"`                         // 佣金编码
	CommName                         string                `json:"commName"`                         // 佣金方案名称
	OrderSource                      string                `json:"orderSource"`                      // 订单来源
	AfterSaleInfo                    []AfterSaleDetailInfo `json:"afterSaleInfo"`                    // 商品售后信息
	SizeId                           string                `json:"sizeId"`                           // 商品尺码
	Status                           int16                 `json:"status"`                           // 商品状态
	BrandStoreSn                     string                `json:"brandStoreSn"`                     // 品牌编号
	BrandStoreName                   string                `json:"brandStoreName"`                   // 品牌名称
	SpuId                            string                `json:"spuId"`                            // 商品spuId
	GoodsFinalPrice                  string                `json:"goodsFinalPrice"`                  // 商品成交价
	IsSubsidyTaskOrder               bool                  `json:"isSubsidyTaskOrder"`               // 商品是否命中超级补贴
	SubsidyTaskOrderStatus           int                   `json:"subsidyTaskOrderStatus"`           // 补贴结算的状态
	SubsidyTaskGoodsAward            string                `json:"subsidyTaskGoodsAward"`            // 补贴商品金额
	Cat1Code                         string                `json:"cat1Code"`                         // 商品一级品类code
	Cat1Name                         string                `json:"cat1Name"`                         // 商品一级品类名称
	Cat2Code                         string                `json:"cat2Code"`                         // 商品二级品类code
	Cat2Name                         string                `json:"cat2Name"`                         // 商品二级品类名称
	Cat3Code                         string                `json:"cat3Code"`                         // 商品三级品类code
	Cat3Name                         string                `json:"cat3Name"`                         // 商品三级品类名称
	BaseCommissionSettleRate         string                `json:"baseCommissionSettleRate"`         // 基础佣金结算比例
	DeductionCommission              string                `json:"DeductionCommission"`              // 扣减佣金
	EstimateCommissionAfterDeduction string                `json:"estimateCommissionAfterDeduction"` // 扣减后预估佣金
	IsCampaignCommission             int                   `json:"isCampaignCommission"`             // 是否命中商单补贴
	CampaignCommissionAmt            string                `json:"campaignCommissionAmt"`            // 商单补贴金额
	CampaignCommissionStatus         int16                 `json:"campaignCommissionStatus"`         // 商单补贴状态
}

// AfterSaleDetailInfo 售后详情信息
type AfterSaleDetailInfo struct {
	AfterSaleChangedCommission string `json:"afterSaleChangedCommission"` // 商品佣金售后变动
	AfterSaleChangedGoodsCount int    `json:"afterSaleChangedGoodsCount"` // 商品数量售后变动
	AfterSaleSn                string `json:"afterSaleSn"`                // 商品售后单号
	AfterSaleStatus            int    `json:"afterSaleStatus"`            // 商品售后状态
	AfterSaleType              int    `json:"afterSaleType"`              // 售后类型
	AfterSaleFinishTime        int64  `json:"afterSaleFinishTime"`        // 售后完成时间
	NewOrderSn                 string `json:"newOrderSn"`                 // 售后换货新单号
}
