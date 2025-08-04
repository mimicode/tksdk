package request

import (
	"net/url"
)

// AlibabaCpsParseUrlOrTokenRequest 解析推广物料信息 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.parseUrlOrToken-1&aopApiCategory=category_new
type AlibabaCpsParseUrlOrTokenRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsParseUrlOrTokenRequest 创建AlibabaCpsParseUrlOrTokenRequest实例
func NewAlibabaCpsParseUrlOrTokenRequest() *AlibabaCpsParseUrlOrTokenRequest {
	return &AlibabaCpsParseUrlOrTokenRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsParseUrlOrTokenRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil || r.Parameters.Get("urlOrToken") == "" {
		panic("urlOrToken is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsParseUrlOrTokenRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetUrlOrToken 设置待解析内容
// 示例值: 111
func (r *AlibabaCpsParseUrlOrTokenRequest) SetUrlOrToken(urlOrToken string) {
	r.AddParameter("urlOrToken", urlOrToken)
}

// GetUrlOrToken 获取待解析内容
func (r *AlibabaCpsParseUrlOrTokenRequest) GetUrlOrToken() string {
	if r.Parameters == nil {
		return ""
	}
	return r.Parameters.Get("urlOrToken")
}

// GetApiName 获取API名称
func (r *AlibabaCpsParseUrlOrTokenRequest) GetApiName() string {
	return "alibaba.cps.parseUrlOrToken"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsParseUrlOrTokenRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsParseUrlOrTokenRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsParseUrlOrTokenRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}