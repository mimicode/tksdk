package jdopensdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mimicode/tksdk/jdopensdk/request"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopencoupongiftget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopencoupongiftstop"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsbigfieldquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsjingfenquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodspromotiongoodsinfoquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenorderrowquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpositioncreate"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpositionquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotionbysubunionidget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotioncommonget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenstatisticsgiftcouponquery"
)

var (
	appKey, appSecret, sessionKey, key string
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Jd struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
					Key        string `json:"key"`
				} `json:"jd"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Jd.AppKey
				appSecret = data.Jd.AppSecret
				sessionKey = data.Jd.SessionKey
				key = data.Jd.Key
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

func TestJdUnionOpenGoodsJingfenQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsJingfenQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"goodsReq":{"eliteId":"1","pid":"","sortName":"commission","sort":"desc","pageSize":"50","pageIndex":"1"}}`)
	var getResponse DefaultResponse = &jdunionopengoodsjingfenquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodsjingfenquery.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPromotionCommonGetRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPromotionCommonGetRequest{}
	getRequest.AddParameter("360buy_param_json", `{"promotionCodeReq":{"materialId":"https://item.jd.com/10030840282202.html","siteId":"223997188","positionId":"1631770896","subUnionId":"user_1","couponUrl":"","giftCouponKey":"","ext1":"user_1"}}`)
	var getResponse DefaultResponse = &jdunionopenpromotioncommonget.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpromotioncommonget.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPromotionBysubunionidGetRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPromotionBysubunionidGetRequest{}
	getRequest.AddParameter("360buy_param_json", `{"promotionCodeReq":{"materialId":"https://item.jd.com/10030840282202.html", "positionId":"1631770896","subUnionId":"user_1","couponUrl":"","giftCouponKey":"","chainType":"3"}}`)
	var getResponse DefaultResponse = &jdunionopenpromotionbysubunionidget.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpromotionbysubunionidget.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenOrderRowQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenOrderRowQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"orderReq":{"pageIndex":"1", "pageSize":"20","type":"3","startTime":"2021-08-02 11:45:00","endTime":"2021-08-02 12:45:00"}}`)
	var getResponse DefaultResponse = &jdunionopenorderrowquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenorderrowquery.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenGoodsPromotiongoodsinfoQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsPromotiongoodsinfoQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{skuIds:"10030840282202"}`)
	var getResponse DefaultResponse = &jdunionopengoodspromotiongoodsinfoquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodspromotiongoodsinfoquery.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenGoodsQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsQueryRequest{}
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"skuIds":[10045819630288],"fields":"videoInfo"}}`)
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"iPhone+13"}}`)
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"小米+13"}}`)
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"民法典"}}`)
	getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"空调"}}`)
	var getResponse DefaultResponse = &jdunionopengoodsquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodsquery.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPositionQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPositionQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"positionReq":{"unionId":103431087,"pageIndex":1,"pageSize":100,"key":"`+key+`","unionType":4}}`)
	var getResponse DefaultResponse = &jdunionopenpositionquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpositionquery.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPositionCreateRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPositionCreateRequest{}
	getRequest.AddParameter("360buy_param_json", `{"positionReq":{"unionId":103431087,"key":"`+key+`","unionType":3,"type":4,"spaceNameList":["测试专用-01"]}}`)
	var getResponse DefaultResponse = &jdunionopenpositioncreate.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpositioncreate.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
		// {"jd_union_open_position_create_responce":{"code":"0","createResult":"{\"code\":200,\"data\":{\"resultList\":{\"测试专用-01\":21001775614},\"siteId\":0,\"type\":4,\"unionId\":103431087},\"message\":\"success\",\"requestId\":\"o_0b126c7d_l531g96m_4147537\"}"}}
	}
}

func TestJdUnionOpenCouponGiftGetRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenCouponGiftGetRequest{}
	var param = map[string]interface{}{
		"couponReq": map[string]interface{}{
			"couponTitle":      "礼金测试01",
			"skuMaterialId":    "https://item.jd.com/10053130147978.html",
			"discount":         1,
			"amount":           20, //最低20个
			"receiveStartTime": "2023-09-16 07",
			"receiveEndTime":   "2023-09-16 23",
			"expireType":       2,
			"useStartTime":     "2023-09-16",
			"useEndTime":       "2023-09-16",
			"share":            1,
			"contentMatch":     0,
			"isSpu":            1,
			//"targetType":"" //定向推广类型，默认为1，向【运营申请-真操蛋，啥玩意儿都要申请】定向功能后才能入参4,5,6并生效； 1.不定向推广 4:本账号推广-定向PID 5：合作账号推广-定向联盟ID 6:合作账号推广-定向PID
			//"childPromoters": "[{ 'targetName': '达人A', 'targetValue': '1002132107_4100446224_3003707911' }]",
		},
	}
	marshal, _ := json.Marshal(param)
	getRequest.AddParameter("360buy_param_json", string(marshal))
	var getResponse DefaultResponse = &jdunionopencoupongiftget.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopencoupongiftget.Response)
		fmt.Println(commonGetResponse.IsError())
		//{"jd_union_open_coupon_gift_get_responce":{"code":"0","getResult":"{\"code\":200,\"data\":{\"giftCouponKey\":\"9f622b45acaff312\"},\"message\":\"success\",\"requestId\":\"o_0b64a07b_lmlzrjqb_153015728\"}"}}
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenStatisticsGiftcouponQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenStatisticsGiftcouponQueryRequest{}
	var param = map[string]interface{}{
		"effectDataReq": map[string]interface{}{
			"giftCouponKey": "9f622b45acaff312",
		},
	}
	marshal, _ := json.Marshal(param)
	getRequest.AddParameter("360buy_param_json", string(marshal))
	var getResponse DefaultResponse = &jdunionopenstatisticsgiftcouponquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenstatisticsgiftcouponquery.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
		//{"jd_union_open_statistics_giftcoupon_query_responce":{"code":"0","queryResult":"{\"code\":200,\"data\":[{\"amount\":20,\"contentMatch\":0,\"contentMatchMedias\":[-1],\"costAmount\":0.00,\"costNum\":0,\"couponTitle\":\"礼金测试01\",\"denomination\":1.0,\"effectiveDays\":0,\"expireType\":2,\"giftCouponKey\":\"9f622b45acaff312\",\"promoterList\":[{\"amount\":20,\"pid\":\"\",\"unionId\":103431087}],\"receiveEndTime\":\"2023-09-16 23:59:59\",\"receiveNum\":0,\"receiveStartTime\":\"2023-09-16 07:00:00\",\"share\":1,\"showInMedias\":0,\"showStatus\":3,\"skuIdList\":[10053130147978,10053130147979,10080660235974,10053130147977],\"status\":1,\"type\":2,\"useEndTime\":\"2023-09-16 23:59:59\",\"useStartTime\":\"2023-09-16 00:00:00\",\"ygCommission\":0.00}],\"message\":\"success\",\"requestId\":\"o_0b9180ed_lmlzyq8g_172473927\"}"}}
	}
}

func TestJdUnionOpenCouponGiftStopRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenCouponGiftStopRequest{}
	var param = map[string]interface{}{
		"couponReq": map[string]interface{}{
			"giftCouponKey": "9f622b45acaff312",
		},
	}
	marshal, _ := json.Marshal(param)
	getRequest.AddParameter("360buy_param_json", string(marshal))
	var getResponse DefaultResponse = &jdunionopencoupongiftstop.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopencoupongiftstop.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenGoodsBigfieldQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsBigfieldQueryRequest{}
	var param = map[string]interface{}{
		"goodsReq": map[string]interface{}{
			"skuIds":  []interface{}{},
			"itemIds": []interface{}{"VgDXlT9hVVVmDDiCbofTFhV7_VIfTFhV7VVyGGPNs"},
			"sceneId": 1,
		},
	}
	marshal, _ := json.Marshal(param)
	getRequest.AddParameter("360buy_param_json", string(marshal))
	var getResponse DefaultResponse = &jdunionopengoodsbigfieldquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodsbigfieldquery.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}
