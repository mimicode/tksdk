package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenPromotionBysubunionidGetRequest jd.union.open.promotion.bysubunionid.get 通过商品链接、领券链接、活动链接获取普通推广链接或优惠券二合一推广链接
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.bysubunionid.get
type JdUnionOpenPromotionBysubunionidGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenPromotionBysubunionidGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenPromotionBysubunionidGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenPromotionBysubunionidGetRequest) GetApiName() string {
	return "jd.union.open.promotion.bysubunionid.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenPromotionBysubunionidGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
