package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbItemPromotionRequest alibaba.alsc.union.kb.item.promotion( 本地生活媒体平台口碑选品 )
// https://open.taobao.com/api.htm?docId=58766&docType=2&scopeId=24408
type AlibabaAlscUnionKbItemPromotionRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbItemPromotionRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page_number"), "page_number")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils.CheckNumber(tk.Parameters.Get("settle_type"), "settle_type")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbItemPromotionRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbItemPromotionRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.item.promotion"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbItemPromotionRequest) GetParameters() url.Values {
	return *tk.Parameters
}
