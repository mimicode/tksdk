package alibabaalscunionkbitempromotion

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.item.promotion( 本地生活媒体平台口碑选品 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_item_promotion_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

// ResultResponse 响应结果
type ResultResponse struct {
	BizErrorDesc string `json:"biz_error_desc"`
	BizErrorCode int    `json:"biz_error_code"`
	Items        Items  `json:"items"`
	RequestID    string `json:"request_id"`
	SessionID    string `json:"session_id"`
	TotalCount   int64  `json:"total_count"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Items struct {
	KBItemPromotionDTO []KBItemPromotionDTO `json:"kb_item_promotion_d_t_o"`
}

type KBItemPromotionDTO struct {
	ApplyShopNum       int64       `json:"apply_shop_num"`
	Commission         string      `json:"commission"`
	Discount           string      `json:"discount"`
	ImageURL           string      `json:"image_url"`
	ItemID             int64       `json:"item_id"`
	ItemSaleEndTime    int64       `json:"item_sale_end_time"`
	ItemSaleStartTime  int64       `json:"item_sale_start_time"`
	ItemType           int64       `json:"item_type"`
	OriginalPrice      string      `json:"original_price"`
	Price              string      `json:"price"`
	ThirtySoldQuantity int64       `json:"thirty_sold_quantity"`
	Title              string      `json:"title"`
	TotalSales         int64       `json:"total_sales"`
	ValidCities        ValidCities `json:"valid_cities"`
	WxAppID            string      `json:"wx_app_id"`
	WxPath             string      `json:"wx_path"`
}

type ValidCities struct {
	String []string `json:"string"`
}
