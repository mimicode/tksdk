package request

import (
	"net/url"
	"strconv"
)

// AlibabaCpsTradeBillListRequest 联盟账单列表查询 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.tradeBillList-1&aopApiCategory=category_new
type AlibabaCpsTradeBillListRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsTradeBillListRequest 创建AlibabaCpsTradeBillListRequest实例
func NewAlibabaCpsTradeBillListRequest() *AlibabaCpsTradeBillListRequest {
	return &AlibabaCpsTradeBillListRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsTradeBillListRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil || r.Parameters.Get("queryOrderType") == "" {
		panic("queryOrderType is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsTradeBillListRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetQueryOrderType 设置订单状态查询类型
// 全部订单:orderAll订单结算:orderSettle 订单付款:orderPay 主订单号查询:orderMain 子订单号查询:orderSub
func (r *AlibabaCpsTradeBillListRequest) SetQueryOrderType(queryOrderType string) {
	r.AddParameter("queryOrderType", queryOrderType)
}

// SetQueryTimeType 设置订单时间查询类型
// 非订单号查询时必传 *1.0 订单创建时间查询类型:gmtCreateTime(订单创建)* 1.1 订单成功时间查询类型:confirmTime(订单成功) * 1.2 订单结算时间查询类型:settleTime(订单结算) * 2.1 维权创建时间查询类型:rightsStartTime * 2.2 维权完成时间查询类型:rightsEndTime * 3 订单结算时间查询类型:settleTime
func (r *AlibabaCpsTradeBillListRequest) SetQueryTimeType(queryTimeType string) {
	r.AddParameter("queryTimeType", queryTimeType)
}

// SetQueryStartTime 设置查询开始时间
// 非订单号查询时必传 2017-12-01
func (r *AlibabaCpsTradeBillListRequest) SetQueryStartTime(queryStartTime string) {
	r.AddParameter("queryStartTime", queryStartTime)
}

// SetQueryEndTime 设置查询结束时间
// 非订单号查询时必传 2017-12-31
func (r *AlibabaCpsTradeBillListRequest) SetQueryEndTime(queryEndTime string) {
	r.AddParameter("queryEndTime", queryEndTime)
}

// SetPageNo 设置页码
// 1
func (r *AlibabaCpsTradeBillListRequest) SetPageNo(pageNo string) {
	r.AddParameter("pageNo", pageNo)
}

// SetPageSize 设置页面大小
// 100
func (r *AlibabaCpsTradeBillListRequest) SetPageSize(pageSize string) {
	r.AddParameter("pageSize", pageSize)
}

// SetOrderId 设置订单号
// 订单号查询时必传
func (r *AlibabaCpsTradeBillListRequest) SetOrderId(orderId string) {
	r.AddParameter("orderId", orderId)
}

// SetBizId 设置主订单id
func (r *AlibabaCpsTradeBillListRequest) SetBizId(bizId int64) {
	r.AddParameter("bizId", strconv.FormatInt(bizId, 10))
}

// SetBizSubId 设置子订单id
func (r *AlibabaCpsTradeBillListRequest) SetBizSubId(bizSubId int64) {
	r.AddParameter("bizSubId", strconv.FormatInt(bizSubId, 10))
}

// SetOrderState 设置订单状态
func (r *AlibabaCpsTradeBillListRequest) SetOrderState(orderState int) {
	r.AddParameter("orderState", strconv.Itoa(orderState))
}

// SetSettleState 设置结算状态
func (r *AlibabaCpsTradeBillListRequest) SetSettleState(settleState int) {
	r.AddParameter("settleState", strconv.Itoa(settleState))
}

// SetRightsState 设置维权状态
func (r *AlibabaCpsTradeBillListRequest) SetRightsState(rightsState int) {
	r.AddParameter("rightsState", strconv.Itoa(rightsState))
}

// GetApiName 获取API名称
func (r *AlibabaCpsTradeBillListRequest) GetApiName() string {
	return "alibaba.cps.tradeBillList"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsTradeBillListRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsTradeBillListRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsTradeBillListRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}