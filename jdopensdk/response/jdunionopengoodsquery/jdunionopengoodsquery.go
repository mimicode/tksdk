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
	Code              int64    `json:"code"`
	Data              []Datum  `json:"data"`
	Message           string   `json:"message"`
	RequestID         string   `json:"requestId"`
	TotalCount        int64    `json:"totalCount"`
	HotWords          string   `json:"hotWords"`          // 日常top10的热搜词，按小时更新
	SimilarSkuList    []int64  `json:"similarSkuList"`    // 相似推荐商品skuId集合
	SimilarItemIdList []string `json:"similarItemIdList"` // 相似推荐联盟商品ID集合
	SearchPosition    string   `json:"searchPosition"`    // 查询索引位，再次查询须入参该值
}

type Datum struct {
	BrandCode          string            `json:"brandCode"`
	BrandName          string            `json:"brandName"`
	CategoryInfo       CategoryInfo      `json:"categoryInfo"`
	Comments           int64             `json:"comments"`
	CommissionInfo     CommissionInfo    `json:"commissionInfo"`
	CouponInfo         CouponInfo        `json:"couponInfo"`
	DeliveryType       int64             `json:"deliveryType"`
	ForbidTypes        []int64           `json:"forbidTypes"` // 0普通商品，10微信京东购物小程序禁售，11微信京喜小程序禁售
	GoodCommentsShare  float64           `json:"goodCommentsShare"`
	ImageInfo          ImageInfo         `json:"imageInfo"`
	InOrderComm30Days  float64           `json:"inOrderComm30Days"`
	InOrderCount30Days int64             `json:"inOrderCount30Days"`
	IsHot              int64             `json:"isHot"`
	IsJdSale           int64             `json:"isJdSale"`
	ItemId             string            `json:"itemId"`
	MaterialURL        string            `json:"materialUrl"`
	Owner              string            `json:"owner"`
	PinGouInfo         PinGouInfo        `json:"pinGouInfo"`
	PriceInfo          PriceInfo         `json:"priceInfo"`
	PurchasePriceInfo  PurchasePriceInfo `json:"purchasePriceInfo"`
	ShopInfo           ShopInfo          `json:"shopInfo"`
	SkuID              int64             `json:"skuId"`
	SkuName            string            `json:"skuName"`
	Spuid              int64             `json:"spuid"`
	VideoInfo          VideoInfo         `json:"videoInfo"`
	BookInfo           BookInfo          `json:"bookInfo"`
	JxFlags            []int64           `json:"jxFlags"`   // 京喜商品类型（1京喜；2京喜工厂直供；3京喜优选（包含3时可在京东APP购买））
	EliteType          []int64           `json:"eliteType"` // 资源位17：极速版商品
	DocumentInfo       DocumentInfo      `json:"documentInfo"`
	SeckillInfo        SeckillInfo       `json:"seckillInfo"` // 秒杀信息
	CommentInfo        CommentInfo       `json:"commentInfo"` // 评价信息
	SpecInfo           SpecInfo          `json:"specInfo"`    // 扩展信息
	StockState         int64             `json:"stockState"`  // 库存状态：1有货、0无货（供tob选品场景参考，toc场景不适用）
	SkuLabelInfo       SkuLabelInfo      `json:"skuLabelInfo"`
	// 商品标签
	PromotionLabelInfoList []PromotionLabelInfo `json:"promotionLabelInfoList"` // 商品促销标签集
	SecondPriceInfoList    []SecondPriceInfo    `json:"secondPriceInfoList"`    // 双价格
	PreSaleInfo            PreSaleInfo          `json:"preSaleInfo"`            // 预售信息
	ReserveInfo            ReserveInfo          `json:"reserveInfo"`            // 预约信息
	IsOversea              int64                `json:"isOversea"`              // 是否全球购商品 1：是
	CompanyType            int64                `json:"companyType"`            // 2：POP自然人小店
	ActivityCardInfo       ActivityCardInfo     `json:"activityCardInfo"`       // 超市购物卡明细
	SmartDocumentInfoList  []SmartDocumentInfo  `json:"smartDocumentInfoList"`  // GPT算法智能生成的推广文案集合,内含多个风格
	KaAdowner              int64                `json:"kaAdowner"`              // 是否星选商家商品。1：是
	OriItemId              string               `json:"oriItemId"`              // 原始入参ItemId
	CallerItemId           string               `json:"callerItemId"`           // 工具商联盟商品ID
	SkuTagList             []SkuTagInfo         `json:"skuTagList"`             // 联盟标签
	InOrderCount30DaysSku  int64                `json:"inOrderCount30DaysSku"`  // 30天引单数量(sku维度)
	SpecialSkuUrlInfo      SpecialSkuUrlInfo    `json:"specialSkuUrlInfo"`      // 频道页信息
	SearchPageInfo         SearchPageInfo       `json:"searchPageInfo"`         // 该字段信息已禁用，请勿获取和使用
	MultiPurchasePriceInfo PurchasePriceInfo    `json:"multiPurchasePriceInfo"` // 普通多件到手价
	BonusInfoList          []BonusInfo          `json:"bonusInfoList"`          // 联盟奖励活动集合
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
	Isbn            string `json:"isbn"`            // 图书编号
	PublisherName   string `json:"publisherName"`   // 出版商名称
	AuthorName      string `json:"authorName"`      // 作者名称
	BookDesc        string `json:"bookDesc"`        // 内容简介
	BookName        string `json:"bookName"`        // 中文图书名称
	ForeignBookName string `json:"foreignBookName"` // 英文图书名称
}

