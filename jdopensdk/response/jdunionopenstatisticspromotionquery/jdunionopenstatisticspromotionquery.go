package jdunionopenstatisticspromotionquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.statistics.promotion.query 推广效果数据查询接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_statistics_promotion_query_responce"`
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
	Code    int64                     `json:"code"`
	Data    []PromotionEffectDataResp `json:"data"`
	Message string                    `json:"message"`
}

// PromotionEffectDataResp 推广效果数据明细
type PromotionEffectDataResp struct {
	UnionID             int64   `json:"unionId"`             //联盟ID
	SkuID               int64   `json:"skuId"`               //商品ID
	ActivityURL         string  `json:"activityUrl"`         //活动链接
	TimeType            int64   `json:"timeType"`            //数据的日期范围（1：今日实时）
	DataType            int64   `json:"dataType"`            //返回的数据类型（1：汇总数据）
	Time                string  `json:"time"`                //数据时间，timeType=1且dataType=1时为当日0点至当前小时的汇总数据
	ClickPv             int64   `json:"clickPv"`             //点击量
	EstimateValidOrders int64   `json:"estimateValidOrders"` //有效订单量
	EstimateValidFee    float64 `json:"estimateValidFee"`    //预估收入
	EstimateValidGmv    float64 `json:"estimateValidGmv"`    //有效订单金额
	RefundOrders        int64   `json:"refundOrders"`        //退款订单量：当日下单付款后又取消的订单量
	CompleteOrders      int64   `json:"completeOrders"`      //完成订单量
	CompleteGmv         float64 `json:"completeGmv"`         //完成订单金额
	ActualFee           float64 `json:"actualFee"`           //实际收入
	ItemID              string  `json:"itemId"`              //联盟商品ID
}
