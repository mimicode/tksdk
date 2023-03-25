package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbItemPromotionFilterListRequest alibaba.alsc.union.kb.item.promotion.filter.list( 本地生活媒体平台口碑选品筛选项集合 )
// https://open.taobao.com/api.htm?docId=59273&docType=2&scopeId=24408
type AlibabaAlscUnionKbItemPromotionFilterListRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbItemPromotionFilterListRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("filter_type"), "filter_type")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbItemPromotionFilterListRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbItemPromotionFilterListRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.item.promotion.filter.list"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbItemPromotionFilterListRequest) GetParameters() url.Values {
	return *tk.Parameters
}
