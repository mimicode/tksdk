package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionElemePromotionOfficialactivityGetRequest alibaba.alsc.union.eleme.promotion.officialactivity.get( 本地联盟饿了么推广官方活动查询 )
// https://open.taobao.com/api.htm?docId=60449&docType=2&scopeId=24408
type AlibabaAlscUnionElemePromotionOfficialactivityGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")
	var subFields = make(map[string]string)
	subFields = tk.parseSubParameters(tk.Parameters.Get("query_request"))
	utils.CheckNotNull(tk.getMapVal(subFields, "pid"), "pid")
	utils.CheckNotNull(tk.getMapVal(subFields, "activity_id"), "activity_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) GetApiName() string {
	return "alibaba.alsc.union.eleme.promotion.officialactivity.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
