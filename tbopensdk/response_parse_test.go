package tbopensdk

import (
	"testing"

	"github.com/mimicode/tksdk/tbopensdk/response/taobaotbkdgpunishorderget"
	"github.com/mimicode/tksdk/tbopensdk/response/taobaotbkscgenerallinkconvert"
)

// 离线解析测试：按官方文档(docId=42050)完整响应示例形状构造JSON，
// 验证 WrapResult 字段映射（无需网络与凭证）
func TestParseDgPunishOrderGetResponseShape(t *testing.T) {
	body := `{"tbk_dg_punish_order_get_response":{"result":{"data":{"results":{"result":[{"union_id":"","special_id":123,"relation_id":122344,"settle_month":201812,"punish_status":"0","violation_type":"店铺淘宝客","tk_trade_create_time":"2018-11-11 00:01:01","tb_trade_id":581552321112183166,"tb_trade_parent_id":581552321112183165,"tk_adzone_id":1212,"tk_site_id":1212,"tk_pub_id":"114469792"}]},"page_no":1,"page_size":10,"total_count":100},"biz_error_desc":"","biz_error_code":0,"result_msg":"ok","result_code":200}}}`
	r := &taobaotbkdgpunishorderget.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Msg, r.Body)
	}
	res := r.TbkDgPunishOrderGetResponse.Result
	if res.ResultCode != 200 || res.ResultMsg != "ok" || res.BizErrorCode != 0 {
		t.Fatalf("rpc result mismatch: %+v", res)
	}
	d := res.Data
	if d.PageNo != 1 || d.PageSize != 10 || d.TotalCount != 100 {
		t.Fatalf("page result mismatch: %+v", d)
	}
	o := d.Results.Result[0]
	if o.RelationID != 122344 || o.SettleMonth != 201812 ||
		o.PunishStatus != "0" || o.ViolationType != "店铺淘宝客" ||
		o.TkTradeCreateTime != "2018-11-11 00:01:01" ||
		o.TbTradeID != 581552321112183166 || o.TbTradeParentID != 581552321112183165 ||
		o.TkAdzoneID != 1212 || o.TkSiteID != 1212 || o.TkPubID != "114469792" ||
		o.UnionID != "" || o.SpecialID != 123 {
		t.Fatalf("order fields mismatch: %+v", o)
	}
}

// 离线解析测试：error_response 形状
func TestParseDgPunishOrderGetErrorResponseShape(t *testing.T) {
	body := `{"error_response":{"msg":"Remote service error","code":50,"sub_msg":"非法参数","sub_code":"isv.invalid-parameter","request_id":"req_1"}}`
	r := &taobaotbkdgpunishorderget.Response{}
	r.WrapResult(body)
	if !r.IsError() {
		t.Fatalf("should be error response: %s", r.Body)
	}
	if r.ErrorResponse.Code != 50 || r.ErrorResponse.SubCode != "isv.invalid-parameter" {
		t.Fatalf("error info mismatch: %+v", r.ErrorResponse)
	}
}