// CommentInfo 评价信息
type CommentInfo struct {
	CommentList []Comment `json:"commentList"`
}

// Comment 评价列表
type Comment struct {
	Content   string      `json:"content"`   // 评价内容
	ImageList []ImageList `json:"imageList"` // 图片集合【废弃】
}

// SpecInfo 扩展信息
type SpecInfo struct {
	Size           string `json:"size"`           // 尺寸
	Color          string `json:"color"`          // 颜色
	Spec           string `json:"spec"`           // 自定义属性
	SpecName       string `json:"specName"`       // 自定义属性名称
	IsFreeShipping int64  `json:"isFreeShipping"` // 是否包邮(1:是,0:否,2:自营商品遵从主站包邮规则)
}

// SkuLabelInfo 商品标签
type SkuLabelInfo struct {
	Is7ToReturn    int64                       `json:"is7ToReturn"`    // 0：不支持； 1或null：支持7天无理由退货； 2：支持90天无理由退货； 4：支持15天无理由退货； 6：支持30天无理由退货
	Fxg            int64                       `json:"fxg"`            // 1：放心购商品
	FxgServiceList []CharacteristicServiceInfo `json:"fxgServiceList"` // 放心购商品子标签集合
}

// CharacteristicServiceInfo 放心购商品子标签，此字段值可能为空
type CharacteristicServiceInfo struct {
	ServiceName string `json:"serviceName"` // 服务名称
}

// SecondPriceInfo 双价格信息
type SecondPriceInfo struct {
	SecondPriceType int64   `json:"secondPriceType"` // 双价格类型：2:plus会员价格，9:学生价，18:新人价
	SecondPrice     float64 `json:"secondPrice"`     // 价格
}

