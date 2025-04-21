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

func (r *ProductSearchRequest) CheckParameters() {

}

func (r *ProductSearchRequest) AddParameter(key, value string) {
	r.Parameters[key] = value
}
