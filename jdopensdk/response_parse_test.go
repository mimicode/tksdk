package jdopensdk

import (
	"testing"

	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsbigfieldquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsrankquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenorderagentquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotionintelligencequery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotiontoolsintelligencequery"
)

// 临时验证：按官方文档字段表构造的真实响应形状能否被 WrapResult 正确解析
func TestParseBigFieldResponseShape(t *testing.T) {
	body := `{"jd_union_open_goods_bigfield_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"requestId\":\"req_1\",\"data\":[{\"skuId\":1111,\"skuName\":\"手机\",\"categoryInfo\":{\"cid1\":6144,\"cid1Name\":\"珠宝首饰\",\"cid2\":12041,\"cid2Name\":\"木手串/把件\",\"cid3\":12052,\"cid3Name\":\"其他\"},\"imageInfo\":{\"imageList\":[{\"url\":\"http://img14.360buyimg.com/a.jpg\"}],\"whiteImage\":\"https://img14.360buyimg.com/w.png\"},\"baseBigFieldInfo\":{\"wdis\":\"<p>介绍</p>\",\"propCode\":\"无\",\"wareQD\":\"清单\",\"propGroups\":\"[]\"},\"bookBigFieldInfo\":{\"comments\":\"无\",\"image\":\"无\",\"contentDesc\":\"无\",\"relatedProducts\":\"无\",\"editerDesc\":\"无\",\"catalogue\":\"无\",\"bookAbstract\":\"无\",\"authorDesc\":\"无\",\"introduction\":\"无\",\"productFeatures\":\"无\"},\"videoBigFieldInfo\":{\"comments\":\"无\",\"image\":\"无\",\"contentDesc\":\"无\",\"editerDesc\":\"无\",\"catalogue\":\"无\",\"box_Contents\":\"无\",\"material_Description\":\"无\",\"manual\":\"无\",\"productFeatures\":\"无\"},\"mainSkuId\":6161111,\"productId\":1236547,\"skuStatus\":1,\"owner\":\"g\",\"detailImages\":\"http://img30.360buyimg.com/d.jpg\",\"itemId\":\"Q9Z2ZdyM\",\"callerItemId\":\"UOWr2TlQ\"}]}"}}`
	r := &jdunionopengoodsbigfieldquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("bigfield parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.SkuID != 1111 || g.SkuName != "手机" || g.CategoryInfo.Cid1Name != "珠宝首饰" ||
		g.ImageInfo.ImageList[0].URL != "http://img14.360buyimg.com/a.jpg" ||
		g.BaseBigFieldInfo.Wdis != "<p>介绍</p>" ||
		g.VideoBigFieldInfo.BoxContents != "无" ||
		g.ItemID != "Q9Z2ZdyM" {
		t.Fatalf("bigfield fields mismatch: %+v", g)
	}
}

func TestParseRankResponseShape(t *testing.T) {
	body := `{"jd_union_open_goods_rank_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"totalCount\":100,\"requestId\":\"req_2\",\"data\":[{\"itemId\":\"Q9Z2ZdyM\",\"skuId\":26898778009,\"skuName\":\"笔记本\",\"imageUrl\":\"https://img14.360buyimg.com/p.png\",\"imgList\":[\"https://img14.360buyimg.com/p.png\"],\"wlprice\":100.0,\"commissionShare\":5.0,\"comments\":30,\"goodComments\":15,\"goodCommentsShare\":99.0,\"skuTagList\":[{\"type\":3,\"index\":1,\"name\":\"7天无理由退货\"}],\"purchasePriceInfo\":{\"purchasePrice\":95.0,\"promotionLabelInfoList\":[{\"promotionLabelId\":\"5000125161\",\"labelName\":\"满折\"}],\"couponList\":[{\"link\":\"http://coupon.jd.com/x\",\"discount\":30.0,\"quota\":39.0,\"remainCnt\":1000,\"couponStyle\":0,\"couponStatus\":-1,\"timeCouponInfoList\":[{\"timeCouponBegin\":\"00:00:00\",\"timeCouponEnd\":\"08:59:59\"}]}]}}]}"}}`
	r := &jdunionopengoodsrankquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("rank parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.SkuID != 26898778009 || g.SkuName != "笔记本" || g.Wlprice != 100 ||
		g.SkuTagList[0].Name != "7天无理由退货" ||
		g.PurchasePriceInfo.PurchasePrice != 95 ||
		g.PurchasePriceInfo.CouponList[0].Discount != 30 ||
		g.PurchasePriceInfo.CouponList[0].TimeCouponInfoList[0].TimeCouponEnd != "08:59:59" ||
		r.Responce.QueryResult.TotalCount != 100 {
		t.Fatalf("rank fields mismatch: %+v", g)
	}
}

func TestParseOrderAgentResponseShape(t *testing.T) {
	body := `{"jd_union_open_order_agent_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"hasMore\":true,\"data\":[{\"id\":\"415900297816660001\",\"orderId\":108618000005,\"parentId\":0,\"orderTime\":\"2020-01-02 15:50:16\",\"finishTime\":\"2020-01-03 15:59:16\",\"modifyTime\":\"2020-01-02 16:01:03\",\"orderEmt\":2,\"plus\":0,\"unionId\":1000618618,\"skuId\":44303679033,\"skuName\":\"盲盒\",\"skuNum\":1,\"skuReturnNum\":0,\"skuFrozenNum\":0,\"price\":69.0,\"commissionRate\":5.0,\"subSideRate\":90.0,\"subsidyRate\":0.0,\"finalRate\":90.0,\"estimateCosPrice\":54.0,\"estimateFee\":2.43,\"actualCosPrice\":0.0,\"actualFee\":0.0,\"validCode\":16,\"traceType\":2,\"positionId\":0,\"siteId\":61866,\"unionAlias\":\"**平台\",\"pid\":\"618_618_618\",\"cid1\":1620,\"cid2\":11158,\"cid3\":11969,\"subUnionId\":\"331967\",\"unionTag\":\"0000000000000001\",\"popId\":709982,\"ext1\":\"hello_world\",\"payMonth\":\"20200120\",\"cpActId\":0,\"unionRole\":1,\"giftCouponOcsAmount\":0,\"giftCouponKey\":\"xxx_coupon_key\",\"sign\":\"B44C0FC3F104167FEC8A53DFD2B26E40\",\"proPriceAmount\":6.18,\"goodsInfo\":{\"imageUrl\":\"http://img14.360buyimg.com/a.jpg\",\"owner\":\"g\",\"mainSkuId\":6161111,\"productId\":1236547,\"shopName\":\"XXXX旗舰店\",\"shopId\":45619},\"categoryInfo\":{\"cid1\":1,\"cid2\":2,\"cid3\":3,\"cid1Name\":\"珠宝首饰\",\"cid2Name\":\"木手串/把件\",\"cid3Name\":\"其他\"},\"expressStatus\":10,\"channelId\":1,\"skuTag\":\"0000000000000000000000000000000000000000000000000000000000000001\",\"rid\":37843,\"itemId\":\"Q9Z2ZdyM\",\"callerItemId\":\"UOWr2TlQ\",\"subCpUnionId\":1001}]}"}}`
	r := &jdunionopenorderagentquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("agent parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ID != "415900297816660001" || g.OrderID != 108618000005 || g.SkuID != 44303679033 ||
		g.EstimateFee != 2.43 || g.ValidCode != 16 || g.PayMonth != "20200120" ||
		g.GoodsInfo.Owner != "g" || g.GoodsInfo.ShopID != 45619 ||
		g.CategoryInfo.Cid1Name != "珠宝首饰" ||
		g.CallerItemID != "UOWr2TlQ" || g.SubCpUnionID != 1001 ||
		!r.Responce.QueryResult.HasMore {
		t.Fatalf("agent fields mismatch: %+v", g)
	}
}

func TestParseIntelligenceResponseShape(t *testing.T) {
	body := `{"jd_union_open_promotion_intelligence_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"data\":[{\"reportContent\":\"特价\",\"type\":1,\"cid1List\":[52253,75691],\"status\":1,\"essence\":1,\"startTime\":\"2021-11-11 12:00:00\",\"endTime\":\"2021-11-18 12:00:00\"}]}"}}`
	r := &jdunionopenpromotionintelligencequery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("intelligence parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ReportContent != "特价" || g.Type != 1 || g.Status != 1 || g.Essence != 1 ||
		g.Cid1List[0] != 52253 || g.Cid1List[1] != 75691 ||
		g.StartTime != "2021-11-11 12:00:00" || g.EndTime != "2021-11-18 12:00:00" {
		t.Fatalf("intelligence fields mismatch: %+v", g)
	}
}

func TestParseToolsIntelligenceResponseShape(t *testing.T) {
	body := `{"jd_union_open_promotion_tools_intelligence_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"data\":[{\"reportContent\":\"特价\",\"type\":2,\"cid1List\":[52253],\"status\":2,\"essence\":0,\"startTime\":\"2021-11-11 12:00:00\",\"endTime\":\"2021-11-18 12:00:00\"}]}"}}`
	r := &jdunionopenpromotiontoolsintelligencequery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("tools intelligence parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ReportContent != "特价" || g.Type != 2 || g.Status != 2 || g.Essence != 0 ||
		g.Cid1List[0] != 52253 || g.EndTime != "2021-11-18 12:00:00" {
		t.Fatalf("tools intelligence fields mismatch: %+v", g)
	}
}
