package comvipadpapiopenserviceunionorderv2servicevirtualorderlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionOrderV2Service
// virtualOrderListWithOauth 2.0.0 获取虚拟订单列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=virtualOrderListWithOauth
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

// Success 虚拟订单响应结构
type Success struct {
	OrderInfoList []VirtualOrderInfo `json:"orderInfoList"` // 业绩查询响应结果
	Total         int                `json:"total"`         // 业绩总条数
	Page          int                `json:"page"`          // 当前页码
	PageSize      int                `json:"pageSize"`      // 页面大小
}

// VirtualOrderInfo 虚拟订单信息
type VirtualOrderInfo struct {
	OrderSn                   string              `json:"orderSn"`                   // 订单号
	Status                    int16               `json:"status"`                    // 订单状态:0-不合格，1-待定，2-已完结
	ChannelTag                string              `json:"channelTag"`                // channelTag=chanTag=pid，即推广位标识
	OrderTime                 int64               `json:"orderTime"`                 // 下单时间 时间戳 单位毫秒
	SettledTime               int64               `json:"settledTime"`               // 结算时间 时间戳 单位毫秒
	GoodsId                   string              `json:"goodsId"`                   // 商品id
	GoodsName                 string              `json:"goodsName"`                 // 商品名称
	GoodsThumb                string              `json:"goodsThumb"`                // 商品缩略图
	GoodsType                 int                 `json:"goodsType"`                 // 商品类型304:超级VIP,302:电子唯品卡,700:MP通用卡券
	LastUpdateTime            int64               `json:"lastUpdateTime"`            // 订单上次更新时间 时间戳 单位毫秒
	Settled                   int16               `json:"settled"`                   // 订单结算状态 0-未结算,1-已结算
	SelfBuy                   int                 `json:"selfBuy"`                   // 是否自推自买 0-否，1-是
	OrderSubStatusName        string              `json:"orderSubStatusName"`        // 订单子状态
	CommissionTotalCost       string              `json:"commissionTotalCost"`       // 订单计佣金额(元,两位小数)
	Commission                string              `json:"commission"`                // 订单佣金(元,两位小数)
	CommissionRate            string              `json:"commissionRate"`            // 订单佣金比例(%)
	CommissionEnterTime       int64               `json:"commissionEnterTime"`       // 入账时间，时间戳，单位毫秒
	OrderSource               string              `json:"orderSource"`               // 订单来源
	TotalCost                 string              `json:"totalCost"`                 // 订单支付金额:单位元
	AfterSaleChangeCommission string              `json:"afterSaleChangeCommission"` // 售后订单佣金变动
	Pid                       string              `json:"pid"`                       // 推广PID
	StatParam                 string              `json:"statParam"`                 // 自定义统计参数
	OrderTrackReason          int                 `json:"orderTrackReason"`          // 订单归因方式
	AppKey                    string              `json:"appKey"`                    // 开发者调用的appKey
	OpenId                    string              `json:"openId"`                    // 渠道用户唯一标识
	AdCode                    string              `json:"adCode"`                    // 标识获取推广物料的来源
	RefundOrderDetails        []RefundOrderDetail `json:"refundOrderDetails"`        // 维权订单明细
	FdsBlackOrder             int                 `json:"fdsBlackOrder"`             // 是否异常订单,1：是,0：否
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
