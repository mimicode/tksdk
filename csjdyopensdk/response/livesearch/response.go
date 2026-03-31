package livesearch

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 直播列表响应
type Response struct {
	response.TopResponse
	Data ResponseData `json:"data"`
}

func (r *Response) WrapResult(result string) {
	r.Body = result
	err := json.Unmarshal([]byte(result), r)
	if err != nil {
		r.Code = -1
		r.Desc = err.Error()
	}
}

// ResponseData 响应数据
type ResponseData struct {
	Total     int64      `json:"total"`      // 符合条件的最大直播间数量
	LiveInfos []LiveInfo `json:"live_infos"` // 直播间信息列表
}

// LiveInfo 直播间信息
type LiveInfo struct {
	AuthorOpenid          string         `json:"author_openid"`           // 主播openid
	AuthorBuyinId         string         `json:"author_buyin_id"`          // 主播百应id
	AuthorName            string         `json:"author_name"`             // 昵称
	AuthorPic             string         `json:"author_pic"`              // 头像
	AuthorLevel           int            `json:"author_level"`             // 达人等级
	ProductCategory       []string       `json:"product_category"`        // 商品类别
	AverageGmv            string         `json:"average_gmv"`             // 场均GMV
	FansNum               int            `json:"fans_num"`                // 粉丝数量
	AverageCommissionRate string         `json:"average_commission_rate"` // 平均佣金率
	RoomId                string         `json:"room_id"`                 // 直播间id
	Gender                string         `json:"gender"`                  // 性别
	Ext                   string         `json:"ext"`                     // 内部加密字段（转链时需要）
	Products              []LiveProduct  `json:"products"`                // 直播间挂车商品
	LotteryProducts       []interface{}  `json:"lottery_products"`       // 福袋商品活动信息
	OnlineNum             int            `json:"online_num"`              // 查询时刻在线人数
	CreateTime            int64          `json:"create_time"`             // 开播时间戳
}

// LiveProduct 直播间商品
type LiveProduct struct {
	ProductId       int64              `json:"product_id"`        // 商品id
	Title           string             `json:"title"`            // 商品名称
	Price           int64              `json:"price"`            // 商品价格（单位：分）
	InStock         bool               `json:"in_stock"`         // 是否上架
	FirstCid        int64              `json:"first_cid"`        // 一级类目
	FirstCname      string             `json:"first_cname"`      // 一级类目名称
	SecondCid       int64              `json:"second_cid"`       // 二级类目
	SecondCname     string             `json:"second_cname"`     // 二级类目名称
	ThirdCid        int64              `json:"third_cid"`        // 三级类目
	ThirdCname      string             `json:"third_cname"`      // 三级类目名称
	Sales           int                `json:"sales"`            // 历史销量
	Cover           string             `json:"cover"`           // 商品主图
	DetailUrl       string             `json:"detail_url"`      // 商品链接
	CouponPrice     int64              `json:"coupon_price"`    // 直播间最终价格（单位：分）
	ShopId          int64              `json:"shop_id"`         // 商铺id
	ShopName        string             `json:"shop_name"`       // 商铺名称
	CosRatio        int                `json:"cos_ratio"`        // 新客分佣比例
	CosFee          int                `json:"cos_fee"`         // 新客佣金金额（单位：分）
	OldFansCosRatio int                `json:"old_fans_cos_ratio"` // 老客分佣比例
	OldFansCosFee   int                `json:"old_fans_cos_fee"`   // 老客佣金金额（单位：分）
	IsExplaining    bool               `json:"is_explaining"`    // 是否正在讲解
	HasGivenProduct bool               `json:"has_given_product"` // 是否有赠品
	AvailableCoupons []LiveCoupon      `json:"available_coupons"`  // 优惠券信息列表
	LotteryActivityInfo *LotteryInfo   `json:"lottery_activity_info"` // 福袋活动信息
}

// LiveCoupon 优惠券信息
type LiveCoupon struct {
	CouponType   int    `json:"coupon_type"`   // 优惠券类型：1平台券，2店铺券，3达人券
	TypeDesc     string `json:"type_desc"`     // 券类型描述
	DiscountDesc string `json:"discount_desc"` // 优惠内容
}

// LotteryInfo 福袋活动信息
type LotteryInfo struct {
	LotteryActivityId       int64 `json:"lottery_activity_id"`        // 福袋活动id
	LotteryProductStartTime int64 `json:"lottery_product_start_time"` // 福袋活动起始时间戳
	LotteryProductEndTime   int64 `json:"lottery_product_end_time"`   // 福袋活动结束/开奖时间戳
}
