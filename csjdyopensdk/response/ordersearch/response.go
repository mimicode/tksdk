package ordersearch

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 订单查询响应
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
	Cursor string  `json:"cursor"` // 下一页订单查询游标
	Orders []Order `json:"orders"` // 订单列表
}

// Order 订单信息
type Order struct {
	OrderId            string        `json:"order_id"`            // 订单id
	AppId              string        `json:"app_id"`              // 应用id
	ProductId          string        `json:"product_id"`          // 商品id
	ProductName        string        `json:"product_name"`        // 商品名称
	AuthorAccount      string        `json:"author_account"`      // 直播间达人昵称（仅直播间分销订单有）
	AdsAttribution     string        `json:"ads_attribution"`     // 结算方式：intersect跨播，directIn直推
	ProductImg         string        `json:"product_img"`         // 商品图片url
	TotalPayAmount     int64         `json:"total_pay_amount"`    // 总付款金额（单位：分）
	PaySuccessTime     string        `json:"pay_success_time"`    // 支付成功时间
	RefundTime         string        `json:"refund_time"`         // 退款时间
	PayGoodsAmount     int64         `json:"pay_goods_amount"`    // 预估参与结算金额（单位：分）
	RefundAmount       int64         `json:"refund_amount"`       // 退款金额（单位：分）
	EstimatedCommission int64        `json:"estimated_commission"` // 预估佣金（单位：分）
	CommissionType     int           `json:"commission_type"`     // 佣金类型：1商品分销，2直播间分销，3活动类型
	OrderStatus        int           `json:"order_status"`        // 订单状态：0已支付，1已结算，2已退款，3在处理
	RoleId             string        `json:"role_id"`            // 子账号id
	UserId             string        `json:"user_id"`            // 主账号id
	ProductUrl         string        `json:"product_url"`        // 商品链接
	ExternalInfo       string        `json:"external_info"`      // 媒体传递的扩展参数
	ColonelsInfo       []ColonelInfo `json:"colonel_order_info"`  // 团长信息
}

// ColonelInfo 团长信息
type ColonelInfo struct {
	ActivityId     int64  `json:"activity_id"`      // 团长活动id
	InstitutionId  int64  `json:"institution_id"`   // 团长机构id
	InstitutionName string `json:"institution_name"` // 团长机构名称
}
