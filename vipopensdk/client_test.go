package vipopensdk

import (
	"encoding/json"
	"fmt"
	utils2 "github.com/mimicode/tksdk/utils"
	request2 "github.com/mimicode/tksdk/vipopensdk/request"
	"github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsservicegetbygoodsidsv2withoauth"
	comvipadpapiopenserviceuniongoodsservicegetbygoodsidswithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsservicegetbygoodsidswithoauth"
	"github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsservicegoodslistv2withoauth"
	comvipadpapiopenserviceuniongoodsservicegoodslistwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsservicegoodslistwithoauth"
	comvipadpapiopenserviceuniongoodsservicequerywithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsservicequerywithoauth"
	comvipadpapiopenserviceuniongoodsservicesimilarrecommendwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsservicesimilarrecommendwithoauth"
	comvipadpapiopenserviceuniongoodsserviceuserrecommendwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceuniongoodsserviceuserrecommendwithoauth"
	comvipadpapiopenserviceunionorderserviceorderlistwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceunionorderserviceorderlistwithoauth"
	comvipadpapiopenserviceunionorderservicerefundorderlistwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceunionorderservicerefundorderlistwithoauth"
	comvipadpapiopenserviceunionurlservicegenbygoodsidwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceunionurlservicegenbygoodsidwithoauth"
	comvipadpapiopenserviceunionurlservicegenbyvipurlwithoauth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceunionurlservicegenbyvipurlwithoauth"
	comvipadpapiopenserviceunionurlserviceviplinkcheckwithouth2 "github.com/mimicode/tksdk/vipopensdk/response/comvipadpapiopenserviceunionurlserviceviplinkcheckwithouth"
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
				Vip struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
				} `json:"vip"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Vip.AppKey
				appSecret = data.Vip.AppSecret
				sessionKey = data.Vip.SessionKey
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

func TestComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionGoodsServiceGoodsListWithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":                   utils2.GetUUID(),
		"channelType":                 0,
		"page":                        1,
		"pageSize":                    100,
		"queryReputation":             false,
		"queryStoreServiceCapability": false,
		"fieldName":                   "COMMISSION",
		"order":                       0,
		"sourceType":                  1,
		"jxCode":                      "7hfpy0m4",
		"queryStock":                  false,
		"queryActivity":               false,
		"commonParams": map[string]interface{}{
			"plat": 2,
		},
	})

	/*
		   channelType	Integer	否			频道类型:0-超高佣，1-出单爆款; 当请求类型为频道时必传
		   page	Integer	是			页码
		   pageSize	Integer	否			分页大小:默认20，最大100
		   requestId	String	是			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
		   queryReputation	Boolean	否			是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
		   queryStoreServiceCapability	Boolean	否			是否查询店铺服务能力信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
		   sourceType	Integer	否			请求源类型：0-频道，1-组货


				类目                 API组货ID
				女装精选 7hfpy0m4
				男装精选 wj7evz2j
				美妆精选 vd0wbfdx
				数码电子 dpot8m5u
				精选-首饰 szkl4kj7
				鞋包精选 byh9331t
				母婴精选 gkf52p8p
				居家精选 cnrzcs22
				运动户外精选 indvf44e
				家用电器 uggxpyh5


		   jxCode	String	否			精选组货码：当请求源类型为组货时必传
		   fieldName	String	否			排序字段: COMMISSION-佣金，PRICE-价格,COMM_RATIO-佣金比例，DISCOUNT-折扣
		   order	Integer	否			排序顺序：0-正序，1-逆序，默认正序
		   queryStock	Boolean	否			是否查询库存:默认不查询
		   queryActivity	Boolean	否			是否查询商品活动信息:默认不查询
		   queryPrepay	Boolean	否			是否查询商品预付信息:默认不查询

		   commonParams	CommonParams	否			通用参数

				plat	Integer	否			用户平台：1-PC,2-APP,3-小程序,不传默认为APP
				deviceType	String	否			设备号类型：IMEI，IDFA，OAID，有则传入
				deviceValue	String	否			设备号MD5加密后的值，有则传入
				ip	String	否			用户ip地址
				longitude	String	否			经度 如:29.590961456298828
				latitude	String	否			纬度 如:106.51573181152344


		   vendorCode	String	否			工具商code
		   chanTag	String	否			pid
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsservicegoodslistwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsservicegoodslistwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionGoodsServiceGetByGoodsIdsWithOauthRequest{}

	getRequest.AddParameter("requestId", utils2.GetUUID())
	getRequest.AddParameter("chanTag", "user_1")
	getRequest.AddParameter("goodsIdList", []string{
		"6919190827778651268",
	})

	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsservicegetbygoodsidswithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsservicegetbygoodsidswithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionGoodsServiceQueryWithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":                   utils2.GetUUID(),
		"keyword":                     "南孚 7号电池",
		"page":                        1,
		"pageSize":                    100,
		"queryReputation":             false,
		"queryStoreServiceCapability": false,
		"fieldName":                   "PRICE",
		"order":                       0,
		"queryStock":                  false,
		"queryPrepay":                 false,
	})
	/*
		keyword	String	是			关键词
		fieldName	String	否			排序字段
		order	Integer	否			排序顺序：0-正序，1-逆序，默认正序
		page	Integer	是			页码
		pageSize	Integer	否			页面大小：默认20,最大50
		requestId	String	是			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
		priceStart	String	否			价格区间---start
		priceEnd	String	否			价格区间---end
		queryReputation	Boolean	否			是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询，影响接口性能
		queryStoreServiceCapability	Boolean	否			是否查询店铺服务能力信息：默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询，影响接口性能
		queryStock	Boolean	否			是否查询库存:默认不查询
		queryActivity	Boolean	否			是否查询商品活动信息:默认不查询
		queryPrepay	Boolean	否			是否查询商品预付信息:默认不查询
		commonParams	CommonParams	否			通用参数
		vendorCode	String	否			工具商code
		chanTag	String	否			用户pid
	*/

	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsservicequerywithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsservicequerywithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionGoodsServiceSimilarRecommendWithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":                   utils2.GetUUID(),
		"goodsId":                     "2165780149",
		"page":                        1,
		"pageSize":                    100,
		"queryReputation":             false,
		"queryStoreServiceCapability": false,
		"queryStock":                  false,
		"queryPrepay":                 false,
		"commonParams":                map[string]interface{}{},
	})

	/*
		goodsId	String	是			商品id
		page	Integer	是			分页页码：从1开始
		pageSize	Integer	否			分页大小：默认20
		requestId	String	是			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
		commonParams	CommonParams	否			通用参数：能获取到则须传入
			plat	Integer	否			用户平台：1-PC,2-APP,3-小程序,不传默认为APP
			deviceType	String	否			设备号类型：IMEI，IDFA，OAID，有则传入
			deviceValue	String	否			设备号MD5加密后的值，有则传入
			ip	String	否			用户ip地址
			longitude	String	否			经度 如:29.590961456298828
			latitude	String	否			纬度 如:106.51573181152344
		queryReputation	Boolean	否			是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询，影响接口性能
		queryStoreServiceCapability	Boolean	否			是否查询店铺服务能力信息：默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询，影响接口性能
		queryStock	Boolean	否			是否查询库存:默认不查询
		queryActivity	Boolean	否			是否查询商品活动信息:默认不查询
		queryPrepay	Boolean	否			是否查询商品预付信息:默认不查询
		chanTag	String	否			自定义渠道标识
	*/

	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsservicesimilarrecommendwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsservicesimilarrecommendwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionGoodsServiceUserRecommendWithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":    utils2.GetUUID(),
		"page":         1,
		"pageSize":     100,
		"commonParams": map[string]interface{}{},
	})

	/*
		page	Integer	是			分页页码：从1开始
		pageSize	Integer	否			分页大小：默认20
		requestId	String	否			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
		inStock	Integer	否			是否有货 1:有货 0:无货 默认1
		goodsSaleStats	Integer	否			商品售卖状态 1（在售）， 2（预热）， 3（在售+预热） 默认1
		offlineStore	Integer	否			筛选线下店商品：1只返线下店，0或者不传只返回特卖会
		chanTag	String	否			自定义渠道标识
		commonParams	CommonParams	否			通用参数：能获取到则须传入
			plat	Integer	否			用户平台：1-PC,2-APP,3-小程序,不传默认为APP
			deviceType	String	否			设备号类型：IMEI，IDFA，OAID，有则传入
			deviceValue	String	否			设备号MD5加密后的值，有则传入
			ip	String	否			用户ip地址
			longitude	String	否			经度 如:29.590961456298828
			latitude	String	否			纬度 如:106.51573181152344

	*/

	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsserviceuserrecommendwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsserviceuserrecommendwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionUrlServiceGenByGoodsIdWithOauthRequest{}
	getRequest.AddParameter("requestId", utils2.GetUUID())
	getRequest.AddParameter("chanTag", "user_1")
	getRequest.AddParameter("goodsIdList", []string{
		"6919190827778651268",
	})

	/*
	   goodsIdList	List<String>	是			商品id列表
	   chanTag	String	否			渠道标识
	   requestId	String	是			请求id：UUID
	   statParam	String	否			自定义渠道统计参数
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceunionurlservicegenbygoodsidwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceunionurlservicegenbygoodsidwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionUrlServiceGenByVIPUrlWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionUrlServiceGenByVIPUrlWithOauthRequest{}
	getRequest.AddParameter("requestId", utils2.GetUUID())
	getRequest.AddParameter("chanTag", "user_1")
	getRequest.AddParameter("urlList", []string{
		//"https://t.vip.com/Q9wouVgHhw9?&desturl=https%3A%2F%2Fm.vip.com%2Fbrand-100741036-0-0-0-1-0-1-20.html&nmsns=shop_android-7.27.9-link&nst=schedule&nsbc=&nct=link&ncid=0035028e-2abc-3238-8de0-9f903a860d75&nabtid=be318beb41a3ef10b1ea4cb350d6adc5&nuid=MjU3NzU3MjE5&nchl_param=share:0035028e-2abc-3238-8de0-9f903a860d75:",
		//"https://t.vip.com/WrqkH2rrPuA?&desturl=https%3A%2F%2Fm.vip.com%2Fproduct-1710614059-6918870462465231563.html%3Fnmsns%3Dshop_android-7.27.9-link%26nst%3Dproduct%26nsbc%3D%26nct%3Dlink%26ncid%3D0035028e-2abc-3238-8de0-9f903a860d75%26nabtid%3D13%26nuid%3D257757219%26nchl_param%3Dshare%3A0035028e-2abc-3238-8de0-9f903a860d75%3A1599366894023&nmsns=shop_android-7.27.9-link&nst=product&nsbc=&nct=link&ncid=0035028e-2abc-3238-8de0-9f903a860d75&nabtid=13&nuid=257757219&nchl_param=share:0035028e-2abc-3238-8de0-9f903a860d75:1599366894023",
		//"https://t.vip.com/5NwgBeslh96",
		//"https://t.vip.com/iaLxayqVEZ7",
		//"https://t.vip.com/NO9iXRQSnK7",
		"https://click.union.vip.com/fqSIl10Y9L8?",
	})

	/*
		urlList	List<String>	是			商品链接url列表
		chanTag	String	否			渠道标识
		requestId	String	是			请求id：UUID
		statParam	String	否			自定义渠道统计参数
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceunionurlservicegenbyvipurlwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceunionurlservicegenbyvipurlwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionUrlServiceVipLinkCheckWithOuthRequest{}
	//getRequest.AddParameter("requestId", utils.GetUUID())
	//getRequest.AddParameter("chanTag", "user_1")
	getRequest.AddParameter("vipLinkCheckReq", map[string]interface{}{
		"source":  "weixin",
		"content": "很不错的换一个购买 https://www.vip.com/detail-1710614067-6917924644994118739.html 这个也是不错的 https://t.vip.com/Ma2qFqMX5e9?chanTag=user_1&wq=1 aaa https://www.baidu.com/",
	})

	/*
		urlList	List<String>	是			商品链接url列表
		chanTag	String	否			渠道标识
		requestId	String	是			请求id：UUID
		statParam	String	否			自定义渠道统计参数
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceunionurlserviceviplinkcheckwithouth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceunionurlserviceviplinkcheckwithouth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionOrderServiceOrderListWithOauthRequest{}
	//getRequest.AddParameter("requestId", utils.GetUUID())
	//getRequest.AddParameter("chanTag", "user_1")
	getRequest.AddParameter("queryModel", map[string]interface{}{
		"requestId": utils2.GetUUID(),
		"page":      1,
		"pageSize":  20,
		//"orderSnList": []string{
		//	"20090194619309",
		//	"20090179074086",
		//	"20090242691449",
		//},
		"orderTimeStart": 1598952600000,
		"orderTimeEnd":   1598956200000,
		//"updateTimeStart": 1598952600000,
		//"updateTimeEnd":   1598956200000,
	})

	/*
			status	Short	否			订单状态:0-不合格，1-待定，2-已完结，该参数不设置默认代表全部状态
			orderTimeStart	Long	否			订单时间起始 时间戳 单位毫秒
			orderTimeEnd	Long	否			订单时间结束 时间戳 单位毫秒
			page	Integer	是			页码：从1开始
			pageSize	Integer	否			页面大小：默认20
			requestId	String	是			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
			updateTimeStart	Long	否			更新时间-起始 时间戳 单位毫秒
			updateTimeEnd	Long	否			更新时间-结束 时间戳 单位毫秒
			orderSnList	List<String>	否			订单号列表：当传入订单号列表时，订单时间和更新时间区间可不传入
			vendorCode	String	否			vendorCode,工具商方式下会传入 暂时不传入
			chanTag	String	否			渠道标识，即推广位PID
		updateTimeStart -> 1598952201283
		updateTimeEnd   -> 1598953201283
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceunionorderserviceorderlistwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceunionorderserviceorderlistwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenService5nionOrderServiceRefundOrderListWithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":  utils2.GetUUID(),
		"page":       1,
		"pageSize":   20,
		"searchType": 0,
		"orderSns": []string{
			"20090194619309",
			"20090179074086",
			"20090242691449",
		},
		//"searchTimeStart": 1598952600000,
		//"searchTimeEnd":   1598956200000,
	})

	/*
		searchType	Integer	是			查询类型:0-维权发起时间，1-维权完成时间（佣金扣除入账时间），2-佣金扣除结算时间
		searchTimeStart	Long	否			目标时间起始：时间戳，单位毫秒; 当searchType=0时，为维权发起时间， 当searchType=1时，为维权完成时间（佣金扣除入账时间）， 当searchType=2时，为佣金扣除结算时间
		searchTimeEnd	Long	否			目标时间结束：时间戳，单位毫秒; 当searchType=0时，为维权发起时间， 当searchType=1时，为维权完成时间（佣金扣除入账时间）， 当searchType=2时，为佣金扣除结算时间
		orderSns	List<String>	否			目标订单号集合:当指定订单号集合时，目标时间区间可以不传,否则必须指定目标时间区间
		page	Integer	是			当前页
		pageSize	Integer	否		20	分页大小：默认20，最大100
		requestId	String	是			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceunionorderservicerefundorderlistwithoauth2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceunionorderservicerefundorderlistwithoauth2.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenServiceUnionGoodsServiceGetByGoodsIdsV2WithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":   utils2.GetUUID(),
		"queryDetail": true,
		"goodsIds": []string{
			"1624525595900175",
			"1350181608052047",
			"1624531048610063",
		},
		"queryStock":                  true,
		"queryReputation":             true,
		"queryStoreServiceCapability": true,
		"queryPMSAct":                 true,
		"queryPrepay":                 true,
		"extendBySpu":                 true,
		"queryExclusiveCoupon":        true,
		"extendSku":                   true,
		"queryCpsInfo":                3,
		"queryFuturePrice":            true,
		"querySubsidyActFlag":         true,
	})

	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsservicegetbygoodsidsv2withoauth.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsservicegetbygoodsidsv2withoauth.Response)

		fmt.Println(result.Body)

	}
}

func TestComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.ComVipAdpApiOpenServiceUnionGoodsServiceGoodsListV2WithOauthRequest{}

	getRequest.AddParameter("request", map[string]interface{}{
		"requestId":                   utils2.GetUUID(),
		"channelType":                 0,
		"offset":                      1,
		"pageSize":                    50,
		"queryReputation":             false,
		"queryStoreServiceCapability": false,
		"fieldName":                   "COMMISSION",
		"order":                       0,
		"sourceType":                  1,
		"jxCode":                      "dz5d5n7i",
		"queryStock":                  false,
		"queryActivity":               false,
		"commonParams": map[string]interface{}{
			"plat": 2,
		},
	})

	/*

		   channelType	Integer	否			频道类型:0-超高佣，1-出单爆款; 当请求类型为频道时必传
		   page	Integer	是			页码
		   pageSize	Integer	否			分页大小:默认20，最大100
		   requestId	String	是			请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
		   queryReputation	Boolean	否			是否查询商品评价信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
		   queryStoreServiceCapability	Boolean	否			是否查询店铺服务能力信息:默认不查询，该数据在详情页有返回，没有特殊需求，建议不查询
		   sourceType	Integer	否			请求源类型：0-频道，1-组货


				类目                 API组货ID
				女装精选 7hfpy0m4
				男装精选 wj7evz2j
				美妆精选 vd0wbfdx
				数码电子 dpot8m5u
				精选-首饰 szkl4kj7
				鞋包精选 byh9331t
				母婴精选 gkf52p8p
				居家精选 cnrzcs22
				运动户外精选 indvf44e
				家用电器 uggxpyh5

				2小时排行榜  dz5d5n7i   注意获取第二页时要传batchNo参数，防止商品重复
				今日热销榜  7heizb6v    注意获取第二页时要传batchNo参数，防止商品重复

		   jxCode	String	否			精选组货码：当请求源类型为组货时必传
		   fieldName	String	否			排序字段: COMMISSION-佣金，PRICE-价格,COMM_RATIO-佣金比例，DISCOUNT-折扣
		   order	Integer	否			排序顺序：0-正序，1-逆序，默认正序
		   queryStock	Boolean	否			是否查询库存:默认不查询
		   queryActivity	Boolean	否			是否查询商品活动信息:默认不查询
		   queryPrepay	Boolean	否			是否查询商品预付信息:默认不查询

		   commonParams	CommonParams	否			通用参数

				plat	Integer	否			用户平台：1-PC,2-APP,3-小程序,不传默认为APP
				deviceType	String	否			设备号类型：IMEI，IDFA，OAID，有则传入
				deviceValue	String	否			设备号MD5加密后的值，有则传入
				ip	String	否			用户ip地址
				longitude	String	否			经度 如:29.590961456298828
				latitude	String	否			纬度 如:106.51573181152344


		   vendorCode	String	否			工具商code
		   chanTag	String	否			pid
	*/
	//初始化结果类型
	var getResponse DefaultResponse = &comvipadpapiopenserviceuniongoodsservicegoodslistv2withoauth.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*comvipadpapiopenserviceuniongoodsservicegoodslistv2withoauth.Response)

		fmt.Println(result.Body)

	}
}
