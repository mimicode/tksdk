package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbBbtItemPromotionFilterListRequest alibaba.alsc.union.kb.bbt.item.promotion.filter.list( 本地生活爆爆团选品筛选集合 )
// https://open.taobao.com/api.htm?docId=61957&docType=2&scopeId=24987
type AlibabaAlscUnionKbBbtItemPromotionFilterListRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("filter_type"), "filter_type")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.bbt.item.promotion.filter.list"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
