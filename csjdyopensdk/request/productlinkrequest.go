package request

// ProductLinkRequest 商品转链接口
// 获取某个商品的跳转抖音渠道(抖音二维码，抖音口令，抖音deeplink)
type ProductLinkRequest struct {
	Parameters map[string]interface{}
}

func (r *ProductLinkRequest) GetApiName() string {
	return "/product/link"
}

func (r *ProductLinkRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *ProductLinkRequest) CheckParameters() error {
	return nil
}

func (r *ProductLinkRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetProductUrl 设置商品url（必填）
func (r *ProductLinkRequest) SetProductUrl(productUrl string) {
	r.AddParameter("product_url", productUrl)
}

// SetProductExt 设置商品搜索接口返回的product.ext字段
func (r *ProductLinkRequest) SetProductExt(productExt string) {
	r.AddParameter("product_ext", productExt)
}

// SetExternalInfo 设置扩展参数（只能包含字母、数字、下划线，长度不超过40）
func (r *ProductLinkRequest) SetExternalInfo(externalInfo string) {
	r.AddParameter("external_info", externalInfo)
}

// SetShareType 设置转链类型：1抖音deep link，2抖音二维码，3抖音口令，4抖音zlink，5ShareLink
func (r *ProductLinkRequest) SetShareType(shareTypes []int) {
	r.AddParameter("share_type", shareTypes)
}

// SetPlatform 设置平台：0抖音，1抖音极速版
func (r *ProductLinkRequest) SetPlatform(platform int) {
	r.AddParameter("platform", platform)
}

// SetUseCoupon 设置是否返回商品惠后价领券链接
func (r *ProductLinkRequest) SetUseCoupon(useCoupon bool) {
	r.AddParameter("use_coupon", useCoupon)
}
