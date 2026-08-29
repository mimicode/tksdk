package jdunionopengoodsrankquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.rank.query 联盟实时热销榜商品接口，支持榜单Id和排序类型查询榜单商品列表
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_rank_query_responce"`
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
	Code       int64           `json:"code"`
	Data       []RankGoodsResp `json:"data"`
	Message    string          `json:"message"`
	TotalCount int64           `json:"totalCount"`
	RequestID  string          `json:"requestId"`
}

// RankGoodsResp 榜单商品信息
type RankGoodsResp struct {
	ItemID            string                `json:"itemId"`            //联盟商品ID
	SkuID             int64                 `json:"skuId"`             //商品skuId
	SkuName           string                `json:"skuName"`           //商品名称
	ImageURL          string                `json:"imageUrl"`          //商品主图
	ImgList           []string              `json:"imgList"`           //商品图片列表
	Wlprice           float64               `json:"wlprice"`           //基准价
	CommissionShare   float64               `json:"commissionShare"`   //佣金比例
	Comments          int64                 `json:"comments"`          //评论总数
	GoodComments      int64                 `json:"goodComments"`      //好评数
	GoodCommentsShare float64               `json:"goodCommentsShare"` //好评率
	SkuTagList        []SkuTagInfo          `json:"skuTagList"`        //联盟标签
	PurchasePriceInfo RankPurchasePriceInfo `json:"purchasePriceInfo"` //到手价
}

// SkuTagInfo 联盟标签明细
type SkuTagInfo struct {
	Type  int64  `json:"type"`  //标签类型
	Index int64  `json:"index"` //优先级顺序（越小越优先）
	Name  string `json:"name"`  //标签名称
}

// RankPurchasePriceInfo 到手价信息
type RankPurchasePriceInfo struct {
	PurchasePrice          float64                  `json:"purchasePrice"`          //到手价
	PromotionLabelInfoList []RankPromotionLabelInfo `json:"promotionLabelInfoList"` //促销标签集
	CouponList             []RankCoupon             `json:"couponList"`             //优惠券集合
}

// RankPromotionLabelInfo 促销标签
type RankPromotionLabelInfo struct {
	PromotionLabelID string `json:"promotionLabelId"` //促销id
	LabelName        string `json:"labelName"`        //促销名称
}

// RankCoupon 优惠券明细
type RankCoupon struct {
	Link               string           `json:"link"`               //券链接
	Discount           float64          `json:"discount"`           //券面额
	Quota              float64          `json:"quota"`              //券消费限额
	RemainCnt          int64            `json:"remainCnt"`          //券剩余数量
	CouponStyle        int64            `json:"couponStyle"`        //优惠券分类 0：满减券，3：满折券，28：每满减券
	CouponStatus       int64            `json:"couponStatus"`       //领取状态 0：正常可领，-1：不可领取，1：已领取
	TimeCouponInfoList []TimeCouponInfo `json:"timeCouponInfoList"` //时段券信息集合
}

// TimeCouponInfo 时段券信息
type TimeCouponInfo struct {
	TimeCouponBegin string `json:"timeCouponBegin"` //时段券领取开始时间
	TimeCouponEnd   string `json:"timeCouponEnd"`   //时段券领取结束时间
}
