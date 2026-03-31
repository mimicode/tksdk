package request

// LiveSearchRequest 直播列表接口
// 批量获取直播间及带货商品信息
type LiveSearchRequest struct {
	Parameters map[string]interface{}
}

func (r *LiveSearchRequest) GetApiName() string {
	return "/live/search"
}

func (r *LiveSearchRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *LiveSearchRequest) CheckParameters() error {
	return nil
}

func (r *LiveSearchRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetPage 设置分页index，从1开始（必填）
func (r *LiveSearchRequest) SetPage(page int) {
	r.AddParameter("page", page)
}

// SetPageSize 设置分页size，范围(0,100]（必填）
func (r *LiveSearchRequest) SetPageSize(pageSize int) {
	r.AddParameter("page_size", pageSize)
}

// SetSortBy 设置排序字段：1综合，2销量，3佣金率，4粉丝数
func (r *LiveSearchRequest) SetSortBy(sortBy int) {
	r.AddParameter("sort_by", sortBy)
}

// SetSortType 设置排序方式：0降序，1升序
func (r *LiveSearchRequest) SetSortType(sortType int) {
	r.AddParameter("sort_type", sortType)
}

// SetStatus 设置直播间状态筛选：0在播和不在播，1在播，2不在播
func (r *LiveSearchRequest) SetStatus(status int) {
	r.AddParameter("status", status)
}

// SetAuthorInfo 设置达人昵称/达人uid
func (r *LiveSearchRequest) SetAuthorInfo(authorInfo string) {
	r.AddParameter("author_info", authorInfo)
}

// SetNeedProducts 设置是否需要返回商品：1返回，0不返回
func (r *LiveSearchRequest) SetNeedProducts(needProducts int) {
	r.AddParameter("need_products", needProducts)
}
