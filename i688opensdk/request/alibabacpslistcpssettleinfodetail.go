package request

import (
	"net/url"
	"strconv"
)

// AlibabaCpsListCpsSettleInfoDetailRequest 获取联盟结算明细列表(按月) API请求
// https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.listCPSSettleInfoDetail-1&aopApiCategory=category_new
type AlibabaCpsListCpsSettleInfoDetailRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsListCpsSettleInfoDetailRequest 创建AlibabaCpsListCpsSettleInfoDetailRequest实例
func NewAlibabaCpsListCpsSettleInfoDetailRequest() *AlibabaCpsListCpsSettleInfoDetailRequest {
	return &AlibabaCpsListCpsSettleInfoDetailRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil || r.Parameters.Get("pageNo") == "" {
		panic("pageNo is required")
	}
	if r.Parameters == nil || r.Parameters.Get("pageSize") == "" {
		panic("pageSize is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetPageNo 设置页偏移量
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) SetPageNo(pageNo int) {
	r.AddParameter("pageNo", strconv.Itoa(pageNo))
}

// SetPageSize 设置分页大小
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) SetPageSize(pageSize int) {
	r.AddParameter("pageSize", strconv.Itoa(pageSize))
}

// GetApiName 获取API名称
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) GetApiName() string {
	return "alibaba.cps.listCPSSettleInfoDetail"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsListCpsSettleInfoDetailRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}