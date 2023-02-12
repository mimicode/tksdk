package tbkdgmaterialoptional

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.material.optional( 通用物料搜索API（导购） )
type Response struct {
	response.TopResponse
	TbkDgMaterialOptionalResult Result `json:"tbk_dg_material_optional_response"`
}

// 解析输出结果
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
	CouponStartTime        string          `json:"coupon_start_time"`
	CouponEndTime          string          `json:"coupon_end_time"`
	InfoDxjh               string          `json:"info_dxjh"`
	TkTotalSales           string          `json:"tk_total_sales"`
	TkTotalCommi           string          `json:"tk_total_commi"`
	CouponID               string          `json:"coupon_id"`
	NumIid                 string          `json:"num_iid"`
	Title                  string          `json:"title"`
	PictURL                string          `json:"pict_url"`
	SmallImages            SmallImages     `json:"small_images"`
	ReservePrice           string          `json:"reserve_price"`
	ZkFinalPrice           string          `json:"zk_final_price"`
	UserType               int64           `json:"user_type"`
	Provcity               string          `json:"provcity"`
	ItemURL                string          `json:"item_url"`
	IncludeMkt             string          `json:"include_mkt"`
	IncludeDxjh            string          `json:"include_dxjh"`
	CommissionRate         string          `json:"commission_rate"`
	Volume                 int64           `json:"volume"`
	SellerID               int64           `json:"seller_id"`
	ShopTitle              string          `json:"shop_title"`
	CouponTotalCount       int64           `json:"coupon_total_count"`
	CouponRemainCount      int64           `json:"coupon_remain_count"`
	CouponInfo             string          `json:"coupon_info"`
	CommissionType         string          `json:"commission_type"`
	ShopDsr                int64           `json:"shop_dsr"`
	CouponShareURL         string          `json:"coupon_share_url"`
	URL                    string          `json:"url"`
	LevelOneCategoryName   string          `json:"level_one_category_name"`
	LevelOneCategoryID     int64           `json:"level_one_category_id"`
	CategoryName           string          `json:"category_name"`
	CategoryID             int64           `json:"category_id"`
	ShortTitle             string          `json:"short_title"`
	WhiteImage             string          `json:"white_image"`
	Oetime                 string          `json:"oetime"`
	Ostime                 string          `json:"ostime"`
	JddNum                 int64           `json:"jdd_num"`
	JddPrice               string          `json:"jdd_price"`
	UvSumPreSale           int64           `json:"uv_sum_pre_sale"`
	XID                    string          `json:"x_id"`
	CouponStartFee         string          `json:"coupon_start_fee"`
	CouponAmount           string          `json:"coupon_amount"`
	ItemDescription        string          `json:"item_description"`
	Nick                   string          `json:"nick"`
	OrigPrice              string          `json:"orig_price"`
	TotalStock             int64           `json:"total_stock"`
	SellNum                int64           `json:"sell_num"`
	Stock                  int64           `json:"stock"`
	TmallPlayActivityInfo  string          `json:"tmall_play_activity_info"`
	ItemID                 string          `json:"item_id"`
	RealPostFee            string          `json:"real_post_fee"`
	LockRate               string          `json:"lock_rate"`
	LockRateEndTime        int64           `json:"lock_rate_end_time"`
	LockRateStartTime      int64           `json:"lock_rate_start_time"`
	PresaleDiscountFeeText string          `json:"presale_discount_fee_text"`
	PresaleTailEndTime     int64           `json:"presale_tail_end_time"`
	PresaleTailStartTime   int64           `json:"presale_tail_start_time"`
	PresaleEndTime         int64           `json:"presale_end_time"`
	PresaleStartTime       int64           `json:"presale_start_time"`
	PresaleDeposit         string          `json:"presale_deposit"`
	YsylTljSendTime        string          `json:"ysyl_tlj_send_time"`
	YsylCommissionRate     string          `json:"ysyl_commission_rate"`
	YsylTljFace            string          `json:"ysyl_tlj_face"`
	YsylClickURL           string          `json:"ysyl_click_url"`
	YsylTljUseEndTime      string          `json:"ysyl_tlj_use_end_time"`
	YsylTljUseStartTime    string          `json:"ysyl_tlj_use_start_time"`
	SaleBeginTime          string          `json:"sale_begin_time"`
	SaleEndTime            string          `json:"sale_end_time"`
	Distance               string          `json:"distance"`
	UsableShopID           string          `json:"usable_shop_id"`
	UsableShopName         string          `json:"usable_shop_name"`
	SalePrice              string          `json:"sale_price"`
	KuadianPromotionInfo   string          `json:"kuadian_promotion_info"`
	SuperiorBrand          string          `json:"superior_brand"`
	RewardInfo             int64           `json:"reward_info"`
	IsBrandFlashSale       string          `json:"is_brand_flash_sale"`
	LocalizationExtend     string          `json:"localization_extend"`
	MatchScore             string          `json:"match_score"`
	CommiScore             string          `json:"commi_score"`
	HotFlag                string          `json:"hot_flag"`
	TopnInfo               TopnInfo        `json:"topn_info"`
	BybtInfo               BybtInfo        `json:"bybt_info"`
	TtSoldCount            string          `json:"tt_sold_count"`
	MaifanPromotion        MaifanPromotion `json:"maifan_promotion"`
	CPARewardType          string          `json:"cpa_reward_type"`
	CPARewardAmount        string          `json:"cpa_reward_amount"`
	ActivityID             string          `json:"activity_id"`
	RankPageURL            string          `json:"rank_page_url"`
	ItemSearchType         string          `json:"item_search_type"`
	SPCampaignList         SPCampaignList  `json:"sp_campaign_list"`
}

type BybtInfo struct {
	BybtBrandLogo    string      `json:"bybt_brand_logo"`
	BybtPicURL       string      `json:"bybt_pic_url"`
	BybtItemTags     SmallImages `json:"bybt_item_tags"`
	BybtCouponAmount string      `json:"bybt_coupon_amount"`
	BybtShowPrice    string      `json:"bybt_show_price"`
	BybtLowestPrice  string      `json:"bybt_lowest_price"`
	BybtEndTime      string      `json:"bybt_end_time"`
	BybtStartTime    string      `json:"bybt_start_time"`
}

type SmallImages struct {
	String []string `json:"string"`
}

type MaifanPromotion struct {
	MaifanPromotionEndTime   string `json:"maifan_promotion_end_time"`
	MaifanPromotionStartTime string `json:"maifan_promotion_start_time"`
	MaifanPromotionDiscount  string `json:"maifan_promotion_discount"`
	MaifanPromotionCondition string `json:"maifan_promotion_condition"`
}

type SPCampaignList struct {
	SPCampaign []SPCampaign `json:"sp_campaign"`
}

type SPCampaign struct {
	SPCid        string `json:"sp_cid"`
	SPName       string `json:"sp_name"`
	SPRate       string `json:"sp_rate"`
	SPLockStatus string `json:"sp_lock_status"`
	SPApplyLink  string `json:"sp_apply_link"`
	SPStatus     string `json:"sp_status"`
}

type TopnInfo struct {
	TopnQuantity   int64  `json:"topn_quantity"`
	TopnTotalCount int64  `json:"topn_total_count"`
	TopnEndTime    string `json:"topn_end_time"`
	TopnStartTime  string `json:"topn_start_time"`
	TopnRate       string `json:"topn_rate"`
}
