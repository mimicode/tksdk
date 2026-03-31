package request

// pangolin.product.category 商品类目接口
// https://www.csjplatform.com/supportcenter/28733
type ProductCategoryRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *ProductCategoryRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ProductCategoryRequest) GetApiName() string {
	return "/product/category"
}

// 方法名称
func (tk *ProductCategoryRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *ProductCategoryRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *ProductCategoryRequest) CheckParameters() {
	// parent_id 为可选参数
}

// SetParentId 设置上级类目id（查询一级类目填写0）
func (tk *ProductCategoryRequest) SetParentId(parentId int) {
	tk.AddParameter("parent_id", parentId)
}

// SetChannel 设置类目体系：0旧类目，1新类目
func (tk *ProductCategoryRequest) SetChannel(channel int) {
	tk.AddParameter("channel", channel)
}
