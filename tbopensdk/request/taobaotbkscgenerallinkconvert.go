package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
	"strconv"
)

// taobao.tbk.sc.general.link.convert( 淘宝客-服务商-万能转链 )
// https://open.taobao.com/api.htm?docId=65412&docType=2&scopeId=28320
type TbkScGeneralLinkConvertRequest struct {
	Parameters *url.Values
}

// CheckParameters 必填校验：site_id、adzone_id
func (tk *TbkScGeneralLinkConvertRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
}

// AddParameter 添加请求参数（保持向后兼容）
func (tk *TbkScGeneralLinkConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *TbkScGeneralLinkConvertRequest) GetApiName() string {
	return "taobao.tbk.sc.general.link.convert"
}

// GetParameters 返回请求参数
func (tk *TbkScGeneralLinkConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}

// ==================== 顶层字段 Setter ====================

// SetBizSceneID 设置场景ID。1-动态ID转链场景，2-消费者比价场景，4-
func (tk *TbkScGeneralLinkConvertRequest) SetBizSceneID(val string) {
	tk.AddParameter("biz_scene_id", val)
}

// SetPromotionType 设置推广类型。1-自购省，2-推广赚（代理模式专属ID，代理模式必填）
func (tk *TbkScGeneralLinkConvertRequest) SetPromotionType(val string) {
	tk.AddParameter("promotion_type", val)
}

// SetMaterialList 设置物料列表。URL或淘口令，多个用英文逗号拼接
func (tk *TbkScGeneralLinkConvertRequest) SetMaterialList(val string) {
	tk.AddParameter("material_list", val)
}

// SetRelationID 设置渠道管理ID。选品推广场景必须入参，且bizSceneId需入参2
func (tk *TbkScGeneralLinkConvertRequest) SetRelationID(val string) {
	tk.AddParameter("relation_id", val)
}

// SetAdzoneID 设置推广位ID，mm_xx_xx_xx pid三段式中的第三段（必填）
func (tk *TbkScGeneralLinkConvertRequest) SetAdzoneID(val int64) {
	tk.AddParameter("adzone_id", strconv.FormatInt(val, 10))
}

// SetSellerIDList 设置卖家ID列表。多个用英文逗号拼接
func (tk *TbkScGeneralLinkConvertRequest) SetSellerIDList(val string) {
	tk.AddParameter("seller_id_list", val)
}

// SetSiteID 设置推广位ID，mm_xx_xx_xx pid三段式中的第二段（必填）
func (tk *TbkScGeneralLinkConvertRequest) SetSiteID(val int64) {
	tk.AddParameter("site_id", strconv.FormatInt(val, 10))
}

// SetItemIDList 设置商品ID列表。多个用英文逗号拼接
func (tk *TbkScGeneralLinkConvertRequest) SetItemIDList(val string) {
	tk.AddParameter("item_id_list", val)
}

// SetPageIDList 设置会场ID列表。多个用英文逗号拼接
func (tk *TbkScGeneralLinkConvertRequest) SetPageIDList(val string) {
	tk.AddParameter("page_id_list", val)
}

// SetSpecialID 设置会员运营ID
func (tk *TbkScGeneralLinkConvertRequest) SetSpecialID(val string) {
	tk.AddParameter("special_id", val)
}

// SetUvid 设置加密用户标识
func (tk *TbkScGeneralLinkConvertRequest) SetUvid(val string) {
	tk.AddParameter("uvid", val)
}

// SetQmtid 设置启明系统任务ID
func (tk *TbkScGeneralLinkConvertRequest) SetQmtid(val string) {
	tk.AddParameter("qmtid", val)
}

// SetUvidNotAuthorized 设置UV展示授权弹窗。1-不展示，0-展示，2-强提醒弹窗（返利场景）
func (tk *TbkScGeneralLinkConvertRequest) SetUvidNotAuthorized(val int64) {
	tk.AddParameter("uvid_not_authorized", strconv.FormatInt(val, 10))
}

