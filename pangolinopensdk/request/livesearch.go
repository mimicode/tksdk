package request

// pangolin.live.search 直播列表接口
// https://www.csjplatform.com/supportcenter/28733
type LiveSearchRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *LiveSearchRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *LiveSearchRequest) GetApiName() string {
	return "/live/search"
}

// 方法名称
func (tk *LiveSearchRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *LiveSearchRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *LiveSearchRequest) CheckParameters() {
	// page 和 page_size 为必填参数
}

// SetPage 设置分页index，从1开始
func (tk *LiveSearchRequest) SetPage(page int) {
	tk.AddParameter("page", page)
}

// SetPageSize 设置分页size，范围(0,100]
func (tk *LiveSearchRequest) SetPageSize(pageSize int) {
	tk.AddParameter("page_size", pageSize)
}

// SetSortBy 设置排序字段：1综合，2销量，3佣金率，4粉丝数
func (tk *LiveSearchRequest) SetSortBy(sortBy int) {
	tk.AddParameter("sort_by", sortBy)
}

// SetSortType 设置排序方式：0降序，1升序
func (tk *LiveSearchRequest) SetSortType(sortType int) {
	tk.AddParameter("sort_type", sortType)
}

// SetStatus 设置直播间状态筛选：0在播和不在播，1在播，2不在播
func (tk *LiveSearchRequest) SetStatus(status int) {
	tk.AddParameter("status", status)
}

// SetAuthorInfo 设置达人昵称/达人uid
func (tk *LiveSearchRequest) SetAuthorInfo(authorInfo string) {
	tk.AddParameter("author_info", authorInfo)
}

// SetNeedProducts 设置是否需要返回商品：1返回，0不返回
func (tk *LiveSearchRequest) SetNeedProducts(needProducts int) {
	tk.AddParameter("need_products", needProducts)
}
