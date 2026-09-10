package jdunionopengoodsseckillquery

import (
	"encoding/json"

	"github.com/mimicode/tksdk/jdopensdk/response"
)

// Response jd.union.open.goods.seckill.query 秒杀商品查询接口【即将下线】
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_goods_seckill_query_responce"`
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
	Code       int64              `json:"code"`
	Data       []SecKillGoodsResp `json:"data"`
	Message    string             `json:"message"`
	TotalCount int64              `json:"totalCount"` //总数量
}

// SecKillGoodsResp 秒杀商品明细
type SecKillGoodsResp struct {
	SkuName            string  `json:"skuName"`            //商品名称
	SkuID              int64   `json:"skuId"`              //商品id
	ImageURL           string  `json:"imageUrl"`           //图片url
	IsSecKill          int64   `json:"isSecKill"`          //是秒杀。1：是商品 0：非秒杀商品
	OriPrice           float64 `json:"oriPrice"`           //原价
	SecKillPrice       float64 `json:"secKillPrice"`       //秒杀价
	SecKillStartTime   int64   `json:"secKillStartTime"`   //秒杀开始展示时间（时间戳：毫秒）
	SecKillEndTime     int64   `json:"secKillEndTime"`     //秒杀结束时间（时间戳：毫秒）
	Cid1ID             int64   `json:"cid1Id"`             //一级类目id
	Cid2ID             int64   `json:"cid2Id"`             //二级类目id
	Cid3ID             int64   `json:"cid3Id"`             //三级类目id
	Cid1Name           string  `json:"cid1Name"`           //一级类目名称
	Cid2Name           string  `json:"cid2Name"`           //二级类目名称
	Cid3Name           string  `json:"cid3Name"`           //三级类目名称
	CommissionShare    float64 `json:"commissionShare"`    //通用佣金比例，百分比
	Commission         float64 `json:"commission"`         //通用佣金
	Owner              string  `json:"owner"`              //是否自营。g=自营，p=pop
	InOrderCount30Days int64   `json:"inOrderCount30Days"` //30天引入订单量（spu）
	InOrderComm30Days  float64 `json:"inOrderComm30Days"`  //30天支出佣金（spu）
	JdPrice            float64 `json:"jdPrice"`            //京东价
}
