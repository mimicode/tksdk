package comvipadpapiopenserviceunionorderv2servicerefundorderlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionOrderV2Service
// refundOrderListWithOauth 2.0.0 获取维权订单列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=refundOrderListWithOauth
type Response struct {
	response.TopResponse
	Success Success `json:"result"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

// Success 维权订单响应结构
type Success struct {
	RefundOrderInfoList []RefundOrderInfo `json:"refundOrderInfoList"` // 联盟维权订单信息
	Total               int               `json:"total"`               // 总记录数
	Page                int               `json:"page"`                // 当前页
	PageSize            int               `json:"pageSize"`            // 分页大小
}

// RefundOrderInfo 维权订单信息
type RefundOrderInfo struct {
	OrderSn               string              `json:"orderSn"`               // 订单号
	ApplyTime             int64               `json:"applyTime"`             // 申请时间:时间戳，单位毫秒
	RefundType            int                 `json:"refundType"`            // 维权类型：1-退货，2-换货，3-价保，4-仅退款，5-换货重发
	CommissionTotalCost   string              `json:"commissionTotalCost"`   // 订单计佣金额(元,两位小数)
	Commission            string              `json:"commission"`            // 订单佣金(元,两位小数)
	GoodsCount            int                 `json:"goodsCount"`            // 商品数量
	CommissionEnterTxn    string              `json:"commissionEnterTxn"`    // 维权返还佣金的入账流水号
	CommissionEnterTime   int64               `json:"commissionEnterTime"`   // 维权返还佣金的入账时间: 时间戳，单位毫秒
	CommissionSettledTime int64               `json:"commissionSettledTime"` // 维权返回佣金的结算时间：时间戳，单位毫秒
	RefundOrderDetails    []RefundOrderDetail `json:"refundOrderDetails"`    // 维权订单明细
	UserId                int64               `json:"userId"`                // 订单归属人id
	OrderTime             int64               `json:"orderTime"`             // 订单时间:时间戳，单位毫秒
	OriginCommEnterTime   int64               `json:"originCommEnterTime"`   // 订单入账时间：发起维权之前入账时间，时间戳，单位毫秒
	OriginCommEnterTxn    string              `json:"originCommEnterTxn"`    // 订单入账流水：发起维权之前入账流水号
	Settled               int                 `json:"settled"`               // 售后订单结算状态：0-未结算，1-已结算
	AfterSaleSn           string              `json:"afterSaleSn"`           // 售后单号
	AfterSaleStatus       int16               `json:"afterSaleStatus"`       // 售后单状态:1-售后中，2-售后完成，3-售后取消
}

// RefundOrderDetail 维权订单明细
type RefundOrderDetail struct {
	GoodsId    string `json:"goodsId"`    // 商品id
	SizeId     string `json:"sizeId"`     // 商品尺码id
	GoodsPrice string `json:"goodsPrice"` // 商品价格：元
	GoodsCount int    `json:"goodsCount"` // 维权商品数量
	Commission string `json:"commission"` // 维权商品佣金
	TotalCost  string `json:"totalCost"`  // 维权商品计佣金额
}