// SetRequiredLinkType 设置指定出参link类型，如 cps_short_tpwd,coupon_short_url
func (tk *TbkScGeneralLinkConvertRequest) SetRequiredLinkType(val string) {
	tk.AddParameter("required_link_type", val)
}

// SetManagePubID 设置商品库账号ID
func (tk *TbkScGeneralLinkConvertRequest) SetManagePubID(val int64) {
	tk.AddParameter("manage_pub_id", strconv.FormatInt(val, 10))
}

// SetDx 设置链接类型。1-商品转通用计划链接，其他值-最优佣金率（含营销计划）链接
func (tk *TbkScGeneralLinkConvertRequest) SetDx(val string) {
	tk.AddParameter("dx", val)
}

// ==================== 嵌套 DTO Setter ====================

// SetTargetItem 设置会场页面内定坑商品
// itemIDList: 商品ID列表
func (tk *TbkScGeneralLinkConvertRequest) SetTargetItem(itemIDList string) {
	tk.AddParameter("target_item.item_id_list", itemIDList)
}

// SetItemDTO 设置商品转链信息
// itemID: 商品ID
// skuID: SKU ID（传入时会透传至转链结果url中）
// isTargetCoupon: 是否指定券。1-指定 0-不指定（KA自用型工具服务商）
// couponID: 优惠券ID（KA自用型工具服务商）
// externalID: 淘宝客外部用户标记，如自身系统账户ID
// dx: 链接类型。1-通用计划，其他-最优佣金率
// managePubID: 商品库账号ID
func (tk *TbkScGeneralLinkConvertRequest) SetItemDTO(itemID, skuID, isTargetCoupon, couponID, externalID, dx string, managePubID int64) {
	if itemID != "" {
		tk.AddParameter("item_d_t_o.item_id", itemID)
	}
	if skuID != "" {
		tk.AddParameter("item_d_t_o.sku_id", skuID)
	}
	if isTargetCoupon != "" {
		tk.AddParameter("item_d_t_o.is_target_coupon", isTargetCoupon)
	}
	if couponID != "" {
		tk.AddParameter("item_d_t_o.coupon_id", couponID)
	}
	if externalID != "" {
		tk.AddParameter("item_d_t_o.external_id", externalID)
	}
	if dx != "" {
		tk.AddParameter("item_d_t_o.dx", dx)
	}
	if managePubID != 0 {
		tk.AddParameter("item_d_t_o.manage_pub_id", strconv.FormatInt(managePubID, 10))
	}
}

// SetPageDTO 设置会场页面转链（支持多个）
// pageID: 会场ID
// targetItemList: 页面内定坑商品ID，用于素材-坑位还原
func (tk *TbkScGeneralLinkConvertRequest) SetPageDTO(pageID, targetItemList string) {
	tk.AddParameter("page_d_t_o.page_id", pageID)
	if targetItemList != "" {
		tk.AddParameter("page_d_t_o.target_item_list", targetItemList)
	}
}

// SetShopDTO 设置店铺转链
// shopID: 店铺ID
func (tk *TbkScGeneralLinkConvertRequest) SetShopDTO(shopID string) {
	tk.AddParameter("shop_d_t_o.shop_id", shopID)
}

// SetMaterialDTO 设置链接/口令转链（支持多个）
// materialURL: 物料链接（URL或淘口令）
// isTargetCoupon: 是否指定券。1-指定 0-不指定
// couponID: 优惠券ID
func (tk *TbkScGeneralLinkConvertRequest) SetMaterialDTO(materialURL, isTargetCoupon, couponID string) {
	if materialURL != "" {
		tk.AddParameter("material_d_t_o.material_url", materialURL)
	}
	if isTargetCoupon != "" {
		tk.AddParameter("material_d_t_o.is_target_coupon", isTargetCoupon)
	}
	if couponID != "" {
		tk.AddParameter("material_d_t_o.coupon_id", couponID)
	}
}
