package alibabatradegetbuyerview

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 订单详情查看(买家视角)响应
type Response struct {
	response.TopResponse
	// 订单详情信息
	Result *TradeInfo `json:"result,omitempty"`
}

// TradeInfo 订单详情信息
type TradeInfo struct {
	// 订单基础信息
	BaseInfo OrderBaseInfo `json:"baseInfo"`
	// 国内物流信息
	NativeLogistics NativeLogisticsInfo `json:"nativeLogistics"`
	// 商品条目信息
	ProductItems []ProductItemInfo `json:"productItems"`
	// 保障条款
	GuaranteesTerms []GuaranteeTermsInfo `json:"guaranteesTerms"`
	// 交易条款
	TradeTerms []TradeTermsInfo `json:"tradeTerms"`
	// 订单业务信息
	OrderBizInfo OrderBizInfo `json:"orderBizInfo"`
}

// OrderBaseInfo 订单基础信息
type OrderBaseInfo struct {
	// 完全发货时间
	AllDeliveredTime string `json:"allDeliveredTime"`
	// 卖家诚信等级
	SellerCreditLevel string `json:"sellerCreditLevel"`
	// 付款时间，如果有多次付款，这里返回的是首次付款时间
	PayTime string `json:"payTime"`
	// 折扣信息，单位分
	Discount int64 `json:"discount"`
	// 外部支付交易Id
	AlipayTradeId string `json:"alipayTradeId"`
	// 产品总金额(该订单产品明细表中的产品金额的和)，单位元
	SumProductPayment string `json:"sumProductPayment"`
	// 买家留言，不超过500字
	BuyerFeedback string `json:"buyerFeedback"`
	// 4.0交易流程模板code
	FlowTemplateCode string `json:"flowTemplateCode"`
	// 是否自主订单（邀约订单）
	SellerOrder bool `json:"sellerOrder"`
	// 买家loginId，旺旺Id
	BuyerLoginId string `json:"buyerLoginId"`
	// 修改时间
	ModifyTime string `json:"modifyTime"`
	// 买家子账号
	SubBuyerLoginId string `json:"subBuyerLoginId"`
	// 交易id
	Id int64 `json:"id"`
	// 关闭原因。buyerCancel:买家取消订单，sellerGoodsLack:卖家库存不足，other:其它
	CloseReason string `json:"closeReason"`
	// 买家联系人
	BuyerContact TradeContact `json:"buyerContact"`
	// 卖家支付宝id
	SellerAlipayId string `json:"sellerAlipayId"`
	// 完成时间
	CompleteTime string `json:"completeTime"`
	// 卖家loginId，旺旺Id
	SellerLoginId string `json:"sellerLoginId"`
	// 买家主账号id
	BuyerID string `json:"buyerID"`
	// 关闭订单操作类型
	CloseOperateType string `json:"closeOperateType"`
	// 应付款总金额，totalAmount = ∑itemAmount + shippingFee，单位为元
	TotalAmount string `json:"totalAmount"`
	// 卖家主账号id
	SellerID string `json:"sellerID"`
	// 运费，单位为元
	ShippingFee string `json:"shippingFee"`
	// 买家数字id
	BuyerUserId int64 `json:"buyerUserId"`
	// 买家备忘信息
	BuyerMemo string `json:"buyerMemo"`
	// 退款金额，单位为元
	Refund string `json:"refund"`
	// 交易状态
	Status string `json:"status"`
	// 退款金额
	RefundPayment int64 `json:"refundPayment"`
	// 卖家联系人信息
	SellerContact TradeSellerContact `json:"sellerContact"`
	// 红包金额，实付金额（totalAmount）已经计算过红包金额
	CouponFee string `json:"couponFee"`
	// 买家备忘标志
	BuyerRemarkIcon string `json:"buyerRemarkIcon"`
	// 收件人信息
	ReceiverInfo OrderReceiverInfo `json:"receiverInfo"`
	// 订单的售中退款状态
	RefundStatus string `json:"refundStatus"`
	// 备注，1688指下单时的备注
	Remark string `json:"remark"`
	// 预订单ID
	PreOrderId int64 `json:"preOrderId"`
	// 确认时间
	ConfirmedTime string `json:"confirmedTime"`
	// 关闭订单备注
	CloseRemark string `json:"closeRemark"`
	// 交易类型
	TradeType string `json:"tradeType"`
	// 收货时间，这里返回的是完全收货时间
	ReceivingTime string `json:"receivingTime"`
	// 分阶段法务协议地址
	StepAgreementPath string `json:"stepAgreementPath"`
	// 交易id(字符串格式)
	IdOfStr string `json:"idOfStr"`
	// 订单的售后退款状态
	RefundStatusForAs string `json:"refundStatusForAs"`
	// 是否一次性付款
	StepPayAll bool `json:"stepPayAll"`
	// 卖家数字id
	SellerUserId int64 `json:"sellerUserId"`
	// 买家支付宝id
	BuyerAlipayId string `json:"buyerAlipayId"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 业务类型
	BusinessType string `json:"businessType"`
	// 是否海外代发订单
	OverSeaOrder bool `json:"overSeaOrder"`
	// 退款单ID
	RefundId string `json:"refundId"`
	// 下单时指定的交易方式
	TradeTypeDesc string `json:"tradeTypeDesc"`
	// 支付渠道名称列表
	PayChannelList []string `json:"payChannelList"`
	// 下单时指定的交易方式
	TradeTypeCode string `json:"tradeTypeCode"`
	// 支付超时时间，定长情况时单位：秒，目前都是定长
	PayTimeout int64 `json:"payTimeout"`
	// 支付超时TYPE，0：定长，1：固定时间
	PayTimeoutType int `json:"payTimeoutType"`
	// 支付渠道code
	PayChannelCodeList []string `json:"payChannelCodeList"`
	// 供货库存模式，jit（jit模式）或cang（仓发模式）
	InventoryMode string `json:"inventoryMode"`
	// 外部订单号
	OutOrderId string `json:"outOrderId"`
}

// TradeContact 买家联系人
type TradeContact struct {
	// 联系电话
	Phone string `json:"phone"`
	// 传真
	Fax string `json:"fax"`
	// 邮箱
	Email string `json:"email"`
	// 联系人在平台的IM账号
	ImInPlatform string `json:"imInPlatform"`
	// 联系人名称
	Name string `json:"name"`
	// 联系人手机号
	Mobile string `json:"mobile"`
	// 公司名称
	CompanyName string `json:"companyName"`
}

// TradeSellerContact 卖家联系人信息
type TradeSellerContact struct {
	// 联系电话
	Phone string `json:"phone"`
	// 传真
	Fax string `json:"fax"`
	// 邮箱
	Email string `json:"email"`
	// 联系人在平台的IM账号
	ImInPlatform string `json:"imInPlatform"`
	// 联系人名称
	Name string `json:"name"`
	// 联系人手机号
	Mobile string `json:"mobile"`
	// 公司名称
	CompanyName string `json:"companyName"`
	// 发件人名称，在微供等分销场景下由分销商设置
	WgSenderName string `json:"wgSenderName"`
	// 发件人电话，在微供等分销场景下由分销商设置
	WgSenderPhone string `json:"wgSenderPhone"`
	// 旺铺名称
	ShopName string `json:"shopName"`
}

// OrderReceiverInfo 收件人信息
type OrderReceiverInfo struct {
	// 收件人
	ToFullName string `json:"toFullName"`
	// 收货人地址区域编码
	ToDivisionCode string `json:"toDivisionCode"`
	// 收件人移动电话
	ToMobile string `json:"toMobile"`
	// 收件人电话
	ToPhone string `json:"toPhone"`
	// 邮编
	ToPost string `json:"toPost"`
	// 收货人街道或镇区域编码，可能为空
	ToTownCode string `json:"toTownCode"`
	// 收货地址
	ToArea string `json:"toArea"`
}

// NativeLogisticsInfo 国内物流信息
type NativeLogisticsInfo struct {
	// 详细地址
	Address string `json:"address"`
	// 县，区
	Area string `json:"area"`
	// 省市区编码
	AreaCode string `json:"areaCode"`
	// 城市
	City string `json:"city"`
	// 联系人姓名
	ContactPerson string `json:"contactPerson"`
	// 传真
	Fax string `json:"fax"`
	// 手机
	Mobile string `json:"mobile"`
	// 省份
	Province string `json:"province"`
	// 电话
	Telephone string `json:"telephone"`
	// 邮编
	Zip string `json:"zip"`
	// 运单明细
	LogisticsItems []NativeLogisticsItemsInfo `json:"logisticsItems"`
	// 街道或镇区域编码
	TownCode string `json:"townCode"`
	// 街道或镇名称
	Town string `json:"town"`
}

// NativeLogisticsItemsInfo 运单明细
type NativeLogisticsItemsInfo struct {
	// 发货时间
	DeliveredTime string `json:"deliveredTime"`
	// 物流编号
	LogisticsCode string `json:"logisticsCode"`
	// 类型
	Type string `json:"type"`
	// 主键id
	Id int64 `json:"id"`
	// 状态
	Status string `json:"status"`
	// 修改时间
	GmtModified string `json:"gmtModified"`
	// 创建时间
	GmtCreate string `json:"gmtCreate"`
	// 运费(单位为元)
	Carriage string `json:"carriage"`
	// 发货省
	FromProvince string `json:"fromProvince"`
	// 发货市
	FromCity string `json:"fromCity"`
	// 发货区
	FromArea string `json:"fromArea"`
	// 发货街道地址
	FromAddress string `json:"fromAddress"`
	// 发货联系电话
	FromPhone string `json:"fromPhone"`
	// 发货联系手机
	FromMobile string `json:"fromMobile"`
	// 发货地址邮编
	FromPost string `json:"fromPost"`
	// 物流公司Id
	LogisticsCompanyId int64 `json:"logisticsCompanyId"`
	// 物流公司编号
	LogisticsCompanyNo string `json:"logisticsCompanyNo"`
	// 物流公司名称
	LogisticsCompanyName string `json:"logisticsCompanyName"`
	// 物流公司运单号
	LogisticsBillNo string `json:"logisticsBillNo"`
	// 商品明细条目id，如有多个以,分隔
	SubItemIds string `json:"subItemIds"`
}

// ProductItemInfo 商品条目信息
type ProductItemInfo struct {
	// 指定单品货号
	CargoNumber string `json:"cargoNumber"`
	// 描述
	Description string `json:"description"`
	// 实付金额，单位为元
	ItemAmount string `json:"itemAmount"`
	// 商品名称
	Name string `json:"name"`
	// 原始单价，以元为单位
	Price string `json:"price"`
	// 产品ID（非在线产品为空）
	ProductID int64 `json:"productID"`
	// 商品图片url
	ProductImgUrl []string `json:"productImgUrl"`
	// 产品快照url
	ProductSnapshotUrl string `json:"productSnapshotUrl"`
	// 以unit为单位的数量
	Quantity string `json:"quantity"`
	// 退款金额，单位为元
	Refund string `json:"refund"`
	// skuID
	SkuID int64 `json:"skuID"`
	// 排序字段
	Sort int `json:"sort"`
	// 子订单状态
	Status string `json:"status"`
	// 子订单号，或商品明细条目ID
	SubItemID int64 `json:"subItemID"`
	// 类型
	Type string `json:"type"`
	// 售卖单位
	Unit string `json:"unit"`
	// 重量
	Weight string `json:"weight"`
	// 重量单位
	WeightUnit string `json:"weightUnit"`
	// 保障条款
	GuaranteesTerms []GuaranteeTermsInfo `json:"guaranteesTerms"`
	// 指定商品货号
	ProductCargoNumber string `json:"productCargoNumber"`
	// SKU信息
	SkuInfos []SkuItemDesc `json:"skuInfos"`
	// 订单明细涨价或降价的金额
	EntryDiscount int64 `json:"entryDiscount"`
	// 订单销售属性ID
	SpecId string `json:"specId"`
	// 以unit为单位的quantity精度系数
	QuantityFactor string `json:"quantityFactor"`
	// 子订单状态描述
	StatusStr string `json:"statusStr"`
	// 退款状态
	RefundStatus string `json:"refundStatus"`
	// 关闭原因
	CloseReason string `json:"closeReason"`
	// 物流状态
	LogisticsStatus int `json:"logisticsStatus"`
	// 创建时间
	GmtCreate string `json:"gmtCreate"`
	// 修改时间
	GmtModified string `json:"gmtModified"`
	// 明细完成时间
	GmtCompleted string `json:"gmtCompleted"`
	// 库存超时时间
	GmtPayExpireTime string `json:"gmtPayExpireTime"`
	// 售中退款单号
	RefundId string `json:"refundId"`
	// 子订单号(字符串类型)
	SubItemIDString string `json:"subItemIDString"`
	// 售后退款单号
	RefundIdForAs string `json:"refundIdForAs"`
	// 分担邮费，单位为元
	SharePostage string `json:"sharePostage"`
}

// GuaranteeTermsInfo 保障条款
type GuaranteeTermsInfo struct {
	// 保障条款
	AssuranceInfo string `json:"assuranceInfo"`
	// 保障方式
	AssuranceType string `json:"assuranceType"`
	// 质量保证类型
	QualityAssuranceType string `json:"qualityAssuranceType"`
	// 保障条款值
	Value string `json:"value"`
}

// SkuItemDesc SKU信息
type SkuItemDesc struct {
	// 属性名
	Name string `json:"name"`
	// 属性值
	Value string `json:"value"`
}

// TradeTermsInfo 交易条款
type TradeTermsInfo struct {
	// 支付状态
	PayStatus string `json:"payStatus"`
	// 完成阶段支付时间
	PayTime string `json:"payTime"`
	// 支付方式
	PayWay string `json:"payWay"`
	// 付款额
	PhasAmount string `json:"phasAmount"`
	// 阶段单id
	Phase int64 `json:"phase"`
	// 阶段条件
	PhaseCondition string `json:"phaseCondition"`
	// 阶段时间
	PhaseDate string `json:"phaseDate"`
	// 是否银行卡支付
	CardPay bool `json:"cardPay"`
	// 是否快捷支付
	ExpressPay bool `json:"expressPay"`
	// 支付方式
	PayWayDesc string `json:"payWayDesc"`
}

// OrderBizInfo 订单业务信息
type OrderBizInfo struct {
	// 是否采源宝订单
	OdsCyd bool `json:"odsCyd"`
	// 账期交易订单的到账时间
	AccountPeriodTime string `json:"accountPeriodTime"`
	// 是否诚e赊订单
	CreditOrder bool `json:"creditOrder"`
	// 是否加工定制订单
	IsCz bool `json:"isCz"`
	// 是否定制订单
	IsDz bool `json:"isDz"`
	// 是否定制订单
	Dz bool `json:"dz"`
	// 是否dropshipping订单
	Dropshipping bool `json:"dropshipping"`
	// 运费险信息
	ShippingInsurance string `json:"shippingInsurance"`
	// 大店仓发订单
	HyperLinkCangFaOrder bool `json:"hyperLinkCangFaOrder"`
	// 超链一阶段订单
	HyperLinkOrder bool `json:"hyperLinkOrder"`
	// 超链大店二阶段订单的第二阶段订单
	HyperLinkSecondStepOrder bool `json:"hyperLinkSecondStepOrder"`
	// 超链一阶段订单的发货模式
	HyperLinkShipType string `json:"hyperLinkShipType"`
	// 闪电仓订单
	LightningWarehouse bool `json:"lightningWarehouse"`
	// ae上门揽订单
	AeDoorPickUp bool `json:"aeDoorPickUp"`
	// 分账订单
	Fz bool `json:"fz"`
	// 托管官方上门揽订单
	TgOfficialPickUp string `json:"tgOfficialPickUp"`
}

// WrapResult 解析响应结果
func (r *Response) WrapResult(result string) {
	if result == "" {
		return
	}

	// 保存原始结果到Body
	r.Body = result

	// 尝试解析为错误响应
	var errorResp struct {
		ErrorMsg  string `json:"errorMsg"`
		ErrorCode string `json:"errorCode"`
	}

	if err := json.Unmarshal([]byte(result), &errorResp); err == nil {
		if errorResp.ErrorMsg != "" || errorResp.ErrorCode != "" {
			r.ErrorResponse.Code = -1
			r.ErrorResponse.Msg = fmt.Sprintf("errorMsg: %s, errorCode: %s", errorResp.ErrorMsg, errorResp.ErrorCode)
			return
		}
	}

	// 解析正常响应
	if err := json.Unmarshal([]byte(result), r); err != nil {
		r.ErrorResponse.Code = 500
		r.ErrorResponse.Msg = fmt.Sprintf("JSON解析失败: %v, 原始数据: %s", err, result)
	}
}