package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenCouponQueryRequest jd.union.open.coupon.query 通过领券链接查询优惠券的平台、面额、期限、可用状态、剩余数量等详细信息
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.coupon.query
type JdUnionOpenCouponQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenCouponQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenCouponQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenCouponQueryRequest) GetApiName() string {
	return "jd.union.open.coupon.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenCouponQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
