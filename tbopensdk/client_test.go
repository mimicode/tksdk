package tbopensdk

import (
	"fmt"
	request2 "github.com/mimicode/tksdk/tbopensdk/request"
	"github.com/mimicode/tksdk/tbopensdk/response/juitemssearch"
	"github.com/mimicode/tksdk/tbopensdk/response/sellercatslistget"
	"github.com/mimicode/tksdk/tbopensdk/response/shopcatslistget"
	"github.com/mimicode/tksdk/tbopensdk/response/shopgetbytitle"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkactivitylinkget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkcouponget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkdgitemcouponget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkdgmaterialoptional"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkdgnewuserorderget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkdgnewuserordersum"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkdgoptimusmaterial"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkdgvegastljcreate"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkitemconvert"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkitemcouponget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkitemget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkiteminfoget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkitemrecommendget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkjutqgget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkprivilegeget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscactivityinfoget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscactivitylinktoolget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbksccouponbrandrecommend"
	"github.com/mimicode/tksdk/tbopensdk/response/tbksccouponrealtimerecommend"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscinvitecodeget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscmaterialoptional"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscorderget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscpublisherinfoget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscpublisherinfosave"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkscshopconvert"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkshopconvert"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkshopget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbkspreadget"
	"github.com/mimicode/tksdk/tbopensdk/response/tbktpwdconvert"
	response2 "github.com/mimicode/tksdk/tbopensdk/response/tbktpwdcreate"
	"testing"
)

var (
	appKey, appSecret, sessionKey string
)

func init() {
	appKey = ""
	appSecret = ""
	sessionKey = ""
}
func TestTbkiteminfoget(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkItemInfoGetRequest{}
	getRequest.AddParameter("num_iids", "640057354722")
	getRequest.AddParameter("platform", "1")
	getRequest.AddParameter("ip", "11.22.33.43")
	getRequest.AddParameter("get_topn_rate", "1")
	//初始化结果类型
	var getResponse DefaultResponse = &tbkiteminfoget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkiteminfoget.Response)

		if result.IsError() {
			fmt.Println(result.Body)
		} else {
			items := result.TbkItemInfoGetResult.Results.NTbkItem
			for _, v := range items {
				fmt.Println(v.Title)
			}
		}

	}
}

func TestTbkPrivilegeGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	//mm_41521745_519950261_108916600241
	getRequest := &request2.TbkPrivilegeGetRequest{}
	getRequest.AddParameter("adzone_id", "108916600241")
	getRequest.AddParameter("site_id", "519950261")

	getRequest.AddParameter("item_id", "640057354722")

	//getRequest.AddParameter("relation_id", "")
	//getRequest.AddParameter("me", "")
	getRequest.AddParameter("platform", "1")
	getRequest.AddParameter("get_topn_rate", "1")

	//初始化结果类型 19.90
	//券后
	//40.01%
	//前2000件佣金率
	//7.96
	var getResponse DefaultResponse = &tbkprivilegeget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkprivilegeget.Response)

		fmt.Println(result.TbkPrivilegeGetResult.Result.Data)

	}
}

func TestTbkScMaterialOptional(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScMaterialOptionalRequest{}
	getRequest.AddParameter("adzone_id", "108916600241")
	getRequest.AddParameter("site_id", "519950261")

	getRequest.AddParameter("q", "https://h5.m.taobao.com/awp/core/detail.htm?id=658058537271")
	getRequest.AddParameter("platform", "1")
	getRequest.AddParameter("page_no", "1")
	getRequest.AddParameter("page_size", "1")
	//getRequest.AddParameter("get_topn_rate", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkscmaterialoptional.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscmaterialoptional.Response)

		fmt.Println(result.TbkScMaterialOptionalResult)

	}
}

func TestTbkScActivitylinkToolget(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScActivitylinkToolgetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	getRequest.AddParameter("promotion_scene_id", "8479247")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkscactivitylinktoolget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscactivitylinktoolget.Response)

		fmt.Println(result.TbkScActivitylinkToolgetResult)

	}
}

func TestTbkItemConvert(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkItemConvertRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("fields", "num_iid,click_url")

	getRequest.AddParameter("num_iids", "583866215568,578307080718")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkitemconvert.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkitemconvert.Response)

		fmt.Println(result.TbkItemConvertResult.Results.NTbkItem)

	}
}

