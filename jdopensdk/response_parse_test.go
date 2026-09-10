package jdopensdk

import (
	"testing"

	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenactivitybonusquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenchannelinvitecodeget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenchannelrelationget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenchannelrelationquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsbigfieldquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodscombinationpageget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsrankquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsrecommendquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenorderagentquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenorderquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenorderrowsupplyquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotionintelligencequery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotiontoolsintelligencequery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenshpromotionget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenstatisticsactivitybonusquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenstatisticspromotionquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenuserpidget"
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

func TestParseOrderResponseShape(t *testing.T) {
	body := `{"jd_union_open_order_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"成功\",\"hasMore\":false,\"data\":[{\"finishTime\":1529271683000,\"orderEmt\":2,\"orderId\":61861861866,\"orderTime\":1529271683000,\"parentId\":66661861866,\"payMonth\":\"20180618\",\"plus\":1,\"popId\":1000066618,\"unionId\":2018618618,\"ext1\":\"100_618_大促\",\"validCode\":18,\"skuList\":[{\"actualCosPrice\":6.88,\"actualFee\":6.18,\"commissionRate\":2.5,\"estimateCosPrice\":618.18,\"estimateFee\":6.18,\"finalRate\":100.0,\"cid1\":737,\"frozenSkuNum\":0,\"pid\":\"618_618_6018\",\"positionId\":66,\"price\":61.8,\"cid2\":738,\"siteId\":6000618,\"skuId\":5487565,\"skuName\":\"空气净化器\",\"skuNum\":2,\"skuReturnNum\":1,\"subSideRate\":90.0,\"subsidyRate\":10.0,\"cid3\":749,\"unionAlias\":\"**平台\",\"unionTag\":\"00000000000000000000000000000001\",\"unionTrafficGroup\":1,\"validCode\":18,\"subUnionId\":\"2018061\",\"traceType\":3,\"payMonth\":20180618,\"popId\":30871,\"ext1\":\"1_1255413_0_XXXX\",\"cpActId\":618,\"unionRole\":1,\"giftCouponKey\":\"xxx_coupon_key\",\"giftCouponOcsAmount\":100.0,\"proPriceAmount\":6.18,\"itemId\":\"Q9Z2ZdyM\"}]}]}"}}`
	r := &jdunionopenorderquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("order parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	o := r.Responce.QueryResult.Data[0]
	if o.OrderID != 61861861866 || o.PayMonth != "20180618" || o.UnionID != 2018618618 || o.ValidCode != 18 ||
		r.Responce.QueryResult.HasMore {
		t.Fatalf("order fields mismatch: %+v", o)
	}
	s := o.SkuList[0]
	if s.SkuID != 5487565 || s.ActualFee != 6.18 || s.UnionTrafficGroup != 1 ||
		s.PayMonth != 20180618 || s.PopID != 30871 || s.GiftCouponOcsAmount != 100 ||
		s.ItemID != "Q9Z2ZdyM" || s.SubUnionID != "2018061" {
		t.Fatalf("sku fields mismatch: %+v", s)
	}
}

func TestParseGoodsRecommendResponseShape(t *testing.T) {
	body := `{"jd_union_open_goods_recommend_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"data\":{\"matchSkuInfo\":{\"skuId\":26898778009,\"itemId\":\"Q9Z2ZdyM\"},\"recommendSkuInfoList\":[{\"skuId\":111,\"itemId\":\"AAA\",\"type\":2,\"subType\":2001,\"reason\":\"同类品价更优\"},{\"skuId\":222,\"itemId\":\"BBB\",\"type\":3,\"subType\":3002,\"reason\":\"相似品高销\"}]}}"}}`
	r := &jdunionopengoodsrecommendquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("recommend parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	d := r.Responce.QueryResult.Data
	if d.MatchSkuInfo.SkuID != 26898778009 || d.MatchSkuInfo.ItemID != "Q9Z2ZdyM" ||
		len(d.RecommendSkuInfoList) != 2 ||
		d.RecommendSkuInfoList[0].SubType != 2001 || d.RecommendSkuInfoList[0].Reason != "同类品价更优" ||
		d.RecommendSkuInfoList[1].Type != 3 {
		t.Fatalf("recommend fields mismatch: %+v", d)
	}
}

func TestParseCombinationPageResponseShape(t *testing.T) {
	body := `{"jd_union_open_goods_combinationpage_get_responce":{"getResult":"{\"code\":200,\"message\":\"成功\",\"data\":{\"shortURL\":\"https://u.jd.com/Ezv***\",\"clickURL\":\"https://union-click.jd.com/jdc?e=XXX\",\"failSkuList\":[{\"sku\":234345,\"message\":\"为预售商品\"}],\"materialId\":\"https://jingfen.jd.com/item?***\",\"failCouponList\":[{\"url\":\"http://coupon.jd.com/x\",\"message\":\"优惠券校验失败\"}],\"failActivityUrlList\":[]}}"}}`
	r := &jdunionopengoodscombinationpageget.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("combinationpage parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	d := r.Responce.GetResult.Data
	if d.ShortURL != "https://u.jd.com/Ezv***" || d.ClickURL != "https://union-click.jd.com/jdc?e=XXX" ||
		d.MaterialID != "https://jingfen.jd.com/item?***" ||
		d.FailSkuList[0].Sku != 234345 || d.FailSkuList[0].Message != "为预售商品" ||
		d.FailCouponList[0].Message != "优惠券校验失败" {
		t.Fatalf("combinationpage fields mismatch: %+v", d)
	}
}

func TestParseChannelRelationGetResponseShape(t *testing.T) {
	body := `{"jd_union_open_channel_relation_get_responce":{"getResult":"{\"code\":200,\"message\":\"success\",\"data\":{\"channelId\":100001}}"}}`
	r := &jdunionopenchannelrelationget.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("channel relation get parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	if r.Responce.GetResult.Data.ChannelID != 100001 {
		t.Fatalf("channel relation get fields mismatch: %+v", r.Responce.GetResult)
	}
}

func TestParseChannelRelationQueryResponseShape(t *testing.T) {
	body := `{"jd_union_open_channel_relation_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"totalCount\":100,\"data\":[{\"channelId\":100001,\"createTime\":\"2020-12-20 23:59:59\",\"note\":\"渠道备注名\",\"channelNote\":\"合作方备注名\",\"id\":1000010179}]}"}}`
	r := &jdunionopenchannelrelationquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("channel relation query parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ChannelID != 100001 || g.ID != 1000010179 || g.Note != "渠道备注名" ||
		g.ChannelNote != "合作方备注名" || g.CreateTime != "2020-12-20 23:59:59" ||
		r.Responce.QueryResult.TotalCount != 100 {
		t.Fatalf("channel relation query fields mismatch: %+v", g)
	}
}

func TestParseChannelInvitecodeGetResponseShape(t *testing.T) {
	body := `{"jd_union_open_channel_invitecode_get_responce":{"getResult":"{\"code\":200,\"message\":\"success\",\"data\":{\"inviteCode\":\"NCERT\"}}"}}`
	r := &jdunionopenchannelinvitecodeget.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("channel invitecode get parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	if r.Responce.GetResult.Data.InviteCode != "NCERT" {
		t.Fatalf("channel invitecode get fields mismatch: %+v", r.Responce.GetResult)
	}
}

func TestParseStatisticsPromotionResponseShape(t *testing.T) {
	body := `{"jd_union_open_statistics_promotion_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"data\":[{\"unionId\":10001,\"skuId\":10031228062854,\"activityUrl\":\"https://pro.m.jd.com/mall/active/xxx/index.html\",\"timeType\":1,\"dataType\":1,\"time\":\"2021-12-30 11\",\"clickPv\":1000,\"estimateValidOrders\":100,\"estimateValidFee\":100.0,\"estimateValidGmv\":5000.0,\"refundOrders\":10,\"completeOrders\":50,\"completeGmv\":2500.0,\"actualFee\":95.0,\"itemId\":\"Q9Z2ZdyM\"}]}"}}`
	r := &jdunionopenstatisticspromotionquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("statistics promotion parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ClickPv != 1000 || g.EstimateValidOrders != 100 || g.EstimateValidFee != 100 ||
		g.CompleteGmv != 2500 || g.ActualFee != 95 || g.RefundOrders != 10 ||
		g.UnionID != 10001 || g.ItemID != "Q9Z2ZdyM" {
		t.Fatalf("statistics promotion fields mismatch: %+v", g)
	}
}

func TestParseStatisticsActivityBonusResponseShape(t *testing.T) {
	body := `{"jd_union_open_statistics_activity_bonus_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"data\":{\"unionId\":5227,\"activityId\":1,\"estimateValidNum\":1000,\"estimateCosPrice\":10000.0,\"estimateBonus\":1000.0,\"actualValidNum\":1200,\"actualCosPrice\":12000.0,\"actualBonus\":1200.0,\"uv\":1,\"channelId\":100001}}"}}`
	r := &jdunionopenstatisticsactivitybonusquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("statistics activity bonus parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data
	if g.UnionID != 5227 || g.ActivityID != 1 || g.EstimateValidNum != 1000 ||
		g.EstimateBonus != 1000 || g.ActualValidNum != 1200 || g.ActualBonus != 1200 ||
		g.UV != 1 || g.ChannelID != 100001 {
		t.Fatalf("statistics activity bonus fields mismatch: %+v", g)
	}
}

func TestParseActivityBonusResponseShape(t *testing.T) {
	body := `{"jd_union_open_activity_bonus_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"hasMore\":true,\"data\":[{\"activityId\":1,\"activityName\":\"京橙周年庆\",\"prepareTime\":1616603400000,\"beginDate\":1616603400000,\"endDate\":1621180799000,\"payType\":1,\"firstPayTime\":1622390400000,\"firstPayRate\":100.0,\"secondPayTime\":1622390400000,\"secondPayRate\":0.0,\"pcDescUrl\":\"http://union.jd.com/activityRule?id=11180\",\"remark\":\"周年庆\"}]}"}}`
	r := &jdunionopenactivitybonusquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("activity bonus parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ActivityID != 1 || g.ActivityName != "京橙周年庆" || g.PayType != 1 ||
		g.FirstPayRate != 100 || g.SecondPayRate != 0 || g.PrepareTime != 1616603400000 ||
		g.PcDescURL != "http://union.jd.com/activityRule?id=11180" ||
		!r.Responce.QueryResult.HasMore {
		t.Fatalf("activity bonus fields mismatch: %+v", g)
	}
}

func TestParseOrderRowSupplyResponseShape(t *testing.T) {
	body := `{"jd_union_open_order_row_supply_query_responce":{"code":"0","queryResult":"{\"code\":200,\"message\":\"success\",\"hasMore\":false,\"data\":[{\"id\":\"415900297816660001\",\"orderId\":108618000005,\"parentId\":0,\"orderTime\":\"2020-01-02 15:50:16\",\"finishTime\":\"2020-01-03 15:59:16\",\"modifyTime\":\"2020-01-02 16:01:03\",\"unionId\":1000618618,\"skuId\":44303679033,\"skuName\":\"盲盒\",\"skuNum\":1,\"skuReturnNum\":0,\"skuFrozenNum\":0,\"price\":69.0,\"commissionRate\":5.0,\"finalRate\":90.0,\"estimateCosPrice\":54.0,\"estimateFee\":2.43,\"actualCosPrice\":0.0,\"actualFee\":0.0,\"validCode\":16,\"cid1\":1620,\"cid2\":11158,\"cid3\":11969,\"popId\":709982,\"payMonth\":0,\"sign\":\"B44C\",\"proPriceAmount\":6.18,\"goodsInfo\":{\"imageUrl\":\"http://img14.com/a.jpg\",\"owner\":\"g\",\"mainSkuId\":6161111,\"productId\":1236547,\"shopName\":\"XXXX旗舰店\",\"shopId\":45619},\"categoryInfo\":{\"cid1\":1,\"cid2\":2,\"cid3\":3,\"cid1Name\":\"珠宝首饰\",\"cid2Name\":\"木手串\",\"cid3Name\":\"其他\"},\"expressStatus\":10,\"outSideOrderId\":\"41590029781666000\",\"talentName\":\"小名\",\"applyPlatform\":1,\"talentId\":\"53xxx2o\"}]}"}}`
	r := &jdunionopenorderrowsupplyquery.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("order row supply parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.QueryResult.Data[0]
	if g.ID != "415900297816660001" || g.OrderID != 108618000005 || g.ValidCode != 16 ||
		g.PayMonth != 0 || g.EstimateFee != 2.43 || g.GoodsInfo.ShopID != 45619 ||
		g.CategoryInfo.Cid1Name != "珠宝首饰" || g.ApplyPlatform != 1 ||
		g.TalentID != "53xxx2o" || g.OutSideOrderID != "41590029781666000" ||
		r.Responce.QueryResult.HasMore {
		t.Fatalf("order row supply fields mismatch: %+v", g)
	}
}

func TestParseShPromotionGetResponseShape(t *testing.T) {
	body := `{"jd_union_open_sh_promotion_get_responce":{"getResult":"{\"code\":200,\"message\":\"success\",\"data\":{\"clickUrl\":\"https://u.jd.com/XXXXX\",\"impressionMonitorUrl\":\"https://mktm.jd.com/u/impress?taskId=1\",\"clickMonitorUrl\":\"https://mktm.jd.com/u/click?taskId=1\",\"appUrl\":\"openapp.jdmobile://virtual?params=1\"}}"}}`
	r := &jdunionopenshpromotionget.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("sh promotion get parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	g := r.Responce.GetResult.Data
	if g.ClickURL != "https://u.jd.com/XXXXX" ||
		g.ImpressionMonitorURL != "https://mktm.jd.com/u/impress?taskId=1" ||
		g.ClickMonitorURL != "https://mktm.jd.com/u/click?taskId=1" ||
		g.AppURL != "openapp.jdmobile://virtual?params=1" {
		t.Fatalf("sh promotion get fields mismatch: %+v", g)
	}
}

func TestParseUserPidGetResponseShape(t *testing.T) {
	body := `{"jd_union_open_user_pid_get_responce":{"getResult":"{\"code\":200,\"message\":\"success\",\"data\":\"1000618618_0_6186181618\"}"}}`
	r := &jdunionopenuserpidget.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("user pid get parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Message, r.Body)
	}
	if r.Responce.GetResult.Data != "1000618618_0_6186181618" {
		t.Fatalf("user pid get fields mismatch: %+v", r.Responce.GetResult)
	}
}
