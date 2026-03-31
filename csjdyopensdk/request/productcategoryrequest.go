package request

// ProductCategoryRequest 商品类目接口
// 获取商品类目信息
type ProductCategoryRequest struct {
	Parameters map[string]interface{}
}

func (r *ProductCategoryRequest) GetApiName() string {
	return "/product/category"
}

func (r *ProductCategoryRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *ProductCategoryRequest) CheckParameters() error {
	return nil
}

func (r *ProductCategoryRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetParentId 设置上级类目id（查询一级类目填写0）
func (r *ProductCategoryRequest) SetParentId(parentId int) {
	r.AddParameter("parent_id", parentId)
}

// SetChannel 设置类目体系：0旧类目，1新类目
func (r *ProductCategoryRequest) SetChannel(channel int) {
	r.AddParameter("channel", channel)
}
