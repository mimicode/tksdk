package alibabaalscunionkbbbtitemdetailget

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/response"
)

// Response alibaba.alsc.union.kb.bbt.item.detail.get( 本地联盟爆爆团商品详情 )
type Response struct {
	response.TopResponse
	ResultResponse ResultResponse `json:"alibaba_alsc_union_kb_bbt_item_detail_get_response"`
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
	ItemID                string        `json:"item_id"`
	Title                 string        `json:"title"`
	SubTitle              string        `json:"sub_title"`
	MainPicture           string        `json:"main_picture"`
	Images                Images        `json:"images"`
	SaleStartTime         int64         `json:"sale_start_time"`
	SaleEndTime           int64         `json:"sale_end_time"`
	OriginalPriceCent     int64         `json:"original_price_cent"`
	ActivityPriceCent     int64         `json:"activity_price_cent"`
	CommissionRate        string        `json:"commission_rate"`
	TotalSales            int64         `json:"total_sales"`
	Stock                 int64         `json:"stock"`
	ApplyShopCount        int64         `json:"apply_shop_count"`
	Discount              string        `json:"discount"`
	ItemDetail            ItemDetail    `json:"item_detail"`
	ItemSubType           string        `json:"item_sub_type"`
	UseTimes              int64         `json:"use_times"`
	NeedPhone             bool          `json:"need_phone"`
	TBCategory2_ID        string        `json:"tb_category_2_id"`
	TBCategory2_Name      string        `json:"tb_category_2_name"`
	TBCategory3_ID        string        `json:"tb_category_3_id"`
	TBCategory3_Name      string        `json:"tb_category_3_name"`
	BuyLimit              int64         `json:"buy_limit"`
	ShopItemImages        Images        `json:"shop_item_images"`
	ShopEnvironmentImages Images        `json:"shop_environment_images"`
	Brand                 Brand         `json:"brand"`
	TripartiteName        string        `json:"tripartite_name"`
	TripartiteAppkey      string        `json:"tripartite_appkey"`
	TripartiteSiteName    string        `json:"tripartite_site_name"`
	Disclaimer            string        `json:"disclaimer"`
	PurchaseLimit         PurchaseLimit `json:"purchase_limit"`
}

type Brand struct {
	BrandID   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

type Images struct {
	ImageDto []ImageDto `json:"image_dto"`
}

type ImageDto struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ItemDetail struct {
	ItemContent ItemContent `json:"item_content"`
	ItemBuyNote ItemBuyNote `json:"item_buy_note"`
	ItemTicket  ItemTicket  `json:"item_ticket"`
}

type ItemBuyNote struct {
	ShopInfo   ShopInfo   `json:"shop_info"`
	UseNote    UseNote    `json:"use_note"`
	ExtraNotes ExtraNotes `json:"extra_notes"`
}

type ExtraNotes struct {
	TextContent []ExtraNotesTextContent `json:"text_content"`
}

type ExtraNotesTextContent struct {
	Title    string  `json:"title"`
	Desc     string  `json:"desc"`
	Contents Remarks `json:"contents"`
}

type Remarks struct {
	String []string `json:"string"`
}

type ShopInfo struct {
	FreeWifi                bool   `json:"free_wifi"`
	FreePark                bool   `json:"free_park"`
	FreeParkHours           int64  `json:"free_park_hours"`
	ParkFeePerHour          string `json:"park_fee_per_hour"`
	ParkFeeUpperBoundPerDay string `json:"park_fee_upper_bound_per_day"`
	SupplyInvoice           bool   `json:"supply_invoice"`
}

type UseNote struct {
	NeedReserve    bool   `json:"need_reserve"`
	ReserveDesc    string `json:"reserve_desc"`
	LimitUserNum   bool   `json:"limit_user_num"`
	UserNumLimited int64  `json:"user_num_limited"`
}

type ItemContent struct {
	ContentGroups ContentGroups `json:"content_groups"`
	ImageContents ImageContents `json:"image_contents"`
	TextContents  ExtraNotes    `json:"text_contents"`
	Remarks       Remarks       `json:"remarks"`
	Announcement  string        `json:"announcement"`
}

type ContentGroups struct {
	ItemContentGroup []ItemContentGroup `json:"item_content_group"`
}

type ItemContentGroup struct {
	Title          string         `json:"title"`
	ContentDetails ContentDetails `json:"content_details"`
}

type ContentDetails struct {
	ContentDetail []ContentDetail `json:"content_detail"`
}

type ContentDetail struct {
	Name         string       `json:"name"`
	Price        string       `json:"price"`
	Quantity     int64        `json:"quantity"`
	SumPrice     string       `json:"sum_price"`
	Unit         string       `json:"unit"`
	Spec         string       `json:"spec"`
	TextContents TextContents `json:"text_contents"`
}

type TextContents struct {
	TextContent []PurpleTextContent `json:"text_content"`
}

type PurpleTextContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ImageContents struct {
	ImageContent []ImageContent `json:"image_content"`
}

type ImageContent struct {
	Title string  `json:"title"`
	Desc  string  `json:"desc"`
	Urls  Remarks `json:"urls"`
}

type ItemTicket struct {
	TicketPeriod    TicketPeriod    `json:"ticket_period"`
	TicketTimeRules TicketTimeRules `json:"ticket_time_rules"`
}

type TicketPeriod struct {
	PeriodType string `json:"period_type"`
	Period     int64  `json:"period"`
	NatureDay  bool   `json:"nature_day"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
}

type TicketTimeRules struct {
	TicketTimeRule []TicketTimeRule `json:"ticket_time_rule"`
}

type TicketTimeRule struct {
	RuleApplyMode string  `json:"rule_apply_mode"`
	HourMinRules  Remarks `json:"hour_min_rules"`
	WeekRules     Remarks `json:"week_rules"`
	DateRules     Remarks `json:"date_rules"`
}

type PurchaseLimit struct {
	ItemDailyLimitPerUser     int64 `json:"item_daily_limit_per_user"`
	ItemLimitPerUser          int64 `json:"item_limit_per_user"`
	ItemLimitPerUserOrder     int64 `json:"item_limit_per_user_order"`
	ActivityDailyLimitPerUser int64 `json:"activity_daily_limit_per_user"`
	ActivityLimitPerUser      int64 `json:"activity_limit_per_user"`
}
