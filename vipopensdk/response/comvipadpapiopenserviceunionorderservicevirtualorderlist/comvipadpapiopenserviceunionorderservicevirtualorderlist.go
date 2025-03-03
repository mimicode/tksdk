package comvipadpapiopenserviceunionorderservicevirtualorderlist

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionOrderService virtualOrderList 获取订单信息列表
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
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

type Result struct {
	Total         int64           `json:"total"`         // 总记录数
	PageSize      int64           `json:"pageSize"`      // 每页记录数
	Page          int64           `json:"page"`          // 当前页码
	OrderInfoList []OrderInfoList `json:"orderInfoList"` // 订单列表
}

type OrderInfoList struct {
	OrderSn             string       `json:"orderSn"`             // 订单号
	Status              int64        `json:"status"`              // 订单状态:0-不合格，1-待定，2-已完结
	NewCustomer         int64        `json:"newCustomer"`         // 新老客：0-待定，1-新客，2-老客
	ChannelTag          string       `json:"channelTag"`          // 渠道标识
	OrderTime           int64        `json:"orderTime"`           // 下单时间
	SignTime            int64        `json:"signTime"`            // 签收时间
	SettledTime         int64        `json:"settledTime"`         // 结算时间
	DetailList          []DetailList `json:"detailList"`          // 商品明细
	LastUpdateTime      int64        `json:"lastUpdateTime"`      // 订单最后更新时间
	Settled             int64        `json:"settled"`             // 订单结算状态 0-未结算,1-已结算
	SelfBuy             int64        `json:"selfBuy"`             // 是否自推自买 0-否，1-是
	OrderSubStatusName  string       `json:"orderSubStatusName"`  // 订单子状态
	Commission          string       `json:"commission"`          // 商品总佣金
	CommissionEnterTime int64        `json:"commissionEnterTime"` // 入账时间
	OrderSource         string       `json:"orderSource"`         // 订单来源
	PID                 string       `json:"pid"`                 // 推广PID
	IsPrepay            int64        `json:"isPrepay"`            // 是否预付订单:0-否，1-是
	StatParam           string       `json:"statParam"`           // 自定义统计参数
}

type DetailList struct {
	GoodsID             string          `json:"goodsId"`             // 商品ID
	GoodsName           string          `json:"goodsName"`           // 商品名称
	GoodsThumb          string          `json:"goodsThumb"`          // 商品缩略图
	GoodsCount          int64           `json:"goodsCount"`          // 商品数量
	CommissionTotalCost string          `json:"commissionTotalCost"` // 商品计佣金额
	CommissionRate      string          `json:"commissionRate"`      // 商品佣金比例
	Commission          string          `json:"commission"`          // 商品佣金金额
	CommCode            string          `json:"commCode"`            // 佣金编码
	CommName            string          `json:"commName"`            // 佣金方案名称
	OrderSource         string          `json:"orderSource"`         // 订单来源
	AfterSaleInfo       []AfterSaleInfo `json:"afterSaleInfo"`       // 商品售后信息
	SizeID              string          `json:"sizeId"`              // 商品尺码
	Status              int64           `json:"status"`              // 商品状态
	BrandStoreSn        string          `json:"brandStoreSn"`        // 品牌编号
	BrandStoreName      string          `json:"brandStoreName"`      // 品牌名称
}

type AfterSaleInfo struct {
	AfterSaleChangedCommission string `json:"afterSaleChangedCommission"` // 商品佣金售后变动
	AfterSaleChangedGoodsCount int64  `json:"afterSaleChangedGoodsCount"` // 商品数量售后变动
	AfterSaleSn                string `json:"afterSaleSn"`                // 商品售后单号
	AfterSaleStatus            int64  `json:"afterSaleStatus"`            // 商品售后状态
	AfterSaleType              int64  `json:"afterSaleType"`              // 售后类型
	AfterSaleFinishTime        int64  `json:"afterSaleFinishTime"`        // 售后完成时间
}
