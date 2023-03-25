package request

import (
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionMediaZoneAddRequest alibaba.alsc.union.media.zone.add( 本地生活媒体推广位创建 )
// https://open.taobao.com/api.htm?docId=58768&docType=2&scopeId=24408
type AlibabaAlscUnionMediaZoneAddRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionMediaZoneAddRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("zone_name"), "zone_name")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionMediaZoneAddRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionMediaZoneAddRequest) GetApiName() string {
	return "alibaba.alsc.union.media.zone.add"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionMediaZoneAddRequest) GetParameters() url.Values {
	return *tk.Parameters
}
