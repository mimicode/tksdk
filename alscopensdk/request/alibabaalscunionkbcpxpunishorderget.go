package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbcpxPunishOrderGetRequest alibaba.alsc.union.kbcpx.punish.order.get( 本地生活媒体推广用户反作弊订单数据报表 )
// https://open.taobao.com/api.htm?docId=59451&docType=2&scopeId=24408
type AlibabaAlscUnionKbcpxPunishOrderGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbcpxPunishOrderGetRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("date_type"), "date_type")
	utils.CheckNumber(tk.Parameters.Get("biz_unit"), "biz_unit")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNumber(tk.Parameters.Get("page_number"), "page_number")
	utils.CheckNumber(tk.Parameters.Get("start_date"), "start_date")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbcpxPunishOrderGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbcpxPunishOrderGetRequest) GetApiName() string {
	return "alibaba.alsc.union.kbcpx.punish.order.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbcpxPunishOrderGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
