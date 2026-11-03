package tbopensdk

import (
	"testing"

	"github.com/mimicode/tksdk/tbopensdk/response/taobaotbkdgpunishorderget"
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
