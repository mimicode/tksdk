package alibabaalscunionkbbbtitemstorerelationquery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.bbt.item.store.relation.query( 本地联盟爆爆团商品门店关系 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_bbt_item_store_relation_query_response"`
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
	RequestID    string `json:"request_id"`
	Data         Data   `json:"data"`
	ResultCode   int64  `json:"result_code"`
	Message      string `json:"message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	Records Records `json:"records"`
	Total   int64   `json:"total"`
}

type Records struct {
	BbtItemShopRelationDto []BbtItemShopRelationDto `json:"bbt_item_shop_relation_dto"`
}

type BbtItemShopRelationDto struct {
	StoreID string `json:"store_id"`
}
