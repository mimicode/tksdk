package snopensdk

import (
	"encoding/json"
	"fmt"
	request2 "github.com/mimicode/tksdk/snopensdk/request"
	suningnetalliancebacthcustomlinkquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancebacthcustomlinkquery"
	suningnetalliancecommoditydetailquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancecommoditydetailquery"
	suningnetalliancecommodityimagesquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancecommodityimagesquery"
	suningnetalliancecouponinfoquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancecouponinfoquery"
	suningnetalliancecustompromotionurlquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancecustompromotionurlquery"
	suningnetallianceextensionlinkget2 "github.com/mimicode/tksdk/snopensdk/response/suningnetallianceextensionlinkget"
	suningnetalliancehoistinglinkquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancehoistinglinkquery"
	suningnetallianceinverstmentcommodityquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetallianceinverstmentcommodityquery"
	suningnetalliancemorerecommendget2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancemorerecommendget"
	suningnetallianceorderinfoquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetallianceorderinfoquery"
	suningnetalliancerecommendcommodityquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancerecommendcommodityquery"
	suningnetalliancesearchcommoditynewquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancesearchcommoditynewquery"
	suningnetalliancesearchcommodityquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancesearchcommodityquery"
	suningnetallianceselectrecommendcommodityquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetallianceselectrecommendcommodityquery"
	suningnetalliancestorepromotionurlquery2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancestorepromotionurlquery"
	suningnetalliancetoolselleradd2 "github.com/mimicode/tksdk/snopensdk/response/suningnetalliancetoolselleradd"
	suningnetallianceunioninfomationget2 "github.com/mimicode/tksdk/snopensdk/response/suningnetallianceunioninfomationget"
	"io/ioutil"
	"os"
	"testing"
)

