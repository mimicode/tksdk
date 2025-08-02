package request

import (
	"net/url"
	"strconv"
)

// AlibabaCpsListShopPageQueryRequest 获取联盟商家列表 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472ccel58TEV&id=com.alibaba.p4p%3Aalibaba.cps.listShopPageQuery-1&aopApiCategory=category_new
type AlibabaCpsListShopPageQueryRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsListShopPageQueryRequest 创建AlibabaCpsListShopPageQueryRequest实例
func NewAlibabaCpsListShopPageQueryRequest() *AlibabaCpsListShopPageQueryRequest {
	return &AlibabaCpsListShopPageQueryRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsListShopPageQueryRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	if r.Parameters.Get("categoryId") == "" {
		panic("categoryId is required")
	}
	if r.Parameters.Get("pageNo") == "" {
		panic("pageNo is required")
	}
	if r.Parameters.Get("pageSize") == "" {
		panic("pageSize is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsListShopPageQueryRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetSellerId 设置卖家id
func (r *AlibabaCpsListShopPageQueryRequest) SetSellerId(sellerId int64) {
	r.AddParameter("sellerId", strconv.FormatInt(sellerId, 10))
}

// SetCompanyName 设置公司名称
func (r *AlibabaCpsListShopPageQueryRequest) SetCompanyName(companyName string) {
	r.AddParameter("companyName", companyName)
}

// SetCategoryId 设置类目id
func (r *AlibabaCpsListShopPageQueryRequest) SetCategoryId(categoryId int64) {
	r.AddParameter("categoryId", strconv.FormatInt(categoryId, 10))
}

// SetDefineTags 设置属性标签
func (r *AlibabaCpsListShopPageQueryRequest) SetDefineTags(defineTags string) {
	r.AddParameter("defineTags", defineTags)
}

// SetFilterRatioMin 设置佣金比例下限
func (r *AlibabaCpsListShopPageQueryRequest) SetFilterRatioMin(filterRatioMin float64) {
	r.AddParameter("filterRatioMin", strconv.FormatFloat(filterRatioMin, 'f', -1, 64))
}

// SetFilterRatioMax 设置佣金比例上限
func (r *AlibabaCpsListShopPageQueryRequest) SetFilterRatioMax(filterRatioMax float64) {
	r.AddParameter("filterRatioMax", strconv.FormatFloat(filterRatioMax, 'f', -1, 64))
}

// SetSortField 设置排序字段
func (r *AlibabaCpsListShopPageQueryRequest) SetSortField(sortField string) {
	r.AddParameter("sortField", sortField)
}

// SetPageNo 设置页偏移量
func (r *AlibabaCpsListShopPageQueryRequest) SetPageNo(pageNo int) {
	r.AddParameter("pageNo", strconv.Itoa(pageNo))
}

// SetPageSize 设置分页大小
func (r *AlibabaCpsListShopPageQueryRequest) SetPageSize(pageSize int) {
	r.AddParameter("pageSize", strconv.Itoa(pageSize))
}

// GetApiName 获取API名称
func (r *AlibabaCpsListShopPageQueryRequest) GetApiName() string {
	return "alibaba.cps.listShopPageQuery"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsListShopPageQueryRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsListShopPageQueryRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsListShopPageQueryRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}