package jdopensdk

import (
	"fmt"
	"github.com/mimicode/tksdk/jdopensdk/request"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsjingfenquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodspromotiongoodsinfoquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopengoodsquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenorderrowquery"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotionbysubunionidget"
	"github.com/mimicode/tksdk/jdopensdk/response/jdunionopenpromotioncommonget"
	"testing"
)

const (
	appKey     = ""
	appSecret  = ""
	sessionKey = ""
)

func GetClient() *TopClient {
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)
	return client
}

func TestJdUnionOpenGoodsJingfenQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsJingfenQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"goodsReq":{"eliteId":"9999","pid":"","sortName":"commission","sort":"desc","pageSize":"50","pageIndex":"1"}}`)
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