var (
	appKey, appSecret, sessionKey string
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Sn struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
				} `json:"sn"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Sn.AppKey
				appSecret = data.Sn.AppSecret
				sessionKey = data.Sn.SessionKey
			}
		}
	}
}
func GetClient() *TopClient {
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)
	return client
}

func TestSuningNetallianceExtensionlinkGetRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceExtensionlinkGetRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"getExtensionlink": map[string]interface{}{
				"productUrl": "https://product.suning.com/0071258669/12081096229.html",
				"quanUrl":    "https://quan.suning.com/lqzx_recommend.do?activityId=202007210013686847&activitySecretKey=lnatEjIgcmLDkykEafYcOftE",
				"subUser":    "user_1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetallianceextensionlinkget2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetallianceextensionlinkget2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceCommoditydetailQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceCommoditydetailQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryCommoditydetail": map[string]interface{}{
				"commodityStr": "11774643024-0071043123",
				"picWidth":     "400",
				"picHeight":    "400",
				"couponMark":   "1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancecommoditydetailquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancecommoditydetailquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceSelectrecommendcommodityQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceSelectrecommendcommodityQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"querySelectrecommendcommodity": map[string]interface{}{
				"eliteId":     "1",
				"pageIndex":   "1",
				"size":        "10",
				"cityCode":    "025",
				"picWidth":    "400",
				"picHeight":   "400",
				"selfSupport": "0",
				"couponMark":  "1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetallianceselectrecommendcommodityquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetallianceselectrecommendcommodityquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceInverstmentcommodityQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceInverstmentcommodityQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryInverstmentcommodity": map[string]interface{}{
				"categoryId": "C1",
				"pageIndex":  "1",
				"size":       "10",
				"cityCode":   "025",
				"picWidth":   "400",
				"picHeight":  "400",
				"couponMark": "1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetallianceinverstmentcommodityquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetallianceinverstmentcommodityquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceSearchcommoditynewQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceSearchcommoditynewQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"querySearchcommoditynew": map[string]interface{}{
				"picWidth": "200",
				//"startPrice": "10.00",
				"picHeight": "200",
				//"coupon":    "1",
				"cityCode": "025",
				//"saleCategoryCode": "50000",
				//"endPrice": "20.00",
				//"branch": "1",
				//"suningService": "1",
				"size":      "50",
				"pageIndex": "1",
				//"snhwg": "1",
				"sortType": "1",
				//"snfwservice": "1",
				"pgSearch": "1",
				"keyword":  "抽屉",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancesearchcommoditynewquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancesearchcommoditynewquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceUnioninfomationGetRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceUnioninfomationGetRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"getUnionInfomation": map[string]interface{}{
				//"goodsCode": "10877475645",
				"goodsCode": "10642834202",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetallianceunioninfomationget2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetallianceunioninfomationget2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceCommodityimagesQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceCommodityimagesQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryCommodityimages": map[string]interface{}{
				////"goodsCode": "10877475645",
				//"goodsCode": "10642834202",
				"commodityCode": "10877475645",
				"supplierCode":  "0000000000",
				"terminalType":  "0",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancecommodityimagesquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancecommodityimagesquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceToolsellerAddRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceToolsellerAddRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"addToolseller": map[string]interface{}{
				"phone":        "18550502841",
				"channel":      "100001",
				"mediaCustNum": "user1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancetoolselleradd2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancetoolselleradd2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceCouponinfoQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceCouponinfoQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryCouponinfo": map[string]interface{}{
				//"couponNum": "202008240013988498",
				"quanUrl": "https://quan.suning.com/lqzx_recommend.do?activityId=202005130013070778&activitySecretKey=ySdjGtz4bkM7nZJXrI8rPBcT",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancecouponinfoquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancecouponinfoquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceRecommendcommodityQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceRecommendcommodityQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryRecommendcommodity": map[string]interface{}{
				"picWidth":   "200",
				"couponMark": "1",
				"picHeight":  "200",
				"size":       "10",
				"pageIndex":  "1",
				"cityCode":   "025",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancerecommendcommodityquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancerecommendcommodityquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceHoistinglinkQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceHoistinglinkQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryHoistinglink": map[string]interface{}{
				"adBookId":  "725735",
				"backurl":   "iqiyi://",
				"mertCode":  "0000000000",
				"goodsCode": "631698242",
				"subUser":   "user_1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancehoistinglinkquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancehoistinglinkquery2.Response)

		fmt.Println(result.Body)

	}
}

func TestSuningNetallianceOrderinfoQueryRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceOrderinfoQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryOrderinfo": map[string]interface{}{
				"orderId":  "32217250275",
				"pageNo":   "1",
				"pageSize": "10",
				//"startTime": "1599005820000",
				//"endTime":   "1599009420000",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetallianceorderinfoquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetallianceorderinfoquery2.Response)

		fmt.Println(result.Body)

	}
}

//相似商品
func TestSuningNetallianceMorerecommendGetRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceMorerecommendGetRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"getMorerecommend": map[string]interface{}{
				"picWidth":      "800",
				"picHeight":     "800",
				"supplierCode":  "0000000000",
				"commodityCode": "631698242",
				"picLocation":   "1",
				"picType":       "0",
				"cityCode":      "025",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancemorerecommendget2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancemorerecommendget2.Response)

		fmt.Println(result.Body)

	}
}

//相似商品
func TestSuningNetallianceSearchcommodityQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceSearchcommodityQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"querySearchcommodity": map[string]interface{}{
				"picWidth": "200",
				//"startPrice": "10.00",
				"picHeight": "200",
				//"coupon":    "1",
				"cityCode": "025",
				//"saleCategoryCode": "50000",
				//"endPrice": "20.00",
				//"branch": "1",
				//"suningService": "1",
				"size":      "40",
				"pageIndex": "2",
				//"snhwg": "1",
				"sortType": "1",
				//"snfwservice": "1",
				"pgSearch": "1",
				"keyword":  "抽屉",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancesearchcommodityquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancesearchcommodityquery2.Response)

		fmt.Println(result.Body)

	}
}

//相似商品
func TestSuningNetallianceCustompromotionurlQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceCustompromotionurlQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryCustompromotionurl": map[string]interface{}{
				//"visitUrl": "https://product.suning.com/0070661211/11419843713.html?utm_source=union&utm_medium=005999&adtype=5&utm_campaign=b3bd004f-c350-429d-a2cb-3c12e6e41c06&union_place=un",
				"visitUrl": "https://shop.suning.com/70139739/index.html?safp=d488778a.10004.0.09b964498d&safc=shop.0.0&safpn=10008.00045",
				"subUser":  "user_1",
			},
		},
	})

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancecustompromotionurlquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancecustompromotionurlquery2.Response)

		fmt.Println(result.Body)

	}
}

//相似商品
func TestSuningNetallianceStorepromotionurlQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceStorepromotionurlQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryStorepromotionurl": map[string]interface{}{
				"commCode": "",           //商品编码 如果商品编码不为空，则生成商品推广链接，否则生成店铺推广链接
				"mertCode": "0070057321", //商家编码 10位数字
				"subUser":  "Axxx",
				"adBookId": "725735",
			},
		},
	})

	//https://shop.suning.com/70139739/index.html?safp=d488778a.10004.0.09b964498d&safc=shop.0.0&safpn=10008.00045
	//https://shop.m.suning.com/70139739.html?safp=d488778a.10004.0.09b964498d&safc=shop.0.0&safpn=10008.00045

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancestorepromotionurlquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancestorepromotionurlquery2.Response)

		fmt.Println(result.Body)

	}
}

//相似商品
func TestSuningNetallianceBacthcustomlinkQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.SuningNetallianceBacthcustomlinkQueryRequest{}

	getRequest.AddParameter("sn_request", map[string]interface{}{
		"sn_body": map[string]interface{}{
			"queryBacthcustomlink": map[string]interface{}{
				//"extend":  "https://shop.suning.com/70139739/index.html?safp=d488778a.10004.0.09b964498d&safc=shop.0.0&safpn=10008.00045",
				"extend":  "https://product.suning.com/0070661211/11419843713.html?utm_source=union&utm_medium=005999&adtype=5&utm_campaign=b3bd004f-c350-429d-a2cb-3c12e6e41c06&union_place=un",
				"subUser": "1111111",
			},
		},
	})

	//https://shop.suning.com/70139739/index.html?safp=d488778a.10004.0.09b964498d&safc=shop.0.0&safpn=10008.00045
	//https://shop.m.suning.com/70139739.html?safp=d488778a.10004.0.09b964498d&safc=shop.0.0&safpn=10008.00045

	//初始化结果类型
	var getResponse DefaultResponse = &suningnetalliancebacthcustomlinkquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*suningnetalliancebacthcustomlinkquery2.Response)

		fmt.Println(result.Body)

	}
}
