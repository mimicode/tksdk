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

### 类目API
- [alibaba.category.get( 根据类目Id查询类目 )](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product%3Aalibaba.category.get-1)
  - 业务模块：`com.alibaba.product`
  - 接口名：`alibaba.category.get`
  - 参数：categoryID（必需）- 类目ID，为0时查询所有一级类目

### CPS推广API
- [alibaba.cps.listMediaInfo( 获取媒体推广位信息 )](https://open.1688.com/api/cps/listMediaInfo.html)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.listMediaInfo`
  - 参数：mediaId（可选）- 媒体ID

- [alibaba.cps.listActivityPageQuery( 查询活动页面列表 )](https://open.1688.com/api/cps/listActivityPageQuery.html)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.listActivityPageQuery`

- [alibaba.cps.genSearchPjjxIntroduceUrlByKeyword( 根据关键词生成搜索推广链接 )](https://open.1688.com/api/cps/genSearchPjjxIntroduceUrlByKeyword.html)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.genSearchPjjxIntroduceUrlByKeyword`
  - 参数：keyword（必需）- 搜索关键词

- [alibaba.cps.getCpsRecommendSameOfferList( 获取CPS推荐相似商品列表 )](https://open.1688.com/api/cps/getCpsRecommendSameOfferList.html)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.getCpsRecommendSameOfferList`
  - 参数：offerId（必需）- 商品ID

- [alibaba.cps.listOfferPageQuery( 获取联盟offer列表 )](https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472cceFMBCLf&id=com.alibaba.p4p%3Aalibaba.cps.listOfferPageQuery-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.listOfferPageQuery`
  - 参数：
    - categoryId（可选）- 类目id，支持多级类目，用逗号分隔，如：0,1,128
    - feedInfo（可选）- 定向推广信息，格式为：定向推广id_定向推广类型，如：123456_1
    - pageNo（可选）- 页码，从1开始，默认1
    - pageSize（可选）- 每页数量，默认20，最大100
    - defineTags（可选）- 自定义标签，多个标签用逗号分隔
    - filterMinPrice（可选）- 最低价格过滤，单位为元
    - filterMaxPrice（可选）- 最高价格过滤，单位为元
    - sortField（可选）- 排序字段，可选值：gmv_desc(GMV降序)、price_asc(价格升序)、price_desc(价格降序)、commission_rate_desc(佣金比例降序)、commission_desc(佣金降序)、sale_desc(销量降序)
    - filterOldBuyerRatio（可选）- 老买家佣金比例过滤，范围0-1

## 新增接口开发指南

### 开发步骤

#### 1. 创建请求文件
在 `request/` 目录下创建接口请求文件，文件名格式：`{接口名小写无点}.go`

**文件结构模板：**
```go
package request

import (
	"net/url"
	"strconv" // 根据需要导入
)

// {接口名驼峰}Request {接口描述} API请求
// {接口文档URL}
type {接口名驼峰}Request struct {
	Parameters *url.Values //请求参数
}

// New{接口名驼峰}Request 创建{接口名驼峰}Request实例
func New{接口名驼峰}Request() *{接口名驼峰}Request {
	return &{接口名驼峰}Request{}
}

// CheckParameters 检查参数
func (r *{接口名驼峰}Request) CheckParameters() {
	// 检查必需参数
	if r.Parameters == nil || r.Parameters.Get("paramName") == "" {
		panic("paramName is required")
	}
}

// AddParameter 添加参数
func (r *{接口名驼峰}Request) AddParameter(key, val string) {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	r.Parameters.Add(key, val)
}

// Set{参数名} 设置{参数描述}
func (r *{接口名驼峰}Request) Set{参数名}(param string) {
	r.AddParameter("paramName", param)
}

// GetApiName 获取API名称
func (r *{接口名驼峰}Request) GetApiName() string {
	return "{接口原始名称}"
}

// GetApiVersion 获取API版本
func (r *{接口名驼峰}Request) GetApiVersion() string {
	return "param2/1"
}

// GetBusinessModule 获取业务模块
func (r *{接口名驼峰}Request) GetBusinessModule() string {
	return "{业务模块名}"
}

// GetParameters 获取所有参数
func (r *{接口名驼峰}Request) GetParameters() url.Values {
	if r.Parameters == nil {
		r.Parameters = &url.Values{}
	}
	return *r.Parameters
}
```

#### 2. 创建响应文件
在 `response/` 目录下创建接口响应目录和文件：
- 目录：`response/{接口名小写无点}/`
- 文件：`response/{接口名小写无点}/{接口名小写无点}.go`

**文件结构模板：**
```go
package {接口名小写无点}

import (
	"encoding/json"
	"github.com/mimicode/tksdk/i688opensdk/response"
)

// Response {接口描述}响应
type Response struct {
	response.TopResponse
	// 不同接口响应的字段不同大体分为如下几个类型
	// 结构体字段
	"{响应字段1驼峰}" {响应字段1结构体名} `json:"{响应字段1}"`
	// 非结构体字段
	"{响应字段2驼峰}" string `json:"{响应字段2}"`
	// 数组字段
	"{响应字段3驼峰}" []{响应字段3结构体名} `json:"{响应字段3}"`
	// map字段 
	"{响应字段4驼峰}" map[string]interface{} `json:"{响应字段4}"`	
}

// {响应字段1结构体名} {接口描述}响应结果
type {响应字段1结构体名} struct {
	// 根据API文档定义响应字段 
}

// WrapResult 包装结果
func (r *Response) WrapResult(result string) { 
	// 保存原始结果到Body
	r.Body = result 
	// 尝试解析具体的响应结构
	if err := json.Unmarshal([]byte(result), r); err != nil {
		// 如果解析失败，设置错误信息
		r.ErrorResponse.Code = 500
		r.ErrorResponse.Msg = "Failed to parse response: " + err.Error()
		return
	} 
	// 对于存在 errorMsg  errorCode 的响应内容 需要判断这两个值是否为空
	if r.ErrorMsg != "" || r.ErrorCode != "" {
		r.ErrorResponse.Code = -1
		r.ErrorResponse.Msg = fmt.Sprintf("errorMsg: %s, errorCode: %s", r.ErrorMsg, r.ErrorCode)
	} 
}
```

#### 3. 编写测试用例
在 `client_test.go` 文件中添加测试函数：

```go
func Test{接口名驼峰}(t *testing.T) {
	// 创建客户端实例
	client := &i688opensdk.TopClient{}
	// 初始化客户端
	client.Init(appKey, appSecret, sessionKey)
	// 创建请求
	req := &request.{接口名驼峰}Request{}
	// 设置参数
	req.Set{参数名}("参数值")
	// 创建响应
	resp := &{接口名小写无点}.Response{}
	// 执行请求
	err := client.Exec(req, resp)
	// 打印原始响应body
	t.Logf("body: %s", resp.Body)

	if err != nil {
		t.Fatalf("执行请求失败: %v", err)
	}

	// 处理响应
	if resp.IsError() {
		t.Fatalf("API返回错误: %v", resp.ErrorResponse)
	}

	// 打印结果进行验证
	fmt.Printf("响应结果: %+v\n", resp.{接口名驼峰}Response)
}
```

### 重要注意事项

#### 0. 请求字段和响应字段必须遵守
- 字段名必须与1688开放平台文档一致
- 字段类型必须与文档描述一致
- 字段是否必填必须与文档说明一致
- 不允许通过网络搜索文档来获取字段名和类型
- 不允许添加任何文档中未说明的字段

#### 1. JSON解析问题
- **字段名不一致**：1688 API返回的JSON字段名可能与官方文档不一致
- **拼写错误**：如 `success` 字段在某些接口中返回为 `succes`（少一个s）
- **嵌套结构**：实际返回的JSON结构可能比文档描述的更复杂，存在多层嵌套

#### 2. 响应结构处理
在 `WrapResult` 方法中需要特别注意：
- 先保存原始JSON到 `Body` 字段用于调试
- 使用多种解析策略处理不同的JSON结构
- 添加详细的错误处理和日志输出
- 根据实际测试结果调整解析逻辑

#### 3. 参数验证
- 在 `CheckParameters` 方法中验证必需参数
- 为每个参数提供专门的设置方法
- 注意参数类型转换（如int64转string）

#### 4. 业务模块识别
不同接口的业务模块可能不同：
- 系统接口：`system`
- CPS推广接口：`com.alibaba.p4p`
- 商品类目接口：`com.alibaba.product`
- 其他业务接口：参考官方文档

#### 5. 测试和调试
- 使用真实的API密钥进行测试
- 打印原始JSON响应用于调试
- 逐步调整响应结构直到正确解析
- 验证所有字段都能正确映射

#### 6. 代码规范
- 文件名使用小写，去掉接口名中的点号
- 结构体名使用驼峰命名
- 添加详细的注释和文档链接
- 遵循Go语言编码规范

### 开发流程总结
1. 查阅1688开放平台API文档
2. 创建请求文件，定义参数和方法
3. 创建响应文件，定义响应结构
4. 编写测试用例
5. 运行测试，查看原始JSON响应
6. 根据实际响应调整响应结构
7. 重复测试直到正确解析
8. 更新README文档中的API列表