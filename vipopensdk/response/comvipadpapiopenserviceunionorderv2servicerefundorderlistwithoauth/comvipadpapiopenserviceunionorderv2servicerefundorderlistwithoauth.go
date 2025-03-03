package comvipadpapiopenserviceunionorderv2servicerefundorderlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionOrderV2Service refundOrderListWithOauth 获取维权订单列表，订单归属人为授权用户
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
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

type Success struct {
	Code int    `json:"code"` // 状态码，1-成功，0-失败
	Msg  string `json:"msg"`  // 状态信息
	Data Data   `json:"data"` // 数据
}

type Data struct {
	Total               int64                 `json:"total"`               // 总记录数
	RefundOrderInfoList []RefundOrderInfoList `json:"refundOrderInfoList"` // 退款订单信息列表
	PageSize            int64                 `json:"pageSize"`            // 每页条数
	Page                int64                 `json:"page"`                // 页码
}

type RefundOrderInfoList struct {
	OrderSn               string              `json:"orderSn"`               // 订单号
	ApplyTime             int64               `json:"applyTime"`             // 申请时间
	RefundType            int                 `json:"refundType"`            // 维权类型：1-退货，2-换货
	CommissionTotalCost   float64             `json:"commissionTotalCost"`   // 订单计佣金额
	Commission            float64             `json:"commission"`            // 订单佣金
	GoodsCount            int                 `json:"goodsCount"`            // 商品数量
	CommissionEnterTxn    string              `json:"commissionEnterTxn"`    // 维权返还佣金的入账流水号
	CommissionEnterTime   int64               `json:"commissionEnterTime"`   // 维权返还佣金的入账时间
	CommissionSettledTime int64               `json:"commissionSettledTime"` // 维权返回佣金的结算时间
	RefundOrderDetails    []RefundOrderDetail `json:"refundOrderDetails"`    // 维权订单明细
	UserId                int64               `json:"userId"`                // 订单归属人id
	OrderTime             int64               `json:"orderTime"`             // 订单时间
	OriginCommEnterTime   int64               `json:"originCommEnterTime"`   // 订单入账时间
	OriginCommEnterTxn    string              `json:"originCommEnterTxn"`    // 订单入账流水号
	Settled               int                 `json:"settled"`               // 售后订单结算状态：0-未结算，1-已结算
	AfterSaleSn           string              `json:"afterSaleSn"`           // 售后单号
	AfterSaleStatus       int16               `json:"afterSaleStatus"`       // 售后单状态:1-售后中，2-售后完成，3-售后取消
}

type RefundOrderDetail struct {
	GoodsId        string  `json:"goodsId"`        // 商品ID
	GoodsName      string  `json:"goodsName"`      // 商品名称
	GoodsThumbUrl  string  `json:"goodsThumbUrl"`  // 商品缩略图
	GoodsCount     int     `json:"goodsCount"`     // 商品数量
	Commission     float64 `json:"commission"`     // 佣金金额
	CommissionRate float64 `json:"commissionRate"` // 佣金比例
	Price          float64 `json:"price"`          // 商品价格
	SkuId          string  `json:"skuId"`          // SKU ID
	SkuName        string  `json:"skuName"`        // SKU名称
	BrandStoreSn   string  `json:"brandStoreSn"`   // 品牌店铺SN
	BrandName      string  `json:"brandName"`      // 品牌名称
}
