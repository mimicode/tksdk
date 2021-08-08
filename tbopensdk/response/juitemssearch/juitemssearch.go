package juitemssearch

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.ju.items.search( 聚划算商品搜索接口 )
type Response struct {
	response.TopResponse
	JuItemsSearchResult Result `json:"ju_items_search_response"`
}

//解析输出结果
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

type Result struct {
	Result    ResultItem `json:"result"`
	RequestID string     `json:"request_id"`
}

type ResultItem struct {
	CurrentPage int64       `json:"current_page"`
	ModelList   ModelList   `json:"model_list"`
	PageSize    int64       `json:"page_size"`
	Success     bool        `json:"success"`
	TotalItem   int64       `json:"total_item"`
	TotalPage   int64       `json:"total_page"`
	TrackParams TrackParams `json:"track_params"`
}

type ModelList struct {
	Items []Item `json:"items"`
}

type Item struct {
	ActPrice        string `json:"act_price"`
	CategoryName    string `json:"category_name"`
	ItemID          int64  `json:"item_id"`
	ItemUspList     List   `json:"item_usp_list"`
	JuID            int64  `json:"ju_id"`
	OnlineEndTime   int64  `json:"online_end_time"`
	OnlineStartTime int64  `json:"online_start_time"`
	OrigPrice       string `json:"orig_price"`
	PayPostage      bool   `json:"pay_postage"`
	PCURL           string `json:"pc_url"`
	PicURLForPC     string `json:"pic_url_for_p_c"`
	PicURLForWL     string `json:"pic_url_for_w_l"`
	PlatformID      int64  `json:"platform_id"`
	PriceUspList    List   `json:"price_usp_list"`
	ShowEndTime     int64  `json:"show_end_time"`
	ShowStartTime   int64  `json:"show_start_time"`
	TBFirstCatID    int64  `json:"tb_first_cat_id"`
	Title           string `json:"title"`
	UspDescList     List   `json:"usp_desc_list"`
	WAPURL          string `json:"wap_url"`
}

type List struct {
	String []string `json:"string"`
}

type TrackParams struct {
	Empty bool `json:"empty"`
}
