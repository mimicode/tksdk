package jdunionopengoodsquery

import (
	"encoding/json"
	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.query 接口描述：查询商品及优惠券信息，返回的结果可调用转链接口生成单品或二合一推广链接。支持按SKUID、关键词、优惠券基本属性、是否拼购、是否爆款等条件查询，建议不要同时传入SKUID和其他字段，以获得较多的结果。支持按价格、佣金比例、佣金、引单量等维度排序。用优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_query_responce"`
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
	Code       int64   `json:"code"`
	Data       []Datum `json:"data"`
	Message    string  `json:"message"`
	RequestID  string  `json:"requestId"`
	TotalCount int64   `json:"totalCount"`
}

type Datum struct {
	BrandCode          string         `json:"brandCode"`
	BrandName          string         `json:"brandName"`
	CategoryInfo       CategoryInfo   `json:"categoryInfo"`
	Comments           int64          `json:"comments"`
	CommissionInfo     CommissionInfo `json:"commissionInfo"`
	CouponInfo         CouponInfo     `json:"couponInfo"`
	DeliveryType       int64          `json:"deliveryType"`
	ForbidTypes        []int64        `json:"forbidTypes"` // 0普通商品，10微信京东购物小程序禁售，11微信京喜小程序禁售
	GoodCommentsShare  float64        `json:"goodCommentsShare"`
	ImageInfo          ImageInfo      `json:"imageInfo"`
	InOrderComm30Days  float64        `json:"inOrderComm30Days"`
	InOrderCount30Days int64          `json:"inOrderCount30Days"`
	IsHot              int64          `json:"isHot"`
	IsJdSale           int64          `json:"isJdSale"`
	MaterialURL        string         `json:"materialUrl"`
	Owner              string         `json:"owner"`
	PinGouInfo         PinGouInfo     `json:"pinGouInfo"`
	PriceInfo          PriceInfo      `json:"priceInfo"`
	ShopInfo           ShopInfo       `json:"shopInfo"`
	SkuID              int64          `json:"skuId"`
	SkuName            string         `json:"skuName"`
	Spuid              int64          `json:"spuid"`
	VideoInfo          VideoInfo      `json:"videoInfo"`
	BookInfo           BookInfo       `json:"bookInfo"`
	JxFlags            []int64        `json:"jxFlags"`   // 京喜商品类型（1京喜；2京喜工厂直供；3京喜优选（包含3时可在京东APP购买））
	EliteType          []int64        `json:"eliteType"` // 资源位17：极速版商品
	DocumentInfo       DocumentInfo   `json:"documentInfo"`
	SeckillInfo        SeckillInfo    `json:"seckillInfo"` //秒杀信息
}

type SeckillInfo struct {
	SeckillOriPrice  float64 `json:"seckillOriPrice"`
	SeckillPrice     float64 `json:"seckillPrice"`
	SeckillStartTime int64   `json:"seckillStartTime"`
	SeckillEndTime   int64   `json:"seckillEndTime"`
}

type DocumentInfo struct {
	Document string `json:"document"`
	Discount string `json:"discount"`
}

type BookInfo struct {
	Isbn string `json:"isbn"`
}

type VideoInfo struct {
	VideoList []VideoList `json:"videoList"`
}

type VideoList struct {
	Video Video `json:"video"`
}

type Video struct {
	Width     int    `json:"width"`
	High      int    `json:"high"`
	ImageUrl  string `json:"imageUrl"`
	VideoType int    `json:"videoType"`
	PlayType  int    `json:"playType"`
	Duration  int    `json:"duration"`
	PlayUrl   string `json:"playUrl"`
}

type PinGouInfo struct {
	PingouEndTime   int64   `json:"pingouEndTime"`
	PingouPrice     float64 `json:"pingouPrice"`
	PingouStartTime int64   `json:"pingouStartTime"`
	PingouTmCount   int64   `json:"pingouTmCount"`
	PingouURL       string  `json:"pingouUrl"`
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
	BindType      int64   `json:"bindType"`
	Discount      float64 `json:"discount"`
	GetEndTime    int64   `json:"getEndTime"`
	GetStartTime  int64   `json:"getStartTime"`
	HotValue      *int64  `json:"hotValue,omitempty"`
	IsBest        int64   `json:"isBest"`
	IsInputCoupon int64   `json:"isInputCoupon"`
	Link          string  `json:"link"`
	PlatformType  int64   `json:"platformType"`
	Quota         float64 `json:"quota"`
	UseEndTime    int64   `json:"useEndTime"`
	UseStartTime  int64   `json:"useStartTime"`
}

type ImageInfo struct {
	ImageList  []ImageList `json:"imageList"`
	WhiteImage string      `json:"whiteImage"`
}

type ImageList struct {
	URL string `json:"url"`
}

type PriceInfo struct {
	LowestCouponPrice float64 `json:"lowestCouponPrice"`
	LowestPrice       float64 `json:"lowestPrice"`
	LowestPriceType   int64   `json:"lowestPriceType"`
	Price             float64 `json:"price"`
}

type ShopInfo struct {
	AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade"`
	AfterServiceScore             string  `json:"afterServiceScore"`
	CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade"`
	LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade"`
	LogisticsLvyueScore           string  `json:"logisticsLvyueScore"`
	ScoreRankRate                 string  `json:"scoreRankRate"`
	ShopID                        int64   `json:"shopId"`
	ShopLabel                     string  `json:"shopLabel"`
	ShopLevel                     float64 `json:"shopLevel"`
	ShopName                      string  `json:"shopName"`
	UserEvaluateScore             string  `json:"userEvaluateScore"`
}
