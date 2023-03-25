package request

import (
	"encoding/json"
	"github.com/mimicode/tksdk/utils"
	"net/url"
)

// AlibabaAlscUnionElemePromotionStorepromotionGetRequest alibaba.alsc.union.eleme.promotion.storepromotion.get( 本地联盟饿了么单店推广单店铺查询 )
// https://open.taobao.com/api.htm?docId=60448&docType=2&scopeId=24408
type AlibabaAlscUnionElemePromotionStorepromotionGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("query_request"), "query_request")
	var subFields = make(map[string]string)
	subFields = tk.parseSubParameters(tk.Parameters.Get("query_request"))
	utils.CheckNotNull(tk.getMapVal(subFields, "pid"), "pid")
	utils.CheckNotNull(tk.getMapVal(subFields, "shop_id"), "shop_id")

}

// AddParameter 添加请求参数
func (tk *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// GetApiName 返回接口名称
func (tk *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) GetApiName() string {
	return "alibaba.alsc.union.eleme.promotion.storepromotion.get"
}

// GetParameters 返回请求参数
func (tk *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}

func (tk *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) parseSubParameters(val string) (data map[string]string) {
	data = make(map[string]string)
	if val == "" {
		return data
	}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		panic(err)
	}
	return data
}

func (tk *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) getMapVal(data map[string]string, key string) string {
	if len(data) == 0 {
		return ""
	}
	v, ok := data[key]
	if !ok {
		return ""
	}
	return v
}
