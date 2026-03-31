package productcategory

import (
	"encoding/json"

	"github.com/mimicode/tksdk/csjdyopensdk/response"
)

// Response 商品类目响应
type Response struct {
	response.TopResponse
	Data ResponseData `json:"data"`
}

func (r *Response) WrapResult(result string) {
	r.Body = result
	err := json.Unmarshal([]byte(result), r)
	if err != nil {
		r.Code = -1
		r.Desc = err.Error()
	}
}

// ResponseData 响应数据
type ResponseData struct {
	Total        int64      `json:"total"`         // 总类目数
	CategoryList []Category `json:"category_list"` // 类目详情列表
}

// Category 类目信息
type Category struct {
	Id    int64  `json:"id"`    // 类目id
	Name  string `json:"name"`  // 类目名称
	Level int    `json:"level"` // 类目层级（最多三级）
}
