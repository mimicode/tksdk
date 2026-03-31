package request

// pangolin.product.link 商品转链接口
// https://www.csjplatform.com/supportcenter/28733
type ProductLinkRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *ProductLinkRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *ProductLinkRequest) GetApiName() string {
	return "/product/link"
}

// 方法名称
func (tk *ProductLinkRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *ProductLinkRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *ProductLinkRequest) CheckParameters() {
	// product_url 为必填参数
}

// SetProductUrl 设置商品url
func (tk *ProductLinkRequest) SetProductUrl(productUrl string) {
	tk.AddParameter("product_url", productUrl)
}

// SetProductExt 设置商品搜索接口返回的product.ext字段
func (tk *ProductLinkRequest) SetProductExt(productExt string) {
	tk.AddParameter("product_ext", productExt)
}

// SetExternalInfo 设置扩展参数（只能包含字母、数字、下划线，长度不超过40）
func (tk *ProductLinkRequest) SetExternalInfo(externalInfo string) {
	tk.AddParameter("external_info", externalInfo)
}

// SetShareType 设置转链类型：1抖音deep link，2抖音二维码，3抖音口令，4抖音zlink，5ShareLink
func (tk *ProductLinkRequest) SetShareType(shareTypes []int) {
	tk.AddParameter("share_type", shareTypes)
}

// SetPlatform 设置平台：0抖音，1抖音极速版
func (tk *ProductLinkRequest) SetPlatform(platform int) {
	tk.AddParameter("platform", platform)
}

// SetUseCoupon 设置是否返回商品惠后价领券链接
func (tk *ProductLinkRequest) SetUseCoupon(useCoupon bool) {
	tk.AddParameter("use_coupon", useCoupon)
}
