package productlink

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 商品转链响应
type Response struct {
	response.TopResponse
	Data ResponseData `json:"data"`
}

func (r *Response) WrapResult(result string) {
	r.Body = result
	err := json.Unmarshal([]byte(result), r)
	if err != nil {
		r.Code = -1
		r.Desc = err.Error()
	}
}

// ResponseData 响应数据
type ResponseData struct {
	DyPassword                     string          `json:"dy_password"`                      // 抖音口令
	DyQrCode                       *DyQrCode       `json:"dy_qr_code"`                      // 抖音二维码
	DyDeeplink                     string          `json:"dy_deeplink"`                     // 抖音deep link
	DyZlink                        string          `json:"dy_zlink"`                        // 抖音zlink
	DySharelink                    string          `json:"dy_sharelink"`                    // 抖音sharelink
	CouponLink                     *CouponLink     `json:"coupon_link"`                     // 优惠价推广链接
	PublicPlanProductLinkResultInfo json.RawMessage `json:"public_plan_product_link_result_info"` // 公共计划转链信息
}

// DyQrCode 抖音二维码
type DyQrCode struct {
	URL    string `json:"url"`    // 二维码地址
	Width  int    `json:"width"`  // 二维码宽度
	Height int    `json:"height"` // 二维码高度
}

// CouponLink 优惠价推广链接
type CouponLink struct {
	CouponStatus int       `json:"coupon_status"` // 是否有优惠价&优惠券：0有优惠券，1没有优惠券
	QrCode       *QrCode   `json:"qrcode"`        // 二维码
	ShareCommand string    `json:"share_command"`  // 口令
	Deeplink     string    `json:"deeplink"`      // deeplink
	ShareLink    string    `json:"share_link"`    // 站外H5领券链接
}

// QrCode 二维码
type QrCode struct {
	URL    string `json:"url"`    // 二维码地址
	Width  int    `json:"width"`  // 宽度
	Height int    `json:"height"` // 高度
}
