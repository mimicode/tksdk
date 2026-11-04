package tbopensdk

import (
	"strings"
	"testing"

	"github.com/mimicode/tksdk/tbopensdk/request"
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

// 离线解析测试：按官方文档(docId=65412)响应示例形状构造 JSON，
// 验证 required_link_type 指定后返回的 8 个超红链接字段均能映射到
// data.material_url_list[] 与 data.item_url_list[] 的 link_info_dto（无需网络与凭证）
func TestParseScGeneralLinkConvertSuperedLinkFields(t *testing.T) {
	const superedLinkInfo = `{
		"coupon_long_url":"https://uland.taobao.com/coupon/edetail?e=xxx",
		"material_type":1,
		"item_id":"i87a9ja90d8-09qrjcoa7qwl",
		"cps_long_url":"https://s.click.taobao.com/t?e=xxx",
		"cps_short_tpwd":"￥XODHdeIlU9B￥",
		"coupon_short_tpwd":"￥LOcqdeIlU9y￥",
		"cps_short_url":"https://s.click.taobao.com/GvHjOPu",
		"coupon_short_url":"https://s.click.taobao.com/HvHjOPu",
		"coupon_full_tpwd":"48￥ CZ3457 LOcqdeIlU9y￥ https://m.tb.cn/h.UQ4ceEC",
		"cps_full_tpwd":"47￥ CZ3457 XODHdeIlU9B￥ https://m.tb.cn/h.UQ4ceEy",
		"coupon_supered_long_url":"https://uland.taobao.com/coupon/edetail?e=supered_coupon_long",
		"coupon_supered_short_url":"https://s.click.taobao.com/superedCouponShort",
		"coupon_supered_long_tpwd":"47￥ CZ3457 p7lYdesjTCG￥ https://m.tb.cn/h.UQ4ceEA 超红二合一长口令",
		"coupon_supered_short_tpwd":"￥LOcqdeIlU9y￥",
		"cps_supered_long_url":"https://s.click.taobao.com/t?e=supered_cps_long",
		"cps_supered_short_url":"https://s.click.taobao.com/superedCpsShort",
		"cps_supered_long_tpwd":"47￥ CZ3457 XODHdeIlU9B￥ https://m.tb.cn/h.UQ4ceEy 超红cps长口令",
		"cps_supered_short_tpwd":"￥p7lYdesjTCG￥"
	}`

	body := `{"tbk_sc_general_link_convert_response":{"data":{` +
		`"material_url_list":{"material_url_list":[{"input_material_url":"s.click.taobao.com","link_info_dto":` + superedLinkInfo + `}]},` +
		`"item_url_list":{"item_url_list":[{"input_item_id":"i87a9ja90d8-09qrjcoa7qwl","link_info_dto":` + superedLinkInfo + `}]},` +
		`"shop_url_list":{"shop_url_list":[]},"event_url_list":{"event_url_list":[]}` +
		`},"biz_error_desc":1,"result_msg":"1"}}`

	r := &taobaotbkscgenerallinkconvert.Response{}
	r.WrapResult(body)
	if r.IsError() {
		t.Fatalf("parse failed: code=%d msg=%s body=%s", r.ErrorResponse.Code, r.ErrorResponse.Msg, r.Body)
	}
	res := r.TbkScGeneralLinkConvertResponse
	if res.BizErrorDesc != 1 || res.ResultMsg != "1" {
		t.Fatalf("rpc result mismatch: %+v", res)
	}
	materials := res.Data.MaterialURLList.MaterialURLList
	items := res.Data.ItemURLList.ItemURLList
	if len(materials) != 1 || len(items) != 1 {
		t.Fatalf("unexpected result list length: material=%d item=%d", len(materials), len(items))
	}

	// 两处 link_info_dto 均需完整映射 8 个超红字段
	assertSupered := func(name string, dto taobaotbkscgenerallinkconvert.ItemURLListLinkInfoDto) {
		if dto.CouponSuperedLongURL != "https://uland.taobao.com/coupon/edetail?e=supered_coupon_long" ||
			dto.CouponSuperedShortURL != "https://s.click.taobao.com/superedCouponShort" ||
			dto.CouponSuperedLongTpwd != "47￥ CZ3457 p7lYdesjTCG￥ https://m.tb.cn/h.UQ4ceEA 超红二合一长口令" ||
			dto.CouponSuperedShortTpwd != "￥LOcqdeIlU9y￥" ||
			dto.CPSSuperedLongURL != "https://s.click.taobao.com/t?e=supered_cps_long" ||
			dto.CPSSuperedShortURL != "https://s.click.taobao.com/superedCpsShort" ||
			dto.CPSSuperedLongTpwd != "47￥ CZ3457 XODHdeIlU9B￥ https://m.tb.cn/h.UQ4ceEy 超红cps长口令" ||
			dto.CPSSuperedShortTpwd != "￥p7lYdesjTCG￥" {
			t.Fatalf("%s supered fields mismatch: %+v", name, dto)
		}
	}
	assertSupered("material_url_list", materials[0].LinkInfoDto)
	assertSupered("item_url_list", items[0].LinkInfoDto)
}

// 离线解析测试：required_link_type 超红字段名常量与 SetRequiredLinkType 组合入参
func TestScGeneralLinkConvertRequiredLinkTypeSuperedConstants(t *testing.T) {
	all := []string{
		request.RequiredLinkTypeCouponSuperedLongURL,
		request.RequiredLinkTypeCouponSuperedLongTpwd,
		request.RequiredLinkTypeCouponSuperedShortURL,
		request.RequiredLinkTypeCouponSuperedShortTpwd,
		request.RequiredLinkTypeCPSSuperedLongURL,
		request.RequiredLinkTypeCPSSuperedLongTpwd,
		request.RequiredLinkTypeCPSSuperedShortURL,
		request.RequiredLinkTypeCPSSuperedShortTpwd,
	}
	want := "coupon_supered_long_url,coupon_supered_long_tpwd,coupon_supered_short_url,coupon_supered_short_tpwd,cps_supered_long_url,cps_supered_long_tpwd,cps_supered_short_url,cps_supered_short_tpwd"
	if got := strings.Join(all, ","); got != want {
		t.Fatalf("required_link_type constants mismatch: got %s", got)
	}

	req := &request.TbkScGeneralLinkConvertRequest{}
	req.SetRequiredLinkType(want)
	if got := req.GetParameters().Get("required_link_type"); got != want {
		t.Fatalf("required_link_type param mismatch: got %s", got)
	}
}
