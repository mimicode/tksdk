package tbkdgnewuserorderget

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.dg.newuser.tborder.get( 淘宝客新用户订单API--导购 )
type Response struct {
	response.TopResponse
	TbkDgNewuserOrderGetResult Result `json:"tbk_dg_newuser_order_get_response"`
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
	Results   ResponseResults `json:"results"`
	RequestID string          `json:"request_id"`
}

type ResponseResults struct {
	Data ResultsData `json:"data"`
}

type ResultsData struct {
	HasNext  bool                                  `json:"has_next"`
	PageNo   int64                                 `json:"page_no"`
	PageSize int64             `json:"page_size"`
	Results  ResultDataResults `json:"results"`
}

type ResultDataResults struct {
	MapData []ResultMapDatum `json:"map_data"`
}

type ResultMapDatum struct {
	AcceptTime      string `json:"accept_time"`
	ActivityID      string `json:"activity_id"`
	ActivityType    string `json:"activity_type"`
	AdzoneID        int64  `json:"adzone_id"`
	AdzoneName      string `json:"adzone_name"`
	BindTime        string `json:"bind_time"`
	BizDate         string `json:"biz_date"`
	BuyTime         string `json:"buy_time"`
	MemberID        int64  `json:"member_id"`
	Mobile          string `json:"mobile"`
	OrderTkType     int64  `json:"order_tk_type"`
	RegisterTime    string `json:"register_time"`
	SiteID          int64  `json:"site_id"`
	SiteName        string `json:"site_name"`
	Status          int64  `json:"status"`
	TBTradeParentID int64  `json:"tb_trade_parent_id"`
	UnionID         string `json:"union_id"`
}
