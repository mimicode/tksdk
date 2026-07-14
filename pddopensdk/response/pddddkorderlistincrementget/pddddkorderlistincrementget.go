package pddddkorderlistincrementget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

// pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）
type Response struct {
	response2.TopResponse
	OrderListGetResponse OrderListGetResponse `json:"order_list_get_response"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type OrderListGetResponse struct {
	TotalCount int64       `json:"total_count"`
	OrderList  []OrderList `json:"order_list"`
	RequestID  string      `json:"request_id"`
}

type OrderList struct {
	GoodsPrice            int64   `json:"goods_price"`              //单团价格，单位：分
	PromotionRate         int64   `json:"promotion_rate"`           //佣金比例，千分比
	Type                  int64   `json:"type"`                     //订单类型
	OrderStatus           int64   `json:"order_status"`             //订单状态：0-已创建，1-已成团，2-确认收货，3-审核完成，4-已结算，5-已取消
	CatIDS                []int64 `json:"cat_ids"`                  //商品类目ID列表
	OrderCreateTime       int64   `json:"order_create_time"`        //订单创建时间，UNIX秒级时间戳
	IsDirect              int64   `json:"is_direct"`                //是否直推：1-是，0-否
	OrderGroupSuccessTime int64   `json:"order_group_success_time"` //成团时间，UNIX秒级时间戳
	OrderAmount           int64   `json:"order_amount"`             //订单金额，单位：分
	OrderModifyAt         int64   `json:"order_modify_at"`          //订单最后修改时间，UNIX秒级时间戳
	AuthDuoID             int64   `json:"auth_duo_id"`              //授权多多客ID
	CPANew                int64   `json:"cpa_new"`                  //CPA奖励金额
	GoodsName             string  `json:"goods_name"`               //商品标题
	BatchNo               string  `json:"batch_no"`                 //批次号
	GoodsQuantity         int64   `json:"goods_quantity"`           //商品购买数量
	GoodsID               int64   `json:"goods_id"`                 //商品ID
	CustomParameters      string  `json:"custom_parameters"`        //自定义参数，用于区分推广渠道
	GoodsThumbnailURL     string  `json:"goods_thumbnail_url"`      //商品缩略图URL
	PromotionAmount       int64   `json:"promotion_amount"`         //佣金金额，单位：分
	OrderPayTime          int64   `json:"order_pay_time"`           //订单支付时间，UNIX秒级时间戳
	GroupID               float64 `json:"group_id"`                 //拼团ID
	OrderStatusDesc       string  `json:"order_status_desc"`        //订单状态描述
	OrderID               string  `json:"order_id"`                 //订单号
	OrderSn               string  `json:"order_sn"`                 //子订单号
	PID                   string  `json:"p_id"`                     //推广位ID
	ZsDuoID               int64   `json:"zs_duo_id"`                //招商多多客ID
	FailReason            string  `json:"fail_reason"`              //订单失效原因
	GoodsSign             string  `json:"goods_sign"`               //商品签名（GoodsID + "_" + 拼团ID）
	MallName              string  `json:"mall_name"`                //店铺名称
	OrderReceiveTime      int64   `json:"order_receive_time"`       //确认收货时间，UNIX秒级时间戳
	OrderSettleTime       int64   `json:"order_settle_time"`        //订单结算时间，UNIX秒级时间戳
	OrderVerifyTime       int64   `json:"order_verify_time"`        //订单审核（结算）时间，UNIX秒级时间戳
	PlatformDiscount      int64   `json:"platform_discount"`        //平台补贴金额，单位：分
}
