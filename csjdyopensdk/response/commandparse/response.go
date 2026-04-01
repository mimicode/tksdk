package commandparse

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 抖口令解析响应
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
	CommandType                     int              `json:"command_type"`                     // 1=商品，2=直播间/直播间商品，3=活动页
	ProductInfo                     *ProductInfo     `json:"product_info"`                     // command_type=1时返回
	LiveInfo                        *LiveInfo        `json:"live_info"`                        // command_type=2时返回
	ActivityInfo                    *ActivityInfo    `json:"activity_info"`                    // command_type=3时返回
	IsOwnCommand                    bool             `json:"is_own_command"`                    // 是否是自己的抖口令
	PublicPlanProductLinkResultInfo json.RawMessage  `json:"public_plan_product_link_result_info"` // 公共计划转链信息
}

// ProductInfo 商品信息（command_type=1时返回）
type ProductInfo struct {
	ProductId        int64   `json:"product_id"`        // 商品id
	InsActivityParam string  `json:"ins_activity_param"` // 团长参数
	ShareInfo        *ShareInfo `json:"share_info"`    // 分享物料
}

// LiveInfo 直播间信息（command_type=2时返回）
type LiveInfo struct {
	AuthorBuyinId string     `json:"author_buyin_id"` // 直播间buyin_id
	ProductId     int64      `json:"product_id"`     // 商品id（直播间内商品口令时返回）
	ShareInfo      *ShareInfo `json:"share_info"`    // 分享物料
}

// ActivityInfo 活动页信息（command_type=3时返回）
type ActivityInfo struct {
	MaterialId  string     `json:"material_id"`  // 聚合页物料id
	ExtraParams string     `json:"extra_params"` // 活动页自定义参数
	ShareInfo   *ShareInfo `json:"share_info"`   // 分享物料
}

// ShareInfo 分享物料
type ShareInfo struct {
	DeepLink     string   `json:"deep_link"`      // 抖音deep link
	ShareCommand string   `json:"share_command"`  // 抖音口令
	QrCode       *QrCode  `json:"qrcode"`         // 二维码
	Zlink        string   `json:"zlink"`          // 抖音zlink
	ShareLink    string   `json:"share_link"`     // 抖音sharelink
}

// QrCode 二维码
type QrCode struct {
	URL    string `json:"url"`    // 二维码地址
	Width  int    `json:"width"`  // 宽度
	Height int    `json:"height"` // 高度
}
