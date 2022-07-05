package tbkscmaterialoptional

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.sc.material.optional( 通用物料搜索API )
type Response struct {
	response.TopResponse
	TbkScMaterialOptionalResult Result `json:"tbk_sc_material_optional_response"`
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
	ResultList   ResultList `json:"result_list"`
	TotalResults int64      `json:"total_results"`
	RequestID    string     `json:"request_id"`
}

type ResultList struct {
	MapData []MapDatum `json:"map_data"`
}

type MapDatum struct {
	CategoryID             int64       `json:"category_id"`
	CategoryName           string      `json:"category_name"`
	CommissionRate         string      `json:"commission_rate"`
	CommissionType         string      `json:"commission_type"`
	CouponID               string      `json:"coupon_id"`
	CouponInfo             string      `json:"coupon_info"`
	CouponRemainCount      int64       `json:"coupon_remain_count"`
	CouponTotalCount       int64       `json:"coupon_total_count"`
	IncludeDxjh            string      `json:"include_dxjh"`
	IncludeMkt             string      `json:"include_mkt"`
	InfoDxjh               string      `json:"info_dxjh"`
	ItemURL                string      `json:"item_url"`
	LevelOneCategoryID     int64       `json:"level_one_category_id"`
	LevelOneCategoryName   string      `json:"level_one_category_name"`
	NumIid                 interface{} `json:"num_iid"`
	PictURL                string      `json:"pict_url"`
	Provcity               string      `json:"provcity"`
	ReservePrice           string      `json:"reserve_price"`
	SellerID               int64       `json:"seller_id"`
	ShopDsr                int64       `json:"shop_dsr"`
	ShopTitle              string      `json:"shop_title"`
	ShortTitle             string      `json:"short_title"`
	SmallImages            SmallImages `json:"small_images"`
	Title                  string      `json:"title"`
	TkTotalCommi           string      `json:"tk_total_commi"`
	TkTotalSales           string      `json:"tk_total_sales"`
	URL                    string      `json:"url"`
	UserType               int64       `json:"user_type"`
	Volume                 int64       `json:"volume"`
	WhiteImage             string      `json:"white_image"`
	ZkFinalPrice           string      `json:"zk_final_price"`
	PresaleDeposit         string      `json:"presale_deposit"`
	PresaleDiscountFeeText string      `json:"presale_discount_fee_text"`
	PresaleEndTime         int64       `json:"presale_end_time"`
	PresaleStartTime       int64       `json:"presale_start_time"`
	PresaleTailEndTime     int64       `json:"presale_tail_end_time"`
	PresaleTailStartTime   int64       `json:"presale_tail_start_time"`
	TopnInfo               TopnInfo    `json:"topn_info"`
}

type SmallImages struct {
	String []string `json:"string"`
}

type TopnInfo struct {
	TopnEndTime    interface{} `json:"topn_end_time"`
	TopnQuantity   interface{} `json:"topn_quantity"`
	TopnStartTime  interface{} `json:"topn_start_time"`
	TopnTotalCount interface{} `json:"topn_total_count"`
}