// 离线解析测试：按官方文档(docId=65412)完整响应示例形状构造JSON，
// 验证 WrapResult 对全部响应字段的映射（无需网络与凭证）
func TestParseScGeneralLinkConvertResponseShape(t *testing.T) {
	body := `{"tbk_sc_general_link_convert_response":{"request_id":"15rjogd9ee9ki","biz_error_desc":1,"result_msg":"1","data":{
		"material_url_list":{"material_url_list":[{
			"msg":"无权限","code":1001,
			"promotion_info_dto":{"commission_rate":"6.00","commission_type":"MKT","multiple_items_prices_count":2,
				"income_info":{"pure_commission_rate":"23.2","pure_commission_amount":0.5,"subsidy_rate":"11.2","subsidy_amount":0.23,"subsidy_upper_fee":"300.2","subsidy_type":"淘客补贴"}},
			"coupon_info_dto":{"coupon_end_time":"2022-12-20 23:59:59","activity_id":"19d114f1a6464c98a30b01b3cb17dad6","coupon_remain_count":99944,"coupon_amount":"80.00","coupon_start_time":"2022-12-14 00:00:00","coupon_desc":"满129元减80元","coupon_type":0},
			"link_info_dto":{"coupon_long_url":"https://uland.taobao.com/coupon/edetail?e=xxx","material_type":1,"cps_long_url":"https://s.click.taobao.com/t?e=xxx","cps_short_tpwd":"￥XODHdeIlU9B￥","tk_biz_type":2,"coupon_short_tpwd":"￥LOcqdeIlU9y￥","cps_short_url":"https://s.click.taobao.com/GvHjOPu","coupon_short_url":"https://s.click.taobao.com/HvHjOPu","coupon_full_tpwd":"48￥ CZ3457 LOcqdeIlU9y￥ https://m.tb.cn/h.UQ4ceEC","cps_full_tpwd":"47￥ CZ3457 XODHdeIlU9B￥ https://m.tb.cn/h.UQ4ceEy","isv_mktid":"aisj979ld8-8asdbalaiu",
				"coupon_channel_url":"https://uland.taobao.com/coupon_channel","coupon_channel_tpwd":"￥channelCoupon￥","cps_channel_url":"https://s.click.taobao.com/cps_channel","cps_channel_tpwd":"￥channelCPS￥",
				"coupon_supered_long_url":"https://uland.taobao.com/supered_long","coupon_supered_long_tpwd":"￥superedCouponLong￥","coupon_supered_short_url":"https://s.click.taobao.com/superedCouponShort","coupon_supered_short_tpwd":"￥superedCouponShort￥","cps_supered_long_url":"https://s.click.taobao.com/superedCpsLong","cps_supered_long_tpwd":"￥superedCpsLong￥","cps_supered_short_url":"https://s.click.taobao.com/superedCpsShort","cps_supered_short_tpwd":"￥superedCpsShort￥"},
			"input_material_url":"s.click.taobao.com","extra_info":"skuid已失效",
			"multiple_items_coupon_info_list":{"material_multi_coupon_promotion_info_d_t_o":[{"coupon_title":"商品券","coupon_desc":"满7999减1300","coupoon_fee":"1300.00","coupon_start_time":"2022-12-14 00:00:00","coupon_end_time":"2022-12-15 00:00:00","activity_id":"19d114f1a6464c98a30b01b3cb17dad6","coupon_remain_count":99987}]}
		}]},
		"item_url_list":{"item_url_list":[{
			"msg":"无权限","code":1001,
			"promotion_info_dto":{"commission_rate":"2.40","commission_type":"MKT","promotion_price":1349,"multiple_items_prices_count":2,
				"income_info":{"pure_commission_rate":"23.2","pure_commission_amount":0.5,"subsidy_rate":"11.2","subsidy_amount":0.23,"subsidy_upper_fee":"300.2","subsidy_type":"淘客补贴"}},
			"coupon_info_dto":{"coupon_end_time":"2022-12-25 23:59:59","activity_id":"642208f7db6b4fb9880a9d09832d6b8d","coupon_remain_count":99987,"coupon_amount":"170.00","coupon_start_time":"2022-12-20 00:00:00","coupon_desc":"满319元减170元"},
			"link_info_dto":{"coupon_long_url":"https://uland.taobao.com/coupon/edetail?e=xxx","material_type":1,"item_id":"i87a9ja90d8-09qrjcoa7qwl","cps_long_url":"https://s.click.taobao.com/t?e=xxx","cps_short_tpwd":"￥XODHdeIlU9B￥","coupon_short_tpwd":"￥LOcqdeIlU9y￥","cps_short_url":"https://s.click.taobao.com/GvHjOPu","coupon_short_url":"https://s.click.taobao.com/HvHjOPu","coupon_full_tpwd":"48￥ CZ3457 LOcqdeIlU9y￥ https://m.tb.cn/h.UQ4ceEC","cps_full_tpwd":"47￥ CZ3457 XODHdeIlU9B￥ https://m.tb.cn/h.UQ4ceEy","isv_mktid":"cacda-cdcd",
				"coupon_channel_url":"https://uland.taobao.com/coupon_channel","coupon_channel_tpwd":"￥channelCoupon￥","cps_channel_url":"https://s.click.taobao.com/cps_channel","cps_channel_tpwd":"￥channelCPS￥",
				"coupon_supered_long_url":"https://uland.taobao.com/supered_long","coupon_supered_long_tpwd":"￥superedCouponLong￥","coupon_supered_short_url":"https://s.click.taobao.com/superedCouponShort","coupon_supered_short_tpwd":"￥superedCouponShort￥","cps_supered_long_url":"https://s.click.taobao.com/superedCpsLong","cps_supered_long_tpwd":"￥superedCpsLong￥","cps_supered_short_url":"https://s.click.taobao.com/superedCpsShort","cps_supered_short_tpwd":"￥superedCpsShort￥"},
			"input_item_id":"i87a9ja90d8-09qrjcoa7qwl","extra_info":"skuid已失效",
			"multiple_items_coupon_info_list":{"item_multi_coupon_promotion_info_d_t_o":[{"coupon_title":"商品券","coupon_desc":"满7999减1300","coupoon_fee":"1300.00","coupon_start_time":"2022-12-14 00:00:00","coupon_end_time":"2022-12-15 00:00:00","activity_id":"19d114f1a6464c98a30b01b3cb17dad6","coupon_remain_count":99987}]}
		}]},
		"shop_url_list":{"shop_url_list":[{"msg":"无权限","code":1001,"link_info_dto":{"material_type":2,"seller_id":"123456","cps_long_url":"https://s.click.taobao.com/t?e=shop","cps_short_tpwd":"￥EAqldesjTCt￥","cps_short_url":"https://s.click.taobao.com/FvHjOPu","cps_full_tpwd":"58￥ CZ3457 EAqldesjTCt￥ https://m.tb.cn/h.UQ4ceEB"},"input_shop_id":"123"}]},
		"event_url_list":{"event_url_list":[{"msg":"无权限","code":1001,"link_info_dto":{"material_type":3,"page_id":"20150318020014311","cps_long_url":"https://s.click.taobao.com/t?e=event","cps_short_tpwd":"￥PnIodeIlU9z￥","cps_short_url":"https://s.click.taobao.com/EvHjOPu","cps_full_tpwd":"37￥ CZ3457 PnIodeIlU9z￥ https://m.tb.cn/h.UQ4ceEz"},"input_page_id":"203567898765"}]}
	}}}`

	r := &taobaotbkscgenerallinkconvert.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Msg, r.Body)
	}
	res := r.TbkScGeneralLinkConvertResponse
	if res.RequestID != "15rjogd9ee9ki" || res.BizErrorDesc != 1 || res.ResultMsg != "1" {
		t.Fatalf("response root mismatch: %+v", res)
	}

	// material_url_list 全部字段
	m := res.Data.MaterialURLList.MaterialURLList[0]
	if m.Msg != "无权限" || m.Code != 1001 || m.InputMaterialURL != "s.click.taobao.com" || m.ExtraInfo != "skuid已失效" {
		t.Fatalf("material element mismatch: %+v", m)
	}
	if m.CouponInfoDto.CouponType != 0 || m.CouponInfoDto.CouponAmount != "80.00" {
		t.Fatalf("material coupon_info_dto mismatch: %+v", m.CouponInfoDto)
	}
	if m.PromotionInfoDto.CommissionRate != "6.00" || m.PromotionInfoDto.CommissionType != "MKT" ||
		m.PromotionInfoDto.MultipleItemsPricesCount != 2 {
		t.Fatalf("material promotion_info_dto mismatch: %+v", m.PromotionInfoDto)
	}
	mi := m.PromotionInfoDto.IncomeInfo
	if mi.PureCommissionRate != "23.2" || mi.PureCommissionAmount != 0.5 || mi.SubsidyRate != "11.2" ||
		mi.SubsidyAmount != 0.23 || mi.SubsidyUpperFee != "300.2" || mi.SubsidyType != "淘客补贴" {
		t.Fatalf("material income_info mismatch: %+v", mi)
	}
	ml := m.LinkInfoDto
	if ml.ISVMktid != "aisj979ld8-8asdbalaiu" ||
		ml.CouponChannelURL != "https://uland.taobao.com/coupon_channel" || ml.CouponChannelTpwd != "￥channelCoupon￥" ||
		ml.CPSChannelURL != "https://s.click.taobao.com/cps_channel" || ml.CPSChannelTpwd != "￥channelCPS￥" {
		t.Fatalf("material channel fields mismatch: %+v", ml)
	}
	if ml.CouponSuperedLongURL != "https://uland.taobao.com/supered_long" || ml.CouponSuperedLongTpwd != "￥superedCouponLong￥" ||
		ml.CouponSuperedShortURL != "https://s.click.taobao.com/superedCouponShort" || ml.CouponSuperedShortTpwd != "￥superedCouponShort￥" ||
		ml.CPSSuperedLongURL != "https://s.click.taobao.com/superedCpsLong" || ml.CPSSuperedLongTpwd != "￥superedCpsLong￥" ||
		ml.CPSSuperedShortURL != "https://s.click.taobao.com/superedCpsShort" || ml.CPSSuperedShortTpwd != "￥superedCpsShort￥" {
		t.Fatalf("material supered fields mismatch: %+v", ml)
	}
	mm := m.MultipleItemsCouponInfoList.MaterialMultiCouponPromotionInfoDTO[0]
	if mm.CouponTitle != "商品券" || mm.CoupoonFee != "1300.00" || mm.CouponRemainCount != 99987 {
		t.Fatalf("material multi coupon mismatch: %+v", mm)
	}

	// item_url_list 全部字段
	i := res.Data.ItemURLList.ItemURLList[0]
	if i.Msg != "无权限" || i.Code != 1001 || i.InputItemID != "i87a9ja90d8-09qrjcoa7qwl" || i.ExtraInfo != "skuid已失效" {
		t.Fatalf("item element mismatch: %+v", i)
	}
	if i.PromotionInfoDto.PromotionPrice == nil || *i.PromotionInfoDto.PromotionPrice != 1349 ||
		i.PromotionInfoDto.IncomeInfo.SubsidyUpperFee != "300.2" {
		t.Fatalf("item promotion_info_dto mismatch: %+v", i.PromotionInfoDto)
	}
	il := i.LinkInfoDto
	if il.ItemID == nil || *il.ItemID != "i87a9ja90d8-09qrjcoa7qwl" || il.ISVMktid != "cacda-cdcd" ||
		il.CouponSuperedLongURL != "https://uland.taobao.com/supered_long" || il.CouponSuperedLongTpwd != "￥superedCouponLong￥" ||
		il.CouponSuperedShortURL != "https://s.click.taobao.com/superedCouponShort" || il.CouponSuperedShortTpwd != "￥superedCouponShort￥" ||
		il.CPSSuperedLongURL != "https://s.click.taobao.com/superedCpsLong" || il.CPSSuperedLongTpwd != "￥superedCpsLong￥" ||
		il.CPSSuperedShortURL != "https://s.click.taobao.com/superedCpsShort" || il.CPSSuperedShortTpwd != "￥superedCpsShort￥" {
		t.Fatalf("item link_info_dto mismatch: %+v", il)
	}
	im := i.MultipleItemsCouponInfoList.ItemMultiCouponPromotionInfoDTO[0]
	if im.CouponTitle != "商品券" || im.CoupoonFee != "1300.00" || im.CouponRemainCount != 99987 {
		t.Fatalf("item multi coupon mismatch: %+v", im)
	}

	// shop_url_list / event_url_list 的 msg、code
	s := res.Data.ShopURLList.ShopURLList[0]
	if s.Msg != "无权限" || s.Code != 1001 || s.LinkInfoDto.SellerID == nil || *s.LinkInfoDto.SellerID != "123456" {
		t.Fatalf("shop element mismatch: %+v", s)
	}
	e := res.Data.EventURLList.EventURLList[0]
	if e.Msg != "无权限" || e.Code != 1001 || e.LinkInfoDto.PageID == nil || *e.LinkInfoDto.PageID != "20150318020014311" {
		t.Fatalf("event element mismatch: %+v", e)
	}
}
