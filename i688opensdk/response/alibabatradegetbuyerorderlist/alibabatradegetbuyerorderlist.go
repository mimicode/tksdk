package alibabatradegetbuyerorderlist

import (
	"encoding/json"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 订单列表查看(买家视角)响应
type Response struct {
	response.TopResponse
	// 查询返回列表
	Result []TradeInfo `json:"result"`
	// 总记录数
	TotalRecord int64 `json:"totalRecord"`
}

// TradeInfo 订单信息
type TradeInfo struct {
	// 采购单详情列表，为大企业采购订单独有域
	QuoteList []CaigouQuoteInfo `json:"quoteList"`
	// 订单扩展属性
	ExtAttributes []KeyValuePair `json:"extAttributes"`
	// 订单评价信息
	OrderRateInfo OrderRateInfo `json:"orderRateInfo"`
	// 发票信息
	OrderInvoiceInfo OrderInvoiceModel `json:"orderInvoiceInfo"`
	// 交易条款
	TradeTerms []TradeTermsInfo `json:"tradeTerms"`
	// 国内物流
	NativeLogistics NativeLogisticsInfo `json:"nativeLogistics"`
	// 保障条款
	GuaranteesTerms GuaranteeTermsInfo `json:"guaranteesTerms"`
	// 订单基础信息
	BaseInfo OrderBaseInfo `json:"baseInfo"`
	// 订单业务信息
	OrderBizInfo OrderBizInfo `json:"orderBizInfo"`
	// 商品条目信息
	ProductItems []ProductItemInfo `json:"productItems"`
	// 跨境地址扩展信息
	OverseasExtraAddress OverseasExtraAddress `json:"overseasExtraAddress"`
	// 跨境报关信息
	Customs Customs `json:"customs"`
	// 海外收货信息
	OverseaLogisticsInfo OverseaLogisticsInfo `json:"overseaLogisticsInfo"`
}

// CaigouQuoteInfo 采购单详情信息
type CaigouQuoteInfo struct {
	// 根据实际API响应定义字段
}

// KeyValuePair 键值对
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// OrderRateInfo 订单评价信息
type OrderRateInfo struct {
	// 买家评价状态
	BuyerRateStatus int `json:"buyerRateStatus"`
	// 卖家评价状态
	SellerRateStatus int `json:"sellerRateStatus"`
}

// OrderInvoiceModel 发票信息
type OrderInvoiceModel struct {
	// 根据实际API响应定义字段
}

// TradeTermsInfo 交易条款信息
type TradeTermsInfo struct {
	// 支付状态
	PayStatus string `json:"payStatus"`
	// 支付时间
	PayTime string `json:"payTime"`
	// 支付方式
	PayWay string `json:"payWay"`
	// 阶段金额
	PhasAmount float64 `json:"phasAmount"`
	// 阶段
	Phase int64 `json:"phase"`
}

// NativeLogisticsInfo 国内物流信息
type NativeLogisticsInfo struct {
	// 区域
	Area string `json:"area"`
	// 区域代码
	AreaCode string `json:"areaCode"`
	// 城市
	City string `json:"city"`
	// 联系人
	ContactPerson string `json:"contactPerson"`
	// 省份
	Province string `json:"province"`
	// 邮编
	Zip string `json:"zip"`
}

// GuaranteeTermsInfo 保障条款信息
type GuaranteeTermsInfo struct {
	// 根据实际API响应定义字段
}

// OrderBaseInfo 订单基础信息
type OrderBaseInfo struct {
	// 业务类型
	BusinessType string `json:"businessType"`
	// 买家ID
	BuyerID string `json:"buyerID"`
	// 完成时间
	CompleteTime string `json:"completeTime"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 订单ID
	Id int64 `json:"id"`
	// 修改时间
	ModifyTime string `json:"modifyTime"`
	// 退款金额
	Refund float64 `json:"refund"`
	// 卖家ID
	SellerID string `json:"sellerID"`
	// 运费
	ShippingFee float64 `json:"shippingFee"`
	// 状态
	Status string `json:"status"`
	// 总金额
	TotalAmount float64 `json:"totalAmount"`
	// 折扣
	Discount float64 `json:"discount"`
	// 买家联系信息
	BuyerContact ContactInfo `json:"buyerContact"`
	// 卖家联系信息
	SellerContact ContactInfo `json:"sellerContact"`
	// 交易类型
	TradeType string `json:"tradeType"`
	// 退款支付
	RefundPayment float64 `json:"refundPayment"`
	// 订单ID字符串
	IdOfStr string `json:"idOfStr"`
	// 支付宝交易ID
	AlipayTradeId string `json:"alipayTradeId"`
	// 收货人信息
	ReceiverInfo ReceiverInfo `json:"receiverInfo"`
	// 买家登录ID
	BuyerLoginId string `json:"buyerLoginId"`
	// 卖家登录ID
	SellerLoginId string `json:"sellerLoginId"`
	// 买家用户ID
	BuyerUserId int64 `json:"buyerUserId"`
	// 卖家用户ID
	SellerUserId int64 `json:"sellerUserId"`
	// 买家支付宝ID
	BuyerAlipayId string `json:"buyerAlipayId"`
	// 卖家支付宝ID
	SellerAlipayId string `json:"sellerAlipayId"`
	// 关闭原因
	CloseReason string `json:"closeReason"`
	// 商品支付总额
	SumProductPayment float64 `json:"sumProductPayment"`
	// 是否全额分期付款
	StepPayAll bool `json:"stepPayAll"`
}

// ContactInfo 联系信息
type ContactInfo struct {
	// 电话
	Phone string `json:"phone"`
	// 姓名
	Name string `json:"name"`
	// 平台内IM
	ImInPlatform string `json:"imInPlatform"`
	// 公司名称
	CompanyName string `json:"companyName"`
}

// ReceiverInfo 收货人信息
type ReceiverInfo struct {
	// 收货人全名
	ToFullName string `json:"toFullName"`
	// 行政区划代码
	ToDivisionCode string `json:"toDivisionCode"`
	// 邮编
	ToPost string `json:"toPost"`
	// 地区
	ToArea string `json:"toArea"`
}

// OrderBizInfo 订单业务信息
type OrderBizInfo struct {
	// 根据实际API响应定义字段
}

// ProductItemInfo 商品条目信息
type ProductItemInfo struct {
	// 商品金额
	ItemAmount float64 `json:"itemAmount"`
	// 商品名称
	Name string `json:"name"`
	// 价格
	Price float64 `json:"price"`
	// 商品ID
	ProductID int64 `json:"productID"`
	// 商品图片URL
	ProductImgUrl []string `json:"productImgUrl"`
	// 商品快照URL
	ProductSnapshotUrl string `json:"productSnapshotUrl"`
	// 数量
	Quantity int `json:"quantity"`
	// 退款
	Refund float64 `json:"refund"`
	// SKU ID
	SkuID int64 `json:"skuID"`
	// 状态
	Status string `json:"status"`
	// 子项目ID
	SubItemID int64 `json:"subItemID"`
	// 类型
	Type string `json:"type"`
	// 单位
	Unit string `json:"unit"`
	// 保障条款
	GuaranteesTerms []interface{} `json:"guaranteesTerms"`
	// 商品货号
	ProductCargoNumber string `json:"productCargoNumber"`
	// SKU信息
	SkuInfos []SkuInfo `json:"skuInfos"`
	// 条目折扣
	EntryDiscount float64 `json:"entryDiscount"`
	// 规格ID
	SpecId string `json:"specId"`
	// 数量因子
	QuantityFactor int `json:"quantityFactor"`
	// 状态字符串
	StatusStr string `json:"statusStr"`
	// 关闭原因
	CloseReason string `json:"closeReason"`
	// 物流状态
	LogisticsStatus int `json:"logisticsStatus"`
}

// SkuInfo SKU信息
type SkuInfo struct {
	// 名称
	Name string `json:"name"`
	// 值
	Value string `json:"value"`
}

// OverseasExtraAddress 跨境地址扩展信息
type OverseasExtraAddress struct {
	// 根据实际API响应定义字段
}

// Customs 跨境报关信息
type Customs struct {
	// 根据实际API响应定义字段
}

// OverseaLogisticsInfo 海外收货信息
type OverseaLogisticsInfo struct {
	// 海外收货地址
	OverseasTransportUserAddrInfo OverseasTransportUserAddrInfo `json:"overseasTransportUserAddrInfo"`
	// 海外物流单号
	OverseasLogisticsIds []string `json:"overseasLogisticsIds"`
}

// OverseasTransportUserAddrInfo 海外收货地址
type OverseasTransportUserAddrInfo struct {
	// 邮编
	PostCode string `json:"postCode"`
	// 国家
	CountryName string `json:"countryName"`
	// 省份
	ProvinceName string `json:"provinceName"`
	// 城市
	CityName string `json:"cityName"`
	// 区域
	AreaName string `json:"areaName"`
	// 地址详情
	AddressDetail string `json:"addressDetail"`
	// 手机号
	Mobile string `json:"mobile"`
	// 固定电话
	FixedPhone string `json:"fixedPhone"`
	// 仓库联系人
	WarehouseContactName string `json:"warehouseContactName"`
	// 仓库名称
	WarehouseName string `json:"warehouseName"`
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