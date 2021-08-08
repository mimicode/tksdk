package request

import (
	utils2 "github.com/mimicode/tksdk/pddopensdk/utils"
	"net/url"
)

//taobao.tbk.shop.recommend.get( 淘宝客店铺关联推荐查询 )
//http://open.taobao.com/api.htm?docId=24522&docType=2&scopeId=11655
type TbkShopRecommendGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkShopRecommendGetRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils2.CheckNotNull(tk.Parameters.Get("user_id"), "user_id")
	utils2.CheckNumber(tk.Parameters.Get("user_id"), "user_id")

}

//添加请求参数
func (tk *TbkShopRecommendGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkShopRecommendGetRequest) GetApiName() string {
	return "taobao.tbk.shop.recommend.get"
}

//返回请求参数
func (tk *TbkShopRecommendGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
