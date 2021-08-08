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
	request := filepath.Join(dir, "request")
	//fmt.Println(request)
	if matches, err := filepath.Glob(request + `/*\.go`); err != nil {
		fmt.Println(err)
		return
	} else {
		compile := regexp.MustCompile(`//(.+)`)
		for _, fileName := range matches {
			bytes, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			all := compile.FindAllSubmatch(bytes, -1)
			name := string(all[0][1])
			url := string(all[1][1])

			sprintf := fmt.Sprintf("- [%s](%s)", strings.TrimSpace(name), strings.TrimSpace(url))

			fmt.Println(sprintf)

		}
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
			if i == 0 {
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

func createAPI(apiRequest ApiInfo) {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1) + "/vipopensdk"
	tplrequest := filepath.Join(dir, "generate/tplrequest.tpl")

	bytes, err := ioutil.ReadFile(tplrequest)
	if err != nil {
		panic(err)
	}
	//"taobaoopensdk/utils"
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
	requestFile := filepath.Join(dir, "request/"+FILENAME)

	//写入请求
	err = ioutil.WriteFile(requestFile, []byte(fileContent), os.ModePerm)
	fmt.Println(err, "request")
	//写入响应
	responseDir := filepath.Join(dir, "response/"+DIRNAME)

	if _, err = os.Stat(responseDir); os.IsNotExist(err) {
		err = os.MkdirAll(responseDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	responseFile := filepath.Join(responseDir, "/"+FILENAME)
	err = ioutil.WriteFile(responseFile, []byte(response), os.ModePerm)

	fmt.Println(err, "response")

}

func createAPIS() {
	request := ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 获取联盟在推商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "goodsListWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=goodsListWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 获取指定商品id结合的商品信息-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "getByGoodsIdsWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=getByGoodsIdsWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 根据关键词查询商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "queryWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=queryWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 相似推荐商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "similarRecommendWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=similarRecommendWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionGoodsService 猜你喜欢商品列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionGoodsService",
		APIMETHOD:  "userRecommendWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&methodName=userRecommendWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService 根据商品id生成联盟链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "genByGoodsIdWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=genByGoodsIdWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService 根据唯品会链接生成联盟链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "genByVIPUrlWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=genByVIPUrlWithOauth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionUrlService 检测一段文本中是否有唯品会链接-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionUrlService",
		APIMETHOD:  "vipLinkCheckWithOuth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=vipLinkCheckWithOuth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderService 获取订单信息列表-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderService",
		APIMETHOD:  "orderListWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=vipLinkCheckWithOuth",
	}

	request = ApiInfo{
		APIDESC:    "com.vip.adp.api.open.service.UnionOrderService 维权订单总收益需要扣除该接口返回-需要oauth授权",
		APIORGNAME: "com.vip.adp.api.open.service.UnionOrderService",
		APIMETHOD:  "refundOrderListWithOauth",
		APIVERSION: "1.0.0",
		APIURL:     "http://vop.vip.com/apicenter/method?serviceName=com.vip.adp.api.open.service.UnionUrlService-1.0.0&methodName=refundOrderListWithOauth",
	}

	createAPI(request)
}

func main() {
	//
	createAPIS()
	//createReadmeApiList()
}
