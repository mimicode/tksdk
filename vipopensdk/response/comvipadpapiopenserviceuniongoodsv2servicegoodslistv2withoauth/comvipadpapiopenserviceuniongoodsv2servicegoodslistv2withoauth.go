package comvipadpapiopenserviceuniongoodsv2servicegoodslistv2withoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionGoodsV2Service
// goodsListV2WithOauth 2.0.0 获取联盟在推商品列表V2-需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=goodsListV2WithOauth
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
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

// Success 商品列表响应信息
type Success struct {
	GoodsInfoList  []GoodsInfo `json:"goodsInfoList"`  // 精选商品列表
	LastPage       bool        `json:"lastPage"`       // 是否最后一页
	NextPageOffset int64       `json:"nextPageOffset"` // 下一页查询偏移（查询下一页时，offset传该值）
	BatchNo        string      `json:"batchNo"`        // 排行版商品数据批次号：查询下一页数据时batchNo传该字段
	GoodsTips      string      `json:"goodsTips"`      // 商品列表异常情况的提示信息
}

// GoodsInfo 商品信息
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

// ActivityInfo 商品参与活动信息
type ActivityInfo struct {
	ActType           int    `json:"actType"`           // 活动类型:18-唯品快抢
	ActName           string `json:"actName"`           // 活动名称
	BeginTime         int64  `json:"beginTime"`         // 开始时间：时间戳，单位毫秒
	EndTime           int64  `json:"endTime"`           // 结束时间：时间戳，单位毫秒
	ForeShowBeginTime int64  `json:"foreShowBeginTime"` // 预热开始时间：时间戳，单位毫秒
}

// ChannelExclusiveCoupon 渠道专属红包
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
	ActivedAmount     int64  `json:"activedAmount"`     // 劵已激活的数量：查询详情时返回
	ReceiveUrl        string `json:"receiveUrl"`        // 专属领券页地址
}

// 以下是从其他包引入的类型，现在在当前包中定义

// StoreInfo 店铺信息
type StoreInfo struct {
	StoreId   int64  `json:"storeId"`   // 店铺id
	StoreName string `json:"storeName"` // 店铺名称
	StoreType int    `json:"storeType"` // 店铺类型：0-自营，1-MP
}

// GoodsCommentsInfo 商品评价信息
type GoodsCommentsInfo struct {
	CommentTotalCount int    `json:"commentTotalCount"` // 评价总数
	GoodCommentRate   string `json:"goodCommentRate"`   // 好评率
}

// StoreServiceCapability 店铺服务能力评价
type StoreServiceCapability struct {
	StoreScore            string `json:"storeScore"`            // 店铺评分
	StoreRankRate         string `json:"storeRankRate"`         // 店铺评分排名
	DeliverySpeedScore    string `json:"deliverySpeedScore"`    // 发货速度评分
	DescriptionMatchScore string `json:"descriptionMatchScore"` // 描述相符评分
	ServiceAttitudeScore  string `json:"serviceAttitudeScore"`  // 服务态度评分
}

// PrepayInfo 商品预付信息
type PrepayInfo struct {
	IsPrepay  int    `json:"isPrepay"`  // 是否预付商品：0-否，1-是
	PrepayFee string `json:"prepayFee"` // 预付定金金额：单位-元
}

// PMSCouponInfo 红包信息
type PMSCouponInfo struct {
	CouponNo          string `json:"couponNo"`          // 优惠券批次号
	CouponName        string `json:"couponName"`        // 优惠劵名称
	Fav               string `json:"fav"`               // 优惠金额：单位-元
	Buy               string `json:"buy"`               // 使用门槛：单位-元
	ActivateBeginTime int64  `json:"activateBeginTime"` // 券激活开始时间,毫秒级时间戳
	ActivateEndTime   int64  `json:"activateEndTime"`   // 券激活结束时间,毫秒级时间戳
	UseBeginTime      int64  `json:"useBeginTime"`      // 使用开始时间，毫秒级时间戳
	UseEndTime        int64  `json:"useEndTime"`        // 使用结束时间，毫秒级时间戳
	TotalAmount       int64  `json:"totalAmount"`       // 生成劵的总量
	ActivedAmount     int64  `json:"activedAmount"`     // 劵已激活的数量
}