func TestTbkShopConvert(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkShopConvertRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("fields", "user_id,click_url")

	getRequest.AddParameter("user_ids", "188124207,383602586")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkshopconvert.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkshopconvert.Response)

		fmt.Println(result.TbkShopConvertResult.Results)

	}
}

func TestTbkTpwdConvert(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkTpwdConvertRequest{}
	getRequest.AddParameter("adzone_id", "108944100308")
	getRequest.AddParameter("password_content", "(Smh4X47XhR9)")
	//getRequest.AddParameter("password_content", "(ibazcMRJPqp)")
	//https://s.click.taobao.com/t?e=m%3D2%26s%3DEjQLLe%2BWEEUcQipKwQzePDAVflQIoZepK7Vc7tFgwiFRAdhuF14FMejlC3NB5Wk9MMgx22UI05aWZIHuAfb16zR0r%2BMBQW%2B1BJ3M5HVhgpIeDJkG8xs1jBZZRUpbmzNoBuOrqNvh1RvGDmntuH4VtA%3D%3D
	//https://s.click.taobao.com/yArvgtu
	//初始化结果类型
	var getResponse DefaultResponse = &tbktpwdconvert.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbktpwdconvert.Response)

		fmt.Println(result.TbkTpwdConvertResult.Data)

	}
}

func TestTbkScOrderGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScOrderGetRequest{}
	getRequest.AddParameter("fields", "trade_parent_id,trade_id,num_iid,item_title,item_num,price,pay_price,seller_nick,seller_shop_title,commission,commission_rate,unid,create_time,earning_time,tk_status,tk3rd_type,tk3rd_pub_id,order_type,income_rate,pub_share_pre_fee,subsidy_rate,subsidy_type,terminal_type,auction_category,site_id,site_name,adzone_id,adzone_name,alipay_total_price,total_commission_rate,total_commission_fee,subsidy_fee,relation_id,click_time")
	getRequest.AddParameter("start_time", "2019-02-13 20:50:00")
	getRequest.AddParameter("span", "600")
	getRequest.AddParameter("order_query_type", "create_time")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkscorderget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscorderget.Response)

		fmt.Println(result.TbkScOrderGetResult.Results)

	}
}

func TestShopcatsListGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.ShopcatsListGetRequest{}
	getRequest.AddParameter("fields", "cid,parent_cid,name,is_parent")

	//初始化结果类型
	var getResponse DefaultResponse = &shopcatslistget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*shopcatslistget.Response)

		fmt.Println(result.ShopcatsListGetResult.ShopCats)

	}
}

func TestSellercatsListGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.SellercatsListGetRequest{}
	getRequest.AddParameter("nick", "韩都衣舍旗舰店")

	//初始化结果类型
	var getResponse DefaultResponse = &sellercatslistget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*sellercatslistget.Response)

		fmt.Println(result.SellercatsListGetResult.SellerCats)

	}
}

func TestShopGetbytitle(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.ShopGetbytitleRequest{}
	getRequest.AddParameter("title", "小糖医旗舰店")

	//初始化结果类型
	var getResponse DefaultResponse = &shopgetbytitle.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*shopgetbytitle.Response)

		fmt.Println(result.ShopGetbytitleResult.ResultShop)

	}
}
func TestTbkShopGetRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkShopGetRequest{}
	getRequest.AddParameter("fields", "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url")
	getRequest.AddParameter("q", "小糖医旗舰店")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkshopget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkshopget.Response)

		fmt.Println(result.Body)

	}
}

func TestTbkScInvitecodeGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScInvitecodeGetRequest{}
	getRequest.AddParameter("relation_app", "common")
	getRequest.AddParameter("code_type", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkscinvitecodeget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscinvitecodeget.Response)

		fmt.Println(result.TbkScInvitecodeGetResult.Data)

	}
}

func TestTbkScPublisherInfoGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScPublisherInfoGetRequest{}
	getRequest.AddParameter("relation_app", "common")
	getRequest.AddParameter("info_type", "1")
	getRequest.AddParameter("relation_id", "21458414451")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkscpublisherinfoget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscpublisherinfoget.Response)

		fmt.Println(result.TbkScPublisherInfoGetResult.Data.InviterList.MapData)

	}
}

