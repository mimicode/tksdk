package livelink

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 直播转链响应
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
	DyPassword string    `json:"dy_password"` // 抖音口令
	DyQrCode   *DyQrCode `json:"dy_qr_code"` // 抖音二维码
	DyDeeplink string    `json:"dy_deeplink"` // 抖音deep link
	DyZlink    string    `json:"dy_zlink"`   // 抖音zlink
}

// DyQrCode 抖音二维码
type DyQrCode struct {
	URL    string `json:"url"`    // 二维码地址
	Width  int    `json:"width"`  // 宽度
	Height int    `json:"height"` // 高度
}
