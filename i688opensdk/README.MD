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

### 订单API
- [alibaba.trade.getBuyerOrderList( 订单列表查看(买家视角) )](https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.trade%3Aalibaba.trade.getBuyerOrderList-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.trade`
  - 接口名：`alibaba.trade.getBuyerOrderList`
  - 参数：
    - bizTypes（可选）- 业务类型，支持："cn"(普通订单类型), "ws"(大额批发订单类型), "yp"(普通拿样订单类型), "yf"(一分钱拿样订单类型), "fs"(倒批(限时折扣)订单类型), "cz"(加工定制订单类型), "ag"(协议采购订单类型), "hp"(伙拼订单类型), "gc"(国采订单类型), "supply"(供销订单类型), "nyg"(nyg订单类型), "factory"(淘工厂订单类型), "quick"(快订下单), "xiangpin"(享拼订单), "nest"(采购商城-鸟巢), "f2f"(当面付), "cyfw"(存样服务), "sp"(代销订单标记), "wg"(微供订单), "factorysamp"(淘工厂打样订单), "factorybig"(淘工厂大货订单)
    - createEndTime（可选）- 下单结束时间，格式：20180802211113000+0800
    - createStartTime（可选）- 下单开始时间，格式：20180102211113000+0800
    - isHis（可选）- 是否查询历史订单表，默认false
    - modifyEndTime（可选）- 查询修改时间结束，格式：20180802211113000+0800
    - modifyStartTime（可选）- 查询修改时间开始，格式：20180102211113000+0800
    - orderStatus（可选）- 订单状态，值有：success, cancel(交易取消，违约金等交割完毕), waitbuyerpay(等待卖家付款), waitsellersend(等待卖家发货), waitbuyerreceive(等待买家收货)
    - page（可选）- 查询分页页码，从1开始
    - pageSize（可选）- 查询的每页的数量
    - refundStatus（可选）- 退款状态，支持："waitselleragree"(等待卖家同意), "refundsuccess"(退款成功), "refundclose"(退款关闭), "waitbuyermodify"(待买家修改), "waitbuyersend"(等待买家退货), "waitsellerreceive"(等待卖家确认收货)
    - sellerMemberId（可选）- 卖家memberId
    - sellerLoginId（可选）- 卖家loginId
    - sellerRateStatus（可选）- 卖家评价状态 (4:已评价,5:未评价,6;不需要评价)
    - tradeType（可选）- 交易类型：担保交易(1), 预存款交易(2), ETC境外收单交易(3), 即时到帐交易(4), 保障金安全交易(5), 统一交易流程(6), 分阶段交易(7), 货到付款交易(8), 信用凭证支付交易(9), 账期支付交易(10), 1688交易4.0，新分阶段交易(50060), 当面付的交易流程(50070), 服务类的交易流程(50080)
    - productName（可选）- 商品名称
    - needBuyerAddressAndPhone（可选）- 是否需要查询买家的详细地址信息和电话
    - needMemoInfo（可选）- 是否需要查询备注信息
    - outOrderId（可选）- 外部订单号，可用于控制幂等

- [alibaba.trade.get.buyerView( 订单详情查看(买家视角) )](https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.211d55edtSHFBC&id=com.alibaba.trade%3Aalibaba.trade.get.buyerView-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.trade`
  - 接口名：`alibaba.trade.get.buyerView`
  - 参数：
    - webSite（必需）- 站点信息，固定值：1688
    - orderId（必需）- 订单ID
    - includeFields（可选）- 返回字段，多个字段用逗号分隔，如：baseInfo,orderEntries

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

- [alibaba.cps.parseUrlOrToken( 解析推广物料信息 )](https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.parseUrlOrToken-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.parseUrlOrToken`
  - 参数：
    - urlOrToken（必需）- 待解析内容，示例值：111

- [alibaba.cps.tradeBillList( 联盟账单列表查询 )](https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.tradeBillList-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.tradeBillList`
  - 参数：
    - queryOrderType（必需）- 订单状态查询类型，可选值：orderAll(全部订单)、orderSettle(订单结算)、orderPay(订单付款)、orderMain(主订单号查询)、orderSub(子订单号查询)
    - queryTimeType（可选）- 订单时间查询类型，非订单号查询时必传，可选值：gmtCreateTime(订单创建)、confirmTime(订单成功)、settleTime(订单结算)、rightsStartTime(维权创建时间)、rightsEndTime(维权完成时间)
    - queryStartTime（可选）- 查询开始时间，非订单号查询时必传，格式：2017-12-01
    - queryEndTime（可选）- 查询结束时间，非订单号查询时必传，格式：2017-12-31
    - bizId（可选）- 主订单id，订单号查询时必传
    - bizSubId（可选）- 子订单id，订单号查询时必传
    - orderState（可选）- 订单状态，可选值：10(订单创建)、20(已付款)、50(交易成功)、80(交易关闭)
    - settleState（可选）- 结算状态，可选值：1(结算成功状态)、2(结算失败状态)
    - rightsState（可选）- 维权状态，可选值：10(维权中)、30(维权成功)、40(退款关闭)
    - pageNo（可选）- 页码，默认1
    - pageSize（可选）- 页面大小，默认100
    - orderId（可选）- 订单号，订单号查询时必传

- [alibaba.cps.listCPSSettleInfoDetail( 获取联盟结算明细列表(按月) )](https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.listCPSSettleInfoDetail-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.listCPSSettleInfoDetail`
  - 参数：
    - pageNo（必需）- 页偏移量
    - pageSize（必需）- 分页大小

- [alibaba.cps.getCPSSettleSummaryInfo( 获取联盟结算摘要账单 )](https://open.1688.com/api/apidocdetail.htm?spm=a260s.develop-solution-detail.0.0.36ad55edishgBb&id=com.alibaba.p4p%3Aalibaba.cps.getCPSSettleSummaryInfo-1&aopApiCategory=category_new)
  - 业务模块：`com.alibaba.p4p`
  - 接口名：`alibaba.cps.getCPSSettleSummaryInfo`
  - 参数：无需参数
  - 返回：昨日预估收入、本月预估收入、上月预估收入、预估未结算收入、可提现余额
