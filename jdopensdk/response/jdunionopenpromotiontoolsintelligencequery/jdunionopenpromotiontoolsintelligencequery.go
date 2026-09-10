package jdunionopenpromotiontoolsintelligencequery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.promotion.tools.intelligence.query 工具商线报推广【申请】：获取京东商品段子、优惠信息、店铺及类目活动爆料信息，接口出参直接是转链后链接
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_promotion_tools_intelligence_query_responce"`
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
	Code    int64              `json:"code"`
	Data    []IntelligenceResp `json:"data"`
	Message string             `json:"message"`
}

// IntelligenceResp 线报信息
type IntelligenceResp struct {
	ReportContent string  `json:"reportContent"` //线报内容
	Type          int64   `json:"type"`          //线报类型：1 单品；2 店铺活动；3 品类活动；4 其它
	Cid1List      []int64 `json:"cid1List"`      //类目信息
	Status        int64   `json:"status"`        //1 未开始、2 进行中
	Essence       int64   `json:"essence"`       //1 精华，0 非精华
	StartTime     string  `json:"startTime"`     //线报开始时间，格式：yyyy-MM-dd HH:mm:ss
	EndTime       string  `json:"endTime"`       //线报结束时间，格式：yyyy-MM-dd HH:mm:ss
}