// PreSaleInfo 预售信息
type PreSaleInfo struct {
	CurrentPrice     float64 `json:"currentPrice"`     // 预售价格
	Earnest          float64 `json:"earnest"`          // 订金金额（定金不能超过预售总价的20%）
	PreSalePayType   int64   `json:"preSalePayType"`   // 预售支付类型：1.仅全款 2.定金、全款均可 5.一阶梯仅定金
	DiscountType     int64   `json:"discountType"`     // 1: 定金膨胀 2: 定金立减
	DepositWorth     float64 `json:"depositWorth"`     // 定金膨胀金额（定金可抵XXX）【废弃】
	PreAmountDeposit float64 `json:"preAmountDeposit"` // 立减金额
	PreSaleStartTime int64   `json:"preSaleStartTime"` // 定金开始时间
	PreSaleEndTime   int64   `json:"preSaleEndTime"`   // 定金结束时间
	BalanceStartTime int64   `json:"balanceStartTime"` // 尾款开始时间
	BalanceEndTime   int64   `json:"balanceEndTime"`   // 尾款结束时间
	ShipTime         int64   `json:"shipTime"`         // 预计发货时间
	PreSaleStatus    int64   `json:"preSaleStatus"`    // 预售状态（0 未开始；1 预售中；2 预售结束；3 尾款进行中；4 尾款结束）
	AmountDeposit    float64 `json:"amountDeposit"`    // 定金膨胀金额（定金可抵XXX）
}

// ReserveInfo 预约信息
type ReserveInfo struct {
	Price                float64 `json:"price"`                // 预约价格
	Type                 int64   `json:"type"`                 // 预约类型： 1：预约购买资格（仅预约的用户才可以进行购买）； 5：预约抽签（仅中签用户可购买）
	Status               int64   `json:"status"`               // 1：等待预约 2：预约中 3：等待抢购/抽签中 4：抢购中 5：抢购结束
	StartTime            int64   `json:"startTime"`            // 预定开始时间
	EndTime              int64   `json:"endTime"`              // 预定结束时间
	PanicBuyingStartTime int64   `json:"panicBuyingStartTime"` // 抢购开始时间
	PanicBuyingEndTime   int64   `json:"panicBuyingEndTime"`   // 抢购结束时间
}

// ActivityCardInfo 超市购物卡明细
type ActivityCardInfo struct {
	Amount       float64 `json:"amount"`       // 超市卡金额
	ActivityType int64   `json:"activityType"` // 活动类型，购物返卡：3
	ExpireDay    int64   `json:"expireDay"`    // 超市卡有效天数
}

// SmartDocumentInfo 智能文案明细
type SmartDocumentInfo struct {
	DocumentType int64  `json:"documentType"` // 智能文案类型，1：最优推荐，2：知乎风格， 3：小红书风格， 4：社群风格
	DocumentName string `json:"documentName"` // 文案名称
	Document     string `json:"document"`     // 智能文案内容
}

// SkuTagInfo 联盟标签明细
type SkuTagInfo struct {
	Type  int64  `json:"type"`  // 标签类型 1 频道标签 2 商品卖点 3 营销状态 4 价格及热度标签 5 商品及店铺标签 6 会员及国补促销 7 超市权益 8 超级补贴活动 9 优选商品 10 跨境商品 11 经营模式 12 商品服务 13 星选商品 14 推广表现 15 商品表现 16 履约服务 17 京喜自营 18 佣金权益 19 超级补贴优选 20 白条权益 21 政府补贴 22 PLUS权益 23 红包权益
	Index int64  `json:"index"` // 优先级顺序（越小越优先）
	Name  string `json:"name"`  // 标签名称
}

// SpecialSkuUrlInfo 频道页信息
type SpecialSkuUrlInfo struct {
	SkuUrlType int64  `json:"skuUrlType"` // 频道页链接类型，1：秒杀专享价
	SkuUrl     string `json:"skuUrl"`     // 频道页地址
}

// SearchPageInfo 该字段信息已禁用，请勿获取和使用
type SearchPageInfo struct {
	SearchPageUrl string `json:"searchPageUrl"`
}

