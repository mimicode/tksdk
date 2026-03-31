package request

type ProductSearchRequest struct {
	Parameters map[string]interface{}
}

func (r *ProductSearchRequest) GetApiName() string {
	return "/product/search"
}

func (r *ProductSearchRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *ProductSearchRequest) CheckParameters() error {
	return nil
}

func (r *ProductSearchRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetPage 设置页码（必填）
func (r *ProductSearchRequest) SetPage(page int) {
	r.AddParameter("page", page)
}

// SetPageSize 设置每页数量（必填，最大20）
func (r *ProductSearchRequest) SetPageSize(pageSize int) {
	r.AddParameter("page_size", pageSize)
}

// SetTitle 设置商品关键词
func (r *ProductSearchRequest) SetTitle(title string) {
	r.AddParameter("title", title)
}

// SetFirstCids 设置一级类目
func (r *ProductSearchRequest) SetFirstCids(cids []int) {
	r.AddParameter("first_cids", cids)
}

// SetSecondCids 设置二级类目
func (r *ProductSearchRequest) SetSecondCids(cids []int) {
	r.AddParameter("second_cids", cids)
}

// SetThirdCids 设置三级类目
func (r *ProductSearchRequest) SetThirdCids(cids []int) {
	r.AddParameter("third_cids", cids)
}

// SetPriceMin 设置最低价格（单位：分）
func (r *ProductSearchRequest) SetPriceMin(priceMin int) {
	r.AddParameter("price_min", priceMin)
}

// SetPriceMax 设置最高价格（单位：分）
func (r *ProductSearchRequest) SetPriceMax(priceMax int) {
	r.AddParameter("price_max", priceMax)
}

// SetSellNumMin 设置历史销量最小值
func (r *ProductSearchRequest) SetSellNumMin(sellNumMin int) {
	r.AddParameter("sell_num_min", sellNumMin)
}

// SetSellNumMax 设置历史销量最大值
func (r *ProductSearchRequest) SetSellNumMax(sellNumMax int) {
	r.AddParameter("sell_num_max", sellNumMax)
}

// SetSearchType 设置排序类型：0默认排序，1历史销量排序，2价格排序，3佣金排序，4佣金比例排序
func (r *ProductSearchRequest) SetSearchType(searchType int) {
	r.AddParameter("search_type", searchType)
}

// SetOrderType 设置排序方式：0升序，1降序
func (r *ProductSearchRequest) SetOrderType(orderType int) {
	r.AddParameter("order_type", orderType)
}

// SetCosFeeMin 设置最低分佣（单位：分）
func (r *ProductSearchRequest) SetCosFeeMin(cosFeeMin int) {
	r.AddParameter("cos_fee_min", cosFeeMin)
}

// SetCosFeeMax 设置最高分佣（单位：分）
func (r *ProductSearchRequest) SetCosFeeMax(cosFeeMax int) {
	r.AddParameter("cos_fee_max", cosFeeMax)
}

// SetCosRatioMin 设置最低佣金比例（百分比乘以100，如1.1%传110）
func (r *ProductSearchRequest) SetCosRatioMin(cosRatioMin int) {
	r.AddParameter("cos_ratio_min", cosRatioMin)
}

// SetCosRatioMax 设置最高佣金比例（百分比乘以100，如1.1%传110）
func (r *ProductSearchRequest) SetCosRatioMax(cosRatioMax int) {
	r.AddParameter("cos_ratio_max", cosRatioMax)
}

// SetActivityId 设置活动商品标识：1超值购，0全量
func (r *ProductSearchRequest) SetActivityId(activityId int) {
	r.AddParameter("activity_id", activityId)
}

// SetPackageActivityType 设置商品包类型：1爆品，2团长招商品，3两小时榜单，4今日榜单
func (r *ProductSearchRequest) SetPackageActivityType(packageType int) {
	r.AddParameter("package_activity_type", packageType)
}
