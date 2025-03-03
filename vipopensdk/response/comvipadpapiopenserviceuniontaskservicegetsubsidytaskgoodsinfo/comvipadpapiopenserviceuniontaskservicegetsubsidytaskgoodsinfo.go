package comvipadpapiopenserviceuniontaskservicegetsubsidytaskgoodsinfo

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionTaskService getSubsidyTaskGoodsInfo 补贴活动商品列表
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Result struct {
	GoodsInfoList []ChannelSubsidyGoodsInfo `json:"goodsInfoList"` // 精选商品列表
	Page          int                       `json:"page"`          // 当前页码
	PageSize      int                       `json:"pageSize"`      // 分页大小
	TaskBeginTime int64                     `json:"taskBeginTime"` // 活动开始时间
	TaskEndTime   int64                     `json:"taskEndTime"`   // 活动结束时间
}

type ChannelSubsidyGoodsInfo struct {
	GoodsName       string `json:"goodsName"`       // 商品标题
	GoodsImage      string `json:"goodsImage"`      // 商品图片链接
	CouponPrice     string `json:"couponPrice"`     // 到手价（取券后价）
	SaleMarketPrice string `json:"saleMarketPrice"` // 划线价（市场价）
	SelfBuy         string `json:"selfBuy"`         // 购买省
	ShareBuy        string `json:"shareBuy"`        // 分享赚
	Commission      string `json:"commission"`      // 分享赚-佣金
	Subsidy         string `json:"subsidy"`         // 分享赚-补贴
	DetailUrl       string `json:"detailUrl"`       // 详情页链接
	PromoteUrl      string `json:"promoteUrl"`      // 推广链接
	GoodsId         string `json:"goodsId"`         // 商品id
	BrandId         string `json:"brandId"`         // 商品 brandId
}
