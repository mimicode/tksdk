package tbkiteminfoget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//淘宝客商品详情查询（简版） 免费 不需要授权
type Response struct {
	response.TopResponse
	TbkItemInfoGetResult Result `json:"tbk_item_info_get_response"`
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
	Results   TbkItemInfoGeResults `json:"results"`
	RequestID string               `json:"request_id"`
}

type TbkItemInfoGeResults struct {
	NTbkItem []TbkItemInfoGeNTbkItem `json:"n_tbk_item"`
}

type TbkItemInfoGeNTbkItem struct {
	CatLeafName            string                   `json:"cat_leaf_name"`
	CatName                string                   `json:"cat_name"`
	ItemURL                string                   `json:"item_url"`
	MaterialLIBType        string                   `json:"material_lib_type"`
	Nick                   string                   `json:"nick"`
	NumIid                 interface{}              `json:"num_iid"`
	PictURL                string                   `json:"pict_url"`
	Provcity               string                   `json:"provcity"`
	ReservePrice           string                   `json:"reserve_price"`
	SellerID               int64                    `json:"seller_id"`
	SmallImages            TbkItemInfoGeSmallImages `json:"small_images"`
	Title                  string                   `json:"title"`
	UserType               int64                    `json:"user_type"`
	Volume                 int64                    `json:"volume"`
	ZkFinalPrice           string                   `json:"zk_final_price"`
	PresaleDeposit         string                   `json:"presale_deposit"`
	PresaleDiscountFeeText string                   `json:"presale_discount_fee_text"`
	PresaleEndTime         int64                    `json:"presale_end_time"`
	PresaleStartTime       int64                    `json:"presale_start_time"`
	PresaleTailEndTime     int64                    `json:"presale_tail_end_time"`
	PresaleTailStartTime   int64                    `json:"presale_tail_start_time"`
}

type TbkItemInfoGeSmallImages struct {
	String []string `json:"string"`
}
