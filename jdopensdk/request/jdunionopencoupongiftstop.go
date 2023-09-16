package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenCouponGiftStopRequest jd.union.open.coupon.gift.stop 礼金停止
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.coupon.gift.stop
type JdUnionOpenCouponGiftStopRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenCouponGiftStopRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenCouponGiftStopRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenCouponGiftStopRequest) GetApiName() string {
	return "jd.union.open.coupon.gift.stop"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenCouponGiftStopRequest) GetParameters() url.Values {
	return *tk.Parameters
}
