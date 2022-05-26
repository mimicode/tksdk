package suningnetalliancemorerecommendget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.morerecommend.get 关联商品推荐接口
type Response struct {
	response2.TopResponse
	SnResponseContent SnResponseContent `json:"sn_responseContent"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = "-1"
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	} else {
		if t.SnResponseContent.SnError != nil {
			t.ErrorResponse.ErrorCode = t.SnResponseContent.SnError.ErrorCode
			t.ErrorResponse.ErrorMsg = t.SnResponseContent.SnError.ErrorMsg
		}
	}
}

type SnResponseContent struct {
	SnBody  SnBody   `json:"sn_body"`
	SnError *SnError `json:"sn_error"`
}
type SnError struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type SnBody struct {
	GetMorerecommend GetMorerecommend `json:"getMorerecommend"`
}

type GetMorerecommend struct {
	CommodityList []CommodityList `json:"commodityList"`
}

type CommodityList struct {
	IsReserveCommodity   int64       `json:"isReserveCommodity"`
	CommodityName        string      `json:"commodityName"`
	CommodityCode        string      `json:"commodityCode"`
	SupplierCode         string      `json:"supplierCode"`
	ImgVersion           string      `json:"imgVersion"`
	IsOwnCommodity       bool        `json:"isOwnCommodity"`
	CommodityType        int64       `json:"commodityType"`
	Baoyou               int64       `json:"baoyou"`
	SellingPoint         string      `json:"sellingPoint"`
	MonthlySales         int64       `json:"monthlySales"`
	OriginRate           string      `json:"originRate"`
	OriginCommission     string      `json:"originCommission"`
	CommodityPrice       string      `json:"commodityPrice"`
	CommissionPrice      string      `json:"commissionPrice"`
	CommissionRate       string      `json:"commissionRate"`
	SnPrice              string      `json:"snPrice"`
	PriceType            string      `json:"priceType"`
	PriceTypeCode        int64       `json:"priceTypeCode"`
	PGPrice              interface{} `json:"pgPrice"`
	PGActionID           interface{} `json:"pgActionId"`
	PGNum                interface{} `json:"pgNum"`
	CommodityPriceFlag   string      `json:"commodityPriceFlag"`
	SaleStatus           int64       `json:"saleStatus"`
	NoConponCommision    string      `json:"noConponCommision"`
	RebateCommissionRate string      `json:"rebateCommissionRate"`
	OrderActivity        interface{} `json:"orderActivity"`
	CouponAmount         interface{} `json:"couponAmount"`
	CouponSpecialPrice   string      `json:"couponSpecialPrice"`
	CouponActiveID       string      `json:"couponActiveId"`
	ActivitySecretKey    string      `json:"activitySecretKey"`
	CouponPrice          interface{} `json:"couponPrice"`
	MinOrderQuantity     string      `json:"minOrderQuantity"`
	PicList              []PicList   `json:"picList"`
	ProductURL           string      `json:"productUrl"`
}

type PicList struct {
	CmmdtyURL string `json:"cmmdtyUrl"`
}
