package request

import (
	"net/url"
	"strconv"
)

// AlibabaCpsGenClickUrlRequest 批量生成联盟推广url点击信息 API请求
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472ccel58TEV&id=com.alibaba.p4p%3Aalibaba.cps.genClickUrl-1&aopApiCategory=category_new
type AlibabaCpsGenClickUrlRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsGenClickUrlRequest 创建AlibabaCpsGenClickUrlRequest实例
func NewAlibabaCpsGenClickUrlRequest() *AlibabaCpsGenClickUrlRequest {
	return &AlibabaCpsGenClickUrlRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsGenClickUrlRequest) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil {
		panic("parameters is required")
	}
	if r.Parameters.Get("type") == "" {
		panic("type is required")
	}
	if r.Parameters.Get("mediaId") == "" {
		panic("mediaId is required")
	}
	if r.Parameters.Get("mediaZoneId") == "" {
		panic("mediaZoneId is required")
	}
	if r.Parameters.Get("objectValueList") == "" {
		panic("objectValueList is required")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsGenClickUrlRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetType 设置转链接类型(店铺、商品、活动) 0商品;1店铺;2活动;7优惠券; 12 图搜
func (r *AlibabaCpsGenClickUrlRequest) SetType(param int) {
	r.AddParameter("type", strconv.Itoa(param))
}

// SetMediaId 设置媒体id 示例值: 1
func (r *AlibabaCpsGenClickUrlRequest) SetMediaId(param int64) {
	r.AddParameter("mediaId", strconv.FormatInt(param, 10))
}

// SetMediaZoneId 设置媒体推广位id 示例值: 1
func (r *AlibabaCpsGenClickUrlRequest) SetMediaZoneId(param int64) {
	r.AddParameter("mediaZoneId", strconv.FormatInt(param, 10))
}

// SetObjectValueList 设置推广实体 商品：传offerId；活动：传活动id；商铺：卖家sellerId；优惠券：offerId ; 图搜传：图片url
func (r *AlibabaCpsGenClickUrlRequest) SetObjectValueList(param string) {
	r.AddParameter("objectValueList", param)
}

// SetExt 设置其他自定义参数，查询订单会返回该参数 示例值: {"p1":"123","p2":"456","p3":"789"}
func (r *AlibabaCpsGenClickUrlRequest) SetExt(param string) {
	r.AddParameter("ext", param)
}

// GetApiName 获取API名称
func (r *AlibabaCpsGenClickUrlRequest) GetApiName() string {
	return "alibaba.cps.genClickUrl"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsGenClickUrlRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsGenClickUrlRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsGenClickUrlRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}