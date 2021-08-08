package sellercatslistget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.sellercats.list.get( 获取前台展示的店铺内卖家自定义商品类目 )
type Response struct {
	response.TopResponse
	SellercatsListGetResult Result `json:"sellercats_list_get_response"`
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
	SellerCats SellerCats `json:"seller_cats"`
	RequestID  string     `json:"request_id"`
}

type SellerCats struct {
	SellerCat []SellerCat `json:"seller_cat"`
}

type SellerCat struct {
	Cid       int64  `json:"cid"`
	Name      string `json:"name"`
	ParentCid int64  `json:"parent_cid"`
	PicURL    string `json:"pic_url"`
	SortOrder int64  `json:"sort_order"`
	Type      string `json:"type"`
}
