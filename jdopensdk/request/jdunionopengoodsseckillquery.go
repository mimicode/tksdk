package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenGoodsSeckillQueryRequest jd.union.open.goods.seckill.query 秒杀商品查询接口【即将下线】
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.seckill.query
type JdUnionOpenGoodsSeckillQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsSeckillQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsSeckillQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsSeckillQueryRequest) GetApiName() string {
	return "jd.union.open.goods.seckill.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsSeckillQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
