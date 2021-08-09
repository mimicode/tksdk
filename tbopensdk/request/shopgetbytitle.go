package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

//taobao.shop.getbytitle( 根据店铺名称获取店铺信息 )
//http://open.taobao.com/api.htm?docId=24852&docType=2&scopeId=386
type ShopGetbytitleRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *ShopGetbytitleRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("title"), "title")
}

//添加请求参数
func (tk *ShopGetbytitleRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *ShopGetbytitleRequest) GetApiName() string {
	return "taobao.shop.getbytitle"
}

//返回请求参数
func (tk *ShopGetbytitleRequest) GetParameters() url.Values {
	return *tk.Parameters
}
