package request

import (
	"net/url"
	"strconv"
)

// alibaba.cps.listMediaInfo( 获取媒体推广位信息 )
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceEddODt&id=com.alibaba.p4p%3Aalibaba.cps.listMediaInfo-1&aopApiCategory=category_new
type AlibabaCpsListMediaInfoRequest struct {
	Parameters *url.Values //请求参数
}

func (r *AlibabaCpsListMediaInfoRequest) CheckParameters() {
	// 该API无必填参数
}

// 添加请求参数
func (r *AlibabaCpsListMediaInfoRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// /////////////////////// 应用级参数 ///////////////////////////
// SetMediaId 设置媒体id
func (r *AlibabaCpsListMediaInfoRequest) SetMediaId(mediaId int64) {
	r.AddParameter("mediaId", strconv.FormatInt(mediaId, 10))
}

// /////////////////////// 接口信息 ///////////////////////////
// 返回接口名称
func (r *AlibabaCpsListMediaInfoRequest) GetApiName() string {
	return "alibaba.cps.listMediaInfo"
}

// 返回接口版本
func (r *AlibabaCpsListMediaInfoRequest) GetApiVersion() string {
	return "param2/1"
}

// 返回业务模块
func (r *AlibabaCpsListMediaInfoRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// 返回请求参数
func (r *AlibabaCpsListMediaInfoRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}
