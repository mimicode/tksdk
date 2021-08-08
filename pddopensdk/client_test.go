package pddopensdk

import (
	request2 "github.com/mimicode/tksdk/pddopensdk/request"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkallorderlistincrementget"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkgoodsdetail"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkgoodspromotionurlgenerate"
	pddddkgoodsrecommendget2 "github.com/mimicode/tksdk/pddopensdk/response/pddddkgoodsrecommendget"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkgoodssearch"
	pddddkgoodszsuniturlgen2 "github.com/mimicode/tksdk/pddopensdk/response/pddddkgoodszsuniturlgen"
	pddddklotteryurlgen2 "github.com/mimicode/tksdk/pddopensdk/response/pddddklotteryurlgen"
	pddddkmemberauthorityquery2 "github.com/mimicode/tksdk/pddopensdk/response/pddddkmemberauthorityquery"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthcmspromurlgenerate"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodsdetail"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodspidgenerate"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodspidquery"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodspromurlgenerate"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodsrecommendget"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodssearch"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthgoodszsuniturlgen"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthmemberauthorityquery"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthorderdetailget"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthresourceurlgen"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauthrppromurlgenerate"
	"github.com/mimicode/tksdk/pddopensdk/response/pddddkoauththemepromurlgenerate"

	"fmt"
	pddddkorderlistincrementget2 "github.com/mimicode/tksdk/pddopensdk/response/pddddkorderlistincrementget"
	pddddkresourceurlgen2 "github.com/mimicode/tksdk/pddopensdk/response/pddddkresourceurlgen"
	pddddkrppromurlgenerate2 "github.com/mimicode/tksdk/pddopensdk/response/pddddkrppromurlgenerate"
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

func TestPddDdkGoodsDetailRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkGoodsDetailRequest{}
	//getRequest.AddParameter("goods_id_list", `["4004166635"]`)
	getRequest.AddParameter("goods_sign", `c9L2jNyxKEdGV49RwfDYkFgHAcOx_JgDQryG2X`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkgoodsdetail.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkgoodsdetail.Response)

		fmt.Println(result.GoodsDetailResponse.GoodsDetails)

	}
}

func TestPddDdkGoodsPromotionUrlGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkGoodsPromotionUrlGenerateRequest{}
	getRequest.AddParameter("custom_parameters", `user_1`)
	getRequest.AddParameter("p_id", `1949211_49329093`)
	getRequest.AddParameter("goods_sign", `c9L2jNyxKEdGV49RwfDYkFgHAcOx_JgDQryG2X`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkgoodspromotionurlgenerate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkgoodspromotionurlgenerate.Response)

		fmt.Println(result)

	}
}

func TestPddDdkGoodsSearchRequest(t *testing.T) {

	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkGoodsSearchRequest{}
	getRequest.AddParameter("keyword", `女装大码`)
	getRequest.AddParameter("page_size", `100`)
	getRequest.AddParameter("custom_parameters", `user_1`)
	getRequest.AddParameter("pid", `1949211_49329093`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkgoodssearch.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkgoodssearch.Response)

		fmt.Println(result.GoodsSearchResponse.GoodsList)

	}
}

func TestPddDdkOrderListIncrementGetRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOrderListIncrementGetRequest{}
	getRequest.AddParameter("start_update_time", `1594389599`)
	getRequest.AddParameter("end_update_time", `1594396799`)
	getRequest.AddParameter("page_size", "40")
	getRequest.AddParameter("page", "1")
	getRequest.AddParameter("return_count", "true")

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkorderlistincrementget2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkorderlistincrementget2.Response)

		fmt.Println(result.Body)

	}
}

func TestPddddklotteryurlgen(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkLotteryUrlGenRequest{}
	getRequest.AddParameter("pid_list", `["1949211_146717471"]`)
	getRequest.AddParameter("generate_short_url", `true`)
	getRequest.AddParameter("generate_schema_url", `true`)
	getRequest.AddParameter("custom_parameters", `custom_o1`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddklotteryurlgen2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddklotteryurlgen2.Response)

		fmt.Println(result.Body)

	}
}

func TestPddDdkRpPromUrlGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkRpPromUrlGenerateRequest{}
	getRequest.AddParameter("p_id_list", `["1949211_146717471"]`)
	getRequest.AddParameter("generate_short_url", `true`)
	getRequest.AddParameter("generate_schema_url", `true`)
	getRequest.AddParameter("custom_parameters", `custom_o1`)
	getRequest.AddParameter("channel_type", `-1`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkrppromurlgenerate2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkrppromurlgenerate2.Response)

		fmt.Println(result.Body)

	}
}

