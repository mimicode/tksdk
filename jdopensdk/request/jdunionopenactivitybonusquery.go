package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenActivityBonusQueryRequest jd.union.open.activity.bonus.query 奖励活动信息查询接口【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.activity.bonus.query
type JdUnionOpenActivityBonusQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenActivityBonusQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenActivityBonusQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenActivityBonusQueryRequest) GetApiName() string {
	return "jd.union.open.activity.bonus.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenActivityBonusQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
