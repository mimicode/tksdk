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
- 字段有示例的必须把示例也写到字段的注释上
- 不允许通过网络搜索文档来获取字段名和类型
- 不允许添加任何文档中未说明的字段
#### 1. 禁止修改文件
- 禁止修改 i688opensdk/response/topresponse.go

#### 2. JSON解析问题
- **字段名不一致**：1688 API返回的JSON字段名可能与官方文档不一致
- **拼写错误**：如 `success` 字段在某些接口中返回为 `succes`（少一个s）
- **嵌套结构**：实际返回的JSON结构可能比文档描述的更复杂，存在多层嵌套

#### 3. 响应结构处理    
在 `WrapResult` 方法中需要特别注意：
- 先保存原始JSON到 `Body` 字段用于调试
- 使用多种解析策略处理不同的JSON结构
- 添加详细的错误处理和日志输出
- 根据实际测试结果调整解析逻辑

#### 4. 参数验证
- 在 `CheckParameters` 方法中验证必需参数
- 为每个参数提供专门的设置方法
- 注意参数类型转换（如int64转string）

#### 5. 业务模块识别
不同接口的业务模块可能不同：
- 系统接口：`system`
- CPS推广接口：`com.alibaba.p4p`
- 商品类目接口：`com.alibaba.product`
- 其他业务接口：参考官方文档

#### 6. 测试和调试
- 使用真实的API密钥进行测试
- 打印原始JSON响应用于调试
- 逐步调整响应结构直到正确解析
- 验证所有字段都能正确映射

#### 7. 代码规范
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