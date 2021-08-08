package suningnetallianceorderget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.order.get 网盟订单信息单独查询
type Response struct {
	response2.TopResponse
	SnResponseContent SnResponseContent `json:"sn_responseContent"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = "-1"
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	} else {
		if t.SnResponseContent.SnError != nil {
			t.ErrorResponse.ErrorCode = t.SnResponseContent.SnError.ErrorCode
			t.ErrorResponse.ErrorMsg = t.SnResponseContent.SnError.ErrorMsg
		}
	}
}

type SnResponseContent struct {
	SnBody  SnBody   `json:"sn_body"`
	SnError *SnError `json:"sn_error"`
}
type SnError struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
}

type SnBody struct {
	GetOrder GetOrder `json:"getOrder"`
}

type GetOrder struct {
	OrderInfo []OrderInfo `json:"orderInfo"`
	OrderCode string      `json:"orderCode"`
}

type OrderInfo struct {
	GoodsGroupCatalog         string `json:"goodsGroupCatalog"`
	OrderType                 string `json:"orderType"`
	SellName                  string `json:"sellName"`
	PayTime                   string `json:"payTime"`
	ProductThirdCatalog       string `json:"productThirdCatalog"`
	ProductSecondCatalog      string `json:"productSecondCatalog"`
	SaleNum                   string `json:"saleNum"`
	OrderLineStatusDesc       string `json:"orderLineStatusDesc"`
	ProductName               string `json:"productName"`
	OrderLineFlag             string `json:"orderLineFlag"`
	CommissionRatio           string `json:"commissionRatio"`
	SellerCode                string `json:"sellerCode"`
	PayAmount                 string `json:"payAmount"`
	PrePayCommission          string `json:"prePayCommission"`
	OrderLineOrigin           string `json:"orderLineOrigin"`
	OrderSubmitTime           string `json:"orderSubmitTime"`
	ProductFirstCatalog       string `json:"productFirstCatalog"`
	OrderLineNumber           string `json:"orderLineNumber"`
	GoodsNum                  string `json:"goodsNum"`
	OrderLineStatusChangeTime string `json:"orderLineStatusChangeTime"`
	Promotion                 int64  `json:"promotion"`
	ChildAccountID            string `json:"childAccountId"`
}
