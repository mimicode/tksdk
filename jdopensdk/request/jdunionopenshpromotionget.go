package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenShPromotionGetRequest jd.union.open.sh.promotion.get 深海投流账户授权模式获取推广链接接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.sh.promotion.get
type JdUnionOpenShPromotionGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenShPromotionGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenShPromotionGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenShPromotionGetRequest) GetApiName() string {
	return "jd.union.open.sh.promotion.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenShPromotionGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
