package jdunionopenstatisticsactivitybonusquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.statistics.activity.bonus.query 奖励活动奖励金额查询接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_statistics_activity_bonus_query_responce"`
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
	Code    int64               `json:"code"`
	Data    BonusEffectDataResp `json:"data"`
	Message string              `json:"message"`
}

// BonusEffectDataResp 结果对象
type BonusEffectDataResp struct {
	UnionID          int64   `json:"unionId"`          //站长ID
	ActivityID       int64   `json:"activityId"`       //活动ID
	EstimateValidNum int64   `json:"estimateValidNum"` //预估有效订单数
	EstimateCosPrice float64 `json:"estimateCosPrice"` //预估计佣金额(GMV)
	EstimateBonus    float64 `json:"estimateBonus"`    //预估奖励金额
	ActualValidNum   int64   `json:"actualValidNum"`   //实际有效订单数
	ActualCosPrice   float64 `json:"actualCosPrice"`   //实际计佣金额(GMV)
	ActualBonus      float64 `json:"actualBonus"`      //实际奖励金额
	UV               int64   `json:"uv"`               //结算uv
	ChannelID        int64   `json:"channelId"`        //配合入参channelId字段使用，若未入参channelId，则不展示该字段数值
}
