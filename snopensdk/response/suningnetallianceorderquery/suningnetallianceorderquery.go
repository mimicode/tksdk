package suningnetallianceorderquery

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.order.query 网盟订单信息批量查询
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
	QueryOrder []QueryOrder `json:"queryOrder"`
}

type QueryOrder struct {
	//订单号，限时打折活动ID
	OrderCode   string        `json:"orderCode"`
	OrderDetail []OrderDetail `json:"orderDetail"`
}

/*
	payTime	String	2017-10-10 10:00:00	支付时间，格式：yyyy-MM-dd HH:mm:ss
	orderSubmitTime	String	2017-10-10 10:00:00	下单时间，格式：yyyy-MM-dd HH:mm:ss
	orderLineNumber	String	32000012016	订单行项目号
	orderLineStatusDesc	String	支付完成	订单行项目状态
	orderLineStatusChangeTime	String	2017-10-10 10:00:00	行项目状态更新时间，格式：yyyy-MM-dd HH:mm:ss
	orderLineOrigin	String	PC端	订单行来源（PC端、无线端）
	productName	String	洗衣液	商品名称
	saleNum	String	1	商品数量
	payAmount	String	123.00	实付金额
	orderLineFlag	String	普通订单	订单行标记
	childAccountId	String	egewtwexcccw	子推广账号ID(对应sub_user)
	sellName	String	南京苏宁软件技术有限公司	商户名称
	sellerCode	String	7001223	商户编码
	goodsNum	String	134993889	商品编码
	commissionRatio	String	0.03	佣金比例
	prePayCommission	String	10.30	预估佣金
	productFirstCatalog	String	R8219	一级目录
	productSecondCatalog	String	R3510	二级目录
	productThirdCatalog	String	R0102	三级目录
	orderType	String	自营	商品归属
	positionId	String	100365	推广位ID
	goodsGroupCatalog	String	R0000001	商品组目录编码
	saleType	String	链接推广	推广类型
	pictureUrl	String	https://imgservice.suning.cn/uimg1/b2c/image/U6Tx80lyPMvao11bfw7yUw.jpg_200w_200h_4e	商品主图
	promotion	String	1	1.风控订单
	violation	String	1	是否违规,0：否；1：是
	returnCommission	String	1	是否返佣，0：否；1：是
*/
type OrderDetail struct {
	GoodsGroupCatalog         string `json:"goodsGroupCatalog"`
	OrderType                 string `json:"orderType"`
	PayTime                   string `json:"payTime"`
	SaleType                  string `json:"saleType"`
	ProductSecondCatalog      string `json:"productSecondCatalog"`
	OrderLineStatusDesc       string `json:"orderLineStatusDesc"`
	ProductName               string `json:"productName"`
	OrderLineFlag             string `json:"orderLineFlag"`
	SellerCode                string `json:"sellerCode"`
	PayAmount                 string `json:"payAmount"`
	OrderLineOrigin           string `json:"orderLineOrigin"`
	ProductFirstCatalog       string `json:"productFirstCatalog"`
	GoodsNum                  string `json:"goodsNum"`
	SellName                  string `json:"sellName"`
	PictureURL                string `json:"pictureUrl"`
	ProductThirdCatalog       string `json:"productThirdCatalog"`
	SaleNum                   string `json:"saleNum"`
	CommissionRatio           string `json:"commissionRatio"`
	ReturnCommission          int64  `json:"returnCommission"`
	PrePayCommission          string `json:"prePayCommission"`
	PositionID                string `json:"positionId"`
	Violation                 int64  `json:"violation"`
	OrderSubmitTime           string `json:"orderSubmitTime"`
	OrderLineNumber           string `json:"orderLineNumber"`
	OrderLineStatusChangeTime string `json:"orderLineStatusChangeTime"`
	Promotion                 int64  `json:"promotion"`
	ChildAccountID            string `json:"childAccountId"`
}

type SnHead struct {
	TotalSize     int64  `json:"totalSize"`
	PageTotal     int64  `json:"pageTotal"`
	PageNo        int64  `json:"pageNo"`
	ReturnMessage string `json:"returnMessage"`
}
