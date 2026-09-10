package jdunionopenorderquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.order.query 订单查询接口
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_order_query_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析queryResult
		if t.Responce.QueryResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty queryResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.QueryResultStr), &t.Responce.QueryResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.QueryResult.Code
				t.ErrorResponse.Message = t.Responce.QueryResult.Message
			}
		}
	}
	t.Responce.QueryResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"queryResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code    int64       `json:"code"`
	Data    []OrderResp `json:"data"`
	HasMore bool        `json:"hasMore"` //是否还有更多,true：还有数据；false:已查询完毕，没有数据
	Message string      `json:"message"`
}

// OrderResp 订单信息
type OrderResp struct {
	FinishTime int64     `json:"finishTime"` //订单完成时间（购买用户确认收货时间）时间戳，毫秒
	OrderEmt   int64     `json:"orderEmt"`   //下单设备(1:PC,2:无线)
	OrderID    int64     `json:"orderId"`    //订单ID
	OrderTime  int64     `json:"orderTime"`  //下单时间(时间戳，毫秒)
	ParentID   int64     `json:"parentId"`   //主单的订单号：如某个订单因为仓储物流或其它原因拆成多笔订单时，拆分前的原订单号会作为主单号存储在该字段中，拆分出的新订单号作为子单号存储在orderid中，若未发生拆单，该字段为0
	PayMonth   string    `json:"payMonth"`   //订单维度预估结算时间,不建议使用，可以用订单行sku维度paymonth字段参考（格式：yyyyMMdd），0：未结算。账号未通过资质审核或订单发生售后，会影响订单实际结算时间
	Plus       int64     `json:"plus"`       //下单用户是否为PLUS会员 0：否，1：是
	PopID      int64     `json:"popId"`      //订单维度商家ID，不建议使用，可以用订单行sku维度popId参考
	SkuList    []SkuInfo `json:"skuList"`    //订单包含的商品信息列表
	UnionID    int64     `json:"unionId"`    //推客的联盟ID
	Ext1       string    `json:"ext1"`       //订单维度的推客生成推广链接时传入的扩展字段，不建议使用，可以用订单行sku维度ext1参考（需要联系运营开放白名单才能拿到数据）
	ValidCode  int64     `json:"validCode"`  //sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,11.无效-乡村推广员下单,13.违规订单-其他,14.无效-来源与备案网址不符等）
}

// SkuInfo 订单行（商品维度）信息
type SkuInfo struct {
	ActualCosPrice      float64 `json:"actualCosPrice"`      //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新
	ActualFee           float64 `json:"actualFee"`           //推客获得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新
	CommissionRate      float64 `json:"commissionRate"`      //佣金比例
	EstimateCosPrice    float64 `json:"estimateCosPrice"`    //预估计佣金额：由订单的实付金额拆分至每个商品的预估计佣金额，不包括运费，以及京券、东券、E卡、余额等虚拟资产支付的金额。该字段仅为预估值，实际佣金以actualCosPrice为准进行计算
	EstimateFee         float64 `json:"estimateFee"`         //推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新
	FinalRate           float64 `json:"finalRate"`           //最终比例（分成比例+补贴比例）
	Cid1                int64   `json:"cid1"`                //一级类目ID
	FrozenSkuNum        int64   `json:"frozenSkuNum"`        //商品售后中数量
	PID                 string  `json:"pid"`                 //联盟子站长身份标识，格式：子站长ID_子站长网站ID_子站长推广位ID
	PositionID          int64   `json:"positionId"`          //推广位ID,0代表无推广位
	Price               float64 `json:"price"`               //商品单价
	Cid2                int64   `json:"cid2"`                //二级类目ID
	SiteID              int64   `json:"siteId"`              //网站ID，0：无网站
	SkuID               int64   `json:"skuId"`               //商品ID
	SkuName             string  `json:"skuName"`             //商品名称
	SkuNum              int64   `json:"skuNum"`              //商品数量
	SkuReturnNum        int64   `json:"skuReturnNum"`        //商品已退货数量
	SubSideRate         float64 `json:"subSideRate"`         //分成比例
	SubsidyRate         float64 `json:"subsidyRate"`         //补贴比例
	Cid3                int64   `json:"cid3"`                //三级类目ID
	UnionAlias          string  `json:"unionAlias"`          //PID所属母账号平台名称（原第三方服务商来源）
	UnionTag            string  `json:"unionTag"`            //联盟标签数据（32位整型二进制字符串。数据从右向左进行，每一位为1表示符合特征，第1位：红包，第2位：组合推广，第3位：拼购，第5位：有效首次购）
	UnionTrafficGroup   int64   `json:"unionTrafficGroup"`   //渠道组 1：1号店，其他：京东
	ValidCode           int64   `json:"validCode"`           //sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,13.违规订单-其他,16.有效等）
	SubUnionID          string  `json:"subUnionId"`          //子渠道标识，在转链时可自定义传入，格式要求：字母、数字或下划线，最多支持80个字符(需要联系运营开放白名单才能拿到数据)
	TraceType           int64   `json:"traceType"`           //2：同店；3：跨店
	PayMonth            int64   `json:"payMonth"`            //订单行维度预估结算时间（格式：yyyyMMdd），0：未结算。订单'预估结算时间'仅供参考。账号未通过资质审核或订单发生售后，会影响订单实际结算时间
	PopID               int64   `json:"popId"`               //商家ID。'订单行维度'
	Ext1                string  `json:"ext1"`                //推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）。'订单行维度'
	CpActID             int64   `json:"cpActId"`             //招商团活动id：当商品参加了招商团会有该值，为0时表示无活动
	UnionRole           int64   `json:"unionRole"`           //站长角色：1 推客 2 团长 3内容服务商
	GiftCouponKey       string  `json:"giftCouponKey"`       //礼金批次ID：使用礼金的订单会有该值
	GiftCouponOcsAmount float64 `json:"giftCouponOcsAmount"` //礼金分摊金额：使用礼金的订单会有该值
	ProPriceAmount      float64 `json:"proPriceAmount"`      //价保赔付金额：订单申请价保或赔付的金额，实际计佣金额已经减去此金额，您无需处理
	ItemID              string  `json:"itemId"`              //联盟商品ID
}