func TestPddDdkResourceUrlGenRequest(t *testing.T) {
	////4-限时秒杀,39997-充值中心, 39998-转链type，39996-百亿补贴
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkResourceUrlGenRequest{}
	getRequest.AddParameter("pid", `1949211_146717471`)
	getRequest.AddParameter("generate_schema_url", `true`)
	getRequest.AddParameter("custom_parameters", `custom_o1`)
	getRequest.AddParameter("resource_type", `39998`)
	getRequest.AddParameter("url", `https://mobile.yangkeduo.com/attendance.html?_pdd_fs=1&_pdd_tc=ffffff&_pdd_sbs=1&type=1&id=155075`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkresourceurlgen2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkresourceurlgen2.Response)

		fmt.Println(result.Body)

	}
}

//func TestPddDdkResourceUrlGenRequest(t *testing.T) {
//	client := GetClient()
//	//初始化请求接口信息
//	getRequest := &request.PddDdkResourceUrlGenRequest{}
//	getRequest.AddParameter("pid", `1949211_146717471`)
//	getRequest.AddParameter("generate_schema_url", `true`)
//	getRequest.AddParameter("custom_parameters", `custom_o1`)
//	getRequest.AddParameter("resource_type", `39998`)
//	getRequest.AddParameter("url", `https://mobile.yangkeduo.com/promotion_op.html?type=27&id=162895`)
//
//	//初始化结果类型
//	var getResponse DefaultResponse = &pddddkresourceurlgen.Response{}
//	//执行请求接口得到结果
//	err := client.Exec(getRequest, getResponse)
//	if err != nil {
//		t.Log(err)
//	} else {
//		result := getResponse.(*pddddkresourceurlgen.Response)
//
//		fmt.Println(result.Body)
//
//	}
//}

func TestPddDdkMemberAuthorityQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkMemberAuthorityQueryRequest{}
	getRequest.AddParameter("pid", `1949211_176135239`)
	getRequest.AddParameter("custom_parameters", `user_1384`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkmemberauthorityquery2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkmemberauthorityquery2.Response)

		fmt.Println(result.AuthorityQueryResponse.Bind)

	}
}

func TestPddDdkGoodsZsUnitUrlGenRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkGoodsZsUnitUrlGenRequest{}
	getRequest.AddParameter("pid", `1949211_176135239`)
	getRequest.AddParameter("custom_parameters", `user_1384`)
	//getRequest.AddParameter("source_url", `https://p.pinduoduo.com/QalbCK7j`)
	getRequest.AddParameter("source_url", `https://p.pinduoduo.com/sKvbBwgP`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkgoodszsuniturlgen2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkgoodszsuniturlgen2.Response)

		fmt.Println(result)

	}
}

func TestPddDdkGoodsRecommendGetRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkGoodsRecommendGetRequest{}
	getRequest.AddParameter("pid", `1949211_176135239`)
	getRequest.AddParameter("custom_parameters", `user_1384`)
	getRequest.AddParameter("channel_type", `4`)
	getRequest.AddParameter("cat_id", `20100`)
	//getRequest.AddParameter("source_url", `https://p.pinduoduo.com/QalbCK7j`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkgoodsrecommendget2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkgoodsrecommendget2.Response)
		fmt.Println(result)
	}
}

//创建pid
func TestPddDdkOauthGoodsPidGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsPidGenerateRequest{}
	//getRequest.AddParameter("media_id", ``)
	getRequest.AddParameter("number", `1`)
	getRequest.AddParameter("p_id_name_list", `["测试api"]`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodspidgenerate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodspidgenerate.Response)
		fmt.Println(result)
	}
}

//创建pid列表查询
func TestPddDdkOauthGoodsPidQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsPidQueryRequest{}
	getRequest.AddParameter("page", `1`)
	getRequest.AddParameter("page_size", `10`)
	//getRequest.AddParameter("pid_list", `["测试api"]`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodspidquery.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodspidquery.Response)
		fmt.Println(result)
	}
}

