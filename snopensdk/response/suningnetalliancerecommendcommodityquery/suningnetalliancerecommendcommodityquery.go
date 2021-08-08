package suningnetalliancerecommendcommodityquery

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.recommendcommodity.query 小编推荐接口
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
	QueryRecommendcommodity []QueryRecommendcommodity `json:"queryRecommendcommodity"`
}

type QueryRecommendcommodity struct {
	OrderActivity *OrderActivity `json:"orderActivity"`
	CommodityInfo CommodityInfo  `json:"commodityInfo"`
	PGInfo        PGInfo         `json:"pgInfo"`
	CategoryInfo  CategoryInfo   `json:"categoryInfo"`
	CouponInfo    CouponInfo     `json:"couponInfo"`
	AdvanceSale   AdvanceSale    `json:"advanceSale"`
	IsHaveData    *string        `json:"isHaveData,omitempty"`
}

type AdvanceSale struct {
	IsReserveCommodity string `json:"isReserveCommodity"`
	DepositAmount      string `json:"depositAmount"`
	DepositEndTime     string `json:"depositEndTime"`
}

type CategoryInfo struct {
	FirstSaleCategoryID        string `json:"firstSaleCategoryId"`
	FirstSaleCategoryName      string `json:"firstSaleCategoryName"`
	SecondSaleCategoryID       string `json:"secondSaleCategoryId"`
	SecondSaleCategoryName     string `json:"secondSaleCategoryName"`
	ThirdSaleCategoryID        string `json:"thirdSaleCategoryId"`
	ThirdSaleCategoryName      string `json:"thirdSaleCategoryName"`
	FirstPurchaseCategoryID    string `json:"firstPurchaseCategoryId"`
	FirstPurchaseCategoryName  string `json:"firstPurchaseCategoryName"`
	SecondPurchaseCategoryID   string `json:"secondPurchaseCategoryId"`
	SecondPurchaseCategoryName string `json:"secondPurchaseCategoryName"`
	ThirdPurchaseCategoryID    string `json:"thirdPurchaseCategoryId"`
	ThirdPurchaseCategoryName  string `json:"thirdPurchaseCategoryName"`
	GoodsGroupCategoryID       string `json:"goodsGroupCategoryId"`
	GoodsGroupCategoryName     string `json:"goodsGroupCategoryName"`
}

type CommodityInfo struct {
	CommodityName  string       `json:"commodityName"`
	CommodityCode  string       `json:"commodityCode"`
	SupplierCode   string       `json:"supplierCode"`
	SupplierName   string       `json:"supplierName"`
	PictureURL     []PictureURL `json:"pictureUrl"`
	SellingPoint   string       `json:"sellingPoint"`
	MonthSales     interface{}  `json:"monthSales"`
	SnPrice        string       `json:"snPrice"`
	CommodityPrice string       `json:"commodityPrice"`
	CommodityType  int64        `json:"commodityType"`
	PriceType      string       `json:"priceType"`
	PriceTypeCode  int64        `json:"priceTypeCode"`
	Baoyou         int64        `json:"baoyou"`
	Rate           string       `json:"rate"`
	SaleStatus     int64        `json:"saleStatus"`
	ProductURL     *string      `json:"productUrl,omitempty"`
}

type PictureURL struct {
	PicURL     string `json:"picUrl"`
	LocationID int64  `json:"locationId"`
}

type CouponInfo struct {
	CouponURL           string      `json:"couponUrl"`
	ActivityID          string      `json:"activityId"`
	ActivitySecretKey   string      `json:"activitySecretKey"`
	CouponValue         string      `json:"couponValue"`
	CouponCount         *string     `json:"couponCount"`
	CouponStartTime     string      `json:"couponStartTime"`
	CouponEndTime       string      `json:"couponEndTime"`
	StartTime           *string     `json:"startTime"`
	EndTime             *string     `json:"endTime"`
	BounsLimit          string      `json:"bounsLimit"`
	ActivityDescription string      `json:"activityDescription"`
	AfterCouponPrice    interface{} `json:"afterCouponPrice"`
}

type OrderActivity struct {
	ActivityRuleList    []ActivityRuleList `json:"activityRuleList"`
	ActivityDescription string             `json:"activityDescription"`
}

type ActivityRuleList struct {
	SingleMaxRewardAmount string `json:"singleMaxRewardAmount"`
	BaseAddUpType         string `json:"baseAddUpType"`
	RewardQuantifierType  string `json:"rewardQuantifierType"`
	ScopeType             string `json:"scopeType"`
	BaseQuantifierType    string `json:"baseQuantifierType"`
	RewardAmount          string `json:"rewardAmount"`
	Floor                 string `json:"floor"`
	BaseAmount            string `json:"baseAmount"`
}

type PGInfo struct {
	PGNum            string `json:"pgNum"`
	PGPrice          string `json:"pgPrice"`
	MinOrderQuantity string `json:"minOrderQuantity"`
	PGURL            string `json:"pgUrl"`
	PGActionID       string `json:"pgActionId"`
}

type MonthSales struct {
	Integer *int64
	String  *string
}
