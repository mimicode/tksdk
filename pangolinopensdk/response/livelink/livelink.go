package livelink

import (
	"encoding/json"

	"github.com/mimicode/tksdk/pangolinopensdk/response"
)

// Response 直播转链响应
type Response struct {
	response.TopResponse
	Data Data `json:"data"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	t.RawData = result
	if err := json.Unmarshal([]byte(result), t); err != nil {
		t.Code = -1
		t.Desc = err.Error()
	}
}

// Data 响应数据
type Data struct {
	DyPassword string    `json:"dy_password"` // 抖音口令
	DyQRLink   *DyQRLink `json:"dy_qr_code"`  // 抖音二维码
	DyDeeplink string    `json:"dy_deeplink"` // 抖音deep link
	DyZlink    string    `json:"dy_zlink"`    // 抖音zlink
}

// DyQRLink 抖音二维码
type DyQRLink struct {
	URL    string `json:"url"`    // 二维码地址
	Width  int    `json:"width"`  // 宽度
	Height int    `json:"height"` // 高度
}
