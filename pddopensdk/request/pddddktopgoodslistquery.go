package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.top.goods.list.query（多多客获取爆款排行商品接口）
//https://open.pinduoduo.com/#/apidocument/port?portId=pdd.ddk.top.goods.list.query
type PddDdkTopGoodsListQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkTopGoodsListQueryRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid"), "pid")

}

//添加请求参数
func (tk *PddDdkTopGoodsListQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkTopGoodsListQueryRequest) GetApiName() string {
	return "pdd.ddk.top.goods.list.query"
}

//返回请求参数
func (tk *PddDdkTopGoodsListQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
