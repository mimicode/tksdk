package jdunionopengoodsmaterialquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.material.query 猜你喜欢商品推荐
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_material_query_response"`
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
	Code       int64  `json:"code"`
	Data       Data   `json:"data"`
	Message    string `json:"message"`
	TotalCount string `json:"totalCount"`
	RequestID  string `json:"requestId"`
}

type Data struct {
	MaterialGoodsResp MaterialGoodsResp `json:"materialGoodsResp"`
}

type MaterialGoodsResp struct {
	BookInfo               BookInfo                                `json:"bookInfo"`
	MaterialURL            string                                  `json:"materialUrl"`
	ImageInfo              ImageInfo                               `json:"imageInfo"`
	PinGouInfo             PinGouInfo                              `json:"pinGouInfo"`
	ForbidTypes            []int64                                 `json:"forbidTypes"`
	ResourceInfo           ResourceInfo                            `json:"resourceInfo"`
	SkuLabelInfo           SkuLabelInfo                            `json:"skuLabelInfo"`
	AddCartPrice           string                                  `json:"addCartPrice"`
	ActivityCardInfo       ActivityCardInfo                        `json:"activityCardInfo"`
	SkuName                string                                  `json:"skuName"`
	PriceInfo              PriceInfo                               `json:"priceInfo"`
	IsOversea              string                                  `json:"isOversea"`
	Spuid                  string                                  `json:"spuid"`
	CommissionInfo         CommissionInfo                          `json:"commissionInfo"`
	SkuID                  string                                  `json:"skuId"`
	BrandCode              string                                  `json:"brandCode"`
	Owner                  string                                  `json:"owner"`
	KaAdowner              string                                  `json:"kaAdowner"`
	ShopInfo               ShopInfo                                `json:"shopInfo"`
	Comments               string                                  `json:"comments"`
	SeckillInfo            SeckillInfo                             `json:"seckillInfo"`
	CouponInfo             CouponInfo                              `json:"couponInfo"`
	PreSaleInfo            PreSaleInfo                             `json:"preSaleInfo"`
	CompanyType            string                                  `json:"companyType"`
	VideoInfo              VideoInfo                               `json:"videoInfo"`
	SecondPriceInfoList    SecondPriceInfoList                     `json:"secondPriceInfoList"`
	DeliveryType           string                                  `json:"deliveryType"`
	GoodCommentsShare      string                                  `json:"goodCommentsShare"`
	PromotionInfo          PromotionInfo                           `json:"promotionInfo"`
	CategoryInfo           CategoryInfo                            `json:"categoryInfo"`
	InOrderCount30DaysSku  string                                  `json:"inOrderCount30DaysSku"`
	InOrderCount30Days     string                                  `json:"inOrderCount30Days"`
	ReserveInfo            ReserveInfo                             `json:"reserveInfo"`
	ItemID                 string                                  `json:"itemId"`
	PurchasePriceInfo      PurchasePriceInfo                       `json:"purchasePriceInfo"`
	PromotionLabelInfoList MaterialGoodsRespPromotionLabelInfoList `json:"promotionLabelInfoList"`
	IsHot                  string                                  `json:"isHot"`
	JxFlags                []int64                                 `json:"jxFlags"`
}

type ActivityCardInfo struct {
	Amount       string `json:"amount"`
	ExpireDay    string `json:"expireDay"`
	ActivityType string `json:"activityType"`
}

type BookInfo struct {
	Isbn string `json:"isbn"`
}

type CategoryInfo struct {
	Cid1Name string `json:"cid1Name"`
	Cid2Name string `json:"cid2Name"`
	Cid2     string `json:"cid2"`
	Cid3Name string `json:"cid3Name"`
	Cid3     string `json:"cid3"`
	Cid1     string `json:"cid1"`
}

type CommissionInfo struct {
	CommissionShare     string `json:"commissionShare"`
	PlusCommissionShare string `json:"plusCommissionShare"`
	Commission          string `json:"commission"`
	CouponCommission    string `json:"couponCommission"`
}

type CouponInfo struct {
	CouponList CouponInfoCouponList `json:"couponList"`
}

type CouponInfoCouponList struct {
	Coupon PurpleCoupon `json:"coupon"`
}

type PurpleCoupon struct {
	UseEndTime   string `json:"useEndTime"`
	GetEndTime   string `json:"getEndTime"`
	UseStartTime string `json:"useStartTime"`
	Quota        string `json:"quota"`
	BindType     string `json:"bindType"`
	Link         string `json:"link"`
	PlatformType string `json:"platformType"`
	Discount     string `json:"discount"`
	GetStartTime string `json:"getStartTime"`
	IsBest       string `json:"isBest"`
	CouponStyle  string `json:"couponStyle"`
}

type ImageInfo struct {
	WhiteImage string    `json:"whiteImage"`
	ImageList  ImageList `json:"imageList"`
}

type ImageList struct {
	URLInfo URLInfo `json:"urlInfo"`
}

type URLInfo struct {
	URL string `json:"url"`
}

type PinGouInfo struct {
	PingouEndTime   string `json:"pingouEndTime"`
	PingouPrice     string `json:"pingouPrice"`
	PingouTmCount   string `json:"pingouTmCount"`
	PingouStartTime string `json:"pingouStartTime"`
}

