package alibabaalscunionkbbbtitemquery

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.bbt.item.query( 本地联盟爆爆团商品列表 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_bbt_item_query_response"`
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
	ErrorMessage string `json:"error_message"`
}

func (rr ResultResponse) IsError() bool {
	return rr.BizErrorCode != 0
}

func (rr ResultResponse) Error() string {
	return fmt.Sprintf("biz_error_code:%d biz_error_desc:%s", rr.BizErrorCode, rr.BizErrorDesc)
}

type Data struct {
	SessionID  string  `json:"session_id"`
	Records    Records `json:"records"`
	PageNumber int64   `json:"page_number"`
	PageSize   int64   `json:"page_size"`
	Total      int64   `json:"total"`
}

type Records struct {
	BbtItemDto []BbtItemDto `json:"bbt_item_dto"`
}

type BbtItemDto struct {
	ItemID             string `json:"item_id"`
	Title              string `json:"title"`
	MainPicture        string `json:"main_picture"`
	SaleStartTime      int64  `json:"sale_start_time"`
	SaleEndTime        int64  `json:"sale_end_time"`
	OriginalPriceCent  int64  `json:"original_price_cent"`
	ActivityPriceCent  int64  `json:"activity_price_cent"`
	CommissionRate     string `json:"commission_rate"`
	ApplyShopCount     int64  `json:"apply_shop_count"`
	TotalSales         int64  `json:"total_sales"`
	Stock              int64  `json:"stock"`
	ItemSubType        string `json:"item_sub_type"`
	UseTimes           int64  `json:"use_times"`
	NeedPhone          bool   `json:"need_phone"`
	TBCategory2_ID     string `json:"tb_category_2_id"`
	TBCategory2_Name   string `json:"tb_category_2_name"`
	TBCategory3_ID     string `json:"tb_category_3_id"`
	TBCategory3_Name   string `json:"tb_category_3_name"`
	BuyLimit           int64  `json:"buy_limit"`
	Brand              Brand  `json:"brand"`
	TripartiteName     string `json:"tripartite_name"`
	TripartiteAppkey   string `json:"tripartite_appkey"`
	TripartiteSiteName string `json:"tripartite_site_name"`
}

type Brand struct {
	BrandID   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
}
