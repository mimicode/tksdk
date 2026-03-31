package request

// pangolin.product.search 商品列表接口
// https://www.csjplatform.com/supportcenter/28733
type ProductSearchRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *ProductSearchRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ProductSearchRequest) GetApiName() string {
	return "/product/search"
}

// 方法名称
func (tk *ProductSearchRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *ProductSearchRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *ProductSearchRequest) CheckParameters() {
	// page 和 page_size 为必填参数，在实际使用时通过 AddParameter 添加
}

// SetPage 设置页码
func (tk *ProductSearchRequest) SetPage(page int) {
	tk.AddParameter("page", page)
}

// SetPageSize 设置每页数量
func (tk *ProductSearchRequest) SetPageSize(pageSize int) {
	tk.AddParameter("page_size", pageSize)
}

// SetTitle 设置商品关键词
func (tk *ProductSearchRequest) SetTitle(title string) {
	tk.AddParameter("title", title)
}

// SetFirstCids 设置一级类目
func (tk *ProductSearchRequest) SetFirstCids(cids []int) {
	tk.AddParameter("first_cids", cids)
}

// SetSecondCids 设置二级类目
func (tk *ProductSearchRequest) SetSecondCids(cids []int) {
	tk.AddParameter("second_cids", cids)
}

// SetThirdCids 设置三级类目
func (tk *ProductSearchRequest) SetThirdCids(cids []int) {
	tk.AddParameter("third_cids", cids)
}

// SetPriceMin 设置最低价格（单位：分）
func (tk *ProductSearchRequest) SetPriceMin(priceMin int) {
	tk.AddParameter("price_min", priceMin)
}

// SetPriceMax 设置最高价格（单位：分）
func (tk *ProductSearchRequest) SetPriceMax(priceMax int) {
	tk.AddParameter("price_max", priceMax)
}

// SetSellNumMin 设置历史销量最小值
func (tk *ProductSearchRequest) SetSellNumMin(sellNumMin int) {
	tk.AddParameter("sell_num_min", sellNumMin)
}

// SetSellNumMax 设置历史销量最大值
func (tk *ProductSearchRequest) SetSellNumMax(sellNumMax int) {
	tk.AddParameter("sell_num_max", sellNumMax)
}

// SetSearchType 设置排序类型：0默认排序，1历史销量排序，2价格排序，3佣金排序，4佣金比例排序
func (tk *ProductSearchRequest) SetSearchType(searchType int) {
	tk.AddParameter("search_type", searchType)
}

// SetOrderType 设置排序方式：0升序，1降序
func (tk *ProductSearchRequest) SetOrderType(orderType int) {
	tk.AddParameter("order_type", orderType)
}

// SetCosFeeMin 设置最低分佣（单位：分）
func (tk *ProductSearchRequest) SetCosFeeMin(cosFeeMin int) {
	tk.AddParameter("cos_fee_min", cosFeeMin)
}

// SetCosFeeMax 设置最高分佣（单位：分）
func (tk *ProductSearchRequest) SetCosFeeMax(cosFeeMax int) {
	tk.AddParameter("cos_fee_max", cosFeeMax)
}

// SetCosRatioMin 设置最低佣金比例（百分比乘以100，如1.1%传110）
func (tk *ProductSearchRequest) SetCosRatioMin(cosRatioMin int) {
	tk.AddParameter("cos_ratio_min", cosRatioMin)
}

// SetCosRatioMax 设置最高佣金比例（百分比乘以100，如1.1%传110）
func (tk *ProductSearchRequest) SetCosRatioMax(cosRatioMax int) {
	tk.AddParameter("cos_ratio_max", cosRatioMax)
}

// SetActivityID 设置活动商品标识：1超值购，0全量
func (tk *ProductSearchRequest) SetActivityID(activityID int) {
	tk.AddParameter("activity_id", activityID)
}

// SetPackageActivityType 设置商品包类型：1爆品，2团长招商品，3两小时榜单，4今日榜单
func (tk *ProductSearchRequest) SetPackageActivityType(packageType int) {
	tk.AddParameter("package_activity_type", packageType)
}
