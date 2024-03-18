package jdunionopenstatisticsgiftcouponquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.statistics.giftcoupon.query 礼金效果数据
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_statistics_giftcoupon_query_responce"`
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
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
	Data      []Data `json:"data"`
}
type Data struct {
	Amount             int64          `json:"amount"`
	ContentMatch       int64          `json:"contentMatch"`
	ContentMatchMedias []int64        `json:"contentMatchMedias"`
	CostAmount         float64        `json:"costAmount"`
	CostNum            int64          `json:"costNum"`
	CouponTitle        string         `json:"couponTitle"`
	Denomination       float64        `json:"denomination"`
	EffectiveDays      int64          `json:"effectiveDays"`
	ExpireType         int64          `json:"expireType"`
	GiftCouponKey      string         `json:"giftCouponKey"`
	PromoterList       []PromoterList `json:"promoterList"`
	ReceiveEndTime     string         `json:"receiveEndTime"`
	ReceiveNum         int64          `json:"receiveNum"`
	ReceiveStartTime   string         `json:"receiveStartTime"`
	Share              int64          `json:"share"`
	ShowInMedias       int64          `json:"showInMedias"`
	ShowStatus         int64          `json:"showStatus"`
	SkuIDList          []int64        `json:"skuIdList"`
	ItemIdList         []string       `json:"itemIdList"`
	Status             int64          `json:"status"`
	Type               int64          `json:"type"`
	UseEndTime         string         `json:"useEndTime"`
	UseStartTime       string         `json:"useStartTime"`
	YgCommission       float64        `json:"ygCommission"`
	LimitPlatforms     []int64        `json:"limitPlatforms"`
}

type PromoterList struct {
	Amount  int64  `json:"amount"`
	PID     string `json:"pid"`
	UnionID int64  `json:"unionId"`
}
