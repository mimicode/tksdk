package request

import (
	"net/url"
	"strconv"
	"strings"
)

// AlibabaCpsGenerateIntroduceUrlByParamRequest 批量生成推广链接服务(部分条件) API请求
// https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceFVq7hm&id=com.alibaba.p4p%3Aalibaba.cps.generateIntroduceUrlByParam-1&aopApiCategory=category_new
type AlibabaCpsGenerateIntroduceUrlByParamRequest struct {
	Parameters *url.Values //请求参数
}

// NewAlibabaCpsGenerateIntroduceUrlByParamRequest 创建AlibabaCpsGenerateIntroduceUrlByParamRequest实例
func NewAlibabaCpsGenerateIntroduceUrlByParamRequest() *AlibabaCpsGenerateIntroduceUrlByParamRequest {
	return &AlibabaCpsGenerateIntroduceUrlByParamRequest{}
}

// CheckParameters 检查参数
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) CheckParameters() {
	if r.Parameters == nil {
		panic("Parameters is nil")
	}
	if r.Parameters.Get("mediaId") == "" {
		panic("mediaId is required")
	}
	if r.Parameters.Get("mediaZoneId") == "" {
		panic("mediaZoneId is required")
	}
	// offerIds和activityId至少需要存在一个
	if r.Parameters.Get("offerIds") == "" && r.Parameters.Get("activityId") == "" {
		panic("offerIds and activityId exist at least one")
	}
}

// AddParameter 添加参数
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// SetMediaId 设置媒体id
// 示例值: 1
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) SetMediaId(mediaId int64) {
	r.AddParameter("mediaId", strconv.FormatInt(mediaId, 10))
}

// SetMediaZoneId 设置媒体推广位id
// 示例值: 1
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) SetMediaZoneId(mediaZoneId int64) {
	r.AddParameter("mediaZoneId", strconv.FormatInt(mediaZoneId, 10))
}

// SetOfferIds 设置商品id，最大100个，且只支持生成长链接，否则请用genClickUrl接口
// 示例值: "1111,22222,3333"
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) SetOfferIds(offerIds string) {
	r.AddParameter("offerIds", offerIds)
}

// SetOfferIdsFromSlice 设置商品id数组，最大100个，且只支持生成长链接，否则请用genClickUrl接口
// 示例值: [1111,22222,3333]
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) SetOfferIdsFromSlice(offerIds []int64) {
	var strIds []string
	for _, id := range offerIds {
		strIds = append(strIds, strconv.FormatInt(id, 10))
	}
	r.AddParameter("offerIds", strings.Join(strIds, ","))
}

// SetActivityId 设置活动id
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) SetActivityId(activityId int64) {
	r.AddParameter("activityId", strconv.FormatInt(activityId, 10))
}

// SetExt 设置其他自定义参数，查询订单会返回该参数
// 示例值: {"p1":"123","p2":"456","p3":"789"}
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) SetExt(ext string) {
	r.AddParameter("ext", ext)
}

// GetApiName 获取API名称
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) GetApiName() string {
	return "alibaba.cps.generateIntroduceUrlByParam"
}

// GetApiVersion 获取API版本
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) GetBusinessModule() string {
	return "com.alibaba.p4p"
}

// GetParameters 获取所有参数
func (r *AlibabaCpsGenerateIntroduceUrlByParamRequest) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}