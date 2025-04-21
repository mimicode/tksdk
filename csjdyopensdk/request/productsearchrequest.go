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
