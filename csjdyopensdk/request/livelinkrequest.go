package request

// LiveLinkRequest 直播转链接口
// 获取某个直播间或者直播间内某个商品的跳转抖音渠道
type LiveLinkRequest struct {
	Parameters map[string]interface{}
}

func (r *LiveLinkRequest) GetApiName() string {
	return "/live/link"
}

func (r *LiveLinkRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *LiveLinkRequest) CheckParameters() error {
	return nil
}

func (r *LiveLinkRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetAuthorOpenid 设置直播间主播openid（与author_buyin_id二选一）
func (r *LiveLinkRequest) SetAuthorOpenid(authorOpenid string) {
	r.AddParameter("author_openid", authorOpenid)
}

// SetAuthorBuyinId 设置直播间主播百应id（与author_openid二选一）
func (r *LiveLinkRequest) SetAuthorBuyinId(authorBuyinId string) {
	r.AddParameter("author_buyin_id", authorBuyinId)
}

// SetExternalInfo 设置自定义字段（只能包含数字、字母和_，长度不超过40）
func (r *LiveLinkRequest) SetExternalInfo(externalInfo string) {
	r.AddParameter("external_info", externalInfo)
}

// SetLiveExt 设置直播间列表接口下发的live_info.ext
func (r *LiveLinkRequest) SetLiveExt(liveExt string) {
	r.AddParameter("live_ext", liveExt)
}

// SetShareType 设置转链类型：1抖音deep link，2抖音二维码，3抖音口令，4抖音zlink
func (r *LiveLinkRequest) SetShareType(shareTypes []int) {
	r.AddParameter("share_type", shareTypes)
}

// SetProductId 设置直播间内的商品id
func (r *LiveLinkRequest) SetProductId(productId int64) {
	r.AddParameter("product_id", productId)
}

// SetPlatform 设置平台：0抖音，1抖音极速版
func (r *LiveLinkRequest) SetPlatform(platform int) {
	r.AddParameter("platform", platform)
}
