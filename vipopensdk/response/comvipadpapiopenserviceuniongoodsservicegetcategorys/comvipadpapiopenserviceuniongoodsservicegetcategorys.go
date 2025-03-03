package comvipadpapiopenserviceuniongoodsservicegetcategorys

import (
	"encoding/json"

	"github.com/mimicode/tksdk/vipopensdk/response"
)

// com.vip.adp.api.open.service.UnionGoodsService getCategorys 类目查询接口
type Response struct {
	response.TopResponse
	Result Result `json:"result"`
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

type Result struct {
	CategoryList []CategoryInfo `json:"categoryList"` // 类目列表
}

type CategoryInfo struct {
	CategoryId   int    `json:"categoryId"`   // 类目ID
	CategoryName string `json:"categoryName"` // 类目名称
	Level        int    `json:"level"`        // 类目层级
	ParentId     int    `json:"parentId"`     // 父类目ID
	IsLeaf       bool   `json:"isLeaf"`       // 是否叶子节点
	Sort         int    `json:"sort"`         // 排序值
}
