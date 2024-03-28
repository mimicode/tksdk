package jdunionopengoodsjingfenquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.jingfen.query 京粉精选商品查询接口
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_jingfen_query_responce"`
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

type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"queryResult"`
	QueryResult    QueryResult
}

type QueryResult struct {
	Code       int64   `json:"code"`
	Data       []Datum `json:"data"`
	Message    string  `json:"message"`
	RequestID  string  `json:"requestId"`
	TotalCount int64   `json:"totalCount"`
}

type Datum struct {
	BrandCode             string          `json:"brandCode"`
	BrandName             string          `json:"brandName"`
	CategoryInfo          CategoryInfo    `json:"categoryInfo"`
	Comments              int64           `json:"comments"`
	CommissionInfo        CommissionInfo  `json:"commissionInfo"`
	CouponInfo            CouponInfo      `json:"couponInfo"`
	DeliveryType          int64           `json:"deliveryType"`
	ForbidTypes           []int64         `json:"forbidTypes"`
	GoodCommentsShare     float64         `json:"goodCommentsShare"`
	ImageInfo             ImageInfo       `json:"imageInfo"`
	InOrderCount30Days    int64           `json:"inOrderCount30Days"`
	InOrderCount30DaysSku int64           `json:"inOrderCount30DaysSku"`
	IsHot                 int64           `json:"isHot"`
	MaterialURL           string          `json:"materialUrl"`
	Owner                 string          `json:"owner"`
	PinGouInfo            PinGouInfo      `json:"pinGouInfo"`
	PriceInfo             PriceInfo       `json:"priceInfo"`
	ResourceInfo          ResourceInfo    `json:"resourceInfo"`
	ShopInfo              ShopInfo        `json:"shopInfo"`
	SkuID                 int64           `json:"skuId"`
	SkuLabelInfo          SkuLabelInfo    `json:"skuLabelInfo"`
	SkuName               string          `json:"skuName"`
	Spuid                 int64           `json:"spuid"`
	SeckillInfo           *SeckillInfo    `json:"seckillInfo,omitempty"`
	JxFlags               []int64         `json:"jxFlags,omitempty"`
	BonusInfoList         []BonusInfoList `json:"bonusInfoList,omitempty"`
	IsOversea             int64           `json:"isOversea"`
	ItemID                string          `json:"itemId"`
	VideoInfo             *VideoInfo      `json:"videoInfo,omitempty"`
	BookInfo              *BookInfo       `json:"bookInfo,omitempty"`
}

type CategoryInfo struct {
	Cid1     int64  `json:"cid1"`
	Cid1Name string `json:"cid1Name"`
	Cid2     int64  `json:"cid2"`
	Cid2Name string `json:"cid2Name"`
	Cid3     int64  `json:"cid3"`
	Cid3Name string `json:"cid3Name"`
}

type CommissionInfo struct {
	Commission          float64 `json:"commission"`
	CommissionShare     float64 `json:"commissionShare"`
	CouponCommission    float64 `json:"couponCommission"`
	EndTime             int64   `json:"endTime"`
	IsLock              int64   `json:"isLock"`
	PlusCommissionShare float64 `json:"plusCommissionShare"`
	StartTime           int64   `json:"startTime"`
}

type CouponInfo struct {
	CouponList []CouponList `json:"couponList"`
}

type CouponList struct {
	BindType     int64   `json:"bindType"`
	Discount     float64 `json:"discount"`
	GetEndTime   int64   `json:"getEndTime"`
	GetStartTime int64   `json:"getStartTime"`
	HotValue     *int64  `json:"hotValue,omitempty"`
	IsBest       int64   `json:"isBest"`
	Link         string  `json:"link"`
	PlatformType int64   `json:"platformType"`
	Quota        float64 `json:"quota"`
	UseEndTime   int64   `json:"useEndTime"`
	UseStartTime int64   `json:"useStartTime"`
	CouponStyle  int64   `json:"couponStyle"`
}

type ImageList struct {
	URL string `json:"url"`
}

type PinGouInfo struct {
	PingouEndTime   *int64   `json:"pingouEndTime,omitempty"`
	PingouPrice     *float64 `json:"pingouPrice,omitempty"`
	PingouStartTime *int64   `json:"pingouStartTime,omitempty"`
	PingouTmCount   *int64   `json:"pingouTmCount,omitempty"`
	PingouURL       *string  `json:"pingouUrl,omitempty"`
}

type ResourceInfo struct {
	EliteID   int64  `json:"eliteId"`
	EliteName string `json:"eliteName"`
}

type SeckillInfo struct {
	SeckillEndTime   int64   `json:"seckillEndTime"`
	SeckillOriPrice  float64 `json:"seckillOriPrice"`
	SeckillPrice     float64 `json:"seckillPrice"`
	SeckillStartTime int64   `json:"seckillStartTime"`
}

type ShopInfo struct {
	AfsFactorScoreRankGrade       *string `json:"afsFactorScoreRankGrade,omitempty"`
	AfterServiceScore             *string `json:"afterServiceScore,omitempty"`
	CommentFactorScoreRankGrade   *string `json:"commentFactorScoreRankGrade,omitempty"`
	LogisticsFactorScoreRankGrade *string `json:"logisticsFactorScoreRankGrade,omitempty"`
	LogisticsLvyueScore           *string `json:"logisticsLvyueScore,omitempty"`
	ScoreRankRate                 *string `json:"scoreRankRate,omitempty"`
	ShopID                        int64   `json:"shopId"`
	ShopLabel                     string  `json:"shopLabel"`
	ShopLevel                     float64 `json:"shopLevel"`
	ShopName                      string  `json:"shopName"`
	UserEvaluateScore             *string `json:"userEvaluateScore,omitempty"`
}

type SkuLabelInfo struct {
	FxgServiceList []interface{} `json:"fxgServiceList"`
	Is7ToReturn    int64         `json:"is7ToReturn"`
}

type BonusInfoList struct {
	ActivityType  int64   `json:"activityType"`
	BeginTime     int64   `json:"beginTime"`
	BonusSkuNum   int64   `json:"bonusSkuNum"`
	EndTime       int64   `json:"endTime"`
	EstimateBonus float64 `json:"estimateBonus"`
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	State         int64   `json:"state"`
	StockSkuNum   int64   `json:"stockSkuNum"`
}

type BookInfo struct {
	Isbn string `json:"isbn"`
}

type ImageInfo struct {
	ImageList  []ImageList `json:"imageList"`
	WhiteImage *string     `json:"whiteImage,omitempty"`
}

type PriceInfo struct {
	HistoryPriceDay   *int64  `json:"historyPriceDay,omitempty"`
	LowestCouponPrice float64 `json:"lowestCouponPrice"`
	LowestPrice       float64 `json:"lowestPrice"`
	LowestPriceType   int64   `json:"lowestPriceType"`
	Price             float64 `json:"price"`
}

type VideoInfo struct {
	VideoList []VideoList `json:"videoList"`
}

type VideoList struct {
	Duration  int64  `json:"duration"`
	High      int64  `json:"high"`
	ImageURL  string `json:"imageUrl"`
	PlayType  string `json:"playType"`
	PlayURL   string `json:"playUrl"`
	VideoType int64  `json:"videoType"`
	Width     int64  `json:"width"`
}