type PreSaleInfo struct {
	DepositWorth     string `json:"depositWorth"`
	BalanceEndTime   string `json:"balanceEndTime"`
	ShipTime         string `json:"shipTime"`
	PreSalePayType   string `json:"preSalePayType"`
	CurrentPrice     string `json:"currentPrice"`
	PreSaleStartTime string `json:"preSaleStartTime"`
	BalanceStartTime string `json:"balanceStartTime"`
	PreSaleEndTime   string `json:"preSaleEndTime"`
	PreSaleStatus    string `json:"preSaleStatus"`
	AmountDeposit    string `json:"amountDeposit"`
	DiscountType     string `json:"discountType"`
	Earnest          string `json:"earnest"`
	PreAmountDeposit string `json:"preAmountDeposit"`
}

type PriceInfo struct {
	LowestPrice       string `json:"lowestPrice"`
	LowestCouponPrice string `json:"lowestCouponPrice"`
	Price             string `json:"price"`
	LowestPriceType   string `json:"lowestPriceType"`
}

type PromotionInfo struct {
	ClickURL string `json:"clickURL"`
}

type MaterialGoodsRespPromotionLabelInfoList struct {
	PromotionLabelInfo PurplePromotionLabelInfo `json:"promotionLabelInfo"`
}

type PurplePromotionLabelInfo struct {
	PromotionLabel   string `json:"promotionLabel"`
	LableName        string `json:"lableName"`
	PromotionLableID string `json:"promotionLableId"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
}

type PurchasePriceInfo struct {
	ThresholdPrice         string                                  `json:"thresholdPrice"`
	Code                   string                                  `json:"code"`
	BasisPriceType         string                                  `json:"basisPriceType"`
	CouponList             PurchasePriceInfoCouponList             `json:"couponList"`
	PurchasePrice          string                                  `json:"purchasePrice"`
	Message                string                                  `json:"message"`
	PromotionLabelInfoList PurchasePriceInfoPromotionLabelInfoList `json:"promotionLabelInfoList"`
}

type PurchasePriceInfoCouponList struct {
	Coupon FluffyCoupon `json:"coupon"`
}

type FluffyCoupon struct {
	Quota        string `json:"quota"`
	BindType     string `json:"bindType"`
	Link         string `json:"link"`
	PlatformType string `json:"platformType"`
	Discount     string `json:"discount"`
	CouponStyle  string `json:"couponStyle"`
}

type PurchasePriceInfoPromotionLabelInfoList struct {
	PromotionLabelInfo FluffyPromotionLabelInfo `json:"promotionLabelInfo"`
}

type FluffyPromotionLabelInfo struct {
	PromotionLabel   string `json:"promotionLabel"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	LabelName        string `json:"labelName"`
	PromotionLabelID string `json:"promotionLabelId"`
}

type ReserveInfo struct {
	Price                string `json:"price"`
	PanicBuyingEndTime   string `json:"panicBuyingEndTime"`
	StartTime            string `json:"startTime"`
	EndTime              string `json:"endTime"`
	Type                 string `json:"type"`
	Status               string `json:"status"`
	PanicBuyingStartTime string `json:"panicBuyingStartTime"`
}

type ResourceInfo struct {
	EliteID   string `json:"eliteId"`
	EliteName string `json:"eliteName"`
}

type SeckillInfo struct {
	SeckillOriPrice  string `json:"seckillOriPrice"`
	SeckillPrice     string `json:"seckillPrice"`
	SeckillStartTime string `json:"seckillStartTime"`
	SeckillEndTime   string `json:"seckillEndTime"`
}

type SecondPriceInfoList struct {
	SecondPriceInfo SecondPriceInfo `json:"secondPriceInfo"`
}

type SecondPriceInfo struct {
	SecondPriceType string `json:"secondPriceType"`
	SecondPrice     string `json:"secondPrice"`
}

type ShopInfo struct {
	LogisticsLvyueScore           string `json:"logisticsLvyueScore"`
	ShopLevel                     string `json:"shopLevel"`
	UserEvaluateScore             string `json:"userEvaluateScore"`
	ScoreRankRate                 string `json:"scoreRankRate"`
	AfterServiceScore             string `json:"afterServiceScore"`
	ShopName                      string `json:"shopName"`
	ShopLabel                     string `json:"shopLabel"`
	AfsFactorScoreRankGrade       string `json:"afsFactorScoreRankGrade"`
	ShopID                        string `json:"shopId"`
	LogisticsFactorScoreRankGrade string `json:"logisticsFactorScoreRankGrade"`
	CommentFactorScoreRankGrade   string `json:"commentFactorScoreRankGrade"`
}

type SkuLabelInfo struct {
	Is7ToReturn    string         `json:"is7ToReturn"`
	Fxg            string         `json:"fxg"`
	FxgServiceList FxgServiceList `json:"fxgServiceList"`
}

type FxgServiceList struct {
	CharacteristicServiceInfo CharacteristicServiceInfo `json:"characteristicServiceInfo"`
}

type CharacteristicServiceInfo struct {
	ServiceName string `json:"serviceName"`
}

type VideoInfo struct {
	VideoList VideoList `json:"videoList"`
}

type VideoList struct {
	Video Video `json:"video"`
}

type Video struct {
	Duration  string `json:"duration"`
	High      string `json:"high"`
	PlayType  string `json:"playType"`
	VideoType string `json:"videoType"`
	ImageURL  string `json:"imageUrl"`
	Width     string `json:"width"`
	PlayURL   string `json:"playUrl"`
}
