package comvipadpapiopenserviceuniongoodsservicegoodslistv2withoauth

import (
	"encoding/json"
	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService 获取联盟在推商品列表V2-需要oauth授权
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
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

type Result struct {
	LastPage       bool            `json:"lastPage"`
	GoodsInfoList  []GoodsInfoList `json:"goodsInfoList"`
	NextPageOffset int64           `json:"nextPageOffset"`
}

/*
goodsId	String	否			商品id
goodsName	String	否			商品名称
goodsDesc	String	否			商品描述,字段暂不输出
destUrl	String	否			商品落地页
goodsThumbUrl	String	否			商品缩略图
goodsCarouselPictures	List<String>	否			商品轮播图：根据商品id查询时返回，商品列表不返回
goodsMainPicture	String	否			商品主图
categoryId	Long	否			商品三级分类id
categoryName	String	否			商品三级分类
sourceType	Integer	否			商品类型：0-自营，1-MP
marketPrice	String	否			市场价（元）
vipPrice	String	否			唯品价（元）
commissionRate	String	否			佣金比例（%）
commission	String	否			佣金金额（元）
discount	String	否			折扣:唯品价/市场价 保留两位小数字符串
goodsDetailPictures	List<String>	否			商品详情图片：根据商品id查询商品信息时返回，商品列表不返回
cat1stId	Long	否			商品一级分类id
cat1stName	String	否			商品一级分类名称
cat2ndId	Long	否			商品二级分类id
cat2ndName	String	否			商品二级分类名称
brandStoreSn	String	否			商品品牌sn
brandName	String	否			商品品牌名称
brandLogoFull	String	否			商品品牌logo全路径地址
schemeEndTime	Long	否			商品推广计划有效期预估截止时间：仅为预估时间，仅做参考；时间戳，单位：毫秒
sellTimeFrom	Long	否			商品售卖开始时间,时间戳，单位毫秒
sellTimeTo	Long	否			商品售卖结束时间,时间戳, 单位毫秒
weight	Integer	否			推广权重，用于确定推广该商品的优先级，权重值越大，优先级越高
storeInfo	StoreInfo	否			店铺信息
commentsInfo	GoodsCommentsInfo	否			商品评价信息
storeServiceCapability	StoreServiceCapability	否			商品所属店铺服务能力评价
brandId	Long	否			商品所属档期（专场）id
schemeStartTime	Long	否			商品所属推广方案开始时间：时间戳，单位：毫秒
saleStockStatus	Integer	否			商品库存状态：1-已抢光，2-有库存，3-有机会,当列表查询库存或者查询商品详情时返回
status	Integer	否			商品状态：0-下架，1-上架
prepayInfo	PrepayInfo	否			商品预付信息
joinedActivities	List<ActivityInfo>	否			商品参与活动信息:未参与活动集合为空
couponInfo	PMSCouponInfo	否			红包信息
haiTao	Integer	否			是否海淘商品标识：1是 0不是
*/
type GoodsInfoList struct {
	MarketPrice           string      `json:"marketPrice"`
	CommissionRate        string      `json:"commissionRate"`
	GoodsID               string      `json:"goodsId"`
	Discount              string      `json:"discount"`
	CategoryName          string      `json:"categoryName"`
	HaiTAO                int64       `json:"haiTao"`
	GoodsCarouselPictures []string    `json:"goodsCarouselPictures"`
	Cat2NdName            string      `json:"cat2ndName"`
	Cat1StName            string      `json:"cat1stName"`
	VipPrice              string      `json:"vipPrice"`
	Commission            string      `json:"commission"`
	Cat1StID              int64       `json:"cat1stId"`
	GoodsName             string      `json:"goodsName"`
	BrandName             string      `json:"brandName"`
	BrandLogoFull         string      `json:"brandLogoFull"`
	BrandStoreSn          string      `json:"brandStoreSn"`
	SellTimeFrom          int64       `json:"sellTimeFrom"`
	SchemeStartTime       int64       `json:"schemeStartTime"`
	SchemeEndTime         int64       `json:"schemeEndTime"`
	SourceType            int64       `json:"sourceType"`
	SellTimeTo            int64       `json:"sellTimeTo"`
	BrandID               int64       `json:"brandId"`
	GoodsThumbURL         string      `json:"goodsThumbUrl"`
	Cat2NdID              int64       `json:"cat2ndId"`
	StoreInfo             StoreInfo   `json:"storeInfo"`
	GoodsMainPicture      string      `json:"goodsMainPicture"`
	DestURL               string      `json:"destUrl"`
	CategoryID            int64       `json:"categoryId"`
	Status                int64       `json:"status"`
	CouponInfo            *CouponInfo `json:"couponInfo,omitempty"`
}

/*
couponNo	String	否			优惠券批次号
couponName	String	否			优惠劵名称
buy	String	否			使用门槛
fav	String	否			优惠金额
*/
type CouponInfo struct {
	CouponName string `json:"couponName"`
	Buy        string `json:"buy"`
	CouponNo   string `json:"couponNo"`
	Fav        string `json:"fav"`
}

type StoreInfo struct {
	StoreName string `json:"storeName"`
	StoreID   string `json:"storeId"`
}

type SortField struct {
	FieldName string `json:"fieldName"`
	FieldDesc string `json:"fieldDesc"`
}

/*
comments	Integer	否			商品评论数
goodCommentsShare	String	否			商品好评率:百分比，不带百分号
*/
