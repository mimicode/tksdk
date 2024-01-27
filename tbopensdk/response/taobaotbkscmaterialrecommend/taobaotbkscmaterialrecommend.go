package taobaotbkscmaterialrecommend

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.material.recommend( 淘宝客-服务商-物料精选升级版 )
type Response struct {
	response.TopResponse
	TbkScMaterialRecommendResponse TbkScMaterialRecommendResponse `json:"tbk_sc_material_recommend_response"`
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

type TbkScMaterialRecommendResponse struct {
	ResultList ResultList `json:"result_list"`
	IsDefault  string     `json:"is_default"`
	TotalCount int64      `json:"total_count"`
	UvidMsg    string     `json:"uvid_msg"`
}

type ResultList struct {
	MapData []MapDatum `json:"map_data"`
}

type MapDatum struct {
	ItemID             string             `json:"item_id"`
	ItemBasicInfo      ItemBasicInfo      `json:"item_basic_info"`
	PricePromotionInfo PricePromotionInfo `json:"price_promotion_info"`
	PublishInfo        PublishInfo        `json:"publish_info"`
	TmallRankInfo      TmallRankInfo      `json:"tmall_rank_info"`
	PresaleInfo        PresaleInfo        `json:"presale_info"`
	FavoritesInfo      FavoritesInfo      `json:"favorites_info"`
	TopnInfo           TopnInfo           `json:"topn_info"`
	ScopeInfo          ScopeInfo          `json:"scope_info"`
	ISVMktid           string             `json:"isv_mktid"`
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

type ItemBasicInfo struct {
	Title                string `json:"title"`
	ShortTitle           string `json:"short_title"`
	PictURL              string `json:"pict_url"`
	WhiteImage           string `json:"white_image"`
	LevelOneCategoryID   int64  `json:"level_one_category_id"`
	LevelOneCategoryName string `json:"level_one_category_name"`
	CategoryID           int64  `json:"category_id"`
	CategoryName         string `json:"category_name"`
	SellerID             int64  `json:"seller_id"`
	UserType             int64  `json:"user_type"`
	ShopTitle            string `json:"shop_title"`
	Volume               int64  `json:"volume"`
	SubTitle             string `json:"sub_title"`
	BrandName            string `json:"brand_name"`
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
	ReservePrice                    string                          `json:"reserve_price"`
	ZkFinalPrice                    string                          `json:"zk_final_price"`
	FinalPromotionPrice             string                          `json:"final_promotion_price"`
	FinalPromotionPathList          FinalPromotionPathList          `json:"final_promotion_path_list"`
	FutureActivityPromotionPrice    string                          `json:"future_activity_promotion_price"`
	FutureActivityPromotionPathList FutureActivityPromotionPathList `json:"future_activity_promotion_path_list"`
	PromotionTagList                PromotionTagList                `json:"promotion_tag_list"`
	PredictRoundingUpPrice          string                          `json:"predict_rounding_up_price"`
	PredictRoundingUpPriceDesc      string                          `json:"predict_rounding_up_price_desc"`
	MorePromotionList               MorePromotionList               `json:"more_promotion_list"`
	PredictRoundingUpPathList       PredictRoundingUpPathList       `json:"predict_rounding_up_path_list"`
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

type PredictRoundingUpPathList struct {
	PredictRoundingUpPathMapData []PredictRoundingUpPathMapDatum `json:"predict_rounding_up_path_map_data"`
}

type PredictRoundingUpPathMapDatum struct {
	PromotionTitle string `json:"promotion_title"`
	PromotionDesc  string `json:"promotion_desc"`
}

type PromotionTagList struct {
	PromotionTagMapData []PromotionTagMapDatum `json:"promotion_tag_map_data"`
}

type PromotionTagMapDatum struct {
	TagName string `json:"tag_name"`
}

type PublishInfo struct {
	IncomeRate                   string         `json:"income_rate"`
	ClickURL                     string         `json:"click_url"`
	CouponShareURL               string         `json:"coupon_share_url"`
	SPCampaignList               SPCampaignList `json:"sp_campaign_list"`
	FutureActivityCommissionRate string         `json:"future_activity_commission_rate"`
	FutureActivityTime           string         `json:"future_activity_time"`
	IncomeInfo                   IncomeInfo     `json:"income_info"`
	CPARewardType                string         `json:"cpa_reward_type"`
	CPARewardAmount              string         `json:"cpa_reward_amount"`
}

type IncomeInfo struct {
	CommissionRate    string `json:"commission_rate"`
	CommissionAmount  string `json:"commission_amount"`
	SubsidyRate       string `json:"subsidy_rate"`
	SubsidyAmount     string `json:"subsidy_amount"`
	SubsidyUpperLimit string `json:"subsidy_upper_limit"`
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

type ScopeInfo struct {
	SuperiorBrand string `json:"superior_brand"`
}

type TmallRankInfo struct {
	TmallRankText string `json:"tmall_rank_text"`
	TmallRankURL  string `json:"tmall_rank_url"`
}

type TopnInfo struct {
	TopnQuantity   int64  `json:"topn_quantity"`
	TopnTotalCount int64  `json:"topn_total_count"`
	TopnStartTime  string `json:"topn_start_time"`
	TopnEndTime    string `json:"topn_end_time"`
	TopnRate       string `json:"topn_rate"`
}
