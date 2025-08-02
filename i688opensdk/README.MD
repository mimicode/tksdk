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

- [alibaba.cps.listShopPageQuery( 获取联盟商家列表 )](https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472ccel58TEV&id=com.alibaba.p4p%3Aalibaba.cps.listShopPageQuery-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.listShopPageQuery`
  - 参数：
    - sellerId（可选）- 卖家id
    - companyName（可选）- 公司名称
    - categoryId（必需）- 类目id，0表示全部类目
    - defineTags（可选）- 属性标签，如：slsj_flag=true,sdrd_flag=true
    - filterRatioMin（可选）- 佣金比例下限
    - filterRatioMax（可选）- 佣金比例上限
    - sortField（可选）- 排序字段，如：ratio^asc（佣金比例升序）、productCnt^asc（商品数量升序）、tkCnt^asc（30天推广量升序）
    - pageNo（必需）- 页偏移量，从1开始
    - pageSize（必需）- 分页大小，最大100

- [alibaba.cps.genClickUrl( 批量生成联盟推广url点击信息 )](https://open.1688.com/api/apidocdetail.htm?spm=1688open.solution-detail.0.0.1d472ccel58TEV&id=com.alibaba.p4p%3Aalibaba.cps.genClickUrl-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.genClickUrl`
  - 参数：
    - type（必需）- 转链接类型(店铺、商品、活动) 0商品;1店铺;2活动;7优惠券; 12 图搜
    - mediaId（必需）- 媒体id
    - mediaZoneId（必需）- 媒体推广位id
    - objectValueList（必需）- 推广实体 商品：传offerId；活动：传活动id；商铺：卖家sellerId；优惠券：offerId ; 图搜传：图片url
    - ext（可选）- 其他自定义参数，查询订单会返回该参数
