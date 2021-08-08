package tbkdgnewuserordersum

import (
	"encoding/json"
	"github.com/mimicode/tksdk/tbopensdk/response"
)

//taobao.tbk.dg.newuser.tborder.sum( 拉新活动汇总API--导购 )
type Response struct {
	response.TopResponse
	TbkDgNewuserOrderSumResult Result `json:"tbk_dg_newuser_order_sum_response"`
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
	Data Data `json:"data"`
}

type Data struct {
	HasNext  bool                            `json:"has_next"`
	PageNo   int64                           `json:"page_no"`
	PageSize int64       `json:"page_size"`
	Results  DataResults `json:"results"`
}

type DataResults struct {
	Map []Map `json:"map"`
}

type Map struct {
	ActivityID           string `json:"activity_id"`
	AlipayUserCnt        int64  `json:"alipay_user_cnt"`
	AlipayUserCPAPreAmt  string `json:"alipay_user_cpa_pre_amt"`
	BindBuyUserCPAPreAmt string `json:"bind_buy_user_cpa_pre_amt"`
	LoginUserCnt         int64  `json:"login_user_cnt"`
	RcvUserCnt           int64  `json:"rcv_user_cnt"`
	RcvValidUserCnt      int64  `json:"rcv_valid_user_cnt"`
	RegUserCnt           int64  `json:"reg_user_cnt"`
}
