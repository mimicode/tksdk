package jdunionopenchannelrelationget

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.channel.relation.get 渠道关系ID生成接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_channel_relation_get_responce"`
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
			}
		}
	}
	t.Responce.GetResultStr = ""
}

// Responce 响应结果
type Responce struct {
	GetResultStr string `json:"getResult"`
	GetResult    ChannelRelationGetResult
}

// ChannelRelationGetResult 返回结果
type ChannelRelationGetResult struct {
	Code    int64                  `json:"code"`
	Data    ChannelRelationGetResp `json:"data"`
	Message string                 `json:"message"`
}

// ChannelRelationGetResp 数据明细
type ChannelRelationGetResp struct {
	ChannelID int64 `json:"channelId"` //渠道关系ID
}
