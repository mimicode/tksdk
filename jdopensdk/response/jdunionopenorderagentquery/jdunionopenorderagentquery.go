package jdunionopenorderagentquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.order.agent.query 工具商订单行查询接口
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_order_agent_query_responce"`
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
	Code    int64            `json:"code"`
	Data    []OrderAgentResp `json:"data"`
	HasMore bool             `json:"hasMore"` //是否还有更多,true：还有数据；false:已查询完毕，没有数据
	Message string           `json:"message"`
}

// OrderAgentResp 工具商订单行明细
type OrderAgentResp struct {
	ID                  string       `json:"id"`                  //标记唯一订单行
	OrderID             int64        `json:"orderId"`             //订单号
	ParentID            int64        `json:"parentId"`            //主单的订单号：如一个订单拆成多个子订单时，原订单号会作为主单号，拆分的订单号为子单号存储在orderid中。若未发生拆单，该字段为0
	OrderTime           string       `json:"orderTime"`           //下单时间,格式yyyy-MM-dd HH:mm:ss
	FinishTime          string       `json:"finishTime"`          //完成时间（购买用户确认收货时间）,格式yyyy-MM-dd HH:mm:ss
	ModifyTime          string       `json:"modifyTime"`          //更新时间,格式yyyy-MM-dd HH:mm:ss
	OrderEmt            int64        `json:"orderEmt"`            //下单设备 1.pc 2.无线
	Plus                int64        `json:"plus"`                //plus会员 1:是，0:否
	UnionID             int64        `json:"unionId"`             //推客ID
	SkuID               int64        `json:"skuId"`               //商品ID
	SkuName             string       `json:"skuName"`             //商品名称
	SkuNum              int64        `json:"skuNum"`              //商品数量
	SkuReturnNum        int64        `json:"skuReturnNum"`        //商品已退货数量
	SkuFrozenNum        int64        `json:"skuFrozenNum"`        //商品售后中数量
	Price               float64      `json:"price"`               //商品单价
	CommissionRate      float64      `json:"commissionRate"`      //佣金比例(投放的广告主计划比例)
	SubSideRate         float64      `json:"subSideRate"`         //一级分成比例
	SubsidyRate         float64      `json:"subsidyRate"`         //一级补贴比例
	FinalRate           float64      `json:"finalRate"`           //最终分佣比例（单位：%）=分成比例+补贴比例
	EstimateCosPrice    float64      `json:"estimateCosPrice"`    //预估计佣金额：由订单的实付金额拆分至每个商品的预估计佣金额，不包括运费，以及京券、东券、E卡、余额等虚拟资产支付的金额。该字段仅为预估值，实际佣金以actualCosPrice为准进行计算
	EstimateFee         float64      `json:"estimateFee"`         //推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新
	ActualCosPrice      float64      `json:"actualCosPrice"`      //实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新
	ActualFee           float64      `json:"actualFee"`           //推客分得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新
	ValidCode           int64        `json:"validCode"`           //sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,11.无效-乡村推广员下单,13.违规订单-其他,14.无效-来源与备案网址不符,15.无效-流量作弊,16.有效,17.无效-拆单再购,18.无效-练习订单,19.无效-零食召回订单,20.无效-触屏版订单,21.无效-线上导购订单）
	TraceType           int64        `json:"traceType"`           //同跨店：2同店 3跨店
	PositionID          int64        `json:"positionId"`          //推广位ID
	SiteID              int64        `json:"siteId"`              //应用id（网站id、appid、社交媒体id）
	UnionAlias          string       `json:"unionAlias"`          //平台 PID所属母账号平台名称（原第三方服务商来源），两方分佣会有该值
	PID                 string       `json:"pid"`                 //格式:子推客ID_子站长应用ID_子推客推广位ID
	Cid1                int64        `json:"cid1"`                //一级类目id
	Cid2                int64        `json:"cid2"`                //二级类目id
	Cid3                int64        `json:"cid3"`                //三级类目id
	SubUnionID          string       `json:"subUnionId"`          //子联盟ID(需要联系运营开放白名单才能拿到数据)
	UnionTag            string       `json:"unionTag"`            //联盟标签数据（32位整型二进制字符串：00000000000000000000000000000001。数据从右向左进行，每一位为1表示符合特征，第1位：红包，第2位：组合推广，第3位：拼购，第5位：有效首次购）
	PopID               int64        `json:"popId"`               //商家ID
	Ext1                string       `json:"ext1"`                //推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）
	PayMonth            string       `json:"payMonth"`            //预估结算时间，订单完成后才会返回，格式：yyyyMMdd，默认：0。表示最新的预估结算日期。当payMonth为当前的未来时间时，表示该订单可结算；当payMonth为当前的过去时间时，表示该订单已结算
	CpActID             int64        `json:"cpActId"`             //招商团活动id：当商品参加了招商团会有该值，为0时表示无活动
	UnionRole           int64        `json:"unionRole"`           //站长角色：1 推客 2 团长
	GiftCouponOcsAmount float64      `json:"giftCouponOcsAmount"` //礼金分摊金额：使用礼金的订单会有该值
	GiftCouponKey       string       `json:"giftCouponKey"`       //礼金批次ID：使用礼金的订单会有该值
	Sign                string       `json:"sign"`                //数据签名，用来核对出参数据是否被修改，入参fields中写入sign时返回
	ProPriceAmount      float64      `json:"proPriceAmount"`      //价保赔付金额：订单申请价保或赔付的金额，实际计佣金额已经减去此金额，您无需处理
	GoodsInfo           GoodsInfo    `json:"goodsInfo"`           //商品信息，入参传入fields，goodsInfo获取
	CategoryInfo        CategoryInfo `json:"categoryInfo"`        //类目信息,入参传入fields，categoryInfo获取
	ExpressStatus       int64        `json:"expressStatus"`       //发货状态（10：待发货，20：已发货）
	ChannelID           int64        `json:"channelId"`           //渠道关系ID
	SkuTag              string       `json:"skuTag"`              //64位标签字段，数据从右向左进行，64位新标签可以兼容32位unionTag，右32位参考unionTag
	Rid                 int64        `json:"rid"`                 //团长渠道ID，仅限招商团长管理渠道使用，团长开通权限后才可使用
	ItemID              string       `json:"itemId"`              //联盟商品ID
	CallerItemID        string       `json:"callerItemId"`        //工具商联盟商品ID
	SubCpUnionID        int64        `json:"subCpUnionId"`        //【废弃】，请勿使用
}

// GoodsInfo 商品信息
type GoodsInfo struct {
	ImageURL  string `json:"imageUrl"`  //sku主图链接
	Owner     string `json:"owner"`     //g=自营，p=pop
	MainSkuID int64  `json:"mainSkuId"` //自营商品主Id（owner=g取此值）
	ProductID int64  `json:"productId"` //非自营商品主Id（owner=p取此值）
	ShopName  string `json:"shopName"`  //店铺名称（或供应商名称）
	ShopID    int64  `json:"shopId"`    //店铺ID
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
