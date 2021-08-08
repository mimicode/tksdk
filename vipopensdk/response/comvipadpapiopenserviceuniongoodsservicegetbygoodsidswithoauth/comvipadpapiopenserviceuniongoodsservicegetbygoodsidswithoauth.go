package comvipadpapiopenserviceuniongoodsservicegetbygoodsidswithoauth

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/vipopensdk/response"
)

//com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id结合的商品信息-需要oauth授权
type Response struct {
	response2.TopResponse
	Result []Result `json:"result"`
}

//解析输出结果
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
	MarketPrice            string                 `json:"marketPrice"`
	CommissionRate         string                 `json:"commissionRate"`
	GoodsID                string                 `json:"goodsId"`
	Discount               string                 `json:"discount"`
	GoodsCarouselPictures  []string               `json:"goodsCarouselPictures"`
	GoodsDetailPictures    []string               `json:"goodsDetailPictures"`
	CategoryName           string                 `json:"categoryName"`
	HaiTAO                 int64                  `json:"haiTao"`
	PrepayInfo             PrepayInfo             `json:"prepayInfo"`
	Cat2NdName             string                 `json:"cat2ndName"`
	Cat1StName             string                 `json:"cat1stName"`
	VipPrice               string                 `json:"vipPrice"`
	Commission             string                 `json:"commission"`
	Cat1StID               int64                  `json:"cat1stId"`
	GoodsName              string                 `json:"goodsName"`
	StoreServiceCapability StoreServiceCapability `json:"storeServiceCapability"`
	BrandName              string                 `json:"brandName"`
	BrandLogoFull          string                 `json:"brandLogoFull"`
	BrandStoreSn           string                 `json:"brandStoreSn"`
	SellTimeFrom           int64                  `json:"sellTimeFrom"`
	SchemeStartTime        int64                  `json:"schemeStartTime"`
	CommentsInfo           CommentsInfo           `json:"commentsInfo"`
	SaleStockStatus        int64                  `json:"saleStockStatus"`
	SchemeEndTime          int64                  `json:"schemeEndTime"`
	SourceType             int64                  `json:"sourceType"`
	SellTimeTo             int64                  `json:"sellTimeTo"`
	BrandID                int64                  `json:"brandId"`
	GoodsThumbURL          string                 `json:"goodsThumbUrl"`
	Cat2NdID               int64                  `json:"cat2ndId"`
	StoreInfo              StoreInfo              `json:"storeInfo"`
	GoodsMainPicture       string                 `json:"goodsMainPicture"`
	DestURL                string                 `json:"destUrl"`
	CategoryID             int64                  `json:"categoryId"`
	Status                 int64                  `json:"status"`
	CouponInfo             CouponInfo             `json:"couponInfo"`
}

type CommentsInfo struct {
	Comments          int64  `json:"comments"`
	GoodCommentsShare string `json:"goodCommentsShare"`
}

type PrepayInfo struct {
	IsPrepay int64 `json:"isPrepay"`
}

type StoreInfo struct {
	StoreName string `json:"storeName"`
	StoreID   string `json:"storeId"`
}

type StoreServiceCapability struct {
	StoreScore    string `json:"storeScore"`
	StoreRankRate string `json:"storeRankRate"`
}

type CouponInfo struct {
	UseEndTime        int64  `json:"useEndTime"`
	TotalAmount       int64  `json:"totalAmount"`
	CouponName        string `json:"couponName"`
	ActivateEndTime   int64  `json:"activateEndTime"`
	Buy               string `json:"buy"`
	UseBeginTime      int64  `json:"useBeginTime"`
	CouponNo          string `json:"couponNo"`
	Fav               string `json:"fav"`
	ActivateBeginTime int64  `json:"activateBeginTime"`
	ActivedAmount     int64  `json:"activedAmount"`
}
