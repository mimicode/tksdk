package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenPromotionCommonGetRequest jd.union.open.promotion.common.get 网站/APP来获取的推广链接
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.common.get
type JdUnionOpenPromotionCommonGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenPromotionCommonGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenPromotionCommonGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenPromotionCommonGetRequest) GetApiName() string {
	return "jd.union.open.promotion.common.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenPromotionCommonGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
