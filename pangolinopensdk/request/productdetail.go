package request

// pangolin.product.detail 商品详情接口
// https://www.csjplatform.com/supportcenter/28733
type ProductDetailRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *ProductDetailRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ProductDetailRequest) GetApiName() string {
	return "/product/detail"
}

// 方法名称
func (tk *ProductDetailRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *ProductDetailRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *ProductDetailRequest) CheckParameters() {
	// product_ids 为必填参数
}

// SetProductIds 设置商品id列表
func (tk *ProductDetailRequest) SetProductIds(productIds []int64) {
	tk.AddParameter("product_ids", productIds)
}
