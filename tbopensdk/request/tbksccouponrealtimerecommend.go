package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.tbk.sc.coupon.realtime.recommend( 好券直播API【社交】 )
//http://open.taobao.com/api.htm?docId=29820&docType=2&scopeId=12331
type TbkScCouponRealtimeRecommendRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScCouponRealtimeRecommendRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")
}

//添加请求参数
func (tk *TbkScCouponRealtimeRecommendRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScCouponRealtimeRecommendRequest) GetApiName() string {
	return "taobao.tbk.sc.coupon.realtime.recommend"
}

//返回请求参数
func (tk *TbkScCouponRealtimeRecommendRequest) GetParameters() url.Values {
	return *tk.Parameters
}
