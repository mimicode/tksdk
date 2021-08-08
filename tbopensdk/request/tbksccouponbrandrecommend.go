package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.sc.coupon.brand.recommend( 品牌券API【社交】 )
//http://open.taobao.com/api.htm?docId=29823&docType=2&scopeId=12331
type TbkScCouponBrandRecommendRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScCouponBrandRecommendRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils2.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils2.CheckNumber(tk.Parameters.Get("site_id"), "site_id")
	utils2.CheckNumber(tk.Parameters.Get("site_id"), "site_id")
}

//添加请求参数
func (tk *TbkScCouponBrandRecommendRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScCouponBrandRecommendRequest) GetApiName() string {
	return "taobao.tbk.sc.coupon.brand.recommend"
}

//返回请求参数
func (tk *TbkScCouponBrandRecommendRequest) GetParameters() url.Values {
	return *tk.Parameters
}
