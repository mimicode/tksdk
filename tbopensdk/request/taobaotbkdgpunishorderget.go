package request

import (
	"net/url"
)

// taobao.tbk.dg.punish.order.get( 淘宝客-推广者-处罚订单查询 )
// https://open.taobao.com/api.htm?docId=42050&docType=2&scopeId=15736
type TbkDgPunishOrderGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkDgPunishOrderGetRequest) CheckParameters() {
	//业务参数为对象 af_order_option，以JSON字符串传递，键名与官方文档业务参数表一致：
	//relation_id、tb_trade_id、page_size、page_no、span、start_time、special_id(已废弃)、
	//adzone_id、site_id、tb_trade_parent_id(已废弃)、violation_type(已废弃)、punish_status(已废弃)
	//官方文档标注上述子参数均为可选
}

// 添加请求参数
func (tk *TbkDgPunishOrderGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称
func (tk *TbkDgPunishOrderGetRequest) GetApiName() string {
	return "taobao.tbk.dg.punish.order.get"
}

// 返回请求参数
func (tk *TbkDgPunishOrderGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