// BonusInfo 联盟奖励活动
type BonusInfo struct {
	ID            int64   `json:"id"`            // 活动ID
	Name          string  `json:"name"`          // 活动名称
	State         int64   `json:"state"`         // 活动状态， 1 ：已提报， 2 ：预热中， 3： 进行中， 4 ：已结束
	BeginTime     int64   `json:"beginTime"`     // 活动开始时间
	EndTime       int64   `json:"endTime"`       // 活动结束时间
	ActivityType  int64   `json:"activityType"`  // 活动类型，0：其他， 1：超级补贴，2：佣金比例活动，3：前N件奖励，5：超级补贴（新），6：前N件奖励（新）
	EstimateBonus float64 `json:"estimateBonus"` // 预估奖励金额
	BonusSkuNum   int64   `json:"bonusSkuNum"`   // 总奖励件数
	StockSkuNum   int64   `json:"stockSkuNum"`   // 剩余奖励件数
}

type VideoInfo struct {
	VideoList []Video `json:"videoList"`
}

type Video struct {
	Width     int64  `json:"width"`
	High      int64  `json:"high"`
	ImageUrl  string `json:"imageUrl"`
	VideoType int64  `json:"videoType"`
	PlayType  string `json:"playType"` // low：标清，high：高清
	Duration  int64  `json:"duration"` // 时长(单位:s)
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
	Commission                 float64 `json:"commission"`
	CommissionShare            float64 `json:"commissionShare"`
	CouponCommission           float64 `json:"couponCommission"` // 券后佣金，（促销价-优惠券面额）*佣金比例
	EndTime                    int64   `json:"endTime"`
	IsLock                     int64   `json:"isLock"` // 是否锁定佣金比例：1是，0否
	PlusCommissionShare        float64 `json:"plusCommissionShare"`
	StartTime                  int64   `json:"startTime"`
	LhCommissionStockMax       int64   `json:"lhCommissionStockMax"`       // 限量高佣总库存
	LhCommissionStockRemaining int64   `json:"lhCommissionStockRemaining"` // 限量高佣剩余库存
}

type CouponInfo struct {
	CouponList []CouponList `json:"couponList"`
}

type CouponList struct {
	BindType           int64            `json:"bindType"` // 券种类 (优惠券种类：0 - 全品类，1 - 限品类（自营商品），2 - 限店铺，3 - 店铺限商品券)
	Discount           float64          `json:"discount"`
	GetEndTime         int64            `json:"getEndTime"`
	GetStartTime       int64            `json:"getStartTime"`
	HotValue           *int64           `json:"hotValue,omitempty"`
	IsBest             int64            `json:"isBest"`
	IsInputCoupon      int64            `json:"isInputCoupon"`
	Link               string           `json:"link"`
	PlatformType       int64            `json:"platformType"` // 券使用平台 (平台类型：0 - 全平台券，1 - 限平台券)
	Quota              float64          `json:"quota"`
	UseEndTime         int64            `json:"useEndTime"`
	UseStartTime       int64            `json:"useStartTime"`
	CouponStyle        int64            `json:"couponStyle"`  // 优惠券分类 0：满减券，3：满折券，28: 每满减券
	CouponStatus       int64            `json:"couponStatus"` // 领取状态。0： 正常可领， -1 ：不可领取， 1： 已领取
	TimeCouponInfoList []TimeCouponInfo `json:"timeCouponInfoList"`
	LimitPlatforms     []int64          `json:"limitPlatforms"` // 券可用平台。不返回时表示全部可用，返回时表示限返回的平台可用。枚举值：0-PC端、1-京东App、3-M站、4-手机QQ、5-微信、6-QQ浏览器。
}

// TimeCouponInfo 时段券信息
type TimeCouponInfo struct {
	TimeCouponBegin string `json:"timeCouponBegin"` // 时段券领取开始时间
	TimeCouponEnd   string `json:"timeCouponEnd"`   // 时段券领取结束时间
}

type ImageInfo struct {
	ImageList  []ImageList `json:"imageList"`
	WhiteImage string      `json:"whiteImage"`
}

type ImageList struct {
	URL string `json:"url"`
}

