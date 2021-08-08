package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.goods.promotion.url.generate（生成多多进宝推广链接非授权）
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.promotion.url.generate
type PddDdkGoodsPromotionUrlGenerateRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsPromotionUrlGenerateRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("p_id"), "p_id")

}

//添加请求参数
func (tk *PddDdkGoodsPromotionUrlGenerateRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkGoodsPromotionUrlGenerateRequest) GetApiName() string {
	return "pdd.ddk.goods.promotion.url.generate"
}

//返回请求参数
func (tk *PddDdkGoodsPromotionUrlGenerateRequest) GetParameters() url.Values {
	return *tk.Parameters
}
