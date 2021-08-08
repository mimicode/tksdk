package comvipadpapiopenserviceunionorderservicerefundorderlistwithoauth

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/vipopensdk/response"
)

//com.vip.adp.api.open.service.UnionOrderService 维权订单总收益需要扣除该接口返回-需要oauth授权
type Response struct {
	response2.TopResponse
	Result Result `json:"result"`
}

//解析输出结果
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

type Result struct {
	Total               int64                 `json:"total"`
	RefundOrderInfoList []RefundOrderInfoList `json:"refundOrderInfoList"`
	PageSize            int64                 `json:"pageSize"`
	Page                int64                 `json:"page"`
}

/*
	orderSn	String	否			订单号
	applyTime	Long	否			申请时间:时间戳，单位毫秒
	refundType	Integer	否			维权类型：1-退货，2-换货
	commissionTotalCost	String	否			订单计佣金额(元,两位小数)
	commission	String	否			订单佣金(元,两位小数)
	goodsCount	Integer	否			商品数量
	commissionEnterTxn	String	否			维权返还佣金的入账流水号
	commissionEnterTime	Long	否			维权返还佣金的入账时间: 时间戳，单位毫秒
	commissionSettledTime	Long	否			维权返回佣金的结算时间：时间戳，单位毫秒
	refundOrderDetails	List<RefundOrderDetail>	否			维权订单明细
	userId	Long	否			订单归属人id
	orderTime	Long	否			订单时间:时间戳，单位毫秒
	originCommEnterTime	Long	否			订单入账时间：发起维权之前入账时间，时间戳，单位毫秒
	originCommEnterTxn	String	否			订单入账流水：发起维权之前入账流水号
	settled	Integer	否			售后订单结算状态：0-未结算，1-已结算
	afterSaleSn	String	否			售后单号
	afterSaleStatus	Short	否			售后单状态:1-售后中，2-售后完成，3-售后取消

*/
type RefundOrderInfoList struct {
	OrderSn               string              `json:"orderSn"`
	ApplyTime             int64               `json:"applyTime"`
	RefundType            int64               `json:"refundType"`
	CommissionTotalCost   string              `json:"commissionTotalCost"`
	Commission            string              `json:"commission"`
	GoodsCount            int64               `json:"goodsCount"`
	CommissionEnterTxn    string              `json:"commissionEnterTxn"`
	CommissionEnterTime   int64               `json:"commissionEnterTime"`
	CommissionSettledTime int64               `json:"commissionSettledTime"`
	RefundOrderDetails    []RefundOrderDetail `json:"refundOrderDetails"`
	UserId                int64               `json:"userId"`
	OrderTime             int64               `json:"orderTime"`
	OriginCommEnterTime   string              `json:"originCommEnterTime"`
	OriginCommEnterTxn    int64               `json:"originCommEnterTxn"`
	Settled               int64               `json:"settled"`
	AfterSaleSn           string              `json:"afterSaleSn"`
	AfterSaleStatus       int64               `json:"afterSaleStatus"`
}

/*
	goodsId	String	否			商品id
	sizeId	String	否			商品尺码id
	goodsPrice	String	否			商品价格：元
	goodsCount	Integer	否			维权商品数量
	commission	String	否			维权商品佣金
	totalCost	String	否			维权商品计佣金额
*/
type RefundOrderDetail struct {
	GoodsID    string `json:"goodsId"`
	SizeID     string `json:"sizeId"`
	GoodsPrice string `json:"goodsPrice"`
	GoodsCount int64  `json:"goodsCount"`
	Commission string `json:"commission"`
	TotalCost  string `json:"totalCost"`
}
