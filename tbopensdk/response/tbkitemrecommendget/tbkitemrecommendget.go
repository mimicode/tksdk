package tbkitemrecommendget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.item.recommend.get( 淘宝客商品关联推荐查询 )
type Response struct {
	response.TopResponse
	TbkItemRecommendGetResult Result `json:"tbk_item_recommend_get_response"`
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
	Results   Results `json:"results"`
	RequestID string  `json:"request_id"`
}

type Results struct {
	NTbkItem []NTbkItem `json:"n_tbk_item"`
}

type NTbkItem struct {
	ItemURL      string                         `json:"item_url"`
	Nick         string                         `json:"nick"`
	NumIid       int64                          `json:"num_iid"`
	PictURL      string                         `json:"pict_url"`
	Provcity     string                         `json:"provcity"`
	ReservePrice string                         `json:"reserve_price"`
	SellerID    int64       `json:"seller_id"`
	SmallImages SmallImages `json:"small_images"`
	Title       string      `json:"title"`
	UserType     int64                          `json:"user_type"`
	Volume       int64                          `json:"volume"`
	ZkFinalPrice string                         `json:"zk_final_price"`
}

type SmallImages struct {
	String []string `json:"string"`
}
