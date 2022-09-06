package request

import (
	"net/url"
	"github.com/mimicode/tksdk/utils"
)

//taobao.tbk.item.click.extract( 淘宝客-公用-链接解析出商品id )
//https://open.taobao.com/api.htm?docId=43873&docType=2&scopeId=16401
type TbkItemClickExtractRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkItemClickExtractRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("click_url"), "click_url")

}

//添加请求参数
func (tk *TbkItemClickExtractRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemClickExtractRequest) GetApiName() string {
	return "taobao.tbk.item.click.extract"
}

//返回请求参数
func (tk *TbkItemClickExtractRequest) GetParameters() url.Values {
	return *tk.Parameters
}
