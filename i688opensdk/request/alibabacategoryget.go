package request

import (
	"net/url"
	"strconv"
)

// AlibabaCategoryGetRequest 根据类目Id查询类目 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceEddODt&id=com.alibaba.product%3Aalibaba.category.get-1&aopApiCategory=category_new
type AlibabaCategoryGetRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCategoryGetRequest 创建AlibabaCategoryGetRequest实例
func NewAlibabaCategoryGetRequest() *AlibabaCategoryGetRequest {
	return &AlibabaCategoryGetRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCategoryGetRequest) CheckParameters() {
	// categoryID是必需参数
	if r.Parameters == nil || r.Parameters.Get("categoryID") == "" {
		panic("categoryID is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCategoryGetRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetCategoryID 设置类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaCategoryGetRequest) SetCategoryID(categoryID int64) {
	r.AddParameter("categoryID", strconv.FormatInt(categoryID, 10))
}

// GetCategoryID 获取类目ID
func (r *AlibabaCategoryGetRequest) GetCategoryID() int64 {
	if r.Parameters != nil {
		if categoryIDStr := r.Parameters.Get("categoryID"); categoryIDStr != "" {
			if categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64); err == nil {
				return categoryID
			}
		}
	}
	return 0
}

// GetApiName 获取API名称
func (r *AlibabaCategoryGetRequest) GetApiName() string {
	return "alibaba.category.get"
}

// GetApiVersion 获取API版本
func (r *AlibabaCategoryGetRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCategoryGetRequest) GetBusinessModule() string {
	return "com.alibaba.product"
}

// GetParameters 获取所有参数
func (r *AlibabaCategoryGetRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}