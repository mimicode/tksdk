package jdunionopengoodscombinationpageget

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.combinationpage.get 凑单工具页生成接口
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_combinationpage_get_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析getResult
		if t.Responce.GetResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty getResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.GetResultStr), &t.Responce.GetResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.GetResult.Code
				t.ErrorResponse.Message = t.Responce.GetResult.Message
			}
		}
	}
	t.Responce.GetResultStr = ""
}

// Responce 响应结果
type Responce struct {
	GetResultStr string `json:"getResult"`
	GetResult    CombinationGoodsPageResult
}

// CombinationGoodsPageResult 返回结果
type CombinationGoodsPageResult struct {
	Code    int64                    `json:"code"`
	Data    CombinationGoodsPageResp `json:"data"`
	Message string                   `json:"message"`
}

// CombinationGoodsPageResp 数据明细
type CombinationGoodsPageResp struct {
	ShortURL            string                `json:"shortURL"`            //凑单工具页目标链接，以短链接形式，有效期60天
	ClickURL            string                `json:"clickURL"`            //凑单工具页目标的长链，长期有效
	FailSkuList         []FailSkuResp         `json:"failSkuList"`         //失败sku集合
	MaterialID          string                `json:"materialId"`          //推广物料url，即凑单工具页原始url
	FailCouponList      []FailCouponResp      `json:"failCouponList"`      //失败券链接集合
	FailActivityUrlList []FailActivityUrlResp `json:"failActivityUrlList"` //失败会场链接集合
}

// FailSkuResp 失败商品数据
type FailSkuResp struct {
	Sku     int64  `json:"sku"`     //商品sku
	Message string `json:"message"` //sku失败具体消息
}

// FailCouponResp 失败券链接数据
type FailCouponResp struct {
	URL     string `json:"url"`     //券链接
	Message string `json:"message"` //券链接失败具体消息
}

// FailActivityUrlResp 失败会场链接数据
type FailActivityUrlResp struct {
	URL     string `json:"url"`     //活动链接
	Message string `json:"message"` //活动链接失败具体消息
}
