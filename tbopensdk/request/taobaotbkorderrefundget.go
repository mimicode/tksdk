package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.order.refund.get( 淘宝客-推广者-全量售后退款订单查询 )
//https://developer.alibaba.com/docs/api.htm?apiId=65180
type TbkOrderRefundGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkOrderRefundGetRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("order_scene"), "order_scene")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNumber(tk.Parameters.Get("query_type"), "query_type")
	utils.CheckNumber(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNotNull(tk.Parameters.Get("start_time"), "start_time")
	utils.CheckNumber(tk.Parameters.Get("member_type"), "member_type")
	utils.CheckNotNull(tk.Parameters.Get("end_time"), "end_time")

	// 可选参数检查
	if tk.Parameters.Get("jump_type") != "" {
		utils.CheckNumber(tk.Parameters.Get("jump_type"), "jump_type")
	}
}

//添加请求参数
func (tk *TbkOrderRefundGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkOrderRefundGetRequest) GetApiName() string {
	return "taobao.tbk.order.refund.get"
}

//返回请求参数
func (tk *TbkOrderRefundGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
