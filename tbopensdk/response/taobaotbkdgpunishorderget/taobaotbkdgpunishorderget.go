package taobaotbkdgpunishorderget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.punish.order.get( 淘宝客-推广者-处罚订单查询 )
// https://open.taobao.com/api.htm?docId=42050&docType=2&scopeId=15736
type Response struct {
	response.TopResponse
	TbkDgPunishOrderGetResponse TbkDgPunishOrderGetResponse `json:"tbk_dg_punish_order_get_response"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

// TbkDgPunishOrderGetResponse 响应根节点
type TbkDgPunishOrderGetResponse struct {
	Result Result `json:"result"` //查询的对象（官方文档 RpcResult 包装层）
}

// Result 响应结果（官方文档 RpcResult）
type Result struct {
	Data         Data   `json:"data"`           //查询的结果
	BizErrorDesc string `json:"biz_error_desc"` //业务出错的描述
	BizErrorCode int64  `json:"biz_error_code"` //业务出错的状态码
	ResultMsg    string `json:"result_msg"`     //执行结果
	ResultCode   int64  `json:"result_code"`    //执行结果状态码
}

// Data 分页结果（官方文档 PageResult）
type Data struct {
	Results    Results `json:"results"`     //处罚订单列表
	PageNo     int64   `json:"page_no"`     //翻页的pageno
	PageSize   int64   `json:"page_size"`   //翻页的pagesize
	TotalCount int64   `json:"total_count"` //一共能查询出来的结果总数
}

// Results 处罚订单列表
// 官方完整响应示例形状为 {"results":{"result":[...]}}，
// 参照 tbkorderdetailsget(results.publisher_order_dto) 等已上线接口按示例嵌套建模
type Results struct {
	Result []ResultItem `json:"result"` //处罚订单明细集合
}

// ResultItem 处罚订单明细
type ResultItem struct {
	UnionID           string `json:"union_id"`             //淘宝联盟unionid（该字段不再支持）
	SpecialID         int64  `json:"special_id"`           //会员运营id（该字段不再支持）
	RelationID        int64  `json:"relation_id"`          //渠道关系id
	SettleMonth       int64  `json:"settle_month"`         //结算月份
	PunishStatus      string `json:"punish_status"`        //处罚状态。0 冻结，1 解冻
	ViolationType     string `json:"violation_type"`       //处罚类型，目前包括 1.店铺淘宝客 2.订单虚假交易
	TkTradeCreateTime string `json:"tk_trade_create_time"` //淘客订单创建时间
	TbTradeID         int64  `json:"tb_trade_id"`          //子订单号
	TbTradeParentID   int64  `json:"tb_trade_parent_id"`   //父订单号（该字段不再支持）
	TkAdzoneID        int64  `json:"tk_adzone_id"`         //pid里的adzoneid
	TkSiteID          int64  `json:"tk_site_id"`           //pid里的siteid
	TkPubID           string `json:"tk_pub_id"`            //pid里的pubid
}
