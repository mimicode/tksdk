package jdunionopengoodsrecommendquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.recommend.query 相似品推荐接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_recommend_query_responce"`
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
	Data    RecommendGoodsResp `json:"data"`
	Message string             `json:"message"`
}

// RecommendGoodsResp 数据明细
type RecommendGoodsResp struct {
	MatchSkuInfo         MatchSkuInfo       `json:"matchSkuInfo"`         //精准匹配品信息
	RecommendSkuInfoList []RecommendSkuInfo `json:"recommendSkuInfoList"` //推荐品列表信息
}

// MatchSkuInfo 精准匹配品信息
type MatchSkuInfo struct {
	SkuID  int64  `json:"skuId"`  //商品ID
	ItemID string `json:"itemId"` //联盟商品ID
}

// RecommendSkuInfo 推荐品信息
type RecommendSkuInfo struct {
	SkuID   int64  `json:"skuId"`   //商品ID
	ItemID  string `json:"itemId"`  //联盟商品ID
	Type    int64  `json:"type"`    //分类，2：同款推荐，3：相似推荐，4：推荐商品，5：类目榜单，6：其他推荐
	SubType int64  `json:"subType"` //推荐类型：2001：同款价优，2002：同款高佣，2003：同款高销，3001：相似价优，3002：相似高销，4001：推荐价优，4002：推荐高销，4003：推荐高佣，5001：类目优选，6001：其他热销，6002：其他低价
	Reason  string `json:"reason"`  //推荐理由
}
