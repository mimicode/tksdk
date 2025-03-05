package comvipadpapiopenserviceuniongoodsv2servicegetbranddetailwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// Response com.vip.adp.api.open.service.UnionGoodsV2Service
// getBrandDetailWithOauth 2.0.0 通过品牌SN查询品牌详情——需要oauth授权
// http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getBrandDetailWithOauth
type Response struct {
	response.TopResponse
	Success Success `json:"success"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

type Success struct {
	BrandDetail     BrandDetailModel   `json:"brandDetail"`     // 品牌详情信息，为空表示品牌不存在或该品牌不支持查询
	BrandDetailList []BrandDetailModel `json:"brandDetailList"` // 品牌详情信息，为空表示品牌不存在或该品牌不支持查询
}

// BrandDetailModel 品牌详情模型
type BrandDetailModel struct {
	BrandSn         string `json:"brandSn"`         // 品牌Sn
	CnName          string `json:"cnName"`          // 品牌中文名称（与英文名称至少有一个）
	EnName          string `json:"enName"`          // 英文名称（与中文名称至少有一个）
	ShowName        string `json:"showName"`        // 前端展示名称
	Logo            string `json:"logo"`            // 品牌LOGO图地址
	Slogon          string `json:"slogon"`          // 品牌描述
	AtmosphereUrl   string `json:"atmosphereUrl"`   // 品牌氛围图
	Content         string `json:"content"`         // 品牌故事内容
	BrandStoryImage string `json:"brandStoryImage"` // 品牌故事图片地址
}
