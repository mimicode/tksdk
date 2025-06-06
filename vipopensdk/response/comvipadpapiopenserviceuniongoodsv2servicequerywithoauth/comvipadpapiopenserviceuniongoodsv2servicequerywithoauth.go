package comvipadpapiopenserviceuniongoodsv2servicequerywithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionGoodsV2Service
// queryWithOauth 2.0.0 根据关键词查询商品列表-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=queryWithOauth
type Response struct {
	response.TopResponse
	Success Success `json:"result"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Success struct {
	GoodsInfoList    []GoodsInfo `json:"goodsInfoList"`    // 精选商品列表
	Total            int         `json:"total"`            // 总数
	SortFields       []SortField `json:"sortFields"`       // 支持的排序字段,为空时仅支持默认排序
	Page             int         `json:"page"`             // 当前页码
	PageSize         int         `json:"pageSize"`         // 分页大小
	RecommendKeyword string      `json:"recommendKeyword"` // 推荐词搜索时使用的推荐词
	BatchNo          string      `json:"batchNo"`          // 排行版商品数据批次号：查询下一页数据时batchNo传该字段
}

type GoodsInfo struct {
	GoodsId                string                 `json:"goodsId"`                // 商品id
	GoodsName              string                 `json:"goodsName"`              // 商品名称
	GoodsDesc              string                 `json:"goodsDesc"`              // 商品描述,字段暂不输出
	DestUrl                string                 `json:"destUrl"`                // 商品落地页
	GoodsThumbUrl          string                 `json:"goodsThumbUrl"`          // 商品缩略图
	GoodsCarouselPictures  []string               `json:"goodsCarouselPictures"`  // 商品轮播图：根据商品id查询时返回，商品列表不返回
	GoodsMainPicture       string                 `json:"goodsMainPicture"`       // 商品主图
	CategoryId             int64                  `json:"categoryId"`             // 商品三级分类id
	CategoryName           string                 `json:"categoryName"`           // 商品三级分类
	SourceType             int                    `json:"sourceType"`             // 商品类型：0-自营，1-MP
	MarketPrice            string                 `json:"marketPrice"`            // 市场价（元）
	VipPrice               string                 `json:"vipPrice"`               // 唯品价（元）
	CommissionRate         string                 `json:"commissionRate"`         // 佣金比例（%）
	Commission             string                 `json:"commission"`             // 佣金金额（元）
	Discount               string                 `json:"discount"`               // 折扣:唯品价/市场价 保留两位小数字符串
	GoodsDetailPictures    []string               `json:"goodsDetailPictures"`    // 商品详情图片：根据商品id查询商品信息时返回，商品列表不返回
	Cat1stId               int64                  `json:"cat1stId"`               // 商品一级分类id
	Cat1stName             string                 `json:"cat1stName"`             // 商品一级分类名称
	Cat2ndId               int64                  `json:"cat2ndId"`               // 商品二级分类id
	Cat2ndName             string                 `json:"cat2ndName"`             // 商品二级分类名称
	BrandStoreSn           string                 `json:"brandStoreSn"`           // 商品品牌sn
	BrandName              string                 `json:"brandName"`              // 商品品牌名称
	BrandLogoFull          string                 `json:"brandLogoFull"`          // 商品品牌logo全路径地址
	SchemeEndTime          int64                  `json:"schemeEndTime"`          // 商品推广计划有效期预估截止时间：仅为预估时间，仅做参考；时间戳，单位：毫秒
	SellTimeFrom           int64                  `json:"sellTimeFrom"`           // 商品售卖开始时间,时间戳，单位毫秒
	SellTimeTo             int64                  `json:"sellTimeTo"`             // 商品售卖结束时间,时间戳, 单位毫秒
	Weight                 int                    `json:"weight"`                 // 推广权重，用于确定推广该商品的优先级，权重值越大，优先级越高
	StoreInfo              StoreInfo              `json:"storeInfo"`              // 店铺信息
	CommentsInfo           GoodsCommentsInfo      `json:"commentsInfo"`           // 商品评价信息
	StoreServiceCapability StoreServiceCapability `json:"storeServiceCapability"` // 商品所属店铺服务能力评价
	BrandId                int64                  `json:"brandId"`                // 商品所属档期（专场）id
	SchemeStartTime        int64                  `json:"schemeStartTime"`        // 商品所属推广方案开始时间：时间戳，单位：毫秒
	SaleStockStatus        int                    `json:"saleStockStatus"`        // 商品库存状态：1-已抢光，2-有库存，3-有机会,当列表查询库存或者查询商品详情时返回
	Status                 int                    `json:"status"`                 // 商品状态：0-下架，1-上架
	PrepayInfo             PrepayInfo             `json:"prepayInfo"`             // 商品预付信息
	JoinedActivities       []ActivityInfo         `json:"joinedActivities"`       // 商品参与活动信息:未参与活动集合为空
	CouponInfo             PMSCouponInfo          `json:"couponInfo"`             // 红包信息
	HaiTao                 int                    `json:"haiTao"`                 // 是否海淘商品标识：1是  0不是
	SpuId                  string                 `json:"spuId"`                  // 商品spuId
	GoodsIdsWithSameSpu    []string               `json:"goodsIdsWithSameSpu"`    // 同spuId扩展商品id: 新版详情接口返回，其余接口不返回该字段
	SkuInfos               []GoodsSkuInfo         `json:"skuInfos"`               // 对应sku信息
	ExclusiveCoupon        ChannelExclusiveCoupon `json:"exclusiveCoupon"`        // 渠道专属红包：目前仅开放单品券，没有则返回空
	CpsInfo                map[string]interface{} `json:"cpsInfo"`                // cps推广信息：目前只返回小程序链接和通用追踪参数，其他链接请移步转链接口,1-通用推广参数(tra_from)，2-唯品会微信小程序链接
	Sn                     string                 `json:"sn"`                     // 商品货号/商品原编号
	TagNames               []string               `json:"tagNames"`               // 商品标签
	WhiteImage             string                 `json:"whiteImage"`             // 商品透明底图
	FuturePriceMsg         string                 `json:"futurePriceMsg"`         // 商品未来价信息
	IsSubsidyActivityGoods bool                   `json:"isSubsidyActivityGoods"` // 是否为补贴活动商品标识，默认为false
	SubsidyActivityAmount  string                 `json:"subsidyActivityAmount"`  // 补贴活动奖励金额（单位：元）
	SubsidyTaskNo          string                 `json:"subsidyTaskNo"`          // 补贴活动任务编码
	CouponPriceType        int                    `json:"couponPriceType"`        // 券后价类型：0-否，1-公开券后价，2-隐藏券后价
	EstimatePrice          string                 `json:"estimatePrice"`          // 商品预估价格
	GoodsSoldNumDesc       string                 `json:"goodsSoldNumDesc"`       // 商品销量文案(仅2小时榜单和今日榜单会返回)，例如：1万+
	ProductSales           string                 `json:"productSales"`           // 商品销量
	DestUrlPc              string                 `json:"destUrlPc"`              // PC端商品落地页
	AdCode                 string                 `json:"adCode"`                 // 用来标识获取推广物料的来源，请准确保存，并在调用转链接口时传入，防止错用
	GoodsPromotionInfo     GoodsPromotionInfo     `json:"goodsPromotionInfo"`     // 商品促销信息
	IsAllowanceGoods       int                    `json:"isAllowanceGoods"`       // 商品是否命中超级补贴，1：是，0：否
	Allowance              string                 `json:"allowance"`              // 补贴金额，单位元
	AllowanceStartTime     int64                  `json:"allowanceStartTime"`     // 补贴开始时间
	AllowanceEndTime       int64                  `json:"allowanceEndTime"`       // 补贴结束时间
	CampaignCommissionInfo CampaignCommissionInfo `json:"campaignCommissionInfo"` // 商单补贴信息（仅2.0版本服务中支持，1.x版本服务不支持该功能）
}

type StoreInfo struct {
	StoreId   string `json:"storeId"`   // 店铺id
	StoreName string `json:"storeName"` // 店铺名称
}

type GoodsCommentsInfo struct {
	Comments          int    `json:"comments"`          // 商品评论数
	GoodCommentsShare string `json:"goodCommentsShare"` // 商品好评率:百分比，不带百分号
	CommentsText      string `json:"commentsText"`      // 商品第一条评论
}

type StoreServiceCapability struct {
	StoreScore    string `json:"storeScore"`    // 店铺评分：保留两位小数
	StoreRankRate string `json:"storeRankRate"` // 店铺同品类排名比例：例如10表示前10%
}

type PrepayInfo struct {
	IsPrepay             int    `json:"isPrepay"`             // 是否预付商品:0-否，1-是
	PrepayPrice          string `json:"prepayPrice"`          // 预付到手价：元
	FirstAmount          string `json:"firstAmount"`          // 预付首款金额：元
	LastAmount           string `json:"lastAmount"`           // 预付尾款金额：元
	PrepayFavAmount      string `json:"prepayFavAmount"`      // 预付优惠金额: 元
	DeductionPrice       string `json:"deductionPrice"`       // 抵扣价格(首款+优惠金额)：元
	PrepayDiscount       string `json:"prepayDiscount"`       // 预付折扣：(唯品价-优惠金额)/唯品价 保留两位小数的数字字符串
	PrepayFirstStartTime int64  `json:"prepayFirstStartTime"` // 首款支付开始时间:时间戳,单位毫秒
	PrepayFirstEndTime   int64  `json:"prepayFirstEndTime"`   // 首款支付结束时间:时间戳,单位毫秒
	PrepayLastStartTime  int64  `json:"prepayLastStartTime"`  // 尾款支付开始时间:时间戳,单位毫秒
	PrepayLastEndTime    int64  `json:"prepayLastEndTime"`    // 尾款支付结束时间:时间戳,单位毫秒
}

type ActivityInfo struct {
	ActType           int    `json:"actType"`           // 活动类型:18-唯品快抢
	ActName           string `json:"actName"`           // 活动名称
	BeginTime         int64  `json:"beginTime"`         // 开始时间：时间戳，单位毫秒
	EndTime           int64  `json:"endTime"`           // 结束时间：时间戳，单位毫秒
	ForeShowBeginTime int64  `json:"foreShowBeginTime"` // 预热开始时间：时间戳，单位毫秒
}

type PMSCouponInfo struct {
	CouponNo               string  `json:"couponNo"`               // 优惠券批次号
	CouponName             string  `json:"couponName"`             // 优惠劵名称
	Buy                    string  `json:"buy"`                    // 使用门槛
	Fav                    string  `json:"fav"`                    // 优惠金额
	ActivateBeginTime      int64   `json:"activateBeginTime"`      // 券激活开始时间,毫秒级时间戳
	ActivateEndTime        int64   `json:"activateEndTime"`        // 券激活结束时间,毫秒级时间戳
	UseBeginTime           int64   `json:"useBeginTime"`           // 使用开始时间，毫秒级时间戳
	UseEndTime             int64   `json:"useEndTime"`             // 使用结束时间，毫秒级时间戳
	TotalAmount            int64   `json:"totalAmount"`            // 生成劵的总量
	ActivatedAmount        int64   `json:"activedAmount"`          // 劵已激活的数量
	CouponType             byte    `json:"couponType"`             // 券类型,(1:买赠 2:满减 3:折扣 4:免邮 5:多减多减)
	VipPrice               float64 `json:"vipPrice"`               // 唯品价(用于券的展示逻辑)
	HiddenCouponReceiveUrl string  `json:"hiddenCouponReceiveUrl"` // 隐藏红包领券页链接
	Limit                  int64   `json:"limit"`                  // 限制使用的券数量
	Left                   int64   `json:"left"`                   // 剩余券数量（仅限招商商品隐藏红包）
}

type GoodsSkuInfo struct {
	SizeId             string                 `json:"sizeId"`             // 商品尺码id(唯品会体系下和skuId等同)
	SaleProps          map[string]interface{} `json:"saleProps"`          // 售卖属性信息（134 - 颜色,453 - 尺码）
	LeavingStock       int                    `json:"leavingStock"`       // 剩余库存:查询库存时返回
	SaleStockStatus    int                    `json:"saleStockStatus"`    // 商品库存状态：1-已抢光，2-有库存，3-有机会
	VipPrice           string                 `json:"vipPrice"`           // 唯品价:单位元
	MarketPrice        string                 `json:"marketPrice"`        // 市场价:单位元
	SizeSn             string                 `json:"sizeSn"`             // 条码
	GoodsPromotionInfo GoodsPromotionInfo     `json:"goodsPromotionInfo"` // 商品sizeId维度促销信息(目前仅UnionGoodsService-1.3.0#getGoodsDetailMarketing支持返回该字段)
	SquareImages       map[string]interface{} `json:"squareImages"`       // 方形图
}

type GoodsPromotionInfo struct {
	SalePrice         string          `json:"salePrice"`         // 到手价,根据人群等计算出该用户的到手价
	MarketPrice       string          `json:"marketPrice"`       // 市场价
	Discount          string          `json:"discount"`          // 折扣 保留两位小数字符串
	Svip              bool            `json:"svip"`              // 是否是超V价
	SalePriceDetail   string          `json:"salePriceDetail"`   // 到手价明细, 仅商详场景会返回
	SalePriceDesc     string          `json:"salePriceDesc"`     // 到手价说明
	LowPriceTag       string          `json:"lowPriceTag"`       // 低价标签文案，例如7天低价
	PgNewUser         string          `json:"pgNewUser"`         // 新客一口价优惠信息，有新客一口价时返回新客一口价字样
	OnWayActInfo      OnWayActInfo    `json:"onWayActInfo"`      // 进行中活动优惠信息，仅商详场景会返回
	FutureActInfo     FutureActInfo   `json:"futureActInfo"`     // 预告活动优惠信息，仅商详场景会返回
	AllowancePrice    string          `json:"allowancePrice"`    // 补贴活动优惠，有补贴活动时返回对应的补贴活动优惠文案
	CouponInfos       []PMSCouponInfo `json:"couponInfos"`       // 优惠券信息列表
	HiddenCouponInfo  PMSCouponInfo   `json:"hiddenCouponInfo"`  // 隐藏券信息
	NewUserSubsidyFav string          `json:"newUserSubsidyFav"` // 新人补贴优惠金额(仅新人补贴商品情况下才有值)
}

type OnWayActInfo struct {
	ActTips   string `json:"actTips"`   // 活动提示语
	EndTime   int64  `json:"endTime"`   // 活动结束时间，秒级GMT时间戳
	SalePrice string `json:"salePrice"` // 活动价，只有价格类活动才会返回
	Type      int    `json:"type"`      // 活动类型 0-价格类活动 1-折扣类活动
}

type FutureActInfo struct {
	ActTips         string `json:"actTips"`         // 活动提示语
	StartTime       int64  `json:"startTime"`       // 活动开始时间，秒级GMT时间戳
	FutureSalePrice string `json:"futureSalePrice"` // 预告价，只有价格类活动才会返回
	Type            int    `json:"type"`            // 活动类型 0-价格类活动 1-折扣类活动
}

type ChannelExclusiveCoupon struct {
	CouponNo          string `json:"couponNo"`          // 优惠券批次号
	CouponName        string `json:"couponName"`        // 优惠劵名称
	Fav               string `json:"fav"`               // 优惠金额：单位-元,查询详情时返回
	Buy               string `json:"buy"`               // 使用门槛：单位-元，查询详情时返回
	ActivateBeginTime int64  `json:"activateBeginTime"` // 券激活开始时间,毫秒级时间戳
	ActivateEndTime   int64  `json:"activateEndTime"`   // 券激活结束时间,毫秒级时间戳
	UseBeginTime      int64  `json:"useBeginTime"`      // 使用开始时间，毫秒级时间戳
	UseEndTime        int64  `json:"useEndTime"`        // 使用结束时间，毫秒级时间戳
	TotalAmount       int64  `json:"totalAmount"`       // 生成劵的总量：查询详情时返回
	ActivatedAmount   int64  `json:"activedAmount"`     // 劵已激活的数量：查询详情时返回
	ReceiveUrl        string `json:"receiveUrl"`        // 专属领券页地址
}

type CampaignCommissionInfo struct {
	IsCampaignCommission   int    `json:"isCampaignCommission"`   // 是否商单补贴 0-否 1-是
	CampaignCommissionRate string `json:"campaignCommissionRate"` // 商单补贴比例
	StartTime              int64  `json:"startTime"`              // 商单补贴开始时间
	EndTime                int64  `json:"endTime"`                // 商单补贴结束时间
}

type SortField struct {
	FieldName string `json:"fieldName"` // 排序字段
	FieldDesc string `json:"fieldDesc"` // 字段说明
}
