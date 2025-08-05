package request

import (
	"net/url"
	"strconv"
)

// AlibabaTradeGetBuyerViewRequest 订单详情查看(买家视角) API请求
// https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.211d55edtSHFBC&id=com.alibaba.trade%3Aalibaba.trade.get.buyerView-1&aopApiCategory=category_new
type AlibabaTradeGetBuyerViewRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaTradeGetBuyerViewRequest 创建AlibabaTradeGetBuyerViewRequest实例
func NewAlibabaTradeGetBuyerViewRequest() *AlibabaTradeGetBuyerViewRequest {
	return &AlibabaTradeGetBuyerViewRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaTradeGetBuyerViewRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil || r.Parameters.Get("webSite") == "" {
		panic("webSite is required")
	}
	if r.Parameters == nil || r.Parameters.Get("orderId") == "" {
		panic("orderId is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaTradeGetBuyerViewRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetWebSite 设置站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
// 示例值: 1688
func (r *AlibabaTradeGetBuyerViewRequest) SetWebSite(webSite string) {
	r.AddParameter("webSite", webSite)
}

// SetOrderId 设置交易的订单id
// 示例值: 123456
func (r *AlibabaTradeGetBuyerViewRequest) SetOrderId(orderId int64) {
	r.AddParameter("orderId", strconv.FormatInt(orderId, 10))
}

// SetIncludeFields 设置查询结果中包含的域
// GuaranteesTerms：保障条款，NativeLogistics：物流信息，RateDetail：评价详情，OrderInvoice：发票信息
// 默认返回GuaranteesTerms、NativeLogistics、OrderInvoice
// 示例值: GuaranteesTerms,NativeLogistics,RateDetail,OrderInvoice
func (r *AlibabaTradeGetBuyerViewRequest) SetIncludeFields(includeFields string) {
	r.AddParameter("includeFields", includeFields)
}

// SetAttributeKeys 设置垂直表中的attributeKeys
func (r *AlibabaTradeGetBuyerViewRequest) SetAttributeKeys(attributeKeys []string) {
	for _, key := range attributeKeys {
		r.AddParameter("attributeKeys", key)
	}
}

// SetOutOrderId 设置外部订单id，控制幂等
// 示例值: 1556246
func (r *AlibabaTradeGetBuyerViewRequest) SetOutOrderId(outOrderId string) {
	r.AddParameter("outOrderId", outOrderId)
}

// GetApiName 获取API名称
func (r *AlibabaTradeGetBuyerViewRequest) GetApiName() string {
	return "alibaba.trade.get.buyerView"
}

// GetApiVersion 获取API版本
func (r *AlibabaTradeGetBuyerViewRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaTradeGetBuyerViewRequest) GetBusinessModule() string {
	return "com.alibaba.trade"
}

// GetParameters 获取所有参数
func (r *AlibabaTradeGetBuyerViewRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}