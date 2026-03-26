package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.relation.refund( 淘宝客-推广者-售后退款订单查询 )
//https://developer.alibaba.com/docs/api.htm?apiId=40121
type TbkRelationRefundRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkRelationRefundRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNumber(tk.Parameters.Get("search_type"), "search_type")
	utils.CheckNumber(tk.Parameters.Get("refund_type"), "refund_type")
	utils.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")
	utils.CheckNumber(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNumber(tk.Parameters.Get("biz_type"), "biz_type")
}

//添加请求参数
func (tk *TbkRelationRefundRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkRelationRefundRequest) GetApiName() string {
	return "taobao.tbk.relation.refund"
}

//返回请求参数
func (tk *TbkRelationRefundRequest) GetParameters() url.Values {
	return *tk.Parameters
}
