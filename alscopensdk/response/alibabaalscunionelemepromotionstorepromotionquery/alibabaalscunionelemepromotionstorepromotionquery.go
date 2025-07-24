package alibabaalscunionelemepromotionstorepromotionquery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.eleme.promotion.storepromotion.query( 本地联盟饿了么单店推广店铺列表 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_eleme_promotion_storepromotion_query_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"` // 业务错误描述
	BizErrorCode int    `json:"biz_error_code"` // 业务错误码
	RequestID    string `json:"request_id"`     // 请求ID
	Data         Data   `json:"data"`           // 分页数据
	Message      string `json:"message"`        // 返回消息
	ResultCode   int64  `json:"result_code"`    // 返回码
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

// Data 分页数据
type Data struct {
	PageSize  int64   `json:"page_size"`  // 每页数目
	Records   Records `json:"records"`    // 分页记录
	SessionID string  `json:"session_id"` // 会话ID（下次请求作为请求参数，用于标记分页会话自动翻页）
}

// Records 分页记录
type Records struct {
	StorePromotionDto []StorePromotionDto `json:"store_promotion_dto"` // 店铺推广数据列表
}

// StorePromotionDto 店铺推广数据
type StorePromotionDto struct {
	Title                  string            `json:"title"`                    // 门店名称
	ShopLogo               string            `json:"shop_logo"`                // 门店logo
	IndistinctMonthlySales string            `json:"indistinct_monthly_sales"` // 模糊销量
	CommissionRate         string            `json:"commission_rate"`          // 佣金比例
	BizType                string            `json:"biz_type"`                 // 店铺类型（"activityCps":活动cps,"ordinaryCps":基础cps）
	Activity               PromotionActivity `json:"activity"`                 // 奖励金活动数据
	Link                   Link              `json:"link"`                     // 推广链接
	Category1ID            string            `json:"category_1_id"`            // 一级类目ID，高级字段
	DeliveryDistance       int64             `json:"delivery_distance"`        // 配送距离（米），高级字段
	DeliveryTime           int64             `json:"delivery_time"`            // 配送时间（分），高级字段
	DeliveryPrice          string            `json:"delivery_price"`           // 起送价（元），高级字段
	RecommendReasons       RecommendReasons  `json:"recommend_reasons"`        // 推荐理由，高级字段
	ServiceRating          string            `json:"service_rating"`           // 服务评级，高级字段
	Items                  Items             `json:"items"`                    // 推荐商品，高级字段
	ShopID                 string            `json:"shop_id"`                  // 店铺ID（静态加密）
	Category1Name          string            `json:"category_1_name"`          // 一级类目名称，高级字段
	Commission             int64             `json:"commission"`               // 预估佣金（分）
}

// PromotionActivity 奖励金活动数据
type PromotionActivity struct {
	ActivityID              string `json:"activity_id"`               // 活动Id
	ServiceFeeCent          int64  `json:"service_fee_cent"`          // 营销计划服务费（分）
	BonusCent               int64  `json:"bonus_cent"`                // 奖励金红包面额（分）
	DailyQuantity           int64  `json:"daily_quantity"`            // 活动的日库存
	DailySellableQuantity   int64  `json:"daily_sellable_quantity"`   // 活动日剩余库存
	StartTime               int64  `json:"start_time"`                // 起始时间（秒）
	EndTime                 int64  `json:"end_time"`                  // 结束时间（秒）
	BountyMinLimitCent      int64  `json:"bounty_min_limit_cent"`     // 奖励金门槛 (分)
}

// Link 推广链接
type Link struct {
	WxAppid                string `json:"wx_appid"`                  // 小程序appId
	WxPath                 string `json:"wx_path"`                   // 微信小程序path链接
	ReduceWxAppid          string `json:"reduce_wx_appid"`           // 小程序appId-立减活动
	ReduceWxPath           string `json:"reduce_wx_path"`            // 微信小程序path链接-立减活动
	MediaActivityWxAppid   string `json:"media_activity_wx_appid"`   // 小程序appId-媒体出资活动
	MediaActivityWxPath    string `json:"media_activity_wx_path"`    // 微信小程序path链接-媒体出资活动
}

// RecommendReasons 推荐理由
type RecommendReasons struct {
	String []string `json:"string"` // 推荐理由字符串数组
}

// Items 推荐商品
type Items struct {
	PromotionItem []PromotionItem `json:"promotion_item"` // 推广商品列表
}

// PromotionItem 推广商品
type PromotionItem struct {
	Title       string `json:"title"`        // 标题
	OriginPrice string `json:"origin_price"` // 原价
	Price       string `json:"price"`        // 现价
	Picture     string `json:"picture"`      // 图片
}
