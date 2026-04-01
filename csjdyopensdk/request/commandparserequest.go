package request

// CommandParseRequest 抖口令解析接口
// 解析抖口令（商品/直播间/直播间商品），返回对应的转链信息
type CommandParseRequest struct {
	Parameters map[string]interface{}
}

func (r *CommandParseRequest) GetApiName() string {
	return "/command/parse"
}

func (r *CommandParseRequest) GetParameters() map[string]interface{} {
	return r.Parameters
}

func (r *CommandParseRequest) CheckParameters() error {
	return nil
}

func (r *CommandParseRequest) AddParameter(key string, value interface{}) {
	if r.Parameters == nil {
		r.Parameters = map[string]interface{}{}
	}
	r.Parameters[key] = value
}

// SetCommand 设置抖口令（商品/直播间/直播间商品）
func (r *CommandParseRequest) SetCommand(command string) {
	r.AddParameter("command", command)
}

// SetExternalInfo 设置扩展参数（只能包含字母、数字、下划线，长度不超过40）
func (r *CommandParseRequest) SetExternalInfo(externalInfo string) {
	r.AddParameter("external_info", externalInfo)
}

// SetPlatform 设置平台：0抖音，1抖音极速版，默认0
func (r *CommandParseRequest) SetPlatform(platform int) {
	r.AddParameter("platform", platform)
}

// SetShareType 设置转链类型：1-deep link，2-二维码，3-口令，4-zlink，5-ShareLink；不填默认只有1
func (r *CommandParseRequest) SetShareType(shareTypes []int) {
	r.AddParameter("share_type", shareTypes)
}
