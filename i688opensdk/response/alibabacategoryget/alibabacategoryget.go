package alibabacategoryget

import (
	"encoding/json"
	"fmt"

	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response 根据类目Id查询类目响应
type Response struct {
	response.TopResponse
	// 直接解析字段，用于处理实际API返回的结构
	Success      string         `json:"succes"`       // 是否成功
	CategoryInfo []CategoryInfo `json:"categoryInfo"` // 类目列表
	ErrorMsg     string         `json:"errorMsg"`     // 错误信息
	ErrorCode    string         `json:"errorCode"`    // 错误码
}

// CategoryInfo 类目信息
type CategoryInfo struct {
	CategoryID          int64               `json:"categoryID"`          // 类目ID
	Name                string              `json:"name"`                // 类目名称
	Level               int                 `json:"level"`               // 类目层级，1688无此内容
	IsLeaf              bool                `json:"isLeaf"`              // 是否叶子类目（只有叶子类目才能发布商品）
	ParentIDs           []int64             `json:"parentIDs"`           // 父类目ID数组,1688只返回一个父id
	ChildIDs            []int64             `json:"childIDs"`            // 子类目ID数组，1688无此内容
	ChildCategorys      []ChildCategoryInfo `json:"childCategorys"`      // 子类目信息
	MinOrderQuantity    int64               `json:"minOrderQuantity"`    // 最小起订量
	FeatureInfos        []FeatureInfo       `json:"featureInfos"`        // 类目特征信息
	CategoryType        string              `json:"categoryType"`        // 类目的类型，1为1688大市场类目，2为1688工业品专业化类目，3为1688主流商品类目
	IsSupportProcessing bool                `json:"isSupportProcessing"` // 类目是否支持加工定制
}

// ChildCategoryInfo 子类目信息
type ChildCategoryInfo struct {
	ID           int64  `json:"id"`           //子类目ID
	Name         string `json:"name"`         // 子类目名称
	IsLeaf       bool   `json:"isLeaf"`       // 是否叶子类目（只有叶子类目才能发布商品）
	CategoryType string `json:"categoryType"` // 类目的类型：1表示cbu类目，2表示gallop类目
}

// FeatureInfo 类目特征信息
type FeatureInfo struct {
	Key       string `json:"key"`       // 名称
	Value     string `json:"value"`     // 值
	Status    int    `json:"status"`    // 状态
	Hierarchy bool   `json:"hierarchy"` // 是否继承到子元素上
}

// WrapResult 包装结果
func (r *Response) WrapResult(result string) {
	// 保存原始结果到Body
	r.Body = result
	// 尝试解析具体的响应结构
	if err := json.Unmarshal([]byte(result), r); err != nil {
		// 如果解析失败，设置错误信息
		r.ErrorResponse.Code = 500
		r.ErrorResponse.Msg = "Failed to parse response: " + err.Error()
	} else if r.ErrorCode != "" || r.ErrorMsg != "" {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = fmt.Sprintf("Failed to parse response: %s, errorCode: %s", r.ErrorMsg, r.ErrorCode)
	}

}
