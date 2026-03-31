package pangolinopensdk

import (
	"testing"

	"github.com/mimicode/tksdk/pangolinopensdk/request"
	"github.com/mimicode/tksdk/pangolinopensdk/response/livesearch"
	"github.com/mimicode/tksdk/pangolinopensdk/response/ordersearch"
	"github.com/mimicode/tksdk/pangolinopensdk/response/productcategory"
	"github.com/mimicode/tksdk/pangolinopensdk/response/productdetail"
	"github.com/mimicode/tksdk/pangolinopensdk/response/productlink"
	"github.com/mimicode/tksdk/pangolinopensdk/response/productsearch"
)

// TestProductSearch 商品列表接口测试
func TestProductSearch(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.ProductSearchRequest{}
	req.SetPage(1)
	req.SetPageSize(20)
	req.SetTitle("手机")

	var resp productsearch.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("ProductSearch failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, total=%d", resp.Code, resp.Desc, resp.Data.Total)
}

// TestProductLink 商品转链接口测试
func TestProductLink(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.ProductLinkRequest{}
	req.SetProductUrl("https://www.a.com/products/1")
	req.SetShareType([]int{1, 2, 3})
	req.SetPlatform(0)

	var resp productlink.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("ProductLink failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, deeplink=%s", resp.Code, resp.Desc, resp.Data.DyDeeplink)
}

// TestProductDetail 商品详情接口测试
func TestProductDetail(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.ProductDetailRequest{}
	req.SetProductIds([]int64{123456})

	var resp productdetail.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("ProductDetail failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, total=%d", resp.Code, resp.Desc, resp.Data.Total)
}

// TestProductCategory 商品类目接口测试
func TestProductCategory(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.ProductCategoryRequest{}
	req.SetParentId(0) // 查询一级类目
	req.SetChannel(1) // 新类目

	var resp productcategory.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("ProductCategory failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, total=%d", resp.Code, resp.Desc, resp.Data.Total)
}

// TestLiveSearch 直播列表接口测试
func TestLiveSearch(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.LiveSearchRequest{}
	req.SetPage(1)
	req.SetPageSize(20)
	req.SetStatus(1) // 只看在播

	var resp livesearch.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("LiveSearch failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, total=%d", resp.Code, resp.Desc, resp.Data.Total)
}

// TestLiveLink 直播转链接口测试
func TestLiveLink(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.LiveLinkRequest{}
	req.SetAuthorOpenid("author_openid_example")
	req.SetShareType([]int{1, 2, 3})
	req.SetPlatform(0)

	var resp productlink.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("LiveLink failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, deeplink=%s", resp.Code, resp.Desc, resp.Data.DyDeeplink)
}

// TestOrderSearch 订单查询接口测试
func TestOrderSearch(t *testing.T) {
	client := &TopClient{}
	client.Init("your_app_id", "your_user_id", "your_role_id", "your_secure_key")

	req := request.OrderSearchRequest{}
	req.SetPageQuery(50, "0", 1735689600, 1738300800) // 2025-01-01 to 2025-02-01
	req.SetOrderType(1) // 商品分销订单

	var resp ordersearch.Response
	err := client.Execute(req.GetApiName(), &req, &resp)
	if err != nil {
		t.Errorf("OrderSearch failed: %v", err)
	}
	t.Logf("Response: code=%d, desc=%s, cursor=%s", resp.Code, resp.Desc, resp.Data.Cursor)
}
