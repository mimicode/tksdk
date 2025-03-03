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
		if checkAPIExists(dir, apiRequest) {
			fmt.Printf("API已存在,跳过创建: %s.%s %s\n", apiRequest.APIORGNAME, apiRequest.APIMETHOD, apiRequest.APIVERSION)
			continue
		}
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

	// 检查文件是否已存在
	requestFile := filepath.Join(dir, "request", FILENAME)
	if _, err := os.Stat(requestFile); err == nil {
		fmt.Printf("请求文件已存在: %s\n", requestFile)
		return
	}

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
	if apiRequest.APIVERSION == "2.0.0" {
		response = strings.Replace(response, "--RESULT--", "success", -1)
		response = strings.Replace(response, "--RESULT2--", "Success", -1)
	} else {
		response = strings.Replace(response, "--RESULT--", "result", -1)
		response = strings.Replace(response, "--RESULT2--", "Result", -1)
	}

	//写入请求
	err = ioutil.WriteFile(requestFile, []byte(fileContent), os.ModePerm)
	fmt.Printf("创建请求文件: %s, err: %v\n", requestFile, err)

	//写入响应
	responseDir := filepath.Join(dir, "response", DIRNAME)
	if _, err = os.Stat(responseDir); os.IsNotExist(err) {
		err = os.MkdirAll(responseDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	responseFile := filepath.Join(responseDir, FILENAME)
	err = ioutil.WriteFile(responseFile, []byte(response), os.ModePerm)
	fmt.Printf("创建响应文件: %s, err: %v\n", responseFile, err)
}

func createAPIS() {
	var requests []ApiInfo

	request := ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 获取联盟在推商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "goodsListWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=goodsListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id结合的商品信息-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getByGoodsIdsWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=getByGoodsIdsWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 根据关键词查询商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "queryWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=queryWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 相似推荐商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "similarRecommendWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=similarRecommendWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 猜你喜欢商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "userRecommendWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=userRecommendWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService 根据商品id生成联盟链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "genByGoodsIdWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=genByGoodsIdWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService 根据唯品会链接生成联盟链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "genByVIPUrlWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=genByVIPUrlWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService 检测一段文本中是否有唯品会链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "vipLinkCheckWithOuth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=vipLinkCheckWithOuth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderService 获取订单信息列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderService",
		APIMETHOD:  "orderListWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=vipLinkCheckWithOuth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderService 维权订单总收益需要扣除该接口返回-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderService",
		APIMETHOD:  "refundOrderListWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=refundOrderListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 获取联盟在推商品列表V2-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "goodsListV2WithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=goodsListV2WithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id集合的商品信息(新)-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getByGoodsIdsV2WithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=getByGoodsIdsV2WithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPidV2Service 创建推广位PID-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPidV2Service",
		APIMETHOD:  "genPidWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidV2Service-2.0.0/genPidWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlV2Service 渠道根据落地页类型生链接口,需要申请渠道工具商权限包权限",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIMETHOD:  "getChannelUrlByTypeWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getChannelUrlByTypeWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlV2Service 根据openId判断用户是否是渠道用户信息, 注意:如果scene传值为1，是渠道新用户判断，而非唯品会新用户判断",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIMETHOD:  "checkUserWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/checkUserWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "vipapis.oauth.OauthService 刷新令牌服务",
		APIORGNAME: "vipapis.oauth.OauthService",
		APIMETHOD:  "refreshToken",
		APIVERSION: "1.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/vipapis.oauth.OauthService-1.0.0/refreshToken",
	}
	requests = append(requests, request)

	// 新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService goodsList 获取联盟在推商品列表",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "goodsList",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/goodsList",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService query 根据关键词查询商品列表",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "query",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/query",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService vipLinkCheck 进行cps链接的解析",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "vipLinkCheck",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/vipLinkCheck",
	}
	requests = append(requests, request)

	// 添加更多V2版本的API...
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service queryCompetitorMatchingProductWithOauth 获取外网商品匹配结果",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "queryCompetitorMatchingProductWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/queryCompetitorMatchingProductWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUserV2Service unbindOpenIdWithOauth 解绑已授权的openId接口服务(需要工具商权限)",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserV2Service",
		APIMETHOD:  "unbindOpenIdWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/unbindOpenIdWithOauth",
	}
	requests = append(requests, request)

	// 1.3.0 版本API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPromotionMaterialService 获取推广素材列表——需要授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPromotionMaterialService",
		APIMETHOD:  "getPromotionMaterialWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialService-1.3.0/getPromotionMaterialWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getCategorys 类目查询接口",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getCategorys",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getCategorys",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService getWxCodeWithOauth 生成微信小程序码,需要Oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "getWxCodeWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/getWxCodeWithOauth",
	}
	requests = append(requests, request)

	// 2.0.0 版本新API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service getPromotionMaterialDetailWithOauth 获取推广素材详情——需要授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service",
		APIMETHOD:  "getPromotionMaterialDetailWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0/getPromotionMaterialDetailWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderV2Service virtualOrderListWithOauth 获取虚拟订单信息列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderV2Service",
		APIMETHOD:  "virtualOrderListWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0/virtualOrderListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service getBrandSnListWithOauth 精选品牌列表查询-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "getBrandSnListWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/getBrandSnListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service getGoodsDetailMarketingWithOauth 获取指定商品id的商品信息(支持精准计算渠道用户到手价)-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "getGoodsDetailMarketingWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/getGoodsDetailMarketingWithOauth",
	}
	requests = append(requests, request)

	// 1.0.0 版本新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPidService queryPidWithOauth 查询推广位PID-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPidService",
		APIMETHOD:  "queryPidWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidService-1.0.0/queryPidWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUserService userVerifyWithOauth CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserService",
		APIMETHOD:  "userVerifyWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserService-1.0.0/userVerifyWithOauth",
	}
	requests = append(requests, request)

	// 1.3.0 版本新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getBrandDetailWithOauth 通过品牌SN查询品牌详情——需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getBrandDetailWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getBrandDetailWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getBrandSnListWithOauth 精选品牌列表查询-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getBrandSnListWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getBrandSnListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getGoodsDetailMarketingWithOauth 获取指定商品id的商品信息(支持精准计算渠道用户到手价)-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getGoodsDetailMarketingWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getGoodsDetailMarketingWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getGoodsListMarketingWithOauth 获取指定商品id集合的商品信息(支持精准计算渠道用户到手价)-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getGoodsListMarketingWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getGoodsListMarketingWithOauth",
	}
	requests = append(requests, request)

	// 2.0.0 版本新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPidV2Service queryPidWithOauth 查询推广位PID-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPidV2Service",
		APIMETHOD:  "queryPidWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPidV2Service-2.0.0/queryPidWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlV2Service genByVIPUrlWithOauth 根据唯品会链接生成联盟链接-需要申请渠道工具商权限包权限",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIMETHOD:  "genByVIPUrlWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/genByVIPUrlWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlV2Service vipLinkCheckWithOuth 进行cps链接的解析,需要申请渠道包权限进行Oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIMETHOD:  "vipLinkCheckWithOuth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/vipLinkCheckWithOuth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUserV2Service userVerifyWithOauth CPS联盟用户校验-渠道商，判断用户是否新老客、是否领取指定红包等",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserV2Service",
		APIMETHOD:  "userVerifyWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/userVerifyWithOauth",
	}
	requests = append(requests, request)

	// 1.0.0 版本新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionTaskService getSubsidyTaskGoodsInfo 补贴活动商品列表",
		APIORGNAME: "com.vip.adp.api.open.service.UnionTaskService",
		APIMETHOD:  "getSubsidyTaskGoodsInfo",
		APIVERSION: "1.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionTaskService-1.0.0/getSubsidyTaskGoodsInfo",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUserService checkUser 根据openId判断用户是否是渠道新用户",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserService",
		APIMETHOD:  "checkUser",
		APIVERSION: "1.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserService-1.0.0/checkUser",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUserService checkUserWithOauth 根据openId判断用户是否是渠道用户信息",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserService",
		APIMETHOD:  "checkUserWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserService-1.0.0/checkUserWithOauth",
	}
	requests = append(requests, request)

	// 1.3.0 版本新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getByGoodsIdsV2 获取指定商品id集合的商品信息(新)",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getByGoodsIdsV2",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getByGoodsIdsV2",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService queryByBrandSnWithOauth 根据品牌SN查询商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "queryByBrandSnWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/queryByBrandSnWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getBrandListDetailByJxCodeWithOauth 根据组货码获取品牌详情",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getBrandListDetailByJxCodeWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getBrandListDetailByJxCodeWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService getGoodsListByBrandAndJxCodeWithOauth 根据组货码与品牌获取商品列表详情",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getGoodsListByBrandAndJxCodeWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/getGoodsListByBrandAndJxCodeWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService getChannelUrlByType 渠道根据落地页类型生链接口",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "getChannelUrlByType",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/getChannelUrlByType",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderService virtualOrderList 获取订单信息列表",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderService",
		APIMETHOD:  "virtualOrderList",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderService-1.3.0/virtualOrderList",
	}
	requests = append(requests, request)

	// 2.0.0 版本新增API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlV2Service getWxCodeWithOauth 生成微信小程序码,需要Oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIMETHOD:  "getWxCodeWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/getWxCodeWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderV2Service orderListWithOauth 获取订单信息列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderV2Service",
		APIMETHOD:  "orderListWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0/orderListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderV2Service refundOrderListWithOauth 获取维权订单列表，订单归属人为授权用户",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderV2Service",
		APIMETHOD:  "refundOrderListWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionOrderV2Service-2.0.0/refundOrderListWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service queryWithOauth 根据关键词查询商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "queryWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/queryWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service similarRecommendWithOauth 相似推荐商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "similarRecommendWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/similarRecommendWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service userRecommendWithOauth 猜你喜欢商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "userRecommendWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/userRecommendWithOauth",
	}
	requests = append(requests, request)

	// 2.0.0 版本API补充
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service getPromotionMaterialWithOauth 获取推广素材列表——需要授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPromotionMaterialV2Service",
		APIMETHOD:  "getPromotionMaterialWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialV2Service-2.0.0/getPromotionMaterialWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlV2Service genByGoodsIdWithOauth 根据商品id生成联盟链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlV2Service",
		APIMETHOD:  "genByGoodsIdWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlV2Service-2.0.0/genByGoodsIdWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service getByGoodsIdsV2WithOauth 获取指定商品id集合的商品信息(新)-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "getByGoodsIdsV2WithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/getByGoodsIdsV2WithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service getBrandDetailWithOauth 通过品牌SN查询品牌详情——需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "getBrandDetailWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/getBrandDetailWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service goodsListV2WithOauth 获取联盟在推商品列表V2-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "goodsListV2WithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/goodsListV2WithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service queryByBrandSnWithOauth 根据品牌SN查询商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "queryByBrandSnWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/queryByBrandSnWithOauth",
	}
	requests = append(requests, request)

	// 1.3.0 版本补充API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService goodsListV2 获取联盟在推商品列表-V2版本",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "goodsListV2",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsService-1.3.0/goodsListV2",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionPromotionMaterialService getPromotionMaterialDetailWithOauth 获取推广素材详情——需要授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionPromotionMaterialService",
		APIMETHOD:  "getPromotionMaterialDetailWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionPromotionMaterialService-1.3.0/getPromotionMaterialDetailWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService getChannelUrlByTypeWithOauth 渠道根据落地页类型生链接口,需要Oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "getChannelUrlByTypeWithOauth",
		APIVERSION: "1.3.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUrlService-1.3.0/getChannelUrlByTypeWithOauth",
	}
	requests = append(requests, request)

	// 补充遗漏的API
	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsV2Service getGoodsListMarketingWithOauth 获取指定商品id集合的商品信息(支持精准计算渠道用户到手价)-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsV2Service",
		APIMETHOD:  "getGoodsListMarketingWithOauth",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionGoodsV2Service-2.0.0/getGoodsListMarketingWithOauth",
	}
	requests = append(requests, request)

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUserV2Service checkUser 根据openId判断用户是否是渠道新用户",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUserV2Service",
		APIMETHOD:  "checkUser",
		APIVERSION: "2.0.0",
		APIURL:     "https://vop.vip.com/home#/api/method/detail/com.vip.adp.api.open.service.UnionUserV2Service-2.0.0/checkUser",
	}
	requests = append(requests, request)

	createAPIs(requests) // 使用新的批量创建函数
}

func main() {
	//
	createAPIS()
	createReadmeApiList()
}
