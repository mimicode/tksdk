package suningnetallianceordersettlenewget

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.ordersettlenew.get 订单结算信息接口
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
	GetOrdersettlenew GetOrdersettlenew `json:"getOrdersettlenew"`
}

type GetOrdersettlenew struct {
	SettlementInfo []SettlementInfo `json:"settlementInfo"`
}

/*
	orderCode	String	1001434507	订单号
	orderLineNumber	String	10541154444	订单行项目号
	orderLineStatusDesc	String	支付完成	订单行项目状态
	orderLineStatusChangeTime	String	2014-07-19 21:12:22	行项目状态更新时间，格式：yyyy-MM-dd HH:mm:ss
	commissionRatio	String	2%	佣金比例
	payAmount	String	1000	实付金额
	prePayCommission	String	20	预付佣金金额
	needPayCommission	String	20	应付佣金金额
	productName	String	苹果6	商品名称
	productFirstCatalog	String	手机/数码/配件	商品一级目录
	productSecondCatalog	String	手机通讯	商品二级目录
	productThirdCatalog	String	手机	商品三级目录
	productFourthCatalog	String	苹果手机	商品四级目录
	isWholesale	String	0	是否违规：0：否；1：是
	isGrant	String	1	结算状态，0：未结算；1：已结算
	childAccountId	String	100022122	子推广账号ID
	sellerName	String	玉石买买提	商户名称
	sellerCode	String	23232323	商户编码
	goodsNum	String	12323	商品编码
	orderLineFlag	String	1	订单行标记
	returnCommission	String	1	是否返佣，0：否；1：是


*/
type SettlementInfo struct {
	SellName                  string `json:"sellName"`
	NeedPayCommission         string `json:"needPayCommission"`
	ProductThirdCatalog       string `json:"productThirdCatalog"`
	ProductSecondCatalog      string `json:"productSecondCatalog"`
	OrderLineStatusDesc       string `json:"orderLineStatusDesc"`
	ProductName               string `json:"productName"`
	OrderLineFlag             string `json:"orderLineFlag"`
	CommissionRatio           string `json:"commissionRatio"`
	SellerCode                string `json:"sellerCode"`
	ReturnCommission          string `json:"returnCommission"`
	PayAmount                 string `json:"payAmount"`
	PrePayCommission          string `json:"prePayCommission"`
	ProductFirstCatalog       string `json:"productFirstCatalog"`
	OrderCode                 string `json:"orderCode"`
	IsGrant                   string `json:"isGrant"`
	OrderLineNumber           string `json:"orderLineNumber"`
	GoodsNum                  string `json:"goodsNum"`
	OrderLineStatusChangeTime string `json:"orderLineStatusChangeTime"`
	IsWholesale               string `json:"isWholesale"`
	ChildAccountID            string `json:"childAccountId"`
}