//生成营销工具推广链接
func TestPddDdkOauthRpPromUrlGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthRpPromUrlGenerateRequest{}
	//getRequest.AddParameter("channel_type", `10`)
	//	营销工具类型，必填：-1-活动列表，0-红包(需申请推广权限)，2–新人红包，3-刮刮卡，5-员工内购，6-购物车，10-生成绑定备案链接，12-砸金蛋，13-一元购，14-千万补贴B端页面，15-充值中心B端页面；红包推广权限申请流程链接：https://jinbao.pinduoduo.com/qa-system?questionId=289
	getRequest.AddParameter("channel_type", `13`)
	getRequest.AddParameter("generate_we_app", `true`)
	getRequest.AddParameter("generate_short_url", `true`)
	getRequest.AddParameter("generate_schema_url", `true`)
	getRequest.AddParameter("generate_qq_app", `true`)
	//getRequest.AddParameter("goods_sign", `Y9_2n41LTLdOkjGzwvfZ1vtzvZRnI2fG_J2g7666nb`)
	getRequest.AddParameter("p_id_list", `["9561589_216982659"]`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthrppromurlgenerate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthrppromurlgenerate.Response)
		fmt.Println(result)
	}
}

//备案查询
func TestPddDdkOauthMemberAuthorityQueryRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthMemberAuthorityQueryRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthmemberauthorityquery.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthmemberauthorityquery.Response)
		fmt.Println(result)
	}
}

//拼多多主站频道推广接口
func TestPddDdkOauthResourceUrlGenRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthResourceUrlGenRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)
	getRequest.AddParameter("generate_we_app", `true`)
	//频道来源：4-限时秒杀,39997-充值中心, 39998-活动转链，39996-百亿补贴，39999-电器城，40000-领券中心，50005-火车票
	getRequest.AddParameter("resource_type", `39997`)
	getRequest.AddParameter("url", ``)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthresourceurlgen.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthresourceurlgen.Response)
		fmt.Println(result)
	}
}

//多多进宝转链接口
func TestPddDdkOauthGoodsZsUnitUrlGenRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsZsUnitUrlGenRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)
	getRequest.AddParameter("generate_schema_url", `true`)
	//getRequest.AddParameter("source_url", `https://p.pinduoduo.com/OKfa6xI5`)
	getRequest.AddParameter("source_url", `https://mobile.yangkeduo.com/goods.html?goods_id=263240709856`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodszsuniturlgen.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodszsuniturlgen.Response)
		fmt.Println(result)
	}
}

//多多进宝转链接口
func TestPddDdkOauthCmsPromUrlGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthCmsPromUrlGenerateRequest{}
	getRequest.AddParameter("p_id_list", `["9561589_216982659"]`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)
	getRequest.AddParameter("generate_mobile", `true`)
	getRequest.AddParameter("generate_schema_url", `true`)
	getRequest.AddParameter("generate_short_url", `true`)
	getRequest.AddParameter("generate_we_app", `true`)
	//getRequest.AddParameter("multi_group", `true`)
	//getRequest.AddParameter("keyword", `尺子`)
	//	0, "1.9包邮"；1, "今日爆款"； 2, "品牌清仓"； 4,"PC端专属商城"
	getRequest.AddParameter("channel_type", `1`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthcmspromurlgenerate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthcmspromurlgenerate.Response)
		fmt.Println(result)
	}
}

// 多多进宝商品详情查询
func TestPddDdkOauthGoodsDetailRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsDetailRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)
	getRequest.AddParameter("goods_img_type", `0`)
	getRequest.AddParameter("goods_sign", `Y9f2nyAQiHROkjG1wvfZ1nfljXjyFYAB_JFKXVXDZx`)
	getRequest.AddParameter("search_id", ``)
	getRequest.AddParameter("zs_duo_id", ``)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodsdetail.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodsdetail.Response)
		fmt.Println(result)
	}
}

// 生成多多进宝推广链接
func TestPddDdkOauthGoodsPromUrlGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsPromUrlGenerateRequest{}
	getRequest.AddParameter("p_id", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)

	getRequest.AddParameter("force_duo_id", `true`)
	//getRequest.AddParameter("generate_authority_url", `true`)
	getRequest.AddParameter("generate_mall_collect_coupon", `true`)
	getRequest.AddParameter("generate_qq_app", `true`)
	getRequest.AddParameter("generate_schema_url", `true`)
	getRequest.AddParameter("generate_short_url", `true`)
	getRequest.AddParameter("multi_group", `true`)
	getRequest.AddParameter("generate_we_app", `true`)

	getRequest.AddParameter("goods_sign_list", `["Y9_2keuoYAJOkjG1wvTZ1l8QQ69Cn5Px_JQS6y8PhLh"]`)
	getRequest.AddParameter("search_id", ``)
	getRequest.AddParameter("zs_duo_id", ``)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodspromurlgenerate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodspromurlgenerate.Response)
		fmt.Println(result)
	}
}

