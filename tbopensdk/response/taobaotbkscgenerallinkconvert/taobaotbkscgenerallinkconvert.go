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
	t.Body = result
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
	BizErrorDesc int64  `json:"biz_error_desc"` // 业务错误描述
	ResultMsg    string `json:"result_msg"`     // 结果信息
}

type Data struct {
	EventURLList    DataEventURLList    `json:"event_url_list"`     // 活动转链结果
	ItemURLList     DataItemURLList     `json:"item_url_list"`      // 商品转链结果
	MaterialURLList DataMaterialURLList `json:"material_url_list"`  // 物料转链结果
	ShopURLList     DataShopURLList     `json:"shop_url_list"`     // 店铺转链结果
}

// ==================== 活动转链（EventURLList） ====================

type DataEventURLList struct {
	EventURLList []EventURLListElement `json:"event_url_list"`
}

type EventURLListElement struct {
	// 入参的会场ID
	InputPageID string `json:"input_page_id"`
	// 链接信息
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
	// 无权限时返回的错误信息
	Msg string `json:"msg"`
	// 无权限时返回的错误码
	Code int64 `json:"code"`
}

type EventURLListLinkInfoDto struct {
	// CPS链接对应长口令，格式：数字口令 渠道ID 口令原文 商品标题
	CPSFullTpwd string `json:"cps_full_tpwd"`
	// CPS链接长链
	CPSLongURL string `json:"cps_long_url"`
	// CPS链接的短口令
	CPSShortTpwd string `json:"cps_short_tpwd"`
	// CPS链接短链
	CPSShortURL string `json:"cps_short_url"`
	// 物料类型。1—单品 2—店铺 3—会场
	MaterialType int64 `json:"material_type"`
	// 会场ID
	PageID *string `json:"page_id,omitempty"`
	// 卖家ID
	SellerID *string `json:"seller_id,omitempty"`
}

// ==================== 商品转链（ItemURLList） ====================

type DataItemURLList struct {
	ItemURLList []ItemURLListElement `json:"item_url_list"`
}

type ItemURLListElement struct {
	// 券信息
	CouponInfoDto CouponInfoDto `json:"coupon_info_dto"`
	// 入参的商品ID
	InputItemID string `json:"input_item_id"`
	// 链接信息
	LinkInfoDto ItemURLListLinkInfoDto `json:"link_info_dto"`
	// 多件价券信息
	MultipleItemsCouponInfoList ItemURLListMultipleItemsCouponInfoList `json:"multiple_items_coupon_info_list"`
	// 营销信息
	PromotionInfoDto PromotionInfoDto `json:"promotion_info_dto"`
	// 转链成功后的补充说明信息，如"skuid已失效"
	ExtraInfo string `json:"extra_info"`
	// 无权限时返回的错误信息
	Msg string `json:"msg"`
	// 无权限时返回的错误码
	Code int64 `json:"code"`
}

type CouponInfoDto struct {
	// 券ID
	ActivityID string `json:"activity_id"`
	// 券面额，单位元
	CouponAmount string `json:"coupon_amount"`
	// 优惠券信息，格式：满XX元减XX元
	CouponDesc string `json:"coupon_desc"`
	// 优惠券结束时间，格式：yyyy-MM-dd HH:mm:ss
	CouponEndTime string `json:"coupon_end_time"`
	// 优惠券剩余数量
	CouponRemainCount int64 `json:"coupon_remain_count"`
	// 优惠券开始时间，格式：yyyy-MM-dd HH:mm:ss
	CouponStartTime string `json:"coupon_start_time"`
	// 券类型。0-店铺券 1-单品券
	CouponType int64 `json:"coupon_type"`
}

type ItemURLListLinkInfoDto struct {
	// 二合一链接的长口令
	CouponFullTpwd string `json:"coupon_full_tpwd"`
	// 二合一长链接
	CouponLongURL string `json:"coupon_long_url"`
	// 二合一链接的短口令
	CouponShortTpwd string `json:"coupon_short_tpwd"`
	// 二合一短链接
	CouponShortURL string `json:"coupon_short_url"`
	// CPS链接对应长口令
	CPSFullTpwd string `json:"cps_full_tpwd"`
	// CPS链接长链
	CPSLongURL string `json:"cps_long_url"`
	// CPS链接的短口令
	CPSShortTpwd string `json:"cps_short_tpwd"`
	// CPS链接短链
	CPSShortURL string `json:"cps_short_url"`
	// 商品ID
	ItemID *string `json:"item_id,omitempty"`
	// 物料类型。1—单品 2—店铺 3—会场 4-承接开放 5-优惠券 6-直播间 7-淘积木页
	MaterialType int64 `json:"material_type"`
	// 淘宝客口令业务类型。1-联盟口令，2-主站口令。入参物料不为淘口令时返回null
	TkBizType *int64 `json:"tk_biz_type,omitempty"`
	// 频道片二合一链接（需在required_link_type字段入参才可出参，目前支持淘金币）
	CouponChannelURL string `json:"coupon_channel_url"`
	// 频道页二合一口令（需在required_link_type字段入参才可出参，目前支持淘金币）
	CouponChannelTpwd string `json:"coupon_channel_tpwd"`
	// 频道页CPS链接（需在required_link_type字段入参才可出参，目前支持淘金币）
	CPSChannelURL string `json:"cps_channel_url"`
	// 频道页CPS口令（需在required_link_type字段入参才可出参，目前支持淘金币）
	CPSChannelTpwd string `json:"cps_channel_tpwd"`
	// 超红二合一长链接（非超红活动期间返回普通推广链接）
	CouponSuperedLongURL string `json:"coupon_supered_long_url"`
	// 超红二合一短链接（非超红活动期间返回普通推广链接）
	CouponSuperedShortURL string `json:"coupon_supered_short_url"`
	// 超红二合一长口令（非超红活动期间返回普通推广链接）
	CouponSuperedLongTpwd string `json:"coupon_supered_long_tpwd"`
	// 超红二合一短口令（非超红活动期间返回普通推广链接）
	CouponSuperedShortTpwd string `json:"coupon_supered_short_tpwd"`
	// 超红cps长链接（非超红活动期间返回普通推广链接）
	CPSSuperedLongURL string `json:"cps_supered_long_url"`
	// 超红cps短链接（非超红活动期间返回普通推广链接）
	CPSSuperedShortURL string `json:"cps_supered_short_url"`
	// 超红cps长口令（非超红活动期间返回普通推广链接）
	CPSSuperedLongTpwd string `json:"cps_supered_long_tpwd"`
	// 超红cps短口令（非超红活动期间返回普通推广链接）
	CPSSuperedShortTpwd string `json:"cps_supered_short_tpwd"`
	// 根据入参工具服务商账号信息生成的新商品ID
	ISVMktID string `json:"isv_mktid"`
}

