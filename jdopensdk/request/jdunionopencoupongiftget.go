package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenCouponGiftGetRequest jd.union.open.coupon.gift.get 礼金创建
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.coupon.gift.get
type JdUnionOpenCouponGiftGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenCouponGiftGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenCouponGiftGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenCouponGiftGetRequest) GetApiName() string {
	return "jd.union.open.coupon.gift.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenCouponGiftGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
