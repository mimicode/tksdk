package pddddkorderlistincrementget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/pddopensdk/response"
)

//pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）
type Response struct {
	response2.TopResponse
	OrderListGetResponse OrderListGetResponse `json:"order_list_get_response"`
}

//解析输出结果
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
	GoodsPrice            int64   `json:"goods_price"`
	PromotionRate         int64   `json:"promotion_rate"`
	Type                  int64   `json:"type"`
	OrderStatus           int64   `json:"order_status"`
	CatIDS                []int64 `json:"cat_ids"`
	OrderCreateTime       int64   `json:"order_create_time"`
	IsDirect              int64   `json:"is_direct"`
	OrderGroupSuccessTime int64   `json:"order_group_success_time"`
	OrderAmount           int64   `json:"order_amount"`
	OrderModifyAt         int64   `json:"order_modify_at"`
	AuthDuoID             int64   `json:"auth_duo_id"`
	CPANew                int64   `json:"cpa_new"`
	GoodsName             string  `json:"goods_name"`
	BatchNo               string  `json:"batch_no"`
	GoodsQuantity         int64   `json:"goods_quantity"`
	GoodsID               int64   `json:"goods_id"`
	CustomParameters      string  `json:"custom_parameters"`
	GoodsThumbnailURL     string  `json:"goods_thumbnail_url"`
	PromotionAmount       int64   `json:"promotion_amount"`
	OrderPayTime          int64   `json:"order_pay_time"`
	GroupID               float64 `json:"group_id"`
	OrderStatusDesc       string  `json:"order_status_desc"`
	OrderID               string  `json:"order_id"`
	OrderSn               string  `json:"order_sn"`
	PID                   string  `json:"p_id"`
	ZsDuoID               int64   `json:"zs_duo_id"`
}
