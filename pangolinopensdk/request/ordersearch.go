package request

// pangolin.order.search 订单查询接口
// https://www.csjplatform.com/supportcenter/28733
type OrderSearchRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *OrderSearchRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *OrderSearchRequest) GetApiName() string {
	return "/order/search"
}

// 方法名称
func (tk *OrderSearchRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *OrderSearchRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *OrderSearchRequest) CheckParameters() {
	// order_type 为必填参数
}

// 设置分页查询参数（与 order_ids 二选一）
// SetPageQuery 设置分页滚动拉取参数
func (tk *OrderSearchRequest) SetPageQuery(size int, cursor string, startTime int, endTime int) {
	tk.AddParameter("size", size)
	tk.AddParameter("cursor", cursor)
	tk.AddParameter("start_time", startTime)
	tk.AddParameter("end_time", endTime)
}

// 设置订单ID查询参数（与分页查询二选一）
// SetOrderIds 设置订单号列表（最多10个）
func (tk *OrderSearchRequest) SetOrderIds(orderIds []string) {
	tk.AddParameter("order_ids", orderIds)
}

// SetOrderType 设置订单类型：1商品分销订单，2直播间分销订单，3活动类型订单
func (tk *OrderSearchRequest) SetOrderType(orderType int) {
	tk.AddParameter("order_type", orderType)
}

// SetTimeType 设置时间类型：pay按支付时间，update按订单更新时间
func (tk *OrderSearchRequest) SetTimeType(timeType string) {
	tk.AddParameter("time_type", timeType)
}

// SetAccessType 设置接入类型：1API接入，2SDK接入
func (tk *OrderSearchRequest) SetAccessType(accessType int) {
	tk.AddParameter("access_type", accessType)
}
