package jdunionopenactivitybonusquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.activity.bonus.query 奖励活动信息查询接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_activity_bonus_query_responce"`
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
	Data    []BonusActivityResp `json:"data"`
	HasMore bool                `json:"hasMore"` //是否有更多
	Message string              `json:"message"`
}

// BonusActivityResp 奖励活动明细
type BonusActivityResp struct {
	ActivityID    int64   `json:"activityId"`    //活动id
	ActivityName  string  `json:"activityName"`  //活动名称
	PrepareTime   int64   `json:"prepareTime"`   //预热时间，时间戳（ms）
	BeginDate     int64   `json:"beginDate"`     //开始时间，时间戳（ms）
	EndDate       int64   `json:"endDate"`       //结束时间，时间戳（ms）
	PayType       int64   `json:"payType"`       //结算类型：1 一次结算，2 二次结算
	FirstPayTime  int64   `json:"firstPayTime"`  //一次结算时间，时间戳（ms）
	FirstPayRate  float64 `json:"firstPayRate"`  //一次结算比例
	SecondPayTime int64   `json:"secondPayTime"` //二次结算时间，时间戳（ms）
	SecondPayRate float64 `json:"secondPayRate"` //二次计算比例
	PcDescURL     string  `json:"pcDescUrl"`     //活动规则
	Remark        string  `json:"remark"`        //活动概述
}
