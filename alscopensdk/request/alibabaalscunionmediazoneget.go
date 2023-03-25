package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionMediaZoneGetRequest alibaba.alsc.union.media.zone.get( 本地生活媒体推广位查询 )
// https://open.taobao.com/api.htm?docId=59276&docType=2&scopeId=24408
type AlibabaAlscUnionMediaZoneGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionMediaZoneGetRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page"), "page")
	utils.CheckNumber(tk.Parameters.Get("limit"), "limit")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionMediaZoneGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionMediaZoneGetRequest) GetApiName() string {
	return "alibaba.alsc.union.media.zone.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionMediaZoneGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
