# 淘客sdk

- 京东
- 淘宝
- 拼多多
- 苏宁
- 唯品会
## Installation

Download and install it:

```sh
$ go get -u github.com/mimicode/tksdk
```

Import it in your code:

```go
import "github.com/mimicode/tksdk/tbopensdk"
```
Examples
```go
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
```

### 微信交流群
![微信交流群](tksdkqun.png)