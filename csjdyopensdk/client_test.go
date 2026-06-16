package csjdyopensdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mimicode/tksdk/csjdyopensdk/request"
	"github.com/mimicode/tksdk/csjdyopensdk/response/commandparse"
	"github.com/mimicode/tksdk/csjdyopensdk/response/livelink"
	"github.com/mimicode/tksdk/csjdyopensdk/response/livesearch"
	"github.com/mimicode/tksdk/csjdyopensdk/response/ordersearch"
	"github.com/mimicode/tksdk/csjdyopensdk/response/productcategory"
	"github.com/mimicode/tksdk/csjdyopensdk/response/productdetail"
	"github.com/mimicode/tksdk/csjdyopensdk/response/productlink"
	"github.com/mimicode/tksdk/csjdyopensdk/response/productsearch"
)

var (
	appKey, appSecret string
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Csjdy struct {
					AppKey    string `json:"app_key"`
					AppSecret string `json:"app_secret"`
				} `json:"csjdy"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Csjdy.AppKey
				appSecret = data.Csjdy.AppSecret
			}
		}
	}
}

func GetClient() *TopClient {
	client := &TopClient{}
	client.Init(appKey, appSecret, "")
	return client
}

// TestProductSearch 商品列表接口测试
func TestProductSearch(t *testing.T) {
	client := GetClient()
	req := &request.ProductSearchRequest{}
	req.SetPage(1)
	req.SetPageSize(10)
	req.SetTitle("手机")
	req.SetSearchType(1) // 历史销量排序
	req.SetOrderType(1)  // 降序

	var resp productsearch.Response
	err := client.Exec(req, &resp)
	if err != nil {
		t.Errorf("ProductSearch failed: %v", err)
		return
	}

	fmt.Printf("商品列表响应: code=%d, desc=%s, total=%d\n", resp.Code, resp.Desc, resp.Data.Total)
	for i, p := range resp.Data.Products {
		if i >= 3 {
			fmt.Printf("... 共%d个商品\n", len(resp.Data.Products))
			break
		}
		fmt.Printf("  商品%d: id=%d, title=%s, price=%d, cos_fee=%d\n", i+1, p.ProductId, p.Title, p.Price, p.CosFee)
	}
}

// TestProductDetail 商品详情接口测试
func TestProductDetail(t *testing.T) {
	// 先搜索获取商品ID
	client := GetClient()
	searchReq := &request.ProductSearchRequest{}
	searchReq.SetPage(1)
	searchReq.SetPageSize(1)

	var searchResp productsearch.Response
	err := client.Exec(searchReq, &searchResp)
	if err != nil {
		t.Errorf("Search failed: %v", err)
		return
	}

	if len(searchResp.Data.Products) == 0 {
		t.Log("No products found from search, skip detail test")
		return
	}

	productId := searchResp.Data.Products[0].ProductId

	// 查询商品详情
	detailReq := &request.ProductDetailRequest{}
	detailReq.SetProductIds([]int64{productId})

	var detailResp productdetail.Response
	err = client.Exec(detailReq, &detailResp)
	if err != nil {
		t.Errorf("ProductDetail failed: %v", err)
		return
	}

	fmt.Printf("商品详情响应: code=%d, desc=%s, total=%d\n", detailResp.Code, detailResp.Desc, detailResp.Data.Total)
	if len(detailResp.Data.Products) > 0 {
		p := detailResp.Data.Products[0]
		fmt.Printf("  商品: id=%d, title=%s\n", p.ProductId, p.Title)
		fmt.Printf("  价格: %d分, 佣金比例: %d, 佣金金额: %d分\n", p.Price, p.CosRatio, p.CosFee)
		fmt.Printf("  店铺: id=%d, name=%s\n", p.ShopId, p.ShopName)
	}
}

// TestProductCategory 商品类目接口测试
func TestProductCategory(t *testing.T) {
	client := GetClient()

	// 查询一级类目
	req := &request.ProductCategoryRequest{}
	req.SetParentId(0)
	req.SetChannel(1)

	var resp productcategory.Response
	err := client.Exec(req, &resp)
	if err != nil {
		t.Errorf("ProductCategory failed: %v", err)
		return
	}

	fmt.Printf("商品类目响应: code=%d, desc=%s, total=%d\n", resp.Code, resp.Desc, resp.Data.Total)
	for i, c := range resp.Data.CategoryList {
		if i >= 5 {
			fmt.Printf("... 共%d个类目\n", len(resp.Data.CategoryList))
			break
		}
		fmt.Printf("  类目%d: id=%d, name=%s, level=%d\n", i+1, c.Id, c.Name, c.Level)
	}
}

// TestProductLink 商品转链接口测试
func TestProductLink(t *testing.T) {
	// 先搜索获取商品
	client := GetClient()
	searchReq := &request.ProductSearchRequest{}
	searchReq.SetPage(1)
	searchReq.SetPageSize(1)

	var searchResp productsearch.Response
	err := client.Exec(searchReq, &searchResp)
	if err != nil {
		t.Errorf("Search failed: %v", err)
		return
	}

	if len(searchResp.Data.Products) == 0 {
		t.Log("No products found from search, skip link test")
		return
	}

	product := searchResp.Data.Products[0]

	// 商品转链
	linkReq := &request.ProductLinkRequest{}
	linkReq.SetProductUrl(product.DetailUrl)
	linkReq.SetProductExt(product.Ext)
	linkReq.SetShareType([]int{1, 3}) // deep link + 口令
	linkReq.SetPlatform(0)            // 抖音
	linkReq.SetUseCoupon(true)

	var linkResp productlink.Response
	err = client.Exec(linkReq, &linkResp)
	if err != nil {
		t.Errorf("ProductLink failed: %v", err)
		return
	}

	fmt.Printf("商品转链响应: code=%d, desc=%s\n", linkResp.Code, linkResp.Desc)
	fmt.Printf("  deeplink: %s\n", linkResp.Data.DyDeeplink)
	fmt.Printf("  口令: %s\n", linkResp.Data.DyPassword)
	if linkResp.Data.DySharelink != "" {
		fmt.Printf("  sharelink: %s\n", linkResp.Data.DySharelink)
	}
}

// TestLiveSearch 直播列表接口测试
func TestLiveSearch(t *testing.T) {
	client := GetClient()
	req := &request.LiveSearchRequest{}
	req.SetPage(1)
	req.SetPageSize(10)
	req.SetStatus(1)       // 只看在播
	req.SetSortBy(2)       // 按销量排序
	req.SetNeedProducts(1) // 返回商品信息

	var resp livesearch.Response
	err := client.Exec(req, &resp)
	if err != nil {
		t.Errorf("LiveSearch failed: %v", err)
		return
	}

	fmt.Printf("直播列表响应: code=%d, desc=%s, total=%d\n", resp.Code, resp.Desc, resp.Data.Total)
	for i, live := range resp.Data.LiveInfos {
		if i >= 3 {
			fmt.Printf("... 共%d个直播间\n", len(resp.Data.LiveInfos))
			break
		}
		fmt.Printf("  直播间%d: room_id=%s, 主播=%s, 粉丝=%d, 在线=%d\n",
			i+1, live.RoomId, live.AuthorName, live.FansNum, live.OnlineNum)
		fmt.Printf("    平均GMV: %s, 平均佣金率: %s\n", live.AverageGmv, live.AverageCommissionRate)
		if len(live.Products) > 0 {
			fmt.Printf("    商品数量: %d\n", len(live.Products))
		}
	}
}

// TestLiveLink 直播转链接口测试
func TestLiveLink(t *testing.T) {
	// 先获取直播间信息
	client := GetClient()
	searchReq := &request.LiveSearchRequest{}
	searchReq.SetPage(1)
	searchReq.SetPageSize(5)
	searchReq.SetStatus(1)

	var searchResp livesearch.Response
	err := client.Exec(searchReq, &searchResp)
	if err != nil {
		t.Errorf("LiveSearch failed: %v", err)
		return
	}

	if len(searchResp.Data.LiveInfos) == 0 {
		t.Log("No live rooms found, skip link test")
		return
	}

	live := searchResp.Data.LiveInfos[0]

	// 直播转链
	linkReq := &request.LiveLinkRequest{}
	linkReq.SetAuthorOpenid(live.AuthorOpenid)
	linkReq.SetLiveExt(live.Ext)
	linkReq.SetShareType([]int{1, 3}) // deep link + 口令
	linkReq.SetPlatform(0)

	// 如果直播间有商品，也可以设置商品ID
	if len(live.Products) > 0 {
		linkReq.SetProductId(live.Products[0].ProductId)
	}

	var linkResp livelink.Response
	err = client.Exec(linkReq, &linkResp)
	if err != nil {
		t.Errorf("LiveLink failed: %v", err)
		return
	}

	fmt.Printf("直播转链响应: code=%d, desc=%s\n", linkResp.Code, linkResp.Desc)
	fmt.Printf("  deeplink: %s\n", linkResp.Data.DyDeeplink)
	fmt.Printf("  口令: %s\n", linkResp.Data.DyPassword)
	if linkResp.Data.DyZlink != "" {
		fmt.Printf("  zlink: %s\n", linkResp.Data.DyZlink)
	}
}

// TestOrderSearch 订单查询接口测试
func TestOrderSearch(t *testing.T) {
	client := GetClient()

	// 按时间范围查询订单
	req := &request.OrderSearchRequest{}
	// 查询最近7天的订单
	endTime := int(os.Getpid())    // 用当前时间作为示例
	startTime := endTime - 86400*7 // 7天前
	req.SetPageQuery(50, "0", startTime, endTime)
	req.SetOrderType(1) // 商品分销订单
	req.SetTimeType("pay")

	var resp ordersearch.Response
	err := client.Exec(req, &resp)
	if err != nil {
		t.Errorf("OrderSearch failed: %v", err)
		return
	}

	fmt.Printf("订单查询响应: code=%d, desc=%s, cursor=%s\n", resp.Code, resp.Desc, resp.Data.Cursor)
	fmt.Printf("  本次返回订单数: %d\n", len(resp.Data.Orders))
	for i, order := range resp.Data.Orders {
		if i >= 3 {
			fmt.Printf("... 共%d个订单\n", len(resp.Data.Orders))
			break
		}
		fmt.Printf("  订单%d: order_id=%s, 商品=%s, 金额=%d\n",
			i+1, order.OrderId, order.ProductName, order.TotalPayAmount)
	}
}

// TestOrderSearchByIds 按订单ID查询测试
func TestOrderSearchByIds(t *testing.T) {
	client := GetClient()

	// 如果有已知订单ID，可以按ID查询
	req := &request.OrderSearchRequest{}
	req.SetOrderIds([]string{"123456789", "987654321"})
	req.SetOrderType(1)

	var resp ordersearch.Response
	err := client.Exec(req, &resp)
	if err != nil {
		t.Errorf("OrderSearchByIds failed: %v", err)
		return
	}

	fmt.Printf("按ID查询订单响应: code=%d, desc=%s\n", resp.Code, resp.Desc)
	fmt.Printf("  返回订单数: %d\n", len(resp.Data.Orders))
}

// TestCommandParse 抖口令解析接口测试
// 先通过商品搜索转链获取抖口令，再解析该口令
func TestCommandParse(t *testing.T) {
	client := GetClient()

	//// 第一步：搜索商品并转链获取抖口令
	//searchReq := &request.ProductSearchRequest{}
	//searchReq.SetPage(1)
	//searchReq.SetPageSize(1)
	//
	//var searchResp productsearch.Response
	//err := client.Exec(searchReq, &searchResp)
	//if err != nil {
	//	t.Errorf("Search failed: %v", err)
	//	return
	//}
	//
	//if len(searchResp.Data.Products) == 0 {
	//	t.Log("No products found from search, skip CommandParse test")
	//	return
	//}
	//
	//product := searchResp.Data.Products[0]
	//
	//// 商品转链获取口令
	//linkReq := &request.ProductLinkRequest{}
	//linkReq.SetProductUrl(product.DetailUrl)
	//linkReq.SetProductExt(product.Ext)
	//linkReq.SetShareType([]int{1, 3}) // deep link + 口令
	//linkReq.SetPlatform(0)
	//
	//var linkResp productlink.Response
	//err = client.Exec(linkReq, &linkResp)
	//if err != nil {
	//	t.Errorf("ProductLink failed: %v", err)
	//	return
	//}
	//
	//if linkResp.Code != 0 {
	//	t.Logf("ProductLink not succeeded: code=%d, desc=%s", linkResp.Code, linkResp.Desc)
	//	return
	//}
	//
	//dyPassword := linkResp.Data.DyPassword
	//if dyPassword == "" {
	//	t.Log("No dy_password returned from ProductLink, skip CommandParse test")
	//	return
	//}
	//
	//fmt.Printf("获取到抖口令: %s\n", dyPassword)

	// 第二步：解析抖口令
	parseReq := &request.CommandParseRequest{}
	parseReq.SetCommand("9:/ c@n.Qx 07/06 fu致※※K6AxAs1Q859➝➝(🎶dou音+:%")
	parseReq.SetShareType([]int{1, 2, 3, 4, 5}) // 获取所有转链类型
	parseReq.SetPlatform(0)

	var parseResp commandparse.Response
	err := client.Exec(parseReq, &parseResp)
	if err != nil {
		t.Errorf("CommandParse failed: %v", err)
		return
	}

	fmt.Printf("抖口令解析响应: code=%d, desc=%s\n", parseResp.Code, parseResp.Desc)
	fmt.Printf("  command_type: %d (1=商品, 2=直播间, 3=活动页)\n", parseResp.Data.CommandType)
	fmt.Printf("  is_own_command: %v\n", parseResp.Data.IsOwnCommand)

	switch parseResp.Data.CommandType {
	case 1:
		if parseResp.Data.ProductInfo != nil {
			fmt.Printf("  商品信息: product_id=%d\n", parseResp.Data.ProductInfo.ProductId)
			if parseResp.Data.ProductInfo.ShareInfo != nil {
				si := parseResp.Data.ProductInfo.ShareInfo
				fmt.Printf("  分享物料: deeplink=%s, share_command=%s, zlink=%s, share_link=%s\n",
					si.DeepLink, si.ShareCommand, si.Zlink, si.ShareLink)
				if si.QrCode != nil {
					fmt.Printf("  二维码: url=%s, width=%d, height=%d\n",
						si.QrCode.URL, si.QrCode.Width, si.QrCode.Height)
				}
			}
		}
	case 2:
		if parseResp.Data.LiveInfo != nil {
			fmt.Printf("  直播间信息: author_buyin_id=%s, product_id=%d\n",
				parseResp.Data.LiveInfo.AuthorBuyinId, parseResp.Data.LiveInfo.ProductId)
			if parseResp.Data.LiveInfo.ShareInfo != nil {
				si := parseResp.Data.LiveInfo.ShareInfo
				fmt.Printf("  分享物料: deeplink=%s, share_command=%s\n", si.DeepLink, si.ShareCommand)
			}
		}
	case 3:
		if parseResp.Data.ActivityInfo != nil {
			fmt.Printf("  活动页信息: material_id=%s, extra_params=%s\n",
				parseResp.Data.ActivityInfo.MaterialId, parseResp.Data.ActivityInfo.ExtraParams)
		}
	}

	if parseResp.Data.PublicPlanProductLinkResultInfo != nil {
		fmt.Printf("  公共计划转链信息: %s\n", string(parseResp.Data.PublicPlanProductLinkResultInfo))
	}
}
