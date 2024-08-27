package taobaotbkiteminfoupgradeget

import (
	"encoding/json"

	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.item.info.upgrade.get( 淘宝客-公用-淘宝客商品详情查询升级版（简易版） )
type Response struct {
	response.TopResponse
	TbkItemInfoUpgradeGetResponse TbkItemInfoUpgradeGetResponse `json:"tbk_item_info_upgrade_get_response"`
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

type TbkItemInfoUpgradeGetResponse struct {
	Results Results `json:"results"`
}

type Results struct {
	TbkItemDetail []TbkItemDetail `json:"tbk_item_detail"`
}

type TbkItemDetail struct {
	ItemID             string             `json:"item_id"`
	PublishInfo        PublishInfo        `json:"publish_info"`
	PricePromotionInfo PricePromotionInfo `json:"price_promotion_info"`
	InputItemIid       string             `json:"input_item_iid"`
	ItemBasicInfo      ItemBasicInfo      `json:"item_basic_info"`
	PresaleInfo        PresaleInfo        `json:"presale_info"`
}

type ItemBasicInfo struct {
	Title                string `json:"title"`
	PictURL              string `json:"pict_url"`
	LevelOneCategoryName string `json:"level_one_category_name"`
	CategoryName         string `json:"category_name"`
	Provcity             string `json:"provcity"`
	ItemURL              string `json:"item_url"`
	UserType             int64  `json:"user_type"`
	Ratesum              int64  `json:"ratesum"`
	ShopDsr              int64  `json:"shop_dsr"`
	IRfdRate             bool   `json:"i_rfd_rate"`
	HGoodRate            bool   `json:"h_good_rate"`
	HPayRate30           bool   `json:"h_pay_rate30"`
	Volume               int64  `json:"volume"`
	IsPrepay             bool   `json:"is_prepay"`
	SuperiorBrand        string `json:"superior_brand"`
	MaterialLIBType      string `json:"material_lib_type"`
	TmallDescURL         string `json:"tmall_desc_url"`
	TaobaoDescURL        string `json:"taobao_desc_url"`
	ShopTitle            string `json:"shop_title"`
	FreeShipment         bool   `json:"free_shipment"`
	BrandName            string `json:"brand_name"`
	ShortTitle           string `json:"short_title"`
	AnnualVol            string `json:"annual_vol"`
	WhiteImage           string `json:"white_image"`
	SellerID             int64  `json:"seller_id"`
	SmallImages          struct {
		String []string `json:"string"`
	} `json:"small_images"`
}

type PresaleInfo struct {
	PresaleTailEndTime     int64  `json:"presale_tail_end_time"`
	PresaleTailStartTime   int64  `json:"presale_tail_start_time"`
	PresaleEndTime         int64  `json:"presale_end_time"`
	PresaleStartTime       int64  `json:"presale_start_time"`
	PresaleDeposit         string `json:"presale_deposit"`
	PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
}

type PricePromotionInfo struct {
	ReservePrice               string                    `json:"reserve_price"`
	ZkFinalPrice               string                    `json:"zk_final_price"`
	FinalPromotionPrice        string                    `json:"final_promotion_price"`
	FinalPromotionPathList     FinalPromotionPathList    `json:"final_promotion_path_list"`
	PromotionTagList           PromotionTagList          `json:"promotion_tag_list"`
	PredictRoundingUpPrice     string                    `json:"predict_rounding_up_price"`
	PredictRoundingUpPriceDesc string                    `json:"predict_rounding_up_price_desc"`
	PredictRoundingUpPathList  PredictRoundingUpPathList `json:"predict_rounding_up_path_list"`
	MorePromotionList          MorePromotionList         `json:"more_promotion_list"`
	ActivityTagList            ActivityTagList           `json:"activity_tag_list"`
}

type PredictRoundingUpPathList struct {
	PredictRoundingUpPathMapData []PredictRoundingUpPathMapDatum `json:"predict_rounding_up_path_map_data"`
}

type PredictRoundingUpPathMapDatum struct {
	PromotionDesc  string `json:"promotion_desc"`
	PromotionTitle string `json:"promotion_title"`
}

type ActivityTagList struct {
	ActivityTagMapData []TagMapDatum `json:"activity_tag_map_data"`
}

type TagMapDatum struct {
	TagName string `json:"tag_name"`
}

type FinalPromotionPathList struct {
	FinalPromotionPathMapData []MapDatum `json:"final_promotion_path_map_data"`
}

type MapDatum struct {
	PromotionTitle     string  `json:"promotion_title"`
	PromotionDesc      string  `json:"promotion_desc"`
	PromotionFee       *string `json:"promotion_fee,omitempty"`
	PromotionStartTime string  `json:"promotion_start_time"`
	PromotionEndTime   string  `json:"promotion_end_time"`
	PromotionID        *string `json:"promotion_id,omitempty"`
}

type MorePromotionList struct {
	MorePromotionMapData []MapDatum `json:"more_promotion_map_data"`
}

type PromotionTagList struct {
	PromotionTagMapData []TagMapDatum `json:"promotion_tag_map_data"`
}

type PublishInfo struct {
	IncomeRate                   string         `json:"income_rate"`
	TopnInfo                     TopnInfo       `json:"topn_info"`
	ClickURL                     string         `json:"click_url"`
	CouponShareURL               string         `json:"coupon_share_url"`
	FutureActivityCommissionRate string         `json:"future_activity_commission_rate"`
	FutureActivityTime           string         `json:"future_activity_time"`
	SPCampaignList               SPCampaignList `json:"sp_campaign_list"`
	RankPageURL                  string         `json:"rank_page_url"`
	CommissionType               string         `json:"commission_type"`
	CPARewardType                string         `json:"cpa_reward_type"`
	CPARewardAmount              string         `json:"cpa_reward_amount"`
	IncomeInfo                   IncomeInfo     `json:"income_info"`
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

type IncomeInfo struct {
	CommissionAmount string `json:"commission_amount"`
	CommissionRate   string `json:"commission_rate"`
	SubsidyAmount    string `json:"subsidy_amount"`
	SubsidyRate      string `json:"subsidy_rate"`
}

type TopnInfo struct {
	TopnQuantity   int64  `json:"topn_quantity"`
	TopnTotalCount int64  `json:"topn_total_count"`
	TopnEndTime    string `json:"topn_end_time"`
	TopnStartTime  string `json:"topn_start_time"`
	TopnRate       string `json:"topn_rate"`
}
