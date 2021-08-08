package request

import (
	"net/url"
)

//pdd.ddk.oauth.goods.pid.query多多客已生成推广位信息查询
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.pid.query&permissionId=7
type PddDdkOauthGoodsPidQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkOauthGoodsPidQueryRequest) CheckParameters() {

}

//添加请求参数
func (tk *PddDdkOauthGoodsPidQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *PddDdkOauthGoodsPidQueryRequest) GetApiName() string {
	return "pdd.ddk.oauth.goods.pid.query"
}

//返回请求参数
func (tk *PddDdkOauthGoodsPidQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}