// GoodsSkuInfo 商品SKU信息
type GoodsSkuInfo struct {
	SkuId              string `json:"skuId"`              // skuId
	SkuName            string `json:"skuName"`            // sku名称
	SkuPrice           string `json:"skuPrice"`           // sku价格
	SkuThumbUrl        string `json:"skuThumbUrl"`        // sku缩略图
	SkuSalesVolume     int    `json:"skuSalesVolume"`     // sku销量
	SkuSaleStockStatus int    `json:"skuSaleStockStatus"` // sku库存状态：1-已抢光，2-有库存，3-有机会
}

// GoodsPromotionInfo 商品促销信息
type GoodsPromotionInfo struct {
	PromotionType int    `json:"promotionType"` // 促销类型：1-满减，2-满折，3-秒杀，4-限时折扣，5-限时包邮，6-特价，7-买赠，8-满赠，9-满减送，10-满折送，11-定金立减，12-阶梯价，13-单品加价购，14-单品换购，15-第二件半价，16-满件折，17-会员价，18-单品特价，19-尾货价，20-企业内购，21-精品集结，22-单品直降，23-一口价，24-拼团，25-爆款清仓，26-大促活动，27-新人价，28-爆款抢购，29-预售，30-预售优惠，31-特卖，32-特卖单品，33-特卖满减，34-特卖满折，35-特卖满件折，36-特卖满赠，37-特卖满减送，38-特卖满折送，39-特卖买赠，40-特卖加价购，41-特卖换购，42-特卖第二件半价，43-特卖满件减，44-特卖满额减，45-特卖满额折，46-特卖满额赠，47-特卖满额减赠，48-特卖满额折赠，49-特卖满件赠，50-特卖满件减赠，51-特卖满件折赠，52-特卖满件折，53-特卖满额折，54-特卖满额赠，55-特卖满额减赠，56-特卖满额折赠，57-特卖满件赠，58-特卖满件减赠，59-特卖满件折赠，60-特卖满件折，61-特卖满额折，62-特卖满额赠，63-特卖满额减赠，64-特卖满额折赠，65-特卖满件赠，66-特卖满件减赠，67-特卖满件折赠，68-特卖满件折，69-特卖满额折，70-特卖满额赠，71-特卖满额减赠，72-特卖满额折赠，73-特卖满件赠，74-特卖满件减赠，75-特卖满件折赠，76-特卖满件折，77-特卖满额折，78-特卖满额赠，79-特卖满额减赠，80-特卖满额折赠，81-特卖满件赠，82-特卖满件减赠，83-特卖满件折赠，84-特卖满件折，85-特卖满额折，86-特卖满额赠，87-特卖满额减赠，88-特卖满额折赠，89-特卖满件赠，90-特卖满件减赠，91-特卖满件折赠，92-特卖满件折，93-特卖满额折，94-特卖满额赠，95-特卖满额减赠，96-特卖满额折赠，97-特卖满件赠，98-特卖满件减赠，99-特卖满件折赠
	PromotionDesc string `json:"promotionDesc"` // 促销描述
}

// CampaignCommissionInfo 商单补贴信息
type CampaignCommissionInfo struct {
	CampaignType int    `json:"campaignType"` // 商单类型：1-超级单品，2-全店推广，3-超级单品补贴，4-全店推广补贴
	CampaignName string `json:"campaignName"` // 商单名称
	Rate         string `json:"rate"`         // 商单佣金比例（%）
	Commission   string `json:"commission"`   // 商单佣金金额（元）
	Subsidy      string `json:"subsidy"`      // 商单补贴金额（元）
	StartTime    int64  `json:"startTime"`    // 商单开始时间：时间戳，单位：毫秒
	EndTime      int64  `json:"endTime"`      // 商单结束时间：时间戳，单位：毫秒
}
