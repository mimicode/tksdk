package alscopensdk

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/alscopensdk/request"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionelemepromotionofficialactivityget"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionelemepromotionstorepromotionget"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionelemepromotionstorepromotionquery"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionkbitempromotion"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionkbitempromotionfilterlist"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionkbitemquery"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionmediazoneadd"
	"github.com/mimicode/tksdk/alscopensdk/response/alibabaalscunionmediazoneget"

	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	appKey, appSecret, sessionKey string
	pid                           struct {
		AccountID string
		SiteID    string
		AdzoneID  string
		SID       string
		WholePID  string
	}
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Ele struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
					Pid        string `json:"pid"`
					SID        string `json:"sid"`
				} `json:"ele"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Ele.AppKey
				appSecret = data.Ele.AppSecret
				sessionKey = data.Ele.SessionKey
				pid.SID = data.Ele.SID
				split := strings.Split(strings.TrimPrefix(data.Ele.Pid, "alsc_"), "_")
				if len(split) == 3 {
					pid.AccountID = split[0]
					pid.SiteID = split[1]
					pid.AdzoneID = split[2]
					pid.WholePID = data.Ele.Pid
				}

			}
		}
	}
}
func TestAlscUnionKbItemPromotionRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionKbItemPromotionRequest{}
	getRequest.AddParameter("page_number", "1")
	getRequest.AddParameter("page_size", "20")
	getRequest.AddParameter("pid", pid.WholePID)
	getRequest.AddParameter("settle_type", "2")
	getRequest.AddParameter("item_type", "3")
	//biz_scene_id
	//promotion_type
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionkbitempromotion.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionkbitempromotion.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			fmt.Println(result.Body)
		} else {
			items := result.ResultResponse.Items.KBItemPromotionDTO
			for _, v := range items {
				fmt.Println(v.Title)
			}
		}

	}
}

func TestAlibabaAlscUnionMediaZoneAddRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionMediaZoneAddRequest{}
	getRequest.AddParameter("zone_name", "apitest")
	//biz_scene_id
	//promotion_type
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionmediazoneadd.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionmediazoneadd.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {
			items := result.ResultResponse.Result
			t.Logf("name:%s     pid:%s", items.Name, items.Pid)
		}

	}
}

func TestAlibabaAlscUnionKbItemPromotionFilterListRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionKbItemPromotionFilterListRequest{}
	getRequest.AddParameter("filter_type", "category")
	//biz_scene_id
	//promotion_type
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionkbitempromotionfilterlist.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionkbitempromotionfilterlist.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {
			items := result.ResultResponse.Result.FilterTableNameDTO
			for i, item := range items {
				t.Logf("%d - FilterKey: %s - FilterName: %s", i, item.FilterKey, item.FilterName)
			}
		}

	}
}

func TestAlibabaAlscUnionMediaZoneGetRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionMediaZoneGetRequest{}
	getRequest.AddParameter("page", "1")
	getRequest.AddParameter("limit", "30")
	//biz_scene_id
	//promotion_type
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionmediazoneget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionmediazoneget.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {
			items := result.ResultResponse.Result.ZoneInfoDTO
			for i, item := range items {
				t.Logf("%d - Name: %s - Pid: %s", i, item.Name, item.Pid)
			}
		}

	}
}

func TestAlibabaAlscUnionElemePromotionStorepromotionQueryRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionElemePromotionStorepromotionQueryRequest{}
	var data = map[string]string{
		"pid":       pid.WholePID,
		"longitude": "116.77445311758802",
		"latitude":  "39.9494466415974",
	}
	marshal, _ := json.Marshal(data)
	getRequest.AddParameter("query_request", string(marshal))
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionelemepromotionstorepromotionquery.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionelemepromotionstorepromotionquery.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {

		}

	}
}

func TestAlibabaAlscUnionElemePromotionStorepromotionGetRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionElemePromotionStorepromotionGetRequest{}
	var data = map[string]string{
		"pid":     pid.WholePID,
		"shop_id": "E3072025203382751092",
		//"activity_id": "10144",
		"sid": pid.SID,
	}
	marshal, _ := json.Marshal(data)
	getRequest.AddParameter("query_request", string(marshal))
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionelemepromotionstorepromotionget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionelemepromotionstorepromotionget.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {

		}

	}
}

func TestAlibabaAlscUnionElemePromotionOfficialactivityGetRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionElemePromotionOfficialactivityGetRequest{}
	var data = map[string]string{
		"pid":             pid.WholePID,
		"activity_id":     "10144",
		"sid":             pid.SID,
		"include_wx_img":  "true",
		"include_qr_code": "true",
	}
	marshal, _ := json.Marshal(data)
	getRequest.AddParameter("query_request", string(marshal))
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionelemepromotionofficialactivityget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionelemepromotionofficialactivityget.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {

		}

	}
}

func TestAlibabaAlscUnionKbItemQueryRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.AlibabaAlscUnionKbItemQueryRequest{}

	getRequest.AddParameter("page_number", "1")
	getRequest.AddParameter("page_size", "20")
	getRequest.AddParameter("biz_type", "kb_natural")
	getRequest.AddParameter("pid", pid.WholePID)
	//初始化结果类型
	var getResponse DefaultResponse = &alibabaalscunionkbitemquery.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*alibabaalscunionkbitemquery.Response)

		if result.IsError() || result.ResultResponse.IsError() {
			t.Log(result.Body)
		} else {

		}

	}
}
