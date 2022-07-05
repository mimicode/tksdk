package tbkdgoptimusmaterial

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.dg.optimus.material( 淘宝客物料下行-导购 )
type Response struct {
	response.TopResponse
	TbkDgOptimusMaterialResult Result `json:"tbk_dg_optimus_material_response"`
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
	ResultList ResultList `json:"result_list"`
	RequestID  string     `json:"request_id"`
}

type ResultList struct {
	MapData []MapDatum `json:"map_data"`
}

type MapDatum struct {
	CategoryID             int64       `json:"category_id"`
	ClickURL               string      `json:"click_url"`
	CommissionRate         string      `json:"commission_rate"`
	CouponAmount           float64     `json:"coupon_amount"`
	CouponEndTime          string      `json:"coupon_end_time"`
	CouponRemainCount      int64       `json:"coupon_remain_count"`
	CouponStartFee         string      `json:"coupon_start_fee"`
	CouponStartTime        string      `json:"coupon_start_time"`
	CouponTotalCount       int64       `json:"coupon_total_count"`
	ItemDescription        string      `json:"item_description"`
	ItemID                 interface{} `json:"item_id"`
	LevelOneCategoryID     int64       `json:"level_one_category_id"`
	LevelOneCategoryName   string      `json:"level_one_category_name"`
	Nick                   string      `json:"nick"`
	PictURL                string      `json:"pict_url"`
	SellerID               int64       `json:"seller_id"`
	ShortTitle             string      `json:"short_title"`
	SmallImages            SmallImages `json:"small_images"`
	Title                  string      `json:"title"`
	UserType               int64       `json:"user_type"`
	Volume                 int64       `json:"volume"`
	ZkFinalPrice           string      `json:"zk_final_price"`
	CouponClickURL         *string     `json:"coupon_click_url,omitempty"`
	CouponShareURL         *string     `json:"coupon_share_url,omitempty"`
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
