package request

import (
	"net/url"
)

// AlibabaCpsGetCpsSettleSummaryInfoRequest 获取联盟结算摘要账单 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.getCPSSettleSummaryInfo-1&aopApiCategory=category_new
type AlibabaCpsGetCpsSettleSummaryInfoRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsGetCpsSettleSummaryInfoRequest 创建AlibabaCpsGetCpsSettleSummaryInfoRequest实例
func NewAlibabaCpsGetCpsSettleSummaryInfoRequest() *AlibabaCpsGetCpsSettleSummaryInfoRequest {
	return &AlibabaCpsGetCpsSettleSummaryInfoRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsGetCpsSettleSummaryInfoRequest) CheckParameters() {
	// 该接口无必需参数
}

// AddParameter 添加参数
func (r *AlibabaCpsGetCpsSettleSummaryInfoRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// GetApiName 获取API名称
func (r *AlibabaCpsGetCpsSettleSummaryInfoRequest) GetApiName() string {
	return "alibaba.cps.getCPSSettleSummaryInfo"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsGetCpsSettleSummaryInfoRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsGetCpsSettleSummaryInfoRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsGetCpsSettleSummaryInfoRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}