package taobaotbkscmaterialoptionalupgrade

import (
	"encoding/json"

	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.material.optional.upgrade( 淘宝客-服务商-物料搜索升级版 )
type Response struct {
	response.TopResponse
	TbkScMaterialOptionalUpgradeResponse TbkScMaterialOptionalUpgradeResponse `json:"tbk_sc_material_optional_upgrade_response"`
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

type TbkScMaterialOptionalUpgradeResponse struct {
	TotalResults int64      `json:"total_results"`
	ResultList   ResultList `json:"result_list"`
	UvidMsg      string     `json:"uvid_msg"`
}

type ResultList struct {
	MapData []MapDatum `json:"map_data"`
}

type MapDatum struct {
	ItemID             string             `json:"item_id"`
	PricePromotionInfo PricePromotionInfo `json:"price_promotion_info"`
	PublishInfo        PublishInfo        `json:"publish_info"`
	ItemBasicInfo      ItemBasicInfo      `json:"item_basic_info"`
	TmallRankInfo      TmallRankInfo      `json:"tmall_rank_info"`
	PresaleInfo        PresaleInfo        `json:"presale_info"`
	ScopeInfo          ScopeInfo          `json:"scope_info"`
	MgcInfo            MgcInfo            `json:"mgc_info"`
	MatchScore         string             `json:"match_score"`
	CommiScore         string             `json:"commi_score"`
	ISVMktid           string             `json:"isv_mktid"`
}

type SmallImages struct {
	String []string `json:"string"`
}

type ItemBasicInfo struct {
	Title                string      `json:"title"`
	ShortTitle           string      `json:"short_title"`
	PictURL              string      `json:"pict_url"`
	WhiteImage           string      `json:"white_image"`
	LevelOneCategoryID   int64       `json:"level_one_category_id"`
	CategoryID           int64       `json:"category_id"`
	CategoryName         string      `json:"category_name"`
	SellerID             int64       `json:"seller_id"`
	UserType             int64       `json:"user_type"`
	ShopTitle            string      `json:"shop_title"`
	Volume               int64       `json:"volume"`
	SubTitle             string      `json:"sub_title"`
	BrandName            string      `json:"brand_name"`
	LevelOneCategoryName string      `json:"level_one_category_name"`
	TkTotalSales         string      `json:"tk_total_sales"`
	Provcity             string      `json:"provcity"`
	AnnualVol            string      `json:"annual_vol"`
	SmallImages          SmallImages `json:"small_images"`
}

type MgcInfo struct {
	Price            string `json:"price"`
	PriceDesc        string `json:"price_desc"`
	PromotionSummary string `json:"promotion_summary"`
	PublishTime      string `json:"publish_time"`
	ValidTime        string `json:"valid_time"`
}

type PresaleInfo struct {
	PresaleStartTime       int64  `json:"presale_start_time"`
	PresaleEndTime         int64  `json:"presale_end_time"`
	PresaleTailStartTime   int64  `json:"presale_tail_start_time"`
	PresaleTailEndTime     int64  `json:"presale_tail_end_time"`
	PresaleDeposit         string `json:"presale_deposit"`
	PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
}

type PricePromotionInfo struct {
	FinalPromotionPathList          FinalPromotionPathList          `json:"final_promotion_path_list"`
	PredictRoundingUpPrice          string                          `json:"predict_rounding_up_price"`
	PredictRoundingUpPriceDesc      string                          `json:"predict_rounding_up_price_desc"`
	PredictRoundingUpPathList       PredictRoundingUpPathList       `json:"predict_rounding_up_path_list"`
	MorePromotionList               MorePromotionList               `json:"more_promotion_list"`
	ReservePrice                    string                          `json:"reserve_price"`
	ZkFinalPrice                    string                          `json:"zk_final_price"`
	FinalPromotionPrice             string                          `json:"final_promotion_price"`
	FutureActivityPromotionPrice    string                          `json:"future_activity_promotion_price"`
	FutureActivityPromotionPathList FutureActivityPromotionPathList `json:"future_activity_promotion_path_list"`
	PromotionTagList                PromotionTagList                `json:"promotion_tag_list"`
}

type PredictRoundingUpPathList struct {
	PredictRoundingUpPathMapData []PredictRoundingUpPathMapDatum `json:"predict_rounding_up_path_map_data"`
}

type PredictRoundingUpPathMapDatum struct {
	PromotionDesc  string `json:"promotion_desc"`
	PromotionTitle string `json:"promotion_title"`
}

type FinalPromotionPathList struct {
	FinalPromotionPathMapData []FinalPromotionPathMapDatumElement `json:"final_promotion_path_map_data"`
}

type FinalPromotionPathMapDatumElement struct {
	PromotionTitle     string  `json:"promotion_title"`
	PromotionDesc      string  `json:"promotion_desc"`
	PromotionFee       *string `json:"promotion_fee,omitempty"`
	PromotionStartTime string  `json:"promotion_start_time"`
	PromotionEndTime   string  `json:"promotion_end_time"`
	PromotionID        *string `json:"promotion_id,omitempty"`
}

type FutureActivityPromotionPathList struct {
	FutureActivityPromotionPathMapData []FinalPromotionPathMapDatumElement `json:"future_activity_promotion_path_map_data"`
}

type MorePromotionList struct {
	MorePromotionMapData []FinalPromotionPathMapDatumElement `json:"more_promotion_map_data"`
}

type PromotionTagList struct {
	PromotionTagMapData []PromotionTagMapDatum `json:"promotion_tag_map_data"`
}

type PromotionTagMapDatum struct {
	TagName string `json:"tag_name"`
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

type IncomeInfo struct {
	CommissionAmount string `json:"commission_amount"`
	CommissionRate   string `json:"commission_rate"`
	SubsidyAmount    string `json:"subsidy_amount"`
	SubsidyRate      string `json:"subsidy_rate"`
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

type ScopeInfo struct {
	SuperiorBrand int64 `json:"superior_brand"`
}

type TmallRankInfo struct {
	TmallRankText string `json:"tmall_rank_text"`
	TmallRankURL  string `json:"tmall_rank_url"`
}
