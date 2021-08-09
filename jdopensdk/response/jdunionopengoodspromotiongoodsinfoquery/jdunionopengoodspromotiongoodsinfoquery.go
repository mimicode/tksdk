package jdunionopengoodspromotiongoodsinfoquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.promotiongoodsinfo.query 通过SKUID查询推广商品
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_promotiongoodsinfo_query_responce"`
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
	Code      int64   `json:"code"`
	Data      []Datum `json:"data"`
	Message   string  `json:"message"`
	RequestID string  `json:"requestId"`
}

type Datum struct {
	UnitPrice         float64 `json:"unitPrice"`
	MaterialURL       string  `json:"materialUrl"`
	EndDate           int64   `json:"endDate"`
	IsFreeFreightRisk int64   `json:"isFreeFreightRisk"`
	IsFreeShipping    int64   `json:"isFreeShipping"`
	CommisionRatioWl  float64 `json:"commisionRatioWl"`
	CommisionRatioPC  float64 `json:"commisionRatioPc"`
	ImgURL            string  `json:"imgUrl"`
	Vid               int64   `json:"vid"`
	CidName           string  `json:"cidName"`
	WlUnitPrice       float64 `json:"wlUnitPrice"`
	Cid2Name          string  `json:"cid2Name"`
	IsSeckill         int64   `json:"isSeckill"`
	Cid2              int64   `json:"cid2"`
	Cid3Name          string  `json:"cid3Name"`
	InOrderCount      int64   `json:"inOrderCount"`
	Cid3              int64   `json:"cid3"`
	ShopID            int64   `json:"shopId"`
	IsJdSale          int64   `json:"isJdSale"`
	GoodsName         string  `json:"goodsName"`
	SkuID             int64   `json:"skuId"`
	StartDate         int64   `json:"startDate"`
	Cid               int64   `json:"cid"`
}
