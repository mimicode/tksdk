package productcategory

import (
	"encoding/json"

	"github.com/mimicode/tksdk/pangolinopensdk/response"
)

// Response 商品类目响应
type Response struct {
	response.TopResponse
	Data Data `json:"data"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	t.RawData = result
	if err := json.Unmarshal([]byte(result), t); err != nil {
		t.Code = -1
		t.Desc = err.Error()
	}
}

// Data 响应数据
type Data struct {
	Total        int64      `json:"total"`         // 总类目数
	CategoryList []Category `json:"category_list"` // 类目详情列表
}

// Category 类目信息
type Category struct {
	ID    int64  `json:"id"`    // 类目id
	Name  string `json:"name"`  // 类目名称
	Level int    `json:"level"` // 类目层级（最多三级）
}
