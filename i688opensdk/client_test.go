package i688opensdk_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/mimicode/tksdk/i688opensdk"
	"github.com/mimicode/tksdk/i688opensdk/request"
	"github.com/mimicode/tksdk/i688opensdk/response/alibabacpsgensearchpjjxintroduceurlbykeyword"
	"github.com/mimicode/tksdk/i688opensdk/response/alibabacpsgetcpsrecommendsameofferlist"
	"github.com/mimicode/tksdk/i688opensdk/response/alibabacpslistactivitypagequery"
	"github.com/mimicode/tksdk/i688opensdk/response/alibabacpslistmediainfo"
)

var (
	appKey, appSecret, sessionKey, pid string
	mediaID, mediaZoneID               int64
	offerID                            int64 = 761411650121
)

func init() {
	// 优先从env_dev.json读取配置
	if _, err := os.Stat("/Users/zhang/project/golang/tksdk/dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("/Users/zhang/project/golang/tksdk/dev_env.json"); err == nil {
			var data struct {
				I688 struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
					PID        string `json:"pid"`
				} `json:"i688"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.I688.AppKey
				appSecret = data.I688.AppSecret
				sessionKey = data.I688.SessionKey
				pid = data.I688.PID
				if pid != "" {
					// mm_2214357383700_3168003_3273003
					parts := strings.Split(pid, "_")
					if len(parts) == 4 {
						mediaID, _ = strconv.ParseInt(parts[2], 10, 64)
						mediaZoneID, _ = strconv.ParseInt(parts[3], 10, 64)
					}
				}

				fmt.Fprintf(os.Stdout, "appKey: %s, appSecret: %s, sessionKey: %s, mediaID: %d, mediaZoneID: %d\n", appKey, appSecret, sessionKey, mediaID, mediaZoneID)
			}

		}
	}
}

func TestAlibabaCpsListMediaInfo(t *testing.T) {
	// 创建客户端实例
	client := &i688opensdk.TopClient{}
	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)
	// 创建请求
	req := &request.AlibabaCpsListMediaInfoRequest{}
	// 创建响应
	resp := &alibabacpslistmediainfo.Response{}
	// 执行请求
	err := client.Exec(req, resp)
	if err != nil {
		t.Fatalf("执行请求失败: %v", err)
	}

	// 处理响应
	if resp.IsError() {
		t.Fatalf("API返回错误: %v", resp.ErrorResponse)
	}

	// 打印结果
	fmt.Printf("媒体信息数量: %d\n", len(resp.AlibabaCpsListMediaInfoResponse))
	for i, media := range resp.AlibabaCpsListMediaInfoResponse {
		fmt.Printf("媒体%d - ID: %d, 名称: %s, 类型: %d, 状态: %d\n",
			i+1, media.MediaID, media.MediaTitle, media.MediaType, media.AuditState)
	}
}

func TestAlibabaCpsListActivityPageQuery(t *testing.T) {
	// 创建客户端实例
	client := &i688opensdk.TopClient{}
	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)
	// 创建请求
	req := &request.AlibabaCpsListActivityPageQueryRequest{}
	// 创建响应
	resp := &alibabacpslistactivitypagequery.Response{}
	// 执行请求
	err := client.Exec(req, resp)
	if err != nil {
		t.Fatalf("执行请求失败: %v", err)
	}

	// 处理响应
	if resp.IsError() {
		t.Fatalf("API返回错误: %v", resp.ErrorResponse)
	}

	// 打印结果
	fmt.Printf("活动总数: %d\n", resp.TotalRow)
	fmt.Printf("当前页活动数量: %d\n", len(resp.AlibabaCpsListActivityPageQueryResponse))
	for i, activity := range resp.AlibabaCpsListActivityPageQueryResponse {
		fmt.Printf("活动%d - ID: %d, 标题: %s, 类型: %d, 开始时间: %s, 结束时间: %s, 佣金比例: %.2f, 商品数量: %d\n",
			i+1, activity.ID, activity.Title, activity.Type, activity.StartDate, activity.EndDate, activity.Ratio, activity.ProductCnt)
	}
}

func TestAlibabaCpsGenSearchPjjxIntroduceUrlByKeyword(t *testing.T) {
	// 创建客户端实例
	client := &i688opensdk.TopClient{}
	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)
	// 创建请求
	req := &request.AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordRequest{}
	// 设置必填参数
	req.SetKeyword("女装")
	req.SetMediaId(mediaID)
	req.SetMediaZoneId(mediaZoneID)
	// 设置可选参数
	req.SetExt("{\"p1\":\"123\",\"p2\":\"456\",\"p3\":\"789\"}")
	// 创建响应
	resp := &alibabacpsgensearchpjjxintroduceurlbykeyword.Response{}
	// 执行请求
	err := client.Exec(req, resp)
	t.Logf("body: %s", resp.Body)
	if err != nil {
		t.Fatalf("执行请求失败: %v", err)
	}

	// 处理响应
	if resp.IsError() {
		t.Fatalf("API返回错误: %v", resp.ErrorResponse)
	}

	// 打印结果
	if resp.AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse != nil {
		fmt.Printf("结果成功: %t\n", resp.AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse.Success)
		if resp.AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse.Result != nil {
			fmt.Printf("PC推广链接: %s\n", resp.AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse.Result.PcPjjxCpsUrl)
			fmt.Printf("移动推广链接: %s\n", resp.AlibabaCpsGenSearchPjjxIntroduceUrlByKeywordResponse.Result.WapPjjxCpsUrl)
		}
	} else {
		fmt.Println("响应结果为空")
	}
}

func TestAlibabaCpsGetCpsRecommendSameOfferList(t *testing.T) {
	// 创建客户端实例
	client := &i688opensdk.TopClient{}
	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)
	// 创建请求
	req := &request.AlibabaCpsGetCpsRecommendSameOfferListRequest{}
	// 设置必填参数
	req.SetOfferId(offerID)
	req.SetOfferType("1688")
	req.SetMediaZoneId(mediaZoneID)
	req.SetMediaId(mediaID)
	// 设置可选参数
	// req.SetMinPrice(2.0)
	// req.SetMaxPrice(100.12)
	req.SetExt("{}")
	// 创建响应
	resp := &alibabacpsgetcpsrecommendsameofferlist.Response{}
	// 执行请求
	err := client.Exec(req, resp)
	t.Logf("body: %s", resp.Body)
	if err != nil {
		t.Fatalf("执行请求失败: %v", err)
	}

	// 处理响应
	if resp.IsError() {
		t.Fatalf("API返回错误: %v", resp.ErrorResponse)
	}
	// 打印结果
	if resp.AlibabaCpsGetCpsRecommendSameOfferListResponse != nil {
		fmt.Printf("请求成功: %t\n", resp.AlibabaCpsGetCpsRecommendSameOfferListResponse.Success)
		if resp.AlibabaCpsGetCpsRecommendSameOfferListResponse.ErrorCode != "" {
			fmt.Printf("错误码: %s\n", resp.AlibabaCpsGetCpsRecommendSameOfferListResponse.ErrorCode)
			fmt.Printf("错误信息: %s\n", resp.AlibabaCpsGetCpsRecommendSameOfferListResponse.ErrorMsg)
		}
		fmt.Printf("推荐商品数量: %d\n", len(resp.AlibabaCpsGetCpsRecommendSameOfferListResponse.Result))
		for i, offer := range resp.AlibabaCpsGetCpsRecommendSameOfferListResponse.Result {
			fmt.Printf("商品%d - ID: %d, 标题: %s, 价格: %.2f, 月销量: %d\n",
				i+1, offer.OfferId, offer.Title, offer.Price, offer.SaleCount)
			if offer.LongClickUrl != "" {
				fmt.Printf("  推广链接: %s\n", offer.LongClickUrl)
			}
		}
	} else {
		fmt.Println("响应结果为空")
	}
}
