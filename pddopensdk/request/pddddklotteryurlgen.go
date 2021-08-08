package request

import (
	utils2 "github.com/mimicode/tksdk/utils"
	"net/url"
)

//pdd.ddk.lottery.url.gen多多客生成转盘抽免单url
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.lottery.url.gen
type PddDdkLotteryUrlGenRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkLotteryUrlGenRequest) CheckParameters() {
	utils2.CheckNotNull(tk.Parameters.Get("pid_list"), "pid_list")

}

//添加请求参数
func (tk *PddDdkLotteryUrlGenRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkLotteryUrlGenRequest) GetApiName() string {
	return "pdd.ddk.lottery.url.gen"
}

//返回请求参数
func (tk *PddDdkLotteryUrlGenRequest) GetParameters() url.Values {
	return *tk.Parameters
}
