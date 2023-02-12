package tbkdgoptimusmaterial

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.dg.optimus.material( 淘宝客物料下行-导购 )
type Response struct {
	response.TopResponse
	TbkDgOptimusMaterialResult Result `json:"tbk_dg_optimus_material_response"`
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
	ResultList ResultList `json:"result_list"`
	RequestID  string     `json:"request_id"`
	IsDefault  string     `json:"is_default"`
	TotalCount int64      `json:"total_count"`
}
type ResultList struct {
	MapData []MapDatum `json:"map_data"`
}

type MapDatum struct {
	CouponAmount               int64           `json:"coupon_amount"`
	SmallImages                SmallImages     `json:"small_images"`
	ShopTitle                  string          `json:"shop_title"`
	CategoryID                 int64           `json:"category_id"`
	CouponStartFee             string          `json:"coupon_start_fee"`
	ItemID                     string          `json:"item_id"`
	CouponTotalCount           int64           `json:"coupon_total_count"`
	UserType                   int64           `json:"user_type"`
	ZkFinalPrice               string          `json:"zk_final_price"`
	CouponRemainCount          int64           `json:"coupon_remain_count"`
	CommissionRate             string          `json:"commission_rate"`
	CouponStartTime            string          `json:"coupon_start_time"`
	Title                      string          `json:"title"`
	ItemDescription            string          `json:"item_description"`
	SellerID                   int64           `json:"seller_id"`
	Volume                     int64           `json:"volume"`
	CouponEndTime              string          `json:"coupon_end_time"`
	CouponClickURL             string          `json:"coupon_click_url"`
	PictURL                    string          `json:"pict_url"`
	ClickURL                   string          `json:"click_url"`
	Stock                      int64           `json:"stock"`
	SellNum                    int64           `json:"sell_num"`
	TotalStock                 int64           `json:"total_stock"`
	Oetime                     string          `json:"oetime"`
	Ostime                     string          `json:"ostime"`
	JddNum                     int64           `json:"jdd_num"`
	JddPrice                   string          `json:"jdd_price"`
	OrigPrice                  string          `json:"orig_price"`
	LevelOneCategoryName       string          `json:"level_one_category_name"`
	LevelOneCategoryID         int64           `json:"level_one_category_id"`
	CategoryName               string          `json:"category_name"`
	WhiteImage                 string          `json:"white_image"`
	ShortTitle                 string          `json:"short_title"`
	WordList                   WordList        `json:"word_list"`
	TmallPlayActivityInfo      string          `json:"tmall_play_activity_info"`
	UvSumPreSale               int64           `json:"uv_sum_pre_sale"`
	XID                        string          `json:"x_id"`
	NewUserPrice               string          `json:"new_user_price"`
	CouponInfo                 string          `json:"coupon_info"`
	CouponShareURL             string          `json:"coupon_share_url"`
	Nick                       string          `json:"nick"`
	ReservePrice               string          `json:"reserve_price"`
	JuOnlineEndTime            string          `json:"ju_online_end_time"`
	JuOnlineStartTime          string          `json:"ju_online_start_time"`
	MaochaoPlayEndTime         string          `json:"maochao_play_end_time"`
	MaochaoPlayStartTime       string          `json:"maochao_play_start_time"`
	MaochaoPlayConditions      string          `json:"maochao_play_conditions"`
	MaochaoPlayDiscounts       string          `json:"maochao_play_discounts"`
	MaochaoPlayDiscountType    string          `json:"maochao_play_discount_type"`
	MaochaoPlayFreePostFee     string          `json:"maochao_play_free_post_fee"`
	MultiCouponZkRate          string          `json:"multi_coupon_zk_rate"`
	PriceAfterMultiCoupon      string          `json:"price_after_multi_coupon"`
	MultiCouponItemCount       string          `json:"multi_coupon_item_count"`
	LockRate                   string          `json:"lock_rate"`
	LockRateEndTime            int64           `json:"lock_rate_end_time"`
	LockRateStartTime          int64           `json:"lock_rate_start_time"`
	PromotionType              string          `json:"promotion_type"`
	PromotionInfo              string          `json:"promotion_info"`
	PromotionDiscount          string          `json:"promotion_discount"`
	PromotionCondition         string          `json:"promotion_condition"`
	PresaleDiscountFeeText     string          `json:"presale_discount_fee_text"`
	PresaleTailEndTime         int64           `json:"presale_tail_end_time"`
	PresaleTailStartTime       int64           `json:"presale_tail_start_time"`
	PresaleEndTime             int64           `json:"presale_end_time"`
	PresaleStartTime           int64           `json:"presale_start_time"`
	PresaleDeposit             string          `json:"presale_deposit"`
	YsylTljUseStartTime        string          `json:"ysyl_tlj_use_start_time"`
	YsylCommissionRate         string          `json:"ysyl_commission_rate"`
	YsylTljSendTime            string          `json:"ysyl_tlj_send_time"`
	YsylTljFace                string          `json:"ysyl_tlj_face"`
	YsylClickURL               string          `json:"ysyl_click_url"`
	YsylTljUseEndTime          string          `json:"ysyl_tlj_use_end_time"`
	JuPlayEndTime              int64           `json:"ju_play_end_time"`
	JuPlayStartTime            int64           `json:"ju_play_start_time"`
	PlayInfo                   string          `json:"play_info"`
	TmallPlayActivityEndTime   int64           `json:"tmall_play_activity_end_time"`
	TmallPlayActivityStartTime int64           `json:"tmall_play_activity_start_time"`
	JuPreShowEndTime           string          `json:"ju_pre_show_end_time"`
	JuPreShowStartTime         string          `json:"ju_pre_show_start_time"`
	FavoritesInfo              FavoritesInfo   `json:"favorites_info"`
	SalePrice                  string          `json:"sale_price"`
	KuadianPromotionInfo       string          `json:"kuadian_promotion_info"`
	SubTitle                   string          `json:"sub_title"`
	JhsPriceUspList            string          `json:"jhs_price_usp_list"`
	TqgOnlineEndTime           string          `json:"tqg_online_end_time"`
	TqgOnlineStartTime         string          `json:"tqg_online_start_time"`
	TqgSoldCount               int64           `json:"tqg_sold_count"`
	TqgTotalCount              int64           `json:"tqg_total_count"`
	SuperiorBrand              string          `json:"superior_brand"`
	IsBrandFlashSale           string          `json:"is_brand_flash_sale"`
	HotFlag                    string          `json:"hot_flag"`
	TopnInfo                   TopnInfo        `json:"topn_info"`
	BybtInfo                   BybtInfo        `json:"bybt_info"`
	TtSoldCount                string          `json:"tt_sold_count"`
	MaifanPromotion            MaifanPromotion `json:"maifan_promotion"`
	CPARewardType              string          `json:"cpa_reward_type"`
	CPARewardAmount            string          `json:"cpa_reward_amount"`
	ActivityID                 string          `json:"activity_id"`
	SPCampaignList             SPCampaignList  `json:"sp_campaign_list"`
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

type FavoritesInfo struct {
	TotalCount    int64         `json:"total_count"`
	FavoritesList FavoritesList `json:"favorites_list"`
}

type FavoritesList struct {
	FavoritesDetail []FavoritesDetail `json:"favorites_detail"`
}

type FavoritesDetail struct {
	FavoritesID    int64  `json:"favorites_id"`
	FavoritesTitle string `json:"favorites_title"`
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

type WordList struct {
	WordMapData []WordMapDatum `json:"word_map_data"`
}

type WordMapDatum struct {
	URL  string `json:"url"`
	Word string `json:"word"`
}
