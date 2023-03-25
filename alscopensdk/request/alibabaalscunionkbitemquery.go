package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionKbItemQueryRequest alibaba.alsc.union.kb.item.query( 本地联盟口碑商品列表 )
// https://open.taobao.com/api.htm?docId=63314&docType=2&scopeId=24408
type AlibabaAlscUnionKbItemQueryRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionKbItemQueryRequest) CheckParameters() {
	utils.CheckNumber(tk.Parameters.Get("page_number"), "page_number")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNotNull(tk.Parameters.Get("biz_type"), "biz_type")
	utils.CheckNotNull(tk.Parameters.Get("pid"), "pid")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionKbItemQueryRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionKbItemQueryRequest) GetApiName() string {
	return "alibaba.alsc.union.kb.item.query"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionKbItemQueryRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionKbItemQueryRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionKbItemQueryRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
