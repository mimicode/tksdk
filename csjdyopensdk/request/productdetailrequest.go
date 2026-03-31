package request

// ProductDetailRequest 商品详情接口
// 获取某些商品的详细信息
type ProductDetailRequest struct {
	Parameters map[string]interface{}
}

func (r *ProductDetailRequest) GetApiName() string {
	return "/product/detail"
}

func (r *ProductDetailRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *ProductDetailRequest) CheckParameters() error {
	return nil
}

func (r *ProductDetailRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetProductIds 设置商品id列表（必填）
func (r *ProductDetailRequest) SetProductIds(productIds []int64) {
	r.AddParameter("product_ids", productIds)
}
