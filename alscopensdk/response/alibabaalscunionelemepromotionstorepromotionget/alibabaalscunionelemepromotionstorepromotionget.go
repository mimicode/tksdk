package alibabaalscunionelemepromotionstorepromotionget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.eleme.promotion.storepromotion.get( 本地联盟饿了么单店推广单店铺查询 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_eleme_promotion_storepromotion_get_response"`
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
	Data         Data   `json:"data"`           // 店铺数据
	Message      string `json:"message"`        // 返回消息
	ResultCode   int64  `json:"result_code"`    // 返回码
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

// Data 店铺数据
type Data struct {
	Title                  string            `json:"title"`                    // 门店名称
	ShopLogo               string            `json:"shop_logo"`                // 门店logo
	IndistinctMonthlySales string            `json:"indistinct_monthly_sales"` // 模糊销量
	CommissionRate         string            `json:"commission_rate"`          // 佣金比例
	BizType                string            `json:"biz_type"`                 // 店铺类型（"activityCps":活动cps,"ordinaryCps":基础cps）
	Activity               Activity          `json:"activity"`                 // 奖励金活动数据
	Link                   Link              `json:"link"`                     // 推广链接
	Category1ID            string            `json:"category_1_id"`            // 一级类目ID
	DeliveryPrice          string            `json:"delivery_price"`           // 起送价（元）
	RecommendReasons       RecommendReasons  `json:"recommend_reasons"`        // 推荐理由
	ServiceRating          string            `json:"service_rating"`           // 服务评级
	Items                  Items             `json:"items"`                    // 推荐商品
	Category1Name          string            `json:"category_1_name"`          // 一级类目名称
	OverlayCoupon          OverlayCoupon     `json:"overlay_coupon"`           // 叠加优惠券（保留原有字段）
}

// Activity 奖励金活动数据
type Activity struct {
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
	Picture                string `json:"picture"`                   // 图片链接
	ReduceWxAppid          string `json:"reduce_wx_appid"`           // 小程序appId-立减活动
	ReduceWxPath           string `json:"reduce_wx_path"`            // 微信小程序path链接-立减活动
	ReducePicture          string `json:"reduce_picture"`            // 立减活动图片链接
	MiniQrcode             string `json:"mini_qrcode"`               // 小程序二维码
	MediaActivityWxAppid   string `json:"media_activity_wx_appid"`   // 小程序appId-媒体出资活动
	MediaActivityWxPath    string `json:"media_activity_wx_path"`    // 微信小程序path链接-媒体出资活动
	TaobaoSchemeUrl        string `json:"taobao_scheme_url"`         // 淘宝scheme链接
	TaobaoQrcode           string `json:"taobao_qrcode"`             // 淘宝二维码
	H5Url                  string `json:"h5_url"`                    // H5链接
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

// OverlayCoupon 叠加优惠券（保留原有字段）
type OverlayCoupon struct {
	Amount      string  `json:"amount"`       // 优惠券金额
	EndTime     string  `json:"end_time"`     // 结束时间
	StartTime   string  `json:"start_time"`   // 开始时间
	TemplateID  float64 `json:"template_id"`  // 模板ID
	ValidPeriod int64   `json:"valid_period"` // 有效期
}
