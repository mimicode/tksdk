package jdunionopencoupongiftget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.coupon.gift.get 礼金创建
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_coupon_gift_get_responce"`
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
		//解析getResult
		if t.Responce.GetResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty getResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.GetResultStr), &t.Responce.GetResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.GetResult.Code
				t.ErrorResponse.Message = t.Responce.GetResult.Message
				t.ErrorResponse.RequestID = t.Responce.GetResult.RequestID
			}
		}
	}
	t.Responce.GetResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code         string `json:"code"`
	GetResultStr string `json:"getResult"`
	GetResult    GetResult
}

// GetResult 具体内容
type GetResult struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
	Data      Data   `json:"data"`
}

type Data struct {
	GiftCouponKey string `json:"giftCouponKey"`
}
