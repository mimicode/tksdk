package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenStatisticsGiftcouponQueryRequest jd.union.open.statistics.giftcoupon.query 礼金效果数据
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.statistics.giftcoupon.query
type JdUnionOpenStatisticsGiftcouponQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenStatisticsGiftcouponQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenStatisticsGiftcouponQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenStatisticsGiftcouponQueryRequest) GetApiName() string {
	return "jd.union.open.statistics.giftcoupon.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenStatisticsGiftcouponQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
