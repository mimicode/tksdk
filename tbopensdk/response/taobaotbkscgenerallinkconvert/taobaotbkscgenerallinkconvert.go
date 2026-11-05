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
	RequestID    string `json:"request_id"` //平台颁发的每次请求访问的唯一标识
}

type Data struct {
	EventURLList    DataEventURLList    `json:"event_url_list"`
	ItemURLList     DataItemURLList     `json:"item_url_list"`
	MaterialURLList DataMaterialURLList `json:"material_url_list"`
	ShopURLList     DataShopURLList     `json:"shop_url_list"`
}

type DataEventURLList struct {
	EventURLList []EventURLListElement `json:"event_url_list"`
}

type EventURLListElement struct {
	InputPageID string                  `json:"input_page_id"`
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
	Msg         string                  `json:"msg"`  //物料对应错误描述
	Code        int64                   `json:"code"` //物料对应错误码
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

type DataItemURLList struct {
	ItemURLList []ItemURLListElement `json:"item_url_list"`
}

type ItemURLListElement struct {
	CouponInfoDto               CouponInfoDto                          `json:"coupon_info_dto"`
	InputItemID                 string                                 `json:"input_item_id"`
	LinkInfoDto                 ItemURLListLinkInfoDto                 `json:"link_info_dto"`
	MultipleItemsCouponInfoList ItemURLListMultipleItemsCouponInfoList `json:"multiple_items_coupon_info_list"`
	PromotionInfoDto            PromotionInfoDto                       `json:"promotion_info_dto"`
	Msg                         string                                 `json:"msg"`        //物料对应错误描述
	Code                        int64                                  `json:"code"`       //物料对应错误码
	ExtraInfo                   string                                 `json:"extra_info"` //转链成功的场景下，需要补充说明的信息
}

type CouponInfoDto struct {
	ActivityID        string `json:"activity_id"`
	CouponAmount      string `json:"coupon_amount"`
	CouponDesc        string `json:"coupon_desc"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponRemainCount int64  `json:"coupon_remain_count"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponType        int64  `json:"coupon_type"` //0-店铺 1-单品
}

type ItemURLListLinkInfoDto struct {
	CouponFullTpwd         string  `json:"coupon_full_tpwd"`
	CouponLongURL          string  `json:"coupon_long_url"`
	CouponShortTpwd        string  `json:"coupon_short_tpwd"`
	CouponShortURL         string  `json:"coupon_short_url"`
	CPSFullTpwd            string  `json:"cps_full_tpwd"`
	CPSLongURL             string  `json:"cps_long_url"`
	CPSShortTpwd           string  `json:"cps_short_tpwd"`
	CPSShortURL            string  `json:"cps_short_url"`
	ItemID                 *string `json:"item_id,omitempty"`
	MaterialType           int64   `json:"material_type"`
	TkBizType              *int64  `json:"tk_biz_type,omitempty"`
	ISVMktid               string  `json:"isv_mktid"`                 //根据入参工具服务商账号信息生成的新商品ID
	CouponChannelURL       string  `json:"coupon_channel_url"`        //频道片二合一链接，目前支持淘金币
	CouponChannelTpwd      string  `json:"coupon_channel_tpwd"`       //频道页二合一口令，必须在required_link_type字段入参才可出参 目前支持淘金币
	CPSChannelURL          string  `json:"cps_channel_url"`           //频道页CPS链接，必须在required_link_type字段入参才可出参 目前支持淘金币
	CPSChannelTpwd         string  `json:"cps_channel_tpwd"`          //频道页CPS口令，必须在required_link_type字段入参才可出参 目前支持淘金币
	CouponSuperedLongURL   string  `json:"coupon_supered_long_url"`   //超红二合一长链接，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CouponSuperedLongTpwd  string  `json:"coupon_supered_long_tpwd"`  //超红二合一长口令，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CouponSuperedShortURL  string  `json:"coupon_supered_short_url"`  //超红二合一短链接，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CouponSuperedShortTpwd string  `json:"coupon_supered_short_tpwd"` //超红二合一短口令，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CPSSuperedLongURL      string  `json:"cps_supered_long_url"`      //超红cps长链接，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CPSSuperedLongTpwd     string  `json:"cps_supered_long_tpwd"`     //超红cps长口令，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CPSSuperedShortURL     string  `json:"cps_supered_short_url"`     //超红cps短链接，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
	CPSSuperedShortTpwd    string  `json:"cps_supered_short_tpwd"`    //超红cps短口令，必须在required_link_type字段入参才可出参（非超红活动期间，请求该字段返回普通推广链接）
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
	CommissionRate           string     `json:"commission_rate"`
	CommissionType           string     `json:"commission_type"`
	MultipleItemsPricesCount int64      `json:"multiple_items_prices_count"`
	PromotionPrice           *int64     `json:"promotion_price,omitempty"`
	IncomeInfo               IncomeInfo `json:"income_info"` //商品佣金信息
}

type IncomeInfo struct {
	PureCommissionRate   string  `json:"pure_commission_rate"`   //商品佣金比例
	PureCommissionAmount float64 `json:"pure_commission_amount"` //商品佣金金额
	SubsidyRate          string  `json:"subsidy_rate"`           //补贴比例
	SubsidyAmount        float64 `json:"subsidy_amount"`         //补贴金额
	SubsidyUpperFee      string  `json:"subsidy_upper_fee"`      //补贴上限；仅在单笔订单命中补贴上限时返回结果否则出参为空
	SubsidyType          string  `json:"subsidy_type"`           //补贴类型
}

type DataMaterialURLList struct {
	MaterialURLList []MaterialURLListElement `json:"material_url_list"`
}

type MaterialURLListElement struct {
	CouponInfoDto               CouponInfoDto                              `json:"coupon_info_dto"`
	InputMaterialURL            string                                     `json:"input_material_url"`
	LinkInfoDto                 ItemURLListLinkInfoDto                     `json:"link_info_dto"`
	MultipleItemsCouponInfoList MaterialURLListMultipleItemsCouponInfoList `json:"multiple_items_coupon_info_list"`
	PromotionInfoDto            PromotionInfoDto                           `json:"promotion_info_dto"`
	Msg                         string                                     `json:"msg"`        //物料对应错误描述
	Code                        int64                                      `json:"code"`       //物料对应错误码
	ExtraInfo                   string                                     `json:"extra_info"` //转链成功的场景下，需要补充说明的信息
}

type MaterialURLListMultipleItemsCouponInfoList struct {
	MaterialMultiCouponPromotionInfoDTO []MultiCouponPromotionInfoDTO `json:"material_multi_coupon_promotion_info_d_t_o"`
}

type DataShopURLList struct {
	ShopURLList []ShopURLListElement `json:"shop_url_list"`
}

type ShopURLListElement struct {
	InputShopID string                  `json:"input_shop_id"`
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
	Msg         string                  `json:"msg"`  //物料对应错误描述
	Code        int64                   `json:"code"` //物料对应错误码
}
