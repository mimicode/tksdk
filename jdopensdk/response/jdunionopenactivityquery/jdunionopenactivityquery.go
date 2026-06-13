package jdunionopenactivityquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.activity.query 提供联盟官方活动查询
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_activity_query_responce"`
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
	Code       int64          `json:"code"`
	Data       []ActivityResp `json:"data"`
	Message    string         `json:"message"`
	RequestID  string         `json:"requestId"`
	TotalCount int64          `json:"totalCount"`
}

// ActivityResp 活动详情（data 数组元素）
type ActivityResp struct {
	ActStatus          int64              `json:"actStatus"`          // 活动状态，1：未开始；2：进行中；3：已结束
	URLPC              string             `json:"urlPC"`              // 活动落地页-PC
	URLM               string             `json:"urlM"`               // 活动落地页-移动端
	PromotionStartTime int64              `json:"promotionStartTime"` // 主推开始时间
	PromotionEndTime   int64              `json:"promotionEndTime"`   // 主推结束时间
	Advantage          string             `json:"advantage"`          // 活动利益点
	DownloadURL        string             `json:"downloadUrl"`        // 活动素材下载链接
	DownloadCode       string             `json:"downloadCode"`       // 活动素材提取码
	PlatformType       int64              `json:"platformType"`       // 活动推广平台，1：仅支持PC推广；2：仅支持移动端推广；3：PC和移动端均支持推广
	Recommend          int64              `json:"recommend"`          // 推荐指数，从1到5，越高越好
	UpdateTime         int64              `json:"updateTime"`         // 活动更新时间
	Title              string             `json:"title"`              // 活动主题
	Content            string             `json:"content"`            // 活动规则、描述、说明
	EliteID            int64              `json:"eliteId"`            // 商品资源位id，调用京粉精选商品查询接口传入此参数可获取活动SKU
	ID                 int64              `json:"id"`                 // 活动ID
	StartTime          int64              `json:"startTime"`          // 活动开始时间
	EndTime            int64              `json:"endTime"`            // 活动结束时间
	Tag                string             `json:"tag"`                // 活动标签，1：大促活动；2：佣金提升活动；3：常规活动
	CategoryList       []ActivityCategory `json:"categoryList"`       // 类目集
	ImgList            []Image            `json:"imgList"`            // 图片集
}

// ActivityCategory 类目（categoryList 数组元素）
type ActivityCategory struct {
	CategoryID int64 `json:"categoryId"` // 类目Id
	Type       int64 `json:"type"`       // 类目级别（当前京挑客活动全部为一级类目）
}

// Image 图片对象（imgList 数组元素）
type Image struct {
	ImgName     string `json:"imgName"`     // 图片名称
	ImgURL      string `json:"imgUrl"`      // 图片链接
	WidthHeight string `json:"widthHeight"` // 图片尺寸
}
