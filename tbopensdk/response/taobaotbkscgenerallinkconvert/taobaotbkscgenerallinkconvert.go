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
	EventURLList    DataEventURLList    `json:"event_url_list"`
	ItemURLList     DataItemURLList     `json:"item_url_list"`
	MaterialURLList DataMaterialURLList `json:"material_url_list"`
	ShopURLList     DataShopURLList     `json:"shop_url_list"`
}

// ==================== EventURLList ====================

type DataEventURLList struct {
	EventURLList []EventURLListElement `json:"event_url_list"`
}

type EventURLListElement struct {
	InputPageID string                  `json:"input_page_id"`
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
	Msg         string                  `json:"msg"` // 新增
	Code        int64                   `json:"code"` // 新增
}

type EventURLListLinkInfoDto struct {
	CPSFullTpwd  string  `json:"cps_full_tpwd"`
	CPSLongURL   string  `json:"cps_long_url"`
	CPSShortTpwd string  `json:"cps_short_tpwd"`
	CPSShortURL  string  `json:"cps_short_url"`
	MaterialType int64   `json:"material_type"`
	PageID       *string `json:"page_id,omitempty"`
	SellerID     *string `json:"seller_id,omitempty"`
}

// ==================== ItemURLList ====================

type DataItemURLList struct {
	ItemURLList []ItemURLListElement `json:"item_url_list"`
}

type ItemURLListElement struct {
	CouponInfoDto               CouponInfoDto                          `json:"coupon_info_dto"`
	InputItemID                 string                                 `json:"input_item_id"`
	LinkInfoDto                 ItemURLListLinkInfoDto                 `json:"link_info_dto"`
	MultipleItemsCouponInfoList ItemURLListMultipleItemsCouponInfoList `json:"multiple_items_coupon_info_list"`
	PromotionInfoDto            PromotionInfoDto                       `json:"promotion_info_dto"`
	ExtraInfo                   string                                 `json:"extra_info"` // 新增
	Msg                         string                                 `json:"msg"`        // 新增
	Code                        int64                                  `json:"code"`       // 新增
}

type CouponInfoDto struct {
	ActivityID        string `json:"activity_id"`
	CouponAmount      string `json:"coupon_amount"`
	CouponDesc        string `json:"coupon_desc"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponRemainCount int64  `json:"coupon_remain_count"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponType        int64  `json:"coupon_type"` // 新增
}

type ItemURLListLinkInfoDto struct {
	CouponFullTpwd  string  `json:"coupon_full_tpwd"`
	CouponLongURL   string  `json:"coupon_long_url"`
	CouponShortTpwd string  `json:"coupon_short_tpwd"`
	CouponShortURL  string  `json:"coupon_short_url"`
	CPSFullTpwd     string  `json:"cps_full_tpwd"`
	CPSLongURL      string  `json:"cps_long_url"`
	CPSShortTpwd    string  `json:"cps_short_tpwd"`
	CPSShortURL     string  `json:"cps_short_url"`
	ItemID          *string `json:"item_id,omitempty"`
	MaterialType    int64   `json:"material_type"`
	TkBizType       *int64  `json:"tk_biz_type,omitempty"`
	// 新增字段
	CouponChannelURL      string `json:"coupon_channel_url"`       // coupon_channel_url
	CouponChannelTpwd     string `json:"coupon_channel_tpwd"`       // coupon_channel_tpwd
	CPSChannelURL         string `json:"cps_channel_url"`          // cps_channel_url
	CPSChannelTpwd        string `json:"cps_channel_tpwd"`         // cps_channel_tpwd
	CouponSuperedLongURL  string `json:"coupon_supered_long_url"`  // coupon_supered_long_url
	CouponSuperedShortURL string `json:"coupon_supered_short_url"` // coupon_supered_short_url
	CouponSuperedLongTpwd string `json:"coupon_supered_long_tpwd"` // coupon_supered_long_tpwd
	CouponSuperedShortTpwd string `json:"coupon_supered_short_tpwd"` // coupon_supered_short_tpwd
	CPSSuperedLongURL     string `json:"cps_supered_long_url"`      // cps_supered_long_url
	CPSSuperedShortURL    string `json:"cps_supered_short_url"`     // cps_supered_short_url
	CPSSuperedLongTpwd    string `json:"cps_supered_long_tpwd"`     // cps_supered_long_tpwd
	CPSSuperedShortTpwd   string `json:"cps_supered_short_tpwd"`    // cps_supered_short_tpwd
	ISVMktID              string `json:"isv_mktid"`                 // isv_mktid
}

type ItemURLListMultipleItemsCouponInfoList struct {
	ItemMultiCouponPromotionInfoDTO []MultiCouponPromotionInfoDTO `json:"item_multi_coupon_promotion_info_d_t_o"`
}

type MultiCouponPromotionInfoDTO struct {
	ActivityID        string `json:"activity_id"`
	CouponDesc        string `json:"coupon_desc"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponRemainCount int64  `json:"coupon_remain_count"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponTitle       string `json:"coupon_title"`
	CoupoonFee        string `json:"coupoon_fee"`
}

type PromotionInfoDto struct {
	CommissionRate           string           `json:"commission_rate"`
	CommissionType           string           `json:"commission_type"`
	MultipleItemsPricesCount int64            `json:"multiple_items_prices_count"`
	PromotionPrice           *int64           `json:"promotion_price,omitempty"`
	IncomeInfo               *ItemTkIncomeDTO `json:"income_info,omitempty"` // 新增嵌套结构体
}

// 新增 ItemTkIncomeDTO
type ItemTkIncomeDTO struct {
	PureCommissionRate  string `json:"pure_commission_rate"`  // pure_commission_rate
	PureCommissionAmount string `json:"pure_commission_amount"` // pure_commission_amount
	SubsidyRate         string `json:"subsidy_rate"`           // subsidy_rate
	SubsidyAmount       string `json:"subsidy_amount"`         // subsidy_amount
	SubsidyUpperFee     string `json:"subsidy_upper_fee"`     // subsidy_upper_fee
	SubsidyType         string `json:"subsidy_type"`           // subsidy_type
}

// ==================== MaterialURLList ====================

type DataMaterialURLList struct {
	MaterialURLList []MaterialURLListElement `json:"material_url_list"`
}

type MaterialURLListElement struct {
	CouponInfoDto               CouponInfoDto                              `json:"coupon_info_dto"`
	InputMaterialURL            string                                     `json:"input_material_url"`
	LinkInfoDto                 ItemURLListLinkInfoDto                     `json:"link_info_dto"`
	MultipleItemsCouponInfoList MaterialURLListMultipleItemsCouponInfoList `json:"multiple_items_coupon_info_list"`
	PromotionInfoDto            PromotionInfoDto                           `json:"promotion_info_dto"`
	Msg                         string                                     `json:"msg"`  // 新增
	Code                        int64                                      `json:"code"` // 新增
}

type MaterialURLListMultipleItemsCouponInfoList struct {
	MaterialMultiCouponPromotionInfoDTO []MultiCouponPromotionInfoDTO `json:"material_multi_coupon_promotion_info_d_t_o"`
}

// ==================== ShopURLList ====================

type DataShopURLList struct {
	ShopURLList []ShopURLListElement `json:"shop_url_list"`
}

type ShopURLListElement struct {
	InputShopID string                  `json:"input_shop_id"`
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
}
