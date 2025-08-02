package request

import (
	"net/url"
	"strconv"
)

// alibaba.cps.getCpsRecommendSameOfferList( 获取同款推荐的商品(cps品含推广链接) )
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceEddODt&id=com.alibaba.p4p%3Aalibaba.cps.getCpsRecommendSameOfferList-1&aopApiCategory=category_new
type AlibabaCpsGetCpsRecommendSameOfferListRequest struct {
	Parameters *url.Values //请求参数
}

func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) CheckParameters() {
	// 检查必填参数
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	// offerId, offerType, mediaZoneId, mediaId 为必填参数，在调用时需要设置
}

// 添加请求参数
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// /////////////////////// 应用级参数 ///////////////////////////
// SetOfferId 设置商品ID（必填）
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetOfferId(offerId int64) {
	r.AddParameter("offerId", strconv.FormatInt(offerId, 10))
}

// SetOfferType 设置商品类型（必填）标识offerId类型；1688:1688商品;tao:淘宝商品
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetOfferType(offerType string) {
	r.AddParameter("offerType", offerType)
}

// SetMediaZoneId 设置媒体推广位id（必填）
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetMediaZoneId(mediaZoneId int64) {
	r.AddParameter("mediaZoneId", strconv.FormatInt(mediaZoneId, 10))
}

// SetMediaId 设置媒体id（必填）
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetMediaId(mediaId int64) {
	r.AddParameter("mediaId", strconv.FormatInt(mediaId, 10))
}

// SetExt 设置媒体自定义扩展字段（可选）
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetExt(ext string) {
	r.AddParameter("ext", ext)
}

// SetMaxPrice 设置最大价格（可选）
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetMaxPrice(maxPrice float64) {
	r.AddParameter("maxPrice", strconv.FormatFloat(maxPrice, 'f', 2, 64))
}

// SetMinPrice 设置最小价格（可选）
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) SetMinPrice(minPrice float64) {
	r.AddParameter("minPrice", strconv.FormatFloat(minPrice, 'f', 2, 64))
}

// /////////////////////// 接口信息 ///////////////////////////
// 返回接口名称
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) GetApiName() string {
	return "alibaba.cps.getCpsRecommendSameOfferList"
}

// 返回接口版本
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) GetApiVersion() string {
	return "param2/1"
}

// 返回业务模块
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// 返回请求参数
func (r *AlibabaCpsGetCpsRecommendSameOfferListRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}
