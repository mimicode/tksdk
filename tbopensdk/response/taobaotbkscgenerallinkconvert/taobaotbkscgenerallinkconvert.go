package taobaotbkscgenerallinkconvert

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.general.link.convert( 淘宝客-服务商-万能转链 )
type Response struct {
	response.TopResponse
	TbkScGeneralLinkConvertResponse TbkScGeneralLinkConvertResponse `json:"tbk_sc_general_link_convert_response"`
}

// 解析输出结果
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

type Welcome struct {
	TbkScGeneralLinkConvertResponse TbkScGeneralLinkConvertResponse `json:"tbk_sc_general_link_convert_response"`
}

type TbkScGeneralLinkConvertResponse struct {
	Data         Data   `json:"data"`
	BizErrorDesc int64  `json:"biz_error_desc"`
	ResultMsg    string `json:"result_msg"`
}

type Data struct {
	MaterialURLList MaterialURLList `json:"material_url_list"`
	ShopURLList     ShopURLList     `json:"shop_url_list"`
	EventURLList    EventURLList    `json:"event_url_list"`
	ItemURLList     ItemURLList     `json:"item_url_list"`
}

type EventURLList struct {
	EventURLList []EventURLListElement `json:"event_url_list"`
}

type EventURLListElement struct {
	Msg         string                  `json:"msg"`
	Code        int64                   `json:"code"`
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
}

type EventURLListLinkInfoDto struct {
	MaterialType int64  `json:"material_type"`
	PageID       string `json:"page_id"`
	CPSLongURL   string `json:"cps_long_url"`
	CPSShortTpwd string `json:"cps_short_tpwd"`
	CPSShortURL  string `json:"cps_short_url"`
	CPSFullTpwd  string `json:"cps_full_tpwd"`
	SellerID     string `json:"seller_id"`
}

type ItemURLList struct {
	ItemURLList []ItemURLListElement `json:"item_url_list"`
}

type ItemURLListElement struct {
	Msg              string                 `json:"msg"`
	Code             int64                  `json:"code"`
	PromotionInfoDto PromotionInfoDto       `json:"promotion_info_dto"`
	CouponInfoDto    CouponInfoDto          `json:"coupon_info_dto"`
	LinkInfoDto      ItemURLListLinkInfoDto `json:"link_info_dto"`
}

type CouponInfoDto struct {
	CouponEndTime     string `json:"coupon_end_time"`
	ActivityID        string `json:"activity_id"`
	CouponRemainCount int64  `json:"coupon_remain_count"`
	CouponAmount      string `json:"coupon_amount"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponDesc        string `json:"coupon_desc"`
	CouponType        int64  `json:"coupon_type"`
}

type ItemURLListLinkInfoDto struct {
	CouponLongURL   string `json:"coupon_long_url"`
	MaterialType    int64  `json:"material_type"`
	ItemID          string `json:"item_id"`
	CPSLongURL      string `json:"cps_long_url"`
	CPSShortTpwd    string `json:"cps_short_tpwd"`
	CouponShortTpwd string `json:"coupon_short_tpwd"`
	CPSShortURL     string `json:"cps_short_url"`
	CouponShortURL  string `json:"coupon_short_url"`
	CouponFullTpwd  string `json:"coupon_full_tpwd"`
	CPSFullTpwd     string `json:"cps_full_tpwd"`
	TpwdOriginURL   string `json:"tpwd_origin_url"`
	MaterialID      string `json:"material_id"`
	TkBizType       int64  `json:"tk_biz_type"`
}

type PromotionInfoDto struct {
	CommissionRate string `json:"commission_rate"`
}

type MaterialURLList struct {
	MaterialURLList []ItemURLListElement `json:"material_url_list"`
}

type ShopURLList struct {
	ShopURLList []EventURLListElement `json:"shop_url_list"`
}
