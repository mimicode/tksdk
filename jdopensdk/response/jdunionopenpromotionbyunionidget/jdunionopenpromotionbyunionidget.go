package jdunionopenpromotionbyunionidget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.promotion.byunionid.get 工具商媒体帮助子站长获取普通推广链接和优惠券二合一推广链接
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_promotion_byunionid_get_responce"`
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
	QueryResultStr string `json:"getResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code      int64              `json:"code"`
	Message   string             `json:"message"`
	RequestID string             `json:"requestId"`
	Data      *PromotionCodeResp `json:"data"`
}

type PromotionCodeResp struct {
	ShortURL        string `json:"shortURL"`        //生成的推广目标链接，以短链接形式，有效期60天
	ClickURL        string `json:"clickURL"`        //生成推广目标的长链，长期有效
	JCommand        string `json:"jCommand"`        //需要权限申请，京口令（匹配到红包活动有效配置才会返回京口令）
	JShortCommand   string `json:"jShortCommand"`   //需要权限申请，短口令
	WeChatShortLink string `json:"weChatShortLink"` //微信小程序ShortLink，当weChatType入参1时为京小街ShortLink，当weChatType入参2时为京东购物或京东外卖ShortLink
}
