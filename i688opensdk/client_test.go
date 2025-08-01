package i688opensdk_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mimicode/tksdk/i688opensdk"
	"github.com/mimicode/tksdk/i688opensdk/request"
	"github.com/mimicode/tksdk/i688opensdk/response/alibabacpslistmediainfo"
)

var (
	appKey, appSecret, sessionKey, pid string
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
