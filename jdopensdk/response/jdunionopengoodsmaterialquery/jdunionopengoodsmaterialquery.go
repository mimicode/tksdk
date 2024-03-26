package jdunionopengoodsmaterialquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.material.query 猜你喜欢商品推荐
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_material_query_responce"`
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
	Data       []Data `json:"data"`
	Message    string `json:"message"`
	RequestID  string `json:"requestId"`
	TotalCount int64  `json:"totalCount"`
}

type CategoryInfo struct {
	Cid1     int    `json:"cid1"`
	Cid1Name string `json:"cid1Name"`
	Cid2     int    `json:"cid2"`
	Cid2Name string `json:"cid2Name"`
	Cid3     int    `json:"cid3"`
	Cid3Name string `json:"cid3Name"`
}

type CommissionInfo struct {
	Commission          float64 `json:"commission"`
	CommissionShare     float64 `json:"commissionShare"`
	CouponCommission    float64 `json:"couponCommission"`
	PlusCommissionShare float64 `json:"plusCommissionShare"`
}

type CouponList struct {
	BindType     int     `json:"bindType"`
	CouponStyle  int     `json:"couponStyle"`
	Discount     float64 `json:"discount"`
	GetEndTime   int64   `json:"getEndTime"`
	GetStartTime int64   `json:"getStartTime"`
	IsBest       int     `json:"isBest"`
	Link         string  `json:"link"`
	PlatformType int     `json:"platformType"`
	Quota        float64 `json:"quota"`
	UseEndTime   int64   `json:"useEndTime"`
	UseStartTime int64   `json:"useStartTime"`
}

type CouponInfo struct {
	CouponList []CouponList `json:"couponList"`
}

type CardInfo struct {
	ActivityType int `json:"activityType"`
	Amount       int `json:"amount"`
	ExpireDay    int `json:"expireDay"`
}

type VideoInfo struct {
}

type ShopInfo struct {
	AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade,omitempty"`
	AfterServiceScore             string  `json:"afterServiceScore,omitempty"`
	CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade,omitempty"`
	LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade,omitempty"`
	LogisticsLvyueScore           string  `json:"logisticsLvyueScore,omitempty"`
	ScoreRankRate                 string  `json:"scoreRankRate,omitempty"`
	ShopID                        int     `json:"shopId"`
	ShopLabel                     string  `json:"shopLabel"`
	ShopLevel                     float64 `json:"shopLevel"`
	ShopName                      string  `json:"shopName"`
	UserEvaluateScore             string  `json:"userEvaluateScore,omitempty"`
}

type PriceInfo struct {
	LowestCouponPrice float64 `json:"lowestCouponPrice"`
	LowestPrice       float64 `json:"lowestPrice"`
	LowestPriceType   int     `json:"lowestPriceType"`
	Price             float64 `json:"price"`
}

type PromotionInfo struct {
	ClickURL string `json:"clickURL"`
}

type ImageList struct {
	URL string `json:"url"`
}

type ImageInfo struct {
	ImageList  []ImageList `json:"imageList"`
	WhiteImage string      `json:"whiteImage,omitempty"`
}

type PinGouInfo struct {
}

type ResourceInfo struct {
	EliteId   int    `json:"eliteId"`
	EliteName string `json:"eliteName"`
}

type Data struct {
	BrandCode             string         `json:"brandCode"`
	BrandName             string         `json:"brandName"`
	CategoryInfo          CategoryInfo   `json:"categoryInfo"`
	Comments              int64          `json:"comments"`
	CommissionInfo        CommissionInfo `json:"commissionInfo"`
	CouponInfo            CouponInfo     `json:"couponInfo"`
	DeliveryType          int            `json:"deliveryType"`
	ForbidTypes           []int64        `json:"forbidTypes"`
	GoodCommentsShare     float64        `json:"goodCommentsShare"`
	ImageInfo             ImageInfo      `json:"imageInfo"`
	InOrderCount30Days    int            `json:"inOrderCount30Days"`
	InOrderCount30DaysSku int            `json:"inOrderCount30DaysSku"`
	IsHot                 int            `json:"isHot"`
	IsOversea             int            `json:"isOversea"`
	ItemID                string         `json:"itemId"`
	MaterialUrl           string         `json:"materialUrl"`
	Owner                 string         `json:"owner"`
	PinGouInfo            PinGouInfo     `json:"pinGouInfo"`
	PriceInfo             PriceInfo      `json:"priceInfo"`
	PromotionInfo         PromotionInfo  `json:"promotionInfo"`
	ResourceInfo          ResourceInfo   `json:"resourceInfo"`
	ShopInfo              ShopInfo       `json:"shopInfo"`
	SkuID                 int64          `json:"skuId"`
	SkuName               string         `json:"skuName"`
	Spuid                 int64          `json:"spuid"`
	VideoInfo             VideoInfo      `json:"videoInfo"`
	ActivityCardInfo      CardInfo       `json:"activityCardInfo,omitempty"`
	JxFlags               []int64        `json:"jxFlags"`
}
