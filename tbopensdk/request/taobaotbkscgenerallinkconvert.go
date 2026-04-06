package request

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mimicode/tksdk/utils"
)

//taobao.tbk.sc.general.link.convert( 淘宝客-服务商-万能转链 )
//https://open.taobao.com/api.htm?docId=65412&docType=2&scopeId=28320
type TbkScGeneralLinkConvertRequest struct {
	Parameters *url.Values //请求参数

	// 顶层字段（结构体方式构造请求参数）
	BizSceneID           string          // biz_scene_id
	PromotionType        string          // promotion_type
	MaterialList         string          // material_list
	RelationID           string          // relation_id
	AdzoneID             int64           // adzone_id (必填, Number)
	SellerIDList         string          // seller_id_list
	SiteID               int64           // site_id (必填, Number)
	ItemIDList           string          // item_id_list
	PageIDList           string          // page_id_list
	SpecialID            string          // special_id
	Uvid                 string          // uvid
	Qmtid                string          // qmtid
	UvidNotAuthorized    int64           // uvid_not_authorized (Number)
	RequiredLinkType     string          // required_link_type
	ManagePubID          int64           // manage_pub_id (Number)
	Dx                   string          // dx

	// 嵌套 DTO
	TargetItem *TargetItemDTO   // target_item (商品维度转链)
	ItemDTO    *LkItemDTO       // item_d_t_o (商品信息)
	PageDTO    []*LkPageDTO     // page_d_t_o (页面维度转链)
	ShopDTO    []*LkShopDTO     // shop_d_t_o (店铺维度转链)
	MaterialDTO []*LkMaterialDTO // material_d_t_o (物料维度转链)
}

// TargetItemDTO 商品维度转链
type TargetItemDTO struct {
	ItemIDList string `json:"item_id_list"` // item_id_list
}

// LkItemDTO 商品信息
type LkItemDTO struct {
	ItemID          string `json:"item_id"`           // item_id
	SkuID           int64  `json:"sku_id"`            // sku_id (Number)
	IsTargetCoupon  int64  `json:"is_target_coupon"`  // is_target_coupon (Number)
	CouponID        string `json:"coupon_id"`         // coupon_id
	ExternalID      string `json:"external_id"`       // external_id
	Dx              string `json:"dx"`                // dx
	ManagePubID     int64  `json:"manage_pub_id"`     // manage_pub_id (Number)
}

// LkPageDTO 页面维度转链
type LkPageDTO struct {
	PageID         string `json:"page_id"`          // page_id
	TargetItemList string `json:"target_item_list"` // target_item_list
}

// LkShopDTO 店铺维度转链
type LkShopDTO struct {
	ShopID string `json:"shop_id"` // shop_id
}

// LkMaterialDTO 物料维度转链
type LkMaterialDTO struct {
	MaterialURL    string `json:"material_url"`     // material_url
	IsTargetCoupon int64  `json:"is_target_coupon"` // is_target_coupon (Number)
	CouponID       string `json:"coupon_id"`        // coupon_id
}

func (tk *TbkScGeneralLinkConvertRequest) CheckParameters() {
	// 必填校验：site_id、adzone_id
	if tk.SiteID == 0 {
		utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	}
	if tk.AdzoneID == 0 {
		utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	}
}

// 添加请求参数（保持向后兼容）
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

// GetParameters 返回请求参数（优先使用结构体字段，否则回退到 Parameters map）
func (tk *TbkScGeneralLinkConvertRequest) GetParameters() url.Values {
	// 如果没有通过结构体字段设置任何值，则回退到原始透传模式
	if !tk.hasStructFields() {
		return *tk.Parameters
	}
	return tk.StructToURLValues()
}

// hasStructFields 检测是否通过结构体字段设置了值
func (tk *TbkScGeneralLinkConvertRequest) hasStructFields() bool {
	return tk.SiteID != 0 || tk.AdzoneID != 0 ||
		tk.BizSceneID != "" || tk.PromotionType != "" ||
		tk.MaterialList != "" || tk.RelationID != "" ||
		tk.SellerIDList != "" || tk.ItemIDList != "" ||
		tk.PageIDList != "" || tk.SpecialID != "" ||
		tk.Uvid != "" || tk.Qmtid != "" ||
		tk.UvidNotAuthorized != 0 || tk.RequiredLinkType != "" ||
		tk.ManagePubID != 0 || tk.Dx != "" ||
		tk.TargetItem != nil || tk.ItemDTO != nil ||
		len(tk.PageDTO) > 0 || len(tk.ShopDTO) > 0 ||
		len(tk.MaterialDTO) > 0
}

