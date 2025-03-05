package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func createReadmeApiList() {
	dir, _ := os.Getwd()
	sdkPath := filepath.Join(dir, "vipopensdk")
	if matches, err := filepath.Glob(filepath.Join(sdkPath, "request", "*.go")); err != nil {
		fmt.Println(err)
		return
	} else {
		compile := regexp.MustCompile(`//(.+)`)
		var apiList []string
		for _, fileName := range matches {
			bytes, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			all := compile.FindAllSubmatch(bytes, -1)
			name := string(all[0][1])
			url := string(all[1][1])
			apiList = append(apiList, fmt.Sprintf("- [%s](%s)", strings.TrimSpace(name), strings.TrimSpace(url)))
		}
		file, err := os.OpenFile(filepath.Join(sdkPath, "README.MD"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0655)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		file.WriteString("## api列表\n")
		file.WriteString("--\n")
		file.WriteString(strings.Join(apiList, "\n"))
	}
}

type ApiInfo struct {
	APIDESC     string
	APIORGNAME  string
	APIMETHOD   string
	APIVERSION  string
	APIURL      string
	CHECKFIELDS []CheckField
}

type CheckField struct {
	FieldName string

	FieldNotNullCheck  bool //非空检测
	FieldMinCheck      bool //最小值检测
	FieldMaxCheck      bool //最大值检测
	FieldNumberCheck   bool //数字检测
	FieldLenCheck      bool //字符串长度检测
	FieldListSizeCheck bool //集合数量检测

	FieldMin  string
	FieldMax  string
	FieldLen  string
	FieldSize string
}

func strFirstToUpper(str string, step string) string {
	temp := strings.Split(str, step)
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if vv[i] >= 97 && vv[i] <= 122 && i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func createMustCheck(checkFields []CheckField) string {
	checkStr := ""
	for _, v := range checkFields {
		//非空检测
		if v.FieldNotNullCheck {
			checkStr += `utils.CheckNotNull(tk.Parameters.Get("` + v.FieldName + `"), "` + v.FieldName + `")
`
		}
		//最小值检测
		if v.FieldMinCheck {
			checkStr += `utils.CheckMinValue(tk.Parameters.Get("` + v.FieldName + `"),` + v.FieldMin + `, "` + v.FieldName + `")
`
		}
		//最大值检测
		if v.FieldMaxCheck {
			checkStr += `utils.CheckMaxValue(tk.Parameters.Get("` + v.FieldName + `"),` + v.FieldMax + `, "` + v.FieldName + `")
`
		}
		//数值检测
		if v.FieldNumberCheck {
			checkStr += `utils.CheckNumber(tk.Parameters.Get("` + v.FieldName + `"), "` + v.FieldName + `")
`
		}
		//值长度检测
		if v.FieldLenCheck {
			checkStr += `utils.CheckMaxLength(tk.Parameters.Get("` + v.FieldName + `"),` + v.FieldLen + `, "` + v.FieldName + `")
`
		}
		//集合长度检测
		if v.FieldListSizeCheck {
			checkStr += `utils.CheckMaxListSize(tk.Parameters.Get("` + v.FieldName + `"),` + v.FieldSize + `, "` + v.FieldName + `")
`
		}
	}

	return checkStr
}

func checkAPIExists(dir string, apiRequest ApiInfo) bool {
	DIRNAME := strings.ToLower(strings.Replace(apiRequest.APIORGNAME, ".", "", -1)) + strings.ToLower(apiRequest.APIMETHOD)
	FILENAME := DIRNAME + ".go"
	responseDir := filepath.Join(dir, "response", DIRNAME)
	responseFile := filepath.Join(responseDir, FILENAME)

	if _, err := os.Stat(responseFile); err == nil {
		return true
	}
	return false
}

func createAPIs(apiRequests []ApiInfo) {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1) + "/vipopensdk"

	for _, apiRequest := range apiRequests {
		//if checkAPIExists(dir, apiRequest) {
		//	fmt.Printf("API已存在,跳过创建: %s.%s %s\n", apiRequest.APIORGNAME, apiRequest.APIMETHOD, apiRequest.APIVERSION)
		//	continue
		//}
		createAPI(apiRequest)
	}
}

func createAPI(apiRequest ApiInfo) {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1) + "/vipopensdk"
	tplrequest := filepath.Join(dir, "generate/tplrequest.tpl")

	bytes, err := ioutil.ReadFile(tplrequest)
	if err != nil {
		panic(err)
	}

	APIPARAMCHECK := createMustCheck(apiRequest.CHECKFIELDS)
	UTIL := `"github.com/mimicode/tksdk/utils"`
	DIRNAME := strings.ToLower(strings.Replace(apiRequest.APIORGNAME, ".", "", -1)) + strings.ToLower(apiRequest.APIMETHOD)
	FILENAME := DIRNAME + ".go"

	APINAME := strFirstToUpper(apiRequest.APIORGNAME, ".") + strFirstToUpper(apiRequest.APIMETHOD, "-")
	//请求文件
	fileContent := string(bytes)
	fileContent = strings.Replace(fileContent, "--APIDESC--", apiRequest.APIDESC, -1)
	fileContent = strings.Replace(fileContent, "--APIORGNAME--", apiRequest.APIORGNAME, -1)
	fileContent = strings.Replace(fileContent, "--APIURL--", apiRequest.APIURL, -1)
	fileContent = strings.Replace(fileContent, "--APINAME--", APINAME, -1)
	fileContent = strings.Replace(fileContent, "--APIMETHOD--", apiRequest.APIMETHOD, -1)
	fileContent = strings.Replace(fileContent, "--APIVERSION--", apiRequest.APIVERSION, -1)
	if APIPARAMCHECK == "" {
		UTIL = ""
	}
	fileContent = strings.Replace(fileContent, "--APIPARAMCHECK--", APIPARAMCHECK, -1)
	fileContent = strings.Replace(fileContent, "--UTIL--", UTIL, -1)

	tplresponse := filepath.Join(dir, "generate/tplresponse.tpl")
	//响应文件
	responseByte, err := ioutil.ReadFile(tplresponse)
	if err != nil {
		panic(err)
	}
	response := string(responseByte)

	response = strings.Replace(response, "--APIDESC--", apiRequest.APIDESC, -1)
	response = strings.Replace(response, "--PACKNAME--", DIRNAME, -1)
	response = strings.Replace(response, "--APIMETHOD--", apiRequest.APIMETHOD, -1)
	response = strings.Replace(response, "--APIVERSION--", apiRequest.APIVERSION, -1)
	response = strings.Replace(response, "--APIORGNAME--", apiRequest.APIORGNAME, -1)
	response = strings.Replace(response, "--APIURL--", apiRequest.APIURL, -1)

	if apiRequest.APIVERSION == "2.0.0" {
		response = strings.Replace(response, "--RESULT--", "success", -1)
		response = strings.Replace(response, "--RESULT2--", "Success", -1)
	} else {
		response = strings.Replace(response, "--RESULT--", "result", -1)
		response = strings.Replace(response, "--RESULT2--", "Result", -1)
	}

	// 检查文件是否已存在
	requestFile := filepath.Join(dir, "request", FILENAME)
	if _, err := os.Stat(requestFile); err == nil {
		fmt.Printf("请求文件已存在: %s\n", requestFile)
	} else {
		//写入请求
		err = ioutil.WriteFile(requestFile, []byte(fileContent), os.ModePerm)
		fmt.Printf("创建请求文件: %s, err: %v\n", requestFile, err)
	}

	responseDir := filepath.Join(dir, "response", DIRNAME)
	if _, err = os.Stat(responseDir); os.IsNotExist(err) {
		err = os.MkdirAll(responseDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	responseFile := filepath.Join(responseDir, FILENAME)
	if _, err := os.Stat(responseFile); err == nil {
		fmt.Printf("响应文件已存在: %s\n", responseFile)
	} else {
		//写入响应
		err = ioutil.WriteFile(responseFile, []byte(response), os.ModePerm)
		fmt.Printf("创建响应文件: %s, err: %v\n", responseFile, err)
	}
}

func createAPIS() {
	var requests []ApiInfo

	// UnionGoodsV2Service - 商品服务
	requests = append(requests, ApiInfo{
		APIDESC:    "通过品牌SN查询品牌详情——需要oauth授权",
		APIMETHOD:  "getBrandDetailWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getBrandDetailWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "精选品牌列表查询-需要oauth授权",
		APIMETHOD:  "getBrandSnListWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getBrandSnListWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "获取指定商品id集合的商品信息(新)-需要oauth授权",
		APIMETHOD:  "getByGoodsIdsV2WithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getByGoodsIdsV2WithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "获取指定商品id的商品信息(支持精准计算渠道用户到手价)-需要oauth授权",
		APIMETHOD:  "getGoodsDetailMarketingWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=getGoodsDetailMarketingWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "获取联盟在推商品列表V2-需要oauth授权",
		APIMETHOD:  "goodsListV2WithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=goodsListV2WithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "根据关键词查询商品列表-需要oauth授权",
		APIMETHOD:  "queryWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=queryWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "根据品牌SN查询商品列表-需要oauth授权",
		APIMETHOD:  "queryByBrandSnWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=queryByBrandSnWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "猜你喜欢商品列表-需要oauth授权",
		APIMETHOD:  "userRecommendWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=userRecommendWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "相似推荐商品列表-需要oauth授权",
		APIMETHOD:  "similarRecommendWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0&methodName=similarRecommendWithOauth",
	})

	// UnionOrderV2Service - 订单服务
	requests = append(requests, ApiInfo{
		APIDESC:    "获取订单列表V2-需要oauth授权",
		APIMETHOD:  "orderListWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=orderListWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "获取维权订单列表V2-需要oauth授权",
		APIMETHOD:  "refundOrderListWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=refundOrderListWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "获取虚拟订单列表V2-需要oauth授权",
		APIMETHOD:  "virtualOrderListWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0&methodName=virtualOrderListWithOauth",
	})

	// UnionPidV2Service - PID服务
	requests = append(requests, ApiInfo{
		APIDESC:    "生成PID-需要oauth授权",
		APIMETHOD:  "genPidWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPidV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPidV2Service-2.0.0&methodName=genPidWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "查询PID-需要oauth授权",
		APIMETHOD:  "queryPidWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPidV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPidV2Service-2.0.0&methodName=queryPidWithOauth",
	})

	// UnionUserV2Service - 用户服务
	requests = append(requests, ApiInfo{
		APIDESC:    "根据openId判断用户是否是渠道用户信息, 注意:如果scene传值为1，是渠道新用户判断，而非唯品会新用户判断",
		APIMETHOD:  "checkUserWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/checkUserWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "解绑已授权的openId接口服务(需要工具商权限)",
		APIMETHOD:  "unbindOpenIdWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/unbindOpenIdWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等",
		APIMETHOD:  "userVerifyWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/userVerifyWithOauth",
	})

	// UnionUrlV2Service - URL服务
	requests = append(requests, ApiInfo{
		APIDESC:    "根据商品id生成联盟链接-需要oauth授权",
		APIMETHOD:  "genByGoodsIdWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/genByGoodsIdWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "根据唯品会链接生成联盟链接-需要申请渠道工具商权限包权限",
		APIMETHOD:  "genByVIPUrlWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/genByVIPUrlWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "渠道根据落地页类型生链接口,需要申请渠道工具商权限包权限",
		APIMETHOD:  "getChannelUrlByTypeWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getChannelUrlByTypeWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "生成微信小程序码,需要Oauth授权",
		APIMETHOD:  "getWxCodeWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getWxCodeWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "进行cps链接的解析,需要申请渠道包权限进行Oauth授权",
		APIMETHOD:  "vipLinkCheckWithOuth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/vipLinkCheckWithOuth",
	})

	// UnionPromotionMaterialV2Service - 推广素材服务
	requests = append(requests, ApiInfo{
		APIDESC:    "获取推广素材详情——需要授权",
		APIMETHOD:  "getPromotionMaterialDetailWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0&methodName=getPromotionMaterialDetailWithOauth",
	})

	requests = append(requests, ApiInfo{
		APIDESC:    "获取推广素材列表——需要授权",
		APIMETHOD:  "getPromotionMaterialWithOauth",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service",
		APIVERSION: "2.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0&methodName=getPromotionMaterialWithOauth",
	})

	createAPIs(requests)
}

func main() {
	//
	createAPIS()
	createReadmeApiList()
}
