package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbItemPromotionShareCreateRequest alibaba.alsc.union.kb.item.promotion.share.create( 本地生活媒体创建商品推广链接 )
// https://open.taobao.com/api.htm?docId=59620&docType=2&scopeId=24408
type AlibabaAlscUnionKbItemPromotionShareCreateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbItemPromotionShareCreateRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("pid"), "pid")
	utils.CheckNotNull(tk.Parameters.Get("item_id"), "item_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbItemPromotionShareCreateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbItemPromotionShareCreateRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.item.promotion.share.create"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbItemPromotionShareCreateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