// 运营频道商品查询API
func TestPddDdkOauthGoodsRecommendGetRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsRecommendGetRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)

	//活动商品标记数组，例：[4,7]，4-秒杀，7-百亿补贴，10851-千万补贴，10913-招商礼金商品，31-品牌黑标，10564-精选爆品-官方直推爆款，10584-精选爆品-团长推荐，24-品牌高佣，其他的值请忽略
	getRequest.AddParameter("activity_tags", ``)
	//猜你喜欢场景的商品类目，20100-百货，20200-母婴，20300-食品，20400-女装，20500-电器，20600-鞋包，20700-内衣，20800-美妆，20900-男装，21000-水果，21100-家纺，21200-文具,21300-运动,21400-虚拟,21500-汽车,21600-家装,21700-家具,21800-医药;
	getRequest.AddParameter("cat_id", ``)
	//进宝频道推广商品: 1-今日销量榜,3-相似商品推荐,4-猜你喜欢(和进宝网站精选一致),5-实时热销榜,6-实时收益榜。默认值5
	getRequest.AddParameter("channel_type", `4`)
	getRequest.AddParameter("force_auth_duo_id", `true`)
	getRequest.AddParameter("goods_img_type", `0`)
	getRequest.AddParameter("goods_sign_list", `["Y9f2nyAQiHROkjG1wvfZ1nfljXjyFYAB_JFKXVXDZx"]`)
	getRequest.AddParameter("limit", `10`)
	getRequest.AddParameter("offset", `0`)
	getRequest.AddParameter("list_id", ``)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodsrecommendget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodsrecommendget.Response)
		fmt.Println(result)
	}
}

// 多多进宝商品查询
func TestPddDdkOauthGoodsSearchRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthGoodsSearchRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)
	getRequest.AddParameter("keyword", `210715779304`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthgoodssearch.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthgoodssearch.Response)
		fmt.Println(result)
	}
}

// 查询所有授权的多多客订单
func TestPddDdkAllOrderListIncrementGetRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkAllOrderListIncrementGetRequest{}
	getRequest.AddParameter("end_update_time", `1627964478`)
	getRequest.AddParameter("start_update_time", `1627902547`)
	//getRequest.AddParameter("query_order_type", `1`) //订单类型：1-推广订单；2-直播间订单
	getRequest.AddParameter("page", `1`)
	getRequest.AddParameter("page_size", `50`)

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkallorderlistincrementget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkallorderlistincrementget.Response)
		fmt.Println(result)
	}
}

// 获取订单详情
func TestPddDdkOauthOrderDetailGetRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthOrderDetailGetRequest{}
	//getRequest.AddParameter("end_update_time", `1627902247`)
	getRequest.AddParameter("order_sn", `210802-401069834933833`)
	//getRequest.AddParameter("query_order_type", `1`) //订单类型：1-推广订单；2-直播间订单

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauthorderdetailget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauthorderdetailget.Response)
		fmt.Println(result)
	}
}

// 获取订单详情
func TestPddDdkOauthThemePromUrlGenerateRequest(t *testing.T) {
	client := GetClient()
	//初始化请求接口信息
	getRequest := &request2.PddDdkOauthThemePromUrlGenerateRequest{}
	getRequest.AddParameter("pid", `9561589_216982659`)
	getRequest.AddParameter("custom_parameters", `{"uid":"user_1"}`)
	getRequest.AddParameter("generate_short_url", "true")
	getRequest.AddParameter("generate_mobile", "true")
	getRequest.AddParameter("generate_schema_url", "true")
	getRequest.AddParameter("theme_id_list", `[7627]`)
	getRequest.AddParameter("generate_we_app", "true")

	//初始化结果类型
	var getResponse DefaultResponse = &pddddkoauththemepromurlgenerate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*pddddkoauththemepromurlgenerate.Response)
		fmt.Println(result)
	}
}