// StructToURLValues 将结构体字段批量转为 url.Values
func (tk *TbkScGeneralLinkConvertRequest) StructToURLValues() url.Values {
	params := url.Values{}

	// 顶层字段
	if tk.BizSceneID != "" {
		params.Set("biz_scene_id", tk.BizSceneID)
	}
	if tk.PromotionType != "" {
		params.Set("promotion_type", tk.PromotionType)
	}
	if tk.MaterialList != "" {
		params.Set("material_list", tk.MaterialList)
	}
	if tk.RelationID != "" {
		params.Set("relation_id", tk.RelationID)
	}
	if tk.SiteID != 0 {
		params.Set("site_id", strconv.FormatInt(tk.SiteID, 10))
	}
	if tk.AdzoneID != 0 {
		params.Set("adzone_id", strconv.FormatInt(tk.AdzoneID, 10))
	}
	if tk.SellerIDList != "" {
		params.Set("seller_id_list", tk.SellerIDList)
	}
	if tk.ItemIDList != "" {
		params.Set("item_id_list", tk.ItemIDList)
	}
	if tk.PageIDList != "" {
		params.Set("page_id_list", tk.PageIDList)
	}
	if tk.SpecialID != "" {
		params.Set("special_id", tk.SpecialID)
	}
	if tk.Uvid != "" {
		params.Set("uvid", tk.Uvid)
	}
	if tk.Qmtid != "" {
		params.Set("qmtid", tk.Qmtid)
	}
	if tk.UvidNotAuthorized != 0 {
		params.Set("uvid_not_authorized", strconv.FormatInt(tk.UvidNotAuthorized, 10))
	}
	if tk.RequiredLinkType != "" {
		params.Set("required_link_type", tk.RequiredLinkType)
	}
	if tk.ManagePubID != 0 {
		params.Set("manage_pub_id", strconv.FormatInt(tk.ManagePubID, 10))
	}
	if tk.Dx != "" {
		params.Set("dx", tk.Dx)
	}

	// 嵌套 DTO: TargetItemDTO
	if tk.TargetItem != nil {
		if tk.TargetItem.ItemIDList != "" {
			params.Set("target_item.item_id_list", tk.TargetItem.ItemIDList)
		}
	}

	// 嵌套 DTO: LkItemDTO
	if tk.ItemDTO != nil {
		if tk.ItemDTO.ItemID != "" {
			params.Set("item_d_t_o.item_id", tk.ItemDTO.ItemID)
		}
		if tk.ItemDTO.SkuID != 0 {
			params.Set("item_d_t_o.sku_id", strconv.FormatInt(tk.ItemDTO.SkuID, 10))
		}
		if tk.ItemDTO.IsTargetCoupon != 0 {
			params.Set("item_d_t_o.is_target_coupon", strconv.FormatInt(tk.ItemDTO.IsTargetCoupon, 10))
		}
		if tk.ItemDTO.CouponID != "" {
			params.Set("item_d_t_o.coupon_id", tk.ItemDTO.CouponID)
		}
		if tk.ItemDTO.ExternalID != "" {
			params.Set("item_d_t_o.external_id", tk.ItemDTO.ExternalID)
		}
		if tk.ItemDTO.Dx != "" {
			params.Set("item_d_t_o.dx", tk.ItemDTO.Dx)
		}
		if tk.ItemDTO.ManagePubID != 0 {
			params.Set("item_d_t_o.manage_pub_id", strconv.FormatInt(tk.ItemDTO.ManagePubID, 10))
		}
	}

	// 嵌套 DTO: LkPageDTO
	if len(tk.PageDTO) > 0 {
		for i, page := range tk.PageDTO {
			prefix := fmt.Sprintf("page_d_t_o[%d].", i)
			if page.PageID != "" {
				params.Set(prefix+"page_id", page.PageID)
			}
			if page.TargetItemList != "" {
				params.Set(prefix+"target_item_list", page.TargetItemList)
			}
		}
	}

	// 嵌套 DTO: LkShopDTO
	if len(tk.ShopDTO) > 0 {
		for i, shop := range tk.ShopDTO {
			prefix := fmt.Sprintf("shop_d_t_o[%d].", i)
			if shop.ShopID != "" {
				params.Set(prefix+"shop_id", shop.ShopID)
			}
		}
	}

	// 嵌套 DTO: LkMaterialDTO
	if len(tk.MaterialDTO) > 0 {
		for i, material := range tk.MaterialDTO {
			prefix := fmt.Sprintf("material_d_t_o[%d].", i)
			if material.MaterialURL != "" {
				params.Set(prefix+"material_url", material.MaterialURL)
			}
			if material.IsTargetCoupon != 0 {
				params.Set(prefix+"is_target_coupon", strconv.FormatInt(material.IsTargetCoupon, 10))
			}
			if material.CouponID != "" {
				params.Set(prefix+"coupon_id", material.CouponID)
			}
		}
	}

	// 合并通过 AddParameter 添加的额外参数
	if tk.Parameters != nil {
		for k, v := range *tk.Parameters {
			params[k] = v
		}
	}

	return params
}
