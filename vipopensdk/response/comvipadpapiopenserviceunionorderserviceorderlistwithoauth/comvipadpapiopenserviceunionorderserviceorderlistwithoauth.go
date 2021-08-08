package comvipadpapiopenserviceunionorderserviceorderlistwithoauth

import (
	"encoding/json"
	response2 "github.com/mimicode/tksdk/vipopensdk/response"
)

//com.vip.adp.api.open.service.UnionOrderService 获取订单信息列表-需要oauth授权
type Response struct {
	response2.TopResponse
	Result Result `json:"result"`
}

//解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ReturnCode = "-1"
		t.ReturnMessage = unmarshal.Error()
	}
}

/*
	orderInfoList	List<OrderInfo>	否			业绩查询响应结果
	total	Integer	否			业绩总条数
	page	Integer	否			当前页码
	pageSize	Integer	否			页面大小
*/
type Result struct {
	Total         int64           `json:"total"`
	PageSize      int64           `json:"pageSize"`
	Page          int64           `json:"page"`
	OrderInfoList []OrderInfoList `json:"orderInfoList"`
}

/*
	orderSn	String	否			订单号
	status	Short	否			订单状态:0-不合格，1-待定，2-已完结
	newCustomer	Short	否			新老客：0-待定，1-新客，2-老客
	channelTag	String	否			渠道商模式下表示自定义渠道标识；工具商模式下表示pid
	orderTime	Long	否			下单时间 时间戳 单位毫秒
	signTime	Long	否			签收时间 时间戳 单位毫秒
	settledTime	Long	否			结算时间 时间戳 单位毫秒
	detailList	List<OrderDetailInfo>	否			商品明细
	lastUpdateTime	Long	否			订单上次更新时间 时间戳 单位毫秒
	settled	Short	否			订单结算状态 0-未结算,1-已结算
	selfBuy	Integer	否			是否自推自买 0-否，1-是
	orderSubStatusName	String	否			订单子状态：流转状态-支持状态：（已下单、已付款、已签收、待结算、已结算、已失效）
	commission	String	否			商品总佣金:单位元
	afterSaleChangeCommission	String	否			售后订单佣金变动：仅在订单完结之后发生售后行为时返回
	afterSaleChangeGoodsCount	Integer	否			售后订单总商品数量变动：仅在订单完结之后发生售后行为时返回
	commissionEnterTime	Long	否			入账时间，时间戳，单位毫秒
	orderSource	String	否			订单来源
	pid	String	否			推广PID:目前等同于channelTag
	isPrepay	Integer	否			是否预付订单:0-否，1-是
	statParam	String	否			自定义统计参数
*/
type OrderInfoList struct {
	OrderSource               string       `json:"orderSource"`
	OrderSn                   string       `json:"orderSn"`
	Settled                   int64        `json:"settled"`
	CommissionEnterTime       int64        `json:"commissionEnterTime"`
	PID                       string       `json:"pid"`
	OrderSubStatusName        string       `json:"orderSubStatusName"`
	IsPrepay                  int64        `json:"isPrepay"`
	StatParam                 string       `json:"statParam"`
	OrderTime                 int64        `json:"orderTime"`
	NewCustomer               int64        `json:"newCustomer"`
	DetailList                []DetailList `json:"detailList"`
	Commission                string       `json:"commission"`
	ChannelTag                string       `json:"channelTag"`
	AfterSaleChangeCommission string       `json:"afterSaleChangeCommission"`
	AfterSaleChangeGoodsCount string       `json:"afterSaleChangeGoodsCount"`
	SelfBuy                   int64        `json:"selfBuy"`
	Status                    int64        `json:"status"`
	LastUpdateTime            int64        `json:"lastUpdateTime"`
}

/*
	goodsId	String	否			商品id
	goodsName	String	否			商品名称
	goodsThumb	String	否			商品缩略图
	goodsCount	Integer	否			商品数量
	commissionTotalCost	String	否			商品计佣金额(元,保留两位小数)
	commissionRate	String	否			商品佣金比例(%)
	commission	String	否			商品佣金金额(元,保留两位小数)
	commCode	String	否			佣金编码：对应商品二级分类
	commName	String	否			佣金方案名称
	orderSource	String	否			订单来源
	afterSaleInfo	List<AfterSaleDetailInfo>	否			商品售后信息
	sizeId	String	否			商品尺码：2019.01.01之后可用
	status	Short	否			商品状态：0-不合格，1-待定，2-已完结
	brandStoreSn	String	否			品牌编号
	brandStoreName	String	否			品牌名称
*/
type DetailList struct {
	CommissionTotalCost string          `json:"commissionTotalCost"`
	SizeID              string          `json:"sizeId"`
	CommissionRate      string          `json:"commissionRate"`
	OrderSource         string          `json:"orderSource"`
	GoodsID             string          `json:"goodsId"`
	BrandStoreSn        string          `json:"brandStoreSn"`
	BrandStoreName      string          `json:"brandStoreName"`
	CommCode            string          `json:"commCode"`
	GoodsCount          int64           `json:"goodsCount"`
	GoodsThumb          string          `json:"goodsThumb"`
	Commission          string          `json:"commission"`
	CommName            string          `json:"commName"`
	GoodsName           string          `json:"goodsName"`
	Status              int64           `json:"status"`
	AfterSaleInfo       []AfterSaleInfo `json:"afterSaleInfo"`
}

/*
afterSaleChangedCommission	String	否			商品佣金售后变动:仅在订单完结之后发生售后时返回，无售后时为空
afterSaleChangedGoodsCount	Integer	否			商品数量售后变动:仅在订单完结之后发生售后时返回，无售后时为空
afterSaleSn	String	否			商品售后单号，无售后时为空
afterSaleStatus	Integer	否			商品售后状态：1-售后中，2-售后完成，3-售后取消，无售后时为空
afterSaleType	Integer	否			售后类型：1-退货，2-换货,无售后时为空
afterSaleFinishTime	Long	否			售后完成时间，时间戳，单位：毫秒，无售后时为空
*/
type AfterSaleInfo struct {
	AfterSaleChangedCommission string `json:"afterSaleChangedCommission"`
	AfterSaleChangedGoodsCount int64  `json:"afterSaleChangedGoodsCount"`
	AfterSaleSn                string `json:"afterSaleSn"`
	AfterSaleStatus            int64  `json:"afterSaleStatus"`
	AfterSaleType              int64  `json:"afterSaleType"`
	AfterSaleFinishTime        int64  `json:"afterSaleFinishTime"`
}
