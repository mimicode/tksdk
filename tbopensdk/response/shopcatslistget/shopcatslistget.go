package shopcatslistget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.shopcats.list.get( 获取前台展示的店铺类目 )
type Response struct {
	response.TopResponse
	ShopcatsListGetResult Result `json:"shopcats_list_get_response"`
}

//解析输出结果
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

type Result struct {
	ShopCats  ShopCats `json:"shop_cats"`
	RequestID string   `json:"request_id"`
}

type ShopCats struct {
	ShopCat []ShopCat `json:"shop_cat"`
}

type ShopCat struct {
	Cid       int64  `json:"cid"`
	IsParent  bool   `json:"is_parent"`
	Name      string `json:"name"`
	ParentCid int64  `json:"parent_cid"`
}