//(未完成)
func TestTbkScPublisherInfoSave(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScPublisherInfoSaveRequest{}
	getRequest.AddParameter("inviter_code", "QMV5J9")
	getRequest.AddParameter("info_type", "1")
	getRequest.AddParameter("note", "备案测试")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkscpublisherinfosave.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscpublisherinfosave.Response)
		fmt.Println(result.TbkScPublisherInfoSaveResult.Data.RelationID)

	}
}

func TestTbkScCouponRealtimeRecommend(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScCouponRealtimeRecommendRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	//初始化结果类型
	var getResponse DefaultResponse = &tbksccouponrealtimerecommend.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbksccouponrealtimerecommend.Response)

		fmt.Println(result.TbkScCouponRealtimeRecommendResult.Results.TbkCoupon)

	}
}

func TestTbkScCouponBrandRecommend(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScCouponBrandRecommendRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	//初始化结果类型
	var getResponse DefaultResponse = &tbksccouponbrandrecommend.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbksccouponbrandrecommend.Response)

		fmt.Println(result.TbkScCouponBrandRecommendResult.Results.TbkCoupon)

	}
}

func TestTbkSpreadGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkSpreadGetRequest{}
	getRequest.AddParameter("requests", `[{"url":"https://uland.taobao.com/coupon/edetail?e=wog6mTjZv%2FIGQASttHIRqT8%2FTgo%2BDKYNn3MtPDDRCxLuFdGM%2Fy8ADJwJ6dLD4BPh4h5828KZV1siM37WwZ3M460TeAL%2BmcF1hLjdYI0pIJBwx7%2FmoMcv1BemP0hpIIPvjDppvlX%2Bob8NlNJBuapvQ2MDg9t1zp0R8pjV3C9qcwRg5t1bGd7ixKOQJ6Ic4SbV&traceId=0b0b40de15501301853222175e&union_lens=lensId:0b1832c2_0b51_168eaf5847e_8589&xId=vJinjnbsDyrka2dVuwvpFpBvPqgdfSmfEm3bA2Jig86yBLgQear0ZghwNkuOsiTtahUYTf5Kq7I6X8iIlkrpeU&activityId=d1204e96d378436cabe64bfebccc6f74"}]`)

	//初始化结果类型
	var getResponse DefaultResponse = &tbkspreadget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkspreadget.Response)

		fmt.Println(result.TbkSpreadGetResult.Results.TbkSpread)

	}
}

func TestTbkItemCouponGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkItemCouponGetRequest{}
	getRequest.AddParameter("pid", "mm_27437251_7418269_24546980")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkitemcouponget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkitemcouponget.Response)

		fmt.Println(result.TbkItemCouponGetResult.Results.TbkCoupon)

	}
}

func TestTbkitemGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkItemGetRequest{}
	getRequest.AddParameter("fields", "num_iid")
	getRequest.AddParameter("q", "连衣裙")
	getRequest.AddParameter("page_size", "40")
	//初始化结果类型
	var getResponse DefaultResponse = &tbkitemget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkitemget.Response)

		fmt.Println(result.TbkItemGetResult.Results.NTbkItem)

	}
}

func TestTbkItemRecommendGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkItemRecommendGetRequest{}
	getRequest.AddParameter("fields", "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick")
	getRequest.AddParameter("num_iid", "591396538411")
	getRequest.AddParameter("count", "20")
	//初始化结果类型
	var getResponse DefaultResponse = &tbkitemrecommendget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkitemrecommendget.Response)

		fmt.Println(result.TbkItemRecommendGetResult.Results.NTbkItem)

	}

}

func TestJuItemsSearch(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.JuItemsSearchRequest{}
	getRequest.AddParameter("param_top_item_query", `{"current_page":1,"page_size":2,"pid":"mm_27437251_7418269_24546980"}`)

	//初始化结果类型
	var getResponse DefaultResponse = &juitemssearch.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*juitemssearch.Response)

		fmt.Println(result.JuItemsSearchResult.Result)

	}

}

func TestTbkActivitylinkGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkActivitylinkGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")

	getRequest.AddParameter("promotion_scene_id", "8493178")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkactivitylinkget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkactivitylinkget.Response)

		fmt.Println(result.TbkActivitylinkGetResult)

	}
}

func TestTbkDgItemCouponGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkDgItemCouponGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")

	getRequest.AddParameter("cat", "16")
	getRequest.AddParameter("page_size", "2")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkdgitemcouponget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkdgitemcouponget.Response)

		fmt.Println(result.TbkDgItemCouponGetResult.Results.TbkCoupon)

	}
}

func TestTbkDgMaterialOptional(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkDgMaterialOptionalRequest{}
	getRequest.AddParameter("adzone_id", "108944100308")

	//getRequest.AddParameter("cat", "16")
	getRequest.AddParameter("page_size", "2")
	getRequest.AddParameter("page_size", "1")
	getRequest.AddParameter("q", "https://detail.m.tmall.com/item.htm?id=640057354722")

	getRequest.AddParameter("get_topn_rate", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkdgmaterialoptional.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkdgmaterialoptional.Response)

		fmt.Println(result.TbkDgMaterialOptionalResult.ResultList.MapData)

	}
}

func TestTbkDgNewuserOrderGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkDgNewuserOrderGetRequest{}
	getRequest.AddParameter("activity_id", "119013_6")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkdgnewuserorderget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkdgnewuserorderget.Response)

		fmt.Println(result.TbkDgNewuserOrderGetResult.Results.Data.Results.MapData)

	}
}

func TestTbkDgNewuserOrderSum(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkDgNewuserOrderSumRequest{}
	getRequest.AddParameter("activity_id", "119013_6")
	getRequest.AddParameter("page_size", "2")
	getRequest.AddParameter("page_no", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkdgnewuserordersum.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkdgnewuserordersum.Response)

		fmt.Println(result.TbkDgNewuserOrderSumResult.Results.Data.Results)

	}
}

func TestTbkDgOptimusMaterial(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkDgOptimusMaterialRequest{}
	getRequest.AddParameter("adzone_id", "108944100308")
	getRequest.AddParameter("material_id", "44044")

	getRequest.AddParameter("page_size", "20")
	getRequest.AddParameter("page_no", "1")

	getRequest.AddParameter("q", "https://detail.m.tmall.com/item.htm?id=640057354722")

	getRequest.AddParameter("get_topn_rate", "1")

	//智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	//getRequest.AddParameter("device_type", "OAID")
	//智能匹配-设备号加密类型：MD5，类型为OAID时不传
	//getRequest.AddParameter("device_encrypt", "")
	//智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值
	//getRequest.AddParameter("device_value", "f9ec48afbd41669c")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkdgoptimusmaterial.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkdgoptimusmaterial.Response)

		fmt.Println(result.TbkDgOptimusMaterialResult.ResultList.MapData)

	}
}

func TestTbkJuTqgGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkJuTqgGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("fields", "click_url,pic_url,reserve_price,zk_final_price,total_amount,sold_num,title,category_name,start_time,end_time,num_iid")

	getRequest.AddParameter("start_time", "2019-02-15 07:00:00")
	getRequest.AddParameter("end_time", "2019-02-15 15:00:00")

	getRequest.AddParameter("page_size", "2")
	getRequest.AddParameter("page_no", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkjutqgget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkjutqgget.Response)

		fmt.Println(result.TbkJuTqgGetResult.Results.Results)

	}
}

func TestTbkCouponGet(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkCouponGetRequest{}
	getRequest.AddParameter("item_id", "573509331754")

	getRequest.AddParameter("activity_id", "335fa4c30e8a4f5da28596d122b48c28")

	//初始化结果类型
	var getResponse DefaultResponse = &tbkcouponget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkcouponget.Response)

		fmt.Println(result.TbkCouponGetResult.Data)

	}
}

func TestTbkTpwdCreate(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request2.TbkTpwdCreateRequest{}
	getRequest.AddParameter("text", "108916600241")
	getRequest.AddParameter("url", "https://s.click.taobao.com/t?e=m%3D2%26s%3DEjQLLe%2BWEEUcQipKwQzePDAVflQIoZepK7Vc7tFgwiFRAdhuF14FMejlC3NB5Wk9MMgx22UI05aWZIHuAfb16zR0r%2BMBQW%2B1BJ3M5HVhgpIeDJkG8xs1jBZZRUpbmzNoBuOrqNvh1RvGDmntuH4VtA%3D%3D")

	//初始化结果类型
	var getResponse DefaultResponse = &response2.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response2.Response)

		fmt.Println(result.Body)

	}
}

