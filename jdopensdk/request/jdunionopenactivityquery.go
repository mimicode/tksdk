package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenActivityQueryRequest jd.union.open.activity.query 提供联盟官方活动查询
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.activity.query
type JdUnionOpenActivityQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenActivityQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenActivityQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenActivityQueryRequest) GetApiName() string {
	return "jd.union.open.activity.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenActivityQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
