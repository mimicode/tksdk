package jdunionopenorderrowsupplyquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.order.row.supply.query 供开订单行查询接口【申请】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_order_row_supply_query_responce"`
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
	Code    int64                `json:"code"`
	Data    []SupplyOrderRowResp `json:"data"`
	HasMore bool                 `json:"hasMore"` //是否还有更多,true：还有数据；false:已查询完毕，没有数据
	Message string               `json:"message"`
}

// SupplyOrderRowResp 供开订单行明细
type SupplyOrderRowResp struct {
	ID               string       `json:"id"`               //标记唯一订单行：订单+sku维度的唯一标识
	OrderID          int64        `json:"orderId"`          //订单号
	ParentID         int64        `json:"parentId"`         //主单订单号：如一个订单拆成多个子订单时，原订单号会作为主单号，拆分的订单号为子单号存储在orderid中。若未发生拆单，该字段为0
	OrderTime        string       `json:"orderTime"`        //下单时间,格式yyyy-MM-dd HH:mm:ss
	FinishTime       string       `json:"finishTime"`       //完成时间（购买用户确认收货时间）,格式yyyy-MM-dd HH:mm:ss
	ModifyTime       string       `json:"modifyTime"`       //更新时间,格式yyyy-MM-dd HH:mm:ss
	UnionID          int64        `json:"unionId"`          //推客ID
	SkuID            int64        `json:"skuId"`            //商品ID
	SkuName          string       `json:"skuName"`          //商品名称
	SkuNum           int64        `json:"skuNum"`           //商品数量
	SkuReturnNum     int64        `json:"skuReturnNum"`     //商品已退货数量
	SkuFrozenNum     int64        `json:"skuFrozenNum"`     //商品售后中数量
	Price            float64      `json:"price"`            //商品单价
	CommissionRate   float64      `json:"commissionRate"`   //佣金比例(投放的广告主计划比例)
	FinalRate        float64      `json:"finalRate"`        //最终分佣比例（单位：%）
	EstimateCosPrice float64      `json:"estimateCosPrice"` //预估计佣金额：由订单的实付金额拆分至每个商品的预估计佣金额，不包括运费，以及京券、东券、E卡、余额等虚拟资产支付的金额。该字段仅为预估值，实际佣金以actualCosPrice为准进行计算
	EstimateFee      float64      `json:"estimateFee"`      //推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新
	ActualCosPrice   float64      `json:"actualCosPrice"`   //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新
	ActualFee        float64      `json:"actualFee"`        //推客分得的实际佣金（实际计佣金额*佣金比例*最终分佣比例）。如订单完成后发生退款，此金额会更新
	ValidCode        int64        `json:"validCode"`        //sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,5.无效-账号异常,13.违规订单-其他,15.待付款,16.已付款,17.已完成（购买用户确认收货）,27.违规订单-违反京东平台规则等）
	Cid1             int64        `json:"cid1"`             //一级类目id
	Cid2             int64        `json:"cid2"`             //二级类目id
	Cid3             int64        `json:"cid3"`             //三级类目id
	PopID            int64        `json:"popId"`            //商家ID
	PayMonth         int64        `json:"payMonth"`         //预估结算时间，订单完成后才会返回，格式：yyyyMMdd，默认：0。表示最新的预估结算日期
	Sign             string       `json:"sign"`             //数据签名，用来核对出参数据是否被修改，入参fields中写入sign时返回
	ProPriceAmount   float64      `json:"proPriceAmount"`   //价保赔付金额：订单申请价保或赔付的金额，实际计佣金额已经减去此金额，您无需处理
	GoodsInfo        GoodsInfo    `json:"goodsInfo"`        //商品信息，入参传入fields，goodsInfo获取
	CategoryInfo     CategoryInfo `json:"categoryInfo"`     //类目信息,入参传入fields，categoryInfo获取
	ExpressStatus    int64        `json:"expressStatus"`    //发货状态（10：待发货，20：已发货）
	OutSideOrderID   string       `json:"outSideOrderId"`   //抖快外部订单号，若该字段为0，可联系联盟运营
	TalentName       string       `json:"talentName"`       //达人昵称，下单时刻抖快达人的昵称快照，非抖快平台最新的达人昵称
	ApplyPlatform    int64        `json:"applyPlatform"`    //1：抖音平台，2：快手平台
	TalentID         string       `json:"talentId"`         //达人在京东的唯一ID，可在“京东联盟-京红任务-小店达人”中搜索获取，不会随抖快达人昵称的改变而变化
}

// GoodsInfo 商品信息
type GoodsInfo struct {
	ImageURL  string `json:"imageUrl"`  //sku主图链接
	Owner     string `json:"owner"`     //g=自营，p=pop
	MainSkuID int64  `json:"mainSkuId"` //自营商品主Id（owner=g取此值）
	ProductID int64  `json:"productId"` //非自营商品主Id（owner=p取此值）
	ShopName  string `json:"shopName"`  //店铺名称（或供应商名称）
	ShopID    int64  `json:"shopId"`    //店铺Id
}

// CategoryInfo 类目信息
type CategoryInfo struct {
	Cid1     int64  `json:"cid1"`     //一级类目id
	Cid2     int64  `json:"cid2"`     //二级类目id
	Cid3     int64  `json:"cid3"`     //三级类目id
	Cid1Name string `json:"cid1Name"` //一级类目名称
	Cid2Name string `json:"cid2Name"` //二级类目名称
	Cid3Name string `json:"cid3Name"` //三级类目名称
}
