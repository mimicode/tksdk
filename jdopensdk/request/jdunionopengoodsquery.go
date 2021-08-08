package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

// JdUnionOpenGoodsQueryRequest jd.union.open.goods.query 接口描述：查询商品及优惠券信息，返回的结果可调用转链接口生成单品或二合一推广链接。支持按SKUID、关键词、优惠券基本属性、是否拼购、是否爆款等条件查询，建议不要同时传入SKUID和其他字段，以获得较多的结果。支持按价格、佣金比例、佣金、引单量等维度排序。用优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.query
type JdUnionOpenGoodsQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenGoodsQueryRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenGoodsQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenGoodsQueryRequest) GetApiName() string {
	return "jd.union.open.goods.query"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenGoodsQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
