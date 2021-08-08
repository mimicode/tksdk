package suningnetallianceorderinfoquery

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/snopensdk/response"
)

//suning.netalliance.orderinfo.query 订单批量查询接口
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
	QueryOrderinfo []QueryOrderinfo `json:"queryOrderinfo"`
}

/*
orderId	String	38404277469	订单行号
buyerId	String	1578985200000	下单人易购会员编码
orderPrice	String	2770	订单价格，单位:分
payStatus	String	0	付款状态 0、未付款 1、已付款
closeType	String	1	订单关闭原因。1、订单完成关闭2、退款关闭3、纠纷关闭
*/
type QueryOrderinfo struct {
	TimeInfo    TimeInfo    `json:"timeInfo"`
	OrderID     string      `json:"orderId"`
	OrderStatus OrderStatus `json:"orderStatus"`
	OrderPrice  int64       `json:"orderPrice"`
	AmountInfo  AmountInfo  `json:"amountInfo"`
	BuyerID     string      `json:"buyerId"`
	CloseType   int64       `json:"closeType"`
	PayStatus   int64       `json:"payStatus"`
}

/*
totalAmount	String	2770	订单金额，单位为分
commissionRate	String	21	推客佣金比例
payAmount	String	2770	实际支付金额
estimateCosPrice	String	0	推客预估计佣金额，单位为分
platformCommissionRate	String	platformCommissionRate	平台佣金比例,千分比
settleAmount	String	2770	结算金额，单位为分
platformEstimateCosPrice	String	277	平台预估计佣金额，单位为分
*/
type AmountInfo struct {
	TotalAmount              int64 `json:"totalAmount"`
	CommissionRate           int64 `json:"commissionRate"`
	EstimateCosPrice         int64 `json:"estimateCosPrice"`
	PayAmount                int64 `json:"payAmount"`
	PlatformEstimateCosPrice int64 `json:"platformEstimateCosPrice"`
	PlatformCommissionRate   int64 `json:"platformCommissionRate"`
	SettleAmount             int64 `json:"settleAmount"`
}

/*
settleStatus	String	0	0:未结算；1:已结算
parentId	String	38404172416	订单号
pid	String	101220022	推广人易购会员编码
relItemId	String	121307256	商品id
traceId	String	2770	用户追踪标识
itemPicUrl	String	https://uimgpre.cnsuning.com/uimg/b2c/newcatentries/0010001242-000000000121307256_1.jpg_800w_800h_4e	商品图片
itemTitle	String	Apple iPhone 6 Plus-下单后改名	商品名称
itemPrice	String	2770	商品价格，单位:分
itemNum	String	1	商品数量
subSideRate	String	10	分成比例。
promotion	String	0	是否风控（0：未风控，1：风控）
relShopId	String	7018120455	店铺编码
*/
type OrderStatus struct {
	TraceID      string `json:"traceId"`
	ItemPicURL   string `json:"itemPicUrl"`
	SettleStatus int64  `json:"settleStatus"`
	ItemTitle    string `json:"itemTitle"`
	OrderStatus  string `json:"orderStatus"`
	RefundStatus int64  `json:"refundStatus"`
	PID          string `json:"pid"`
	RelShopID    string `json:"relShopId"`
	ParentID     string `json:"parentId"`
	ItemNum      string `json:"itemNum"`
	RelItemID    string `json:"relItemId"`
	ItemPrice    int64  `json:"itemPrice"`
	SubSideRate  int64  `json:"subSideRate"`
	Promotion    int64  `json:"promotion"`
}

/*
orderTime	String	1578987802000	下单时间
payTime	String	1578987814000	付款时间
refundTime	String	1578985200000	退款时间
orderFinishTime	String	1578985200000	确认收货时间
settleTime	String	1578985200000	结算时间
updateDate	String	1578987802000	更新时间
*/
type TimeInfo struct {
	UpdateDate      int64 `json:"updateDate"` //updateDate	String	1578987802000	更新时间
	OrderTime       int64 `json:"orderTime"`  //orderTime	String	1578987802000	下单时间
	SettleTime      int64 `json:"settleTime"` //settleTime	String	1578985200000	结算时间
	PayTime         int64 `json:"payTime"`    //payTime	String	1578987814000	付款时间
	RefundTime      int64 `json:"refundTime"` //refundTime	String	1578985200000	退款时间
	OrderFinishTime int64 `json:"orderFinishTime"`
}

type SnHead struct {
	TotalSize     int64  `json:"totalSize"`
	PageTotal     int64  `json:"pageTotal"`
	PageNo        int64  `json:"pageNo"`
	ReturnMessage string `json:"returnMessage"`
}
