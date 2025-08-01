## 1688开放平台SDK

这是1688开放平台的Go语言SDK，用于对接1688开放平台API。

### 安装

```bash
go get github.com/mimicode/tksdk
```

### 签名规则

1688开放平台的签名规则如下：

1. 构造签名因子：urlPath。URL路径格式为：`param2/1/{业务模块}/{接口名}/{appKey}`
   - 例如：`param2/1/system/currentTime/1000000`
   - 例如：`param2/1/com.alibaba.p4p/alibaba.cps.listMediaInfo/1000000`

2. 构造签名因子：拼装的参数。参数按照首字母排序，然后将key和value拼在一起，最后按顺序拼在一起

3. 合并两个签名因子。把前两步的字符串拼起来

4. 对合并后的签名因子执行hmac_sha1算法。Signature=uppercase(hex(hmac_sha1(s, secretKey)))

### URL结构

完整的API请求URL格式为：
```
http://gw.open.1688.com/openapi/param2/1/{业务模块}/{接口名}/{appKey}?{参数}
```

### 配置文件

您可以在项目根目录下创建`env_dev.json`文件，用于存储API密钥和其他配置信息：

```json
{
  "i688": {
    "app_key": "your_app_key",
    "app_secret": "your_app_secret",
    "session_key": "your_session_key",
	"pid": "your_pid"
  }
}
```

### 使用示例

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/i688opensdk"
	"github.com/mimicode/tksdk/i688opensdk/request"
	"github.com/mimicode/tksdk/i688opensdk/response/system"
	"io/ioutil"
	"os"
)

func main() {
	// 从配置文件读取密钥
	var appKey, appSecret, sessionKey string
	
	// 尝试从env_dev.json读取配置
	if _, err := os.Stat("env_dev.json"); err == nil {
		if bytes, err := ioutil.ReadFile("env_dev.json"); err == nil {
			var data struct {
				I688 struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
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
	 
	// 创建客户端实例
	client := &i688opensdk.TopClient{}

	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)

	// 创建请求
	req := &request.SystemCurrentTimeRequest{}

	// 创建响应
	resp := &system.Response{}

	// 执行请求
	err := client.Exec(req, resp)
	if err != nil {
		fmt.Printf("执行请求失败: %v\n", err)
		return
	}

	// 处理响应
	if resp.IsError() {
		fmt.Printf("API返回错误: %v\n", resp.ErrorResponse)
		return
	}

	// 打印结果
	fmt.Printf("当前时间: %s\n", resp.CurrentTimeResponse.CurrentTime)
	fmt.Printf("请求ID: %s\n", resp.CurrentTimeResponse.RequestID)
}
```

## API列表

### 系统API
- [alibaba.system.currentTime( 获取系统当前时间 )](https://open.1688.com/api/system/currentTime.html)
  - 业务模块：`system`
  - 接口名：`currentTime`

### CPS推广API
- [alibaba.cps.listMediaInfo( 获取媒体推广位信息 )](https://open.1688.com/api/cps/listMediaInfo.html)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.listMediaInfo`
  - 参数：mediaId（可选）- 媒体ID