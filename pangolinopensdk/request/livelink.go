package request

// pangolin.live.link 直播转链接口
// https://www.csjplatform.com/supportcenter/28733
type LiveLinkRequest struct {
	Parameters map[string]interface{}
}

// 添加请求参数
func (tk *LiveLinkRequest) AddParameter(key string, val interface{}) {
	if tk.Parameters == nil {
		tk.Parameters = map[string]interface{}{}
	}
	tk.Parameters[key] = val
}

// 返回接口名称
func (tk *LiveLinkRequest) GetApiName() string {
	return "/live/link"
}

// 方法名称
func (tk *LiveLinkRequest) GetVersion() string {
	return "v1"
}

// 返回请求参数
func (tk *LiveLinkRequest) GetParameters() map[string]interface{} {
	return tk.Parameters
}

// CheckParameters 检测参数
func (tk *LiveLinkRequest) CheckParameters() {
	// author_openid 或 author_buyin_id 二选一必填
}

// SetAuthorOpenid 设置直播间主播openid
func (tk *LiveLinkRequest) SetAuthorOpenid(authorOpenid string) {
	tk.AddParameter("author_openid", authorOpenid)
}

// SetAuthorBuyinId 设置直播间主播百应id
func (tk *LiveLinkRequest) SetAuthorBuyinId(authorBuyinId string) {
	tk.AddParameter("author_buyin_id", authorBuyinId)
}

// SetExternalInfo 设置自定义字段（只能包含数字、字母和_，长度不超过40）
func (tk *LiveLinkRequest) SetExternalInfo(externalInfo string) {
	tk.AddParameter("external_info", externalInfo)
}

// SetLiveExt 设置直播间列表接口下发的live_info.ext
func (tk *LiveLinkRequest) SetLiveExt(liveExt string) {
	tk.AddParameter("live_ext", liveExt)
}

// SetShareType 设置转链类型：1抖音deep link，2抖音二维码，3抖音口令，4抖音zlink
func (tk *LiveLinkRequest) SetShareType(shareTypes []int) {
	tk.AddParameter("share_type", shareTypes)
}

// SetProductId 设置直播间内的商品id
func (tk *LiveLinkRequest) SetProductId(productId int64) {
	tk.AddParameter("product_id", productId)
}

// SetPlatform 设置平台：0抖音，1抖音极速版
func (tk *LiveLinkRequest) SetPlatform(platform int) {
	tk.AddParameter("platform", platform)
}
