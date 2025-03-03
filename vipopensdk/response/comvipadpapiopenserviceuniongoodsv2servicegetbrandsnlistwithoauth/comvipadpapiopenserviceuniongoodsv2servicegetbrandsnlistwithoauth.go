package comvipadpapiopenserviceuniongoodsv2servicegetbrandsnlistwithoauth

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsV2Service getBrandSnListWithOauth 精选品牌列表查询-需要oauth授权
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
	Code int    `json:"code"` // 状态码，1-成功，0-失败
	Msg  string `json:"msg"`  // 状态信息
	Data Data   `json:"data"` // 数据
}

type Data struct {
	Total         int64       `json:"total"`         // 总记录数
	BrandInfoList []BrandInfo `json:"brandInfoList"` // 品牌信息列表
	PageSize      int64       `json:"pageSize"`      // 每页条数
	Page          int64       `json:"page"`          // 页码
}

type BrandInfo struct {
	BrandName    string `json:"brandName"`    // 品牌名称
	BrandStoreSn string `json:"brandStoreSn"` // 品牌店铺SN
	BrandLogoUrl string `json:"brandLogoUrl"` // 品牌logo图片
	BrandDesc    string `json:"brandDesc"`    // 品牌描述
}
