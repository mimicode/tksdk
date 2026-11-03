package jdunionopenorderbonusquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.order.bonus.query 奖励订单查询接口
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_order_bonus_query_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析queryResult
		if t.Responce.QueryResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty queryResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.QueryResultStr), &t.Responce.QueryResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.QueryResult.Code
				t.ErrorResponse.Message = t.Responce.QueryResult.Message
				t.ErrorResponse.RequestID = t.Responce.QueryResult.RequestID
			}
		}
	}
	t.Responce.QueryResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"queryResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code      int64            `json:"code"`
	Message   string           `json:"message"`
	RequestID string           `json:"requestId"`
	Data      []BonusOrderResp `json:"data"`
}

// BonusOrderResp 奖励订单明细
type BonusOrderResp struct {
	ID               int64   `json:"id"`               // 奖励订单ID
	OrderID          int64   `json:"orderId"`          // 订单号
	ParentID         int64   `json:"parentId"`         // 主单的订单号：如一个订单拆成多个子订单时，原订单号会作为主单号，拆分的订单号为子单号存储在orderid中。若未发生拆单，该字段为0
	UnionID          int64   `json:"unionId"`          // 联盟ID
	PositionID       int64   `json:"positionId"`       // 推广位ID
	SkuID            int64   `json:"skuId"`            // skuId
	SkuName          string  `json:"skuName"`          // 商品名称
	ItemID           string  `json:"itemId"`           // 联盟商品ID
	PayPrice         float64 `json:"payPrice"`         // 实际支付金额
	EstimateCosPrice float64 `json:"estimateCosPrice"` // 预估计佣金额
	EstimateFee      float64 `json:"estimateFee"`      // 预估佣金
	ActualCosPrice   float64 `json:"actualCosPrice"`   // 实际计佣金额
	ActualFee        float64 `json:"actualFee"`        // 实际佣金
	EstimateBonusFee float64 `json:"estimateBonusFee"` // 预估奖励金额：查询时间范围内，已付款且奖励有效，满足奖励规则的奖励金额
	ActualBonusFee   float64 `json:"actualBonusFee"`   // 实际奖励金额：查询时间范围内，已付款或已完成（视具体规则），奖励有效且满足奖励规则的奖励金额
	CommissionRate   float64 `json:"commissionRate"`   // 佣金比例，单位：%
	SubSideRate      float64 `json:"subSideRate"`      // 分成比例（单位：%）
	SubsidyRate      float64 `json:"subsidyRate"`      // 补贴比例（单位：%）
	FinalRate        float64 `json:"finalRate"`        // 最终分佣比例（单位：%）=分成比例+补贴比例
	BonusState       int64   `json:"bonusState"`       // 奖励状态，0:无效、1:有效
	BonusText        string  `json:"bonusText"`        // 奖励状态文案
	BonusInvalidCode string  `json:"bonusInvalidCode"` // 无效状态码，-1:无效、2:无效-拆单、3:无效-取消、4:无效-京东帮帮主订单、5:无效-账号异常、6:无效-赠品类目不返佣 等
	BonusInvalidText string  `json:"bonusInvalidText"` // 无效状态码对应的无效状态文案
	OrderState       int64   `json:"orderState"`       // 奖励订单状态，1:已完成、2:已付款、3:待付款
	OrderText        string  `json:"orderText"`        // 奖励订单状态，待付款/已付款/已完成
	OrderTime        int64   `json:"orderTime"`        // 下单时间、时间戳（毫秒）
	FinishTime       int64   `json:"finishTime"`       // 完成时间、时间戳（毫秒）
	SortValue        string  `json:"sortValue"`        // 排序值，按'下单时间'分页查询时使用
	ActivityID       int64   `json:"activityId"`       // 奖励活动ID
	ActivityName     string  `json:"activityName"`     // 活动名称
	ChannelID        int64   `json:"channelId"`        // 渠道关系ID
	SubUnionID       string  `json:"subUnionId"`       // 子联盟ID
	Pid              string  `json:"pid"`              // pid
	Ext1             string  `json:"ext1"`             // 推客生成推广链接时传入的扩展字段
	UnionAlias       string  `json:"unionAlias"`       // 母站长简称
	Account          string  `json:"account"`          // 账户ID
}