func Test_tbkscactivityinfoget(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScActivityInfoGetRequest{}
	//getRequest.AddParameter("activity_material_id", "20150318020002192")
	getRequest.AddParameter("activity_material_id", "1585018034441")
	getRequest.AddParameter("adzone_id", "108916600241")
	getRequest.AddParameter("site_id", "519950261")
	//getRequest.AddParameter("relation_id", "")
	t.Logf("GetScActivityInfoGet : %s", getRequest.GetParameters().Encode())
	//初始化结果类型
	var getResponse DefaultResponse = &tbkscactivityinfoget.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscactivityinfoget.Response)
		t.Log(result.TbkScActivityInfoGetResponseResult.Data.ClickURL)
		fmt.Println(result.Body)

	}
}
func Test_TbkScShopConvertRequest(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request2.TbkScShopConvertRequest{}
	getRequest.AddParameter("fields", "user_id,click_url")
	getRequest.AddParameter("adzone_id", "108916600241")
	getRequest.AddParameter("site_id", "519950261")
	getRequest.AddParameter("platform", "1")
	getRequest.AddParameter("user_ids", "533497499")
	getRequest.AddParameter("relation_id", "")
	t.Logf("GetScActivityInfoGet : %s", getRequest.GetParameters().Encode())
	//初始化结果类型
	var getResponse DefaultResponse = &tbkscshopconvert.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkscshopconvert.Response)
		//t.Log(result.TbkScActivityInfoGetResponseResult.Data.ClickURL)
		fmt.Println(result.Body)

	}
}

func Test_TbkDgVegasTljCreate(t *testing.T) {

	//初始化TopClient
	client := &TopClient{}
	client.Init("27572253", "72c20ce22a8614873eb9ae6b4db96728", "")

	//初始化请求接口信息
	getRequest := &request2.TbkDgVegasTljCreateRequest{}
	getRequest.AddParameter("adzone_id", "108916600241")
	getRequest.AddParameter("item_id", "580973170958")
	getRequest.AddParameter("name", "我是标题啊")
	getRequest.AddParameter("total_num", "1")
	getRequest.AddParameter("user_total_win_num_limit", "1")
	getRequest.AddParameter("security_switch", "false")
	getRequest.AddParameter("per_face", "1")
	getRequest.AddParameter("use_end_time_mode", "2")
	getRequest.AddParameter("send_start_time", "2020-05-27 18:11:51")
	getRequest.AddParameter("send_end_time", "2020-05-28 18:11:51")

	getRequest.AddParameter("use_start_time", "2020-05-27")
	getRequest.AddParameter("use_end_time", "2020-05-28")

	//https://uland.taobao.com/taolijin/edetail?eh=4O5nIa6wOx%2BZuQF0XRz0iAXoB%2BDaBK5LQS0Flu%2FfbSp4QsdWMikAalrisGmre1Id0BFAqRODu13vu9JrnQ8Z7sAi0mUFmpklJRityAi%2FnC%2BcSK3KB2K3X8huzHP8BkZcdovnJEoTD5PBs0FSnlb7iikxaTHDv9oSIN6dAw1GfUxA6nVjU3GMGMmtOyzreyGQF9wdUScG04KEVTzCdvuEkK57cF7NJ%2F6%2FytToUuhFaZvbkp%2FKNzF6UhLvvDepeEhPdUrIY4xEuTv%2BcBVGuURu7Y1J1dE2yfYoRYwpIuVYADcCGruttYDvNg%3D%3D&union_lens=lensId%3A0b195d88_0c94_17255d4d7a9_30fa%3Btraffic_flag%3Dlm
	t.Logf("TbkDgVegasTljCreateRequest : %s", getRequest.GetParameters().Encode())
	//初始化结果类型
	var getResponse DefaultResponse = &tbkdgvegastljcreate.Response{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*tbkdgvegastljcreate.Response)

		fmt.Println(result.Body)

	}
}