type PriceInfo struct {
	LowestPrice       float64 `json:"lowestPrice"`
	LowestPriceType   int64   `json:"lowestPriceType"`   // 促销价类型，1：商品价格；2：拼购价格； 3：秒杀价格； 4：预售价格
	LowestCouponPrice float64 `json:"lowestCouponPrice"` // 券后价（有无券都返回此字段，价格排序以此字段排序）
	Price             float64 `json:"price"`
	HistoryPriceDay   int64   `json:"historyPriceDay"` // 历史最低价天数（例：当前券后价最近180天最低）
}

type ShopInfo struct {
	AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade"`       // 售后服务评级（仅pop店铺有值）
	AfterServiceScore             string  `json:"afterServiceScore"`             // 售后服务评分（仅pop店铺有值）
	CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade"`   // 用户评价评级（仅pop店铺有值）
	LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade"` // 物流履约评级（仅pop店铺有值）
	LogisticsLvyueScore           string  `json:"logisticsLvyueScore"`           // 物流履约评分（仅pop店铺有值）
	ScoreRankRate                 string  `json:"scoreRankRate"`                 // 店铺风向标（仅pop店铺有值）
	ShopID                        int64   `json:"shopId"`
	ShopLabel                     string  `json:"shopLabel"` // 1：京东好店
	ShopLevel                     float64 `json:"shopLevel"`
	ShopName                      string  `json:"shopName"`
	UserEvaluateScore             string  `json:"userEvaluateScore"` // 用户评价评分（仅pop店铺有值）
}

type PurchasePriceInfo struct {
	BasisPriceType         int                  `json:"basisPriceType"` // 依据的价格类型，1、京东价 ,2 Plus价，7 粉丝价，8 新人价，9学生价，10 陪伴计划价（双价格新增）
	Code                   int                  `json:"code"`
	CouponList             []CouponList         `json:"couponList"`
	Message                string               `json:"message"`
	PromotionLabelInfoList []PromotionLabelInfo `json:"promotionLabelInfoList"`
	PurchaseNum            int                  `json:"purchaseNum"`
	PurchasePrice          float64              `json:"purchasePrice"`  // 到手价
	ThresholdPrice         float64              `json:"thresholdPrice"` // 门槛价金额，计算到手价的基准价
	PlusBenefitInfo        PlusBenefitInfo      `json:"plusBenefitInfo"`
}

// PlusBenefitInfo PLUS 权益明细
type PlusBenefitInfo struct {
	Plus95Benefit          float64 `json:"plus95Benefit"`          // plus95折优惠金额
	ParallelSubsidyBenefit float64 `json:"parallelSubsidyBenefit"` // puls专享立减贴金额
	PlusSuperBenefit       float64 `json:"plusSuperBenefit"`       // plus超级津贴优惠金额
}

type PromotionLabelInfo struct {
	EndTime           int64    `json:"endTime"`
	LabelName         string   `json:"labelName"`      // 促销标签名称
	PromotionLabel    string   `json:"promotionLabel"` // 商品促销文案
	PromotionLabelId  int64    `json:"promotionLabelId"`
	StartTime         int64    `json:"startTime"`
	LableName         string   `json:"lableName"`         // 促销标签名称【废弃】
	PromotionLableId  int64    `json:"promotionLableId"`  // 促销ID【废弃】
	ProvinceNameList  []string `json:"provinceNameList"`  // 促销生效区域（国补）--省中文名
	SubType           int64    `json:"subType"`           // 促销类型（国补），9105：以旧换新国补，9107：购新立减国补，9100： 支付立减
	Rebate            float64  `json:"rebate"`            // 促销优惠比例（国补），0到1之间的小数
	TopDiscount       float64  `json:"topDiscount"`       // 最大优惠金额（国补）
	PromotionDiscount float64  `json:"promotionDiscount"` // 首购礼金等明细金额
}
