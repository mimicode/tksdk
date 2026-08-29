package jdopensdk

import (
	"testing"

	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsbigfieldquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsrankquery"
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
