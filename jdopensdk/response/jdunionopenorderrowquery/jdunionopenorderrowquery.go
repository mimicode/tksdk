package jdunionopenorderrowquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.order.row.query 查询推广订单及佣金信息
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_order_row_query_responce"`
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
	HasMore   bool    `json:"hasMore"`
	Message   string  `json:"message"`
	RequestID string  `json:"requestId"`
}

type Datum struct {
	ActualCosPrice      float64      `json:"actualCosPrice"`
	ActualFee           float64      `json:"actualFee"`
	BalanceEXT          string       `json:"balanceExt"`
	CategoryInfo        CategoryInfo `json:"categoryInfo"`
	Cid1                int64        `json:"cid1"`
	Cid2                int64        `json:"cid2"`
	Cid3                int64        `json:"cid3"`
	CommissionRate      float64      `json:"commissionRate"`
	CpActID             int64        `json:"cpActId"`
	EstimateCosPrice    float64      `json:"estimateCosPrice"`
	EstimateFee         float64      `json:"estimateFee"`
	ExpressStatus       int64        `json:"expressStatus"`
	Ext1                string       `json:"ext1"`
	FinalRate           float64      `json:"finalRate"`
	FinishTime          string       `json:"finishTime"`
	GiftCouponKey       string       `json:"giftCouponKey"`
	GiftCouponOcsAmount float64      `json:"giftCouponOcsAmount"`
	GoodsInfo           GoodsInfo    `json:"goodsInfo"`
	ID                  string       `json:"id"`
	ModifyTime          string       `json:"modifyTime"`
	OrderEmt            int64        `json:"orderEmt"`
	OrderID             int64        `json:"orderId"`
	OrderTime           string       `json:"orderTime"`
	ParentID            int64        `json:"parentId"`
	PayMonth            int64        `json:"payMonth"`
	PID                 string       `json:"pid"`
	Plus                int64        `json:"plus"`
	PopID               int64        `json:"popId"`
	PositionID          int64        `json:"positionId"`
	Price               float64      `json:"price"`
	ProPriceAmount      float64      `json:"proPriceAmount"`
	Rid                 int64        `json:"rid"`
	SiteID              int64        `json:"siteId"`
	SkuFrozenNum        int64        `json:"skuFrozenNum"`
	SkuID               int64        `json:"skuId"`
	SkuName             string       `json:"skuName"`
	SkuNum              int64        `json:"skuNum"`
	SkuReturnNum        int64        `json:"skuReturnNum"`
	SubSideRate         float64      `json:"subSideRate"`
	SubUnionID          string       `json:"subUnionId"`
	SubsidyRate         float64      `json:"subsidyRate"`
	TraceType           int64        `json:"traceType"`
	UnionAlias          string       `json:"unionAlias"`
	UnionID             int64        `json:"unionId"`
	UnionRole           int64        `json:"unionRole"`
	UnionTag            string       `json:"unionTag"`
	ValidCode           int64        `json:"validCode"`
}

type CategoryInfo struct {
	Cid1 int64 `json:"cid1"`
	Cid2 int64 `json:"cid2"`
	Cid3 int64 `json:"cid3"`
}

type GoodsInfo struct {
	MainSkuID int64 `json:"mainSkuId"`
	ProductID int64 `json:"productId"`
	ShopID    int64 `json:"shopId"`
}
