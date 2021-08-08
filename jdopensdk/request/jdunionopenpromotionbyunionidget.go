package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenPromotionByunionidGetRequest jd.union.open.promotion.byunionid.get 工具商媒体帮助子站长获取普通推广链接和优惠券二合一推广链接
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.byunionid.get
type JdUnionOpenPromotionByunionidGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenPromotionByunionidGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenPromotionByunionidGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenPromotionByunionidGetRequest) GetApiName() string {
	return "jd.union.open.promotion.byunionid.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenPromotionByunionidGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
