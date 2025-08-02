package request

import (
	"net/url"
	"strconv"
)

// alibaba.cps.genSearchPjjxIntroduceUrlByKeyword( 生成平价精选搜索页的推广链接 )
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceEddODt&id=com.alibaba.p4p%3Aalibaba.cps.genSearchPjjxIntroduceUrlByKeyword-1&aopApiCategory=category_new
type AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest struct {
	Parameters *url.Values //请求参数
}

func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) CheckParameters() {
	// 检查必填参数
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	// keyword, mediaId, mediaZoneId 为必填参数，在调用时需要设置
}

// 添加请求参数
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// /////////////////////// 应用级参数 ///////////////////////////
// SetKeyword 设置搜索关键词（必填）
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) SetKeyword(keyword string) {
	r.AddParameter("keyword", keyword)
}

// SetMediaId 设置媒体id（必填）
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) SetMediaId(mediaId int64) {
	r.AddParameter("mediaId", strconv.FormatInt(mediaId, 10))
}

// SetMediaZoneId 设置媒体推广位id（必填）
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) SetMediaZoneId(mediaZoneId int64) {
	r.AddParameter("mediaZoneId", strconv.FormatInt(mediaZoneId, 10))
}

// SetExt 设置其他自定义参数（可选）
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) SetExt(ext string) {
	r.AddParameter("ext", ext)
}

// /////////////////////// 接口信息 ///////////////////////////
// 返回接口名称
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) GetApiName() string {
	return "alibaba.cps.genSearchPjjxIntroduceUrlByKeyword"
}

// 返回接口版本
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) GetApiVersion() string {
	return "param2/1"
}

// 返回业务模块
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// 返回请求参数
func (r *AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}
