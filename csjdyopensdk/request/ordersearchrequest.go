package request

// OrderSearchRequest 订单查询接口
// 查询某个应用下通过商品转链和直播间转链产生的订单
type OrderSearchRequest struct {
	Parameters map[string]interface{}
}

func (r *OrderSearchRequest) GetApiName() string {
	return "/order/search"
}

func (r *OrderSearchRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *OrderSearchRequest) CheckParameters() error {
	return nil
}

func (r *OrderSearchRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// 设置分页滚动拉取参数（与order_ids二选一）
// SetPageQuery 设置分页查询参数
func (r *OrderSearchRequest) SetPageQuery(size int, cursor string, startTime int, endTime int) {
	r.AddParameter("size", size)
	r.AddParameter("cursor", cursor)
	r.AddParameter("start_time", startTime)
	r.AddParameter("end_time", endTime)
}

// 设置订单ID查询参数（与分页查询二选一）
// SetOrderIds 设置订单号列表（最多10个）
func (r *OrderSearchRequest) SetOrderIds(orderIds []string) {
	r.AddParameter("order_ids", orderIds)
}

// SetOrderType 设置订单类型：1商品分销订单，2直播间分销订单，3活动类型订单（必填）
func (r *OrderSearchRequest) SetOrderType(orderType int) {
	r.AddParameter("order_type", orderType)
}

// SetTimeType 设置时间类型：pay按支付时间，update按订单更新时间
func (r *OrderSearchRequest) SetTimeType(timeType string) {
	r.AddParameter("time_type", timeType)
}

// SetAccessType 设置接入类型：1API接入，2SDK接入
func (r *OrderSearchRequest) SetAccessType(accessType int) {
	r.AddParameter("access_type", accessType)
}