type ItemURLListMultipleItemsCouponInfoList struct {
	ItemMultiCouponPromotionInfoDTO []MultiCouponPromotionInfoDTO `json:"item_multi_coupon_promotion_info_d_t_o"`
}

type MultiCouponPromotionInfoDTO struct {
	// 优惠ID
	ActivityID string `json:"activity_id"`
	// 优惠信息，格式：满XX元减XX元 或 满x件减y元
	CouponDesc string `json:"coupon_desc"`
	// 优惠结束时间，格式：yyyy-MM-dd HH:mm:ss
	CouponEndTime string `json:"coupon_end_time"`
	// 优惠剩余数量
	CouponRemainCount int64 `json:"coupon_remain_count"`
	// 优惠开始时间，格式：yyyy-MM-dd HH:mm:ss
	CouponStartTime string `json:"coupon_start_time"`
	// 优惠名称
	CouponTitle string `json:"coupon_title"`
	// 优惠金额，单位元
	CoupoonFee string `json:"coupoon_fee"`
}

type PromotionInfoDto struct {
	// 商品收入比率(%)=商品佣金比率+补贴比率
	CommissionRate string `json:"commission_rate"`
	// 佣金类型。MKT-营销计划，SP-定向计划，COMMON-通用计划，ZX-自选计划，BRAND_MARKET-营销宝，BRAND_EXCLUSIVE-品牌U享计划，SUPER_LINK-超链计划
	CommissionType string `json:"commission_type"`
	// 多件价需拍件数
	MultipleItemsPricesCount int64 `json:"multiple_items_prices_count"`
	// 商品到手价，单位分
	PromotionPrice *int64 `json:"promotion_price,omitempty"`
	// 商品佣金详细信息
	IncomeInfo *ItemTkIncomeDTO `json:"income_info,omitempty"`
}

// 商品佣金详细信息
type ItemTkIncomeDTO struct {
	// 商品佣金比例(%)
	PureCommissionRate string `json:"pure_commission_rate"`
	// 商品佣金金额
	PureCommissionAmount string `json:"pure_commission_amount"`
	// 补贴比例(%)
	SubsidyRate string `json:"subsidy_rate"`
	// 补贴金额
	SubsidyAmount string `json:"subsidy_amount"`
	// 补贴上限（仅在单笔订单命中补贴上限时返回，否则为空）
	SubsidyUpperFee string `json:"subsidy_upper_fee"`
	// 补贴类型
	SubsidyType string `json:"subsidy_type"`
}

// ==================== 物料转链（MaterialURLList） ====================

type DataMaterialURLList struct {
	MaterialURLList []MaterialURLListElement `json:"material_url_list"`
}

type MaterialURLListElement struct {
	// 券信息
	CouponInfoDto CouponInfoDto `json:"coupon_info_dto"`
	// 入参的物料链接
	InputMaterialURL string `json:"input_material_url"`
	// 链接信息
	LinkInfoDto ItemURLListLinkInfoDto `json:"link_info_dto"`
	// 多件价券信息
	MultipleItemsCouponInfoList MaterialURLListMultipleItemsCouponInfoList `json:"multiple_items_coupon_info_list"`
	// 营销信息
	PromotionInfoDto PromotionInfoDto `json:"promotion_info_dto"`
	// 无权限时返回的错误信息
	Msg string `json:"msg"`
	// 无权限时返回的错误码
	Code int64 `json:"code"`
}

type MaterialURLListMultipleItemsCouponInfoList struct {
	MaterialMultiCouponPromotionInfoDTO []MultiCouponPromotionInfoDTO `json:"material_multi_coupon_promotion_info_d_t_o"`
}

// ==================== 店铺转链（ShopURLList） ====================

type DataShopURLList struct {
	ShopURLList []ShopURLListElement `json:"shop_url_list"`
}

type ShopURLListElement struct {
	// 入参的店铺ID
	InputShopID string `json:"input_shop_id"`
	// 链接信息
	LinkInfoDto EventURLListLinkInfoDto `json:"link_info_dto"`
}
