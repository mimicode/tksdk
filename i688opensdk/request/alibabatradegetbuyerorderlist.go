package request

import (
	"net/url"
	"strconv"
)

// AlibabaTradeGetBuyerOrderListRequest 订单列表查看(买家视角) API请求
// https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.trade%3Aalibaba.trade.getBuyerOrderList-1&aopApiCategory=category_new
type AlibabaTradeGetBuyerOrderListRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaTradeGetBuyerOrderListRequest 创建AlibabaTradeGetBuyerOrderListRequest实例
func NewAlibabaTradeGetBuyerOrderListRequest() *AlibabaTradeGetBuyerOrderListRequest {
	return &AlibabaTradeGetBuyerOrderListRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaTradeGetBuyerOrderListRequest) CheckParameters() {
	// 该接口没有必需参数
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
}

// AddParameter 添加参数
func (r *AlibabaTradeGetBuyerOrderListRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetBizTypes 设置业务类型，支持："cn"(普通订单类型), "ws"(大额批发订单类型), "yp"(普通拿样订单类型), "yf"(一分钱拿样订单类型), "fs"(倒批(限时折扣)订单类型), "cz"(加工定制订单类型), "ag"(协议采购订单类型), "hp"(伙拼订单类型), "gc"(国采订单类型), "supply"(供销订单类型), "nyg"(nyg订单类型), "factory"(淘工厂订单类型), "quick"(快订下单), "xiangpin"(享拼订单), "nest"(采购商城-鸟巢), "f2f"(当面付), "cyfw"(存样服务), "sp"(代销订单标记), "wg"(微供订单), "factorysamp"(淘工厂打样订单), "factorybig"(淘工厂大货订单)
// 示例值: ["cn","ws"]
func (r *AlibabaTradeGetBuyerOrderListRequest) SetBizTypes(bizTypes []string) {
	for _, bizType := range bizTypes {
		r.AddParameter("bizTypes", bizType)
	}
}

// SetCreateEndTime 设置下单结束时间
// 示例值: 20180802211113000+0800
func (r *AlibabaTradeGetBuyerOrderListRequest) SetCreateEndTime(createEndTime string) {
	r.AddParameter("createEndTime", createEndTime)
}

// SetCreateStartTime 设置下单开始时间
// 示例值: 20180102211113000+0800
func (r *AlibabaTradeGetBuyerOrderListRequest) SetCreateStartTime(createStartTime string) {
	r.AddParameter("createStartTime", createStartTime)
}

// SetIsHis 设置是否查询历史订单表,默认查询当前表，即默认值为false
// 示例值: false
func (r *AlibabaTradeGetBuyerOrderListRequest) SetIsHis(isHis bool) {
	r.AddParameter("isHis", strconv.FormatBool(isHis))
}

// SetModifyEndTime 设置查询修改时间结束
// 示例值: 20180802211113000+0800
func (r *AlibabaTradeGetBuyerOrderListRequest) SetModifyEndTime(modifyEndTime string) {
	r.AddParameter("modifyEndTime", modifyEndTime)
}

// SetModifyStartTime 设置查询修改时间开始
// 示例值: 20180102211113000+0800
func (r *AlibabaTradeGetBuyerOrderListRequest) SetModifyStartTime(modifyStartTime string) {
	r.AddParameter("modifyStartTime", modifyStartTime)
}

// SetOrderStatus 设置订单状态，值有 success, cancel(交易取消，违约金等交割完毕), waitbuyerpay(等待卖家付款)， waitsellersend(等待卖家发货), waitbuyerreceive(等待买家收货)
// 示例值: success
func (r *AlibabaTradeGetBuyerOrderListRequest) SetOrderStatus(orderStatus string) {
	r.AddParameter("orderStatus", orderStatus)
}

// SetPage 设置查询分页页码，从1开始
// 示例值: 1
func (r *AlibabaTradeGetBuyerOrderListRequest) SetPage(page int) {
	r.AddParameter("page", strconv.Itoa(page))
}

// SetPageSize 设置查询的每页的数量
// 示例值: 20
func (r *AlibabaTradeGetBuyerOrderListRequest) SetPageSize(pageSize int) {
	r.AddParameter("pageSize", strconv.Itoa(pageSize))
}

// SetRefundStatus 设置退款状态，支持："waitselleragree"(等待卖家同意), "refundsuccess"(退款成功), "refundclose"(退款关闭), "waitbuyermodify"(待买家修改), "waitbuyersend"(等待买家退货), "waitsellerreceive"(等待卖家确认收货)
// 示例值: refundsuccess
func (r *AlibabaTradeGetBuyerOrderListRequest) SetRefundStatus(refundStatus string) {
	r.AddParameter("refundStatus", refundStatus)
}

// SetSellerMemberId 设置卖家memberId
// 示例值: b2b-1624961198
func (r *AlibabaTradeGetBuyerOrderListRequest) SetSellerMemberId(sellerMemberId string) {
	r.AddParameter("sellerMemberId", sellerMemberId)
}

// SetSellerLoginId 设置卖家loginId
// 示例值: alitestforisv02
func (r *AlibabaTradeGetBuyerOrderListRequest) SetSellerLoginId(sellerLoginId string) {
	r.AddParameter("sellerLoginId", sellerLoginId)
}

// SetSellerRateStatus 设置卖家评价状态 (4:已评价,5:未评价,6;不需要评价)
// 示例值: 6
func (r *AlibabaTradeGetBuyerOrderListRequest) SetSellerRateStatus(sellerRateStatus int) {
	r.AddParameter("sellerRateStatus", strconv.Itoa(sellerRateStatus))
}

// SetTradeType 设置交易类型: 担保交易(1), 预存款交易(2), ETC境外收单交易(3), 即时到帐交易(4), 保障金安全交易(5), 统一交易流程(6), 分阶段交易(7), 货到付款交易(8), 信用凭证支付交易(9), 账期支付交易(10), 1688交易4.0，新分阶段交易(50060), 当面付的交易流程(50070), 服务类的交易流程(50080)
// 示例值: 50060
func (r *AlibabaTradeGetBuyerOrderListRequest) SetTradeType(tradeType string) {
	r.AddParameter("tradeType", tradeType)
}

// SetProductName 设置商品名称
// 示例值: 测试商品
func (r *AlibabaTradeGetBuyerOrderListRequest) SetProductName(productName string) {
	r.AddParameter("productName", productName)
}

// SetNeedBuyerAddressAndPhone 设置是否需要查询买家的详细地址信息和电话
// 示例值: false
func (r *AlibabaTradeGetBuyerOrderListRequest) SetNeedBuyerAddressAndPhone(needBuyerAddressAndPhone bool) {
	r.AddParameter("needBuyerAddressAndPhone", strconv.FormatBool(needBuyerAddressAndPhone))
}

// SetNeedMemoInfo 设置是否需要查询备注信息
// 示例值: false
func (r *AlibabaTradeGetBuyerOrderListRequest) SetNeedMemoInfo(needMemoInfo bool) {
	r.AddParameter("needMemoInfo", strconv.FormatBool(needMemoInfo))
}

// SetOutOrderId 设置外部订单号，可用于控制幂等
// 示例值: 90187872898371
func (r *AlibabaTradeGetBuyerOrderListRequest) SetOutOrderId(outOrderId string) {
	r.AddParameter("outOrderId", outOrderId)
}

// GetApiName 获取API名称
func (r *AlibabaTradeGetBuyerOrderListRequest) GetApiName() string {
	return "alibaba.trade.getBuyerOrderList"
}

// GetApiVersion 获取API版本
func (r *AlibabaTradeGetBuyerOrderListRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaTradeGetBuyerOrderListRequest) GetBusinessModule() string {
	return "com.alibaba.trade"
}

// GetParameters 获取所有参数
func (r *AlibabaTradeGetBuyerOrderListRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}