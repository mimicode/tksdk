package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenPositionQueryRequest jd.union.open.position.query 查询推广位【申请】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.position.query
type JdUnionOpenPositionQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenPositionQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenPositionQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenPositionQueryRequest) GetApiName() string {
	return "jd.union.open.position.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenPositionQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
