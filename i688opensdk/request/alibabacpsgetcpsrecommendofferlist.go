package request

import (
	"net/url"
	"strconv"
)

// AlibabaCpsGetCpsRecommendOfferListRequest 获取个性化推荐的cps商品(含推广链接) API请求
// https://open.1688.com/api/apidocdetail.htm?id=alibaba.cps.getCpsRecommendOfferList
type AlibabaCpsGetCpsRecommendOfferListRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsGetCpsRecommendOfferListRequest 创建AlibabaCpsGetCpsRecommendOfferListRequest实例
func NewAlibabaCpsGetCpsRecommendOfferListRequest() *AlibabaCpsGetCpsRecommendOfferListRequest {
	return &AlibabaCpsGetCpsRecommendOfferListRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil || r.Parameters.Get("loginId") == "" {
		panic("loginId is required")
	}
	if r.Parameters == nil || r.Parameters.Get("mediaId") == "" {
		panic("mediaId is required")
	}
	if r.Parameters == nil || r.Parameters.Get("mediaZoneId") == "" {
		panic("mediaZoneId is required")
	}
	if r.Parameters == nil || r.Parameters.Get("pageNo") == "" {
		panic("pageNo is required")
	}
	if r.Parameters == nil || r.Parameters.Get("pageSize") == "" {
		panic("pageSize is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetLoginId 设置下游被推荐用户登录id账号
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) SetLoginId(loginId string) {
	r.AddParameter("loginId", loginId)
}

// SetMediaId 设置媒体id
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) SetMediaId(mediaId int64) {
	r.AddParameter("mediaId", strconv.FormatInt(mediaId, 10))
}

// SetMediaZoneId 设置媒体推广位id
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) SetMediaZoneId(mediaZoneId string) {
	r.AddParameter("mediaZoneId", mediaZoneId)
}

// SetExt 设置媒体自定义扩展字段
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) SetExt(ext string) {
	r.AddParameter("ext", ext)
}

// SetPageNo 设置页码
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) SetPageNo(pageNo int) {
	r.AddParameter("pageNo", strconv.Itoa(pageNo))
}

// SetPageSize 设置每页数量
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) SetPageSize(pageSize int) {
	r.AddParameter("pageSize", strconv.Itoa(pageSize))
}

// GetApiName 获取API名称
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) GetApiName() string {
	return "alibaba.cps.getCpsRecommendOfferList"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsGetCpsRecommendOfferListRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}