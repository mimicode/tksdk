package request

import (
	"net/url"
	"strconv"
)

// alibaba.cps.listMediaInfo( 获取媒体推广位信息 )
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

// SetMediaId 设置媒体id
func (r *AlibabaCpsListMediaInfoRequest) SetMediaId(mediaId int64) {
	r.AddParameter("mediaId", strconv.FormatInt(mediaId, 10))
}

// 返回接口名称
func (r *AlibabaCpsListMediaInfoRequest) GetApiName() string {
	return "alibaba.cps.listMediaInfo"
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