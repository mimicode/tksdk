package comvipadpapiopenserviceunionorderv2serviceorderlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionOrderV2Service orderListWithOauth 获取订单信息列表-需要oauth授权
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
	Total         int64           `json:"total"`         // 总记录数
	OrderInfoList []OrderInfoList `json:"orderInfoList"` // 订单信息列表
	PageSize      int64           `json:"pageSize"`      // 每页条数
	Page          int64           `json:"page"`          // 页码
}

type OrderInfoList struct {
	OrderSn             string        `json:"orderSn"`             // 订单号
	Status              int           `json:"status"`              // 订单状态
	OrderTime           int64         `json:"orderTime"`           // 下单时间
	PayTime             int64         `json:"payTime"`             // 支付时间
	FinishTime          int64         `json:"finishTime"`          // 完成时间
	OrderAmount         float64       `json:"orderAmount"`         // 订单金额
	Commission          float64       `json:"commission"`          // 佣金金额
	CommissionRate      float64       `json:"commissionRate"`      // 佣金比例
	OrderDetails        []OrderDetail `json:"orderDetails"`        // 订单明细
	UserId              int64         `json:"userId"`              // 用户ID
	Settled             int           `json:"settled"`             // 结算状态：0-未结算，1-已结算
	SettleTime          int64         `json:"settleTime"`          // 结算时间
	CommissionEnterTime int64         `json:"commissionEnterTime"` // 佣金入账时间
	CommissionEnterTxn  string        `json:"commissionEnterTxn"`  // 佣金入账流水号
}

type OrderDetail struct {
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
