package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// JdUnionOpenCategoryGoodsGetRequest jd.union.open.category.goods.get 根据商品的父类目id查询子类目id信息
//https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
type JdUnionOpenCategoryGoodsGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *JdUnionOpenCategoryGoodsGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("360buy_param_json"), "360buy_param_json")

}

// AddParameter 添加请求参数
func (tk *JdUnionOpenCategoryGoodsGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *JdUnionOpenCategoryGoodsGetRequest) GetApiName() string {
	return "jd.union.open.category.goods.get"
}

// GetParameters 返回请求参数
func (tk *JdUnionOpenCategoryGoodsGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
