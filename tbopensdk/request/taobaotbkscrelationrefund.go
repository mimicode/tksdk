package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.relation.refund( 淘宝客-服务商-维权退款订单查询 )
//https://open.taobao.com/api.htm?docId=43874&docType=2&scopeId=16322
type TbkScRelationRefundRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScRelationRefundRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNumber(tk.Parameters.Get("search_type"), "search_type")
	utils.CheckNumber(tk.Parameters.Get("refund_type"), "refund_type")
	utils.CheckNumber(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNumber(tk.Parameters.Get("biz_type"), "biz_type")
	utils.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")

}

//添加请求参数
func (tk *TbkScRelationRefundRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScRelationRefundRequest) GetApiName() string {
	return "taobao.tbk.sc.relation.refund"
}

//返回请求参数
func (tk *TbkScRelationRefundRequest) GetParameters() url.Values {
	return *tk.Parameters
}
