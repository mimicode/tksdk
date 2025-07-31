package i688opensdk_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mimicode/tksdk/i688opensdk"
	"github.com/mimicode/tksdk/i688opensdk/request"
	"github.com/mimicode/tksdk/i688opensdk/response/systemcurrenttime"
	"github.com/mimicode/tksdk/i688opensdk/response/alibabacpslistmediainfo"
)

var (
	appKey, appSecret, sessionKey, pid string
)

func init() {
	// 优先从env_dev.json读取配置
	if _, err := os.Stat("../env_dev.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../env_dev.json"); err == nil {
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
			}
		}
	}
}

func TestSystemCurrentTime(t *testing.T) {

	// 创建客户端实例
	client := &i688opensdk.TopClient{}

	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)

	// 创建请求
	req := &request.SystemCurrentTimeRequest{}

	// 创建响应
	resp := &systemcurrenttime.Response{}

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
	fmt.Printf("当前时间: %s\n", resp.CurrentTimeResponse.CurrentTime)
	fmt.Printf("请求ID: %s\n", resp.CurrentTimeResponse.RequestID)
}

func TestAlibabaCpsListMediaInfo(t *testing.T) {
	// 检查是否设置了有效的凭据
	if appKey == "" || appSecret == "" || appKey == "your_app_key" || appSecret == "your_app_secret" {
		t.Skip("跳过测试：未设置有效的I688_APP_KEY或I688_APP_SECRET")
		return
	}

	// 创建客户端实例
	client := &i688opensdk.TopClient{}

	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)

	// 创建请求
	req := &request.AlibabaCpsListMediaInfoRequest{}
	// 可选：设置媒体ID
	// req.SetMediaId(123456)

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
	fmt.Printf("媒体信息数量: %d\n", len(resp.AlibabaCpsListMediaInfoResponse.Result))
	for i, media := range resp.AlibabaCpsListMediaInfoResponse.Result {
		fmt.Printf("媒体%d - ID: %d, 名称: %s, 类型: %s, 状态: %s\n", 
			i+1, media.MediaId, media.MediaName, media.MediaType, media.MediaStatus)
	}
	fmt.Printf("请求ID: %s\n", resp.AlibabaCpsListMediaInfoResponse.RequestID)
}
