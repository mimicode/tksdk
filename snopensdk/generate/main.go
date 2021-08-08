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
	fmt.Println(dir)

	dir = strings.Replace(dir, "\\", "/", -1) + "/snopensdk"

	tplrequest := filepath.Join(dir, "generate/tplrequest.tpl")

	bytes, err := ioutil.ReadFile(tplrequest)
	if err != nil {
		panic(err)
	}
	//"taobaoopensdk/utils"
	APIPARAMCHECK := createMustCheck(apiRequest.CHECKFIELDS)
	UTIL := `"github.com/mimicode/tksdk/utils"`
	DIRNAME := strings.Replace(apiRequest.APIORGNAME, ".", "", -1)
	FILENAME := DIRNAME + ".go"

	APINAME := strFirstToUpper(apiRequest.APIORGNAME, ".")
	//请求文件
	fileContent := string(bytes)
	fileContent = strings.Replace(fileContent, "--APIDESC--", apiRequest.APIDESC, -1)
	fileContent = strings.Replace(fileContent, "--APIORGNAME--", apiRequest.APIORGNAME, -1)
	fileContent = strings.Replace(fileContent, "--APIURL--", apiRequest.APIURL, -1)
	fileContent = strings.Replace(fileContent, "--APINAME--", APINAME, -1)
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
		APIDESC:    "suning.netalliance.extensionlink.get 商品和券二合一接口",
		APIORGNAME: "suning.netalliance.extensionlink.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.extensionlink.get",
	}
	request = ApiInfo{
		APIDESC:    "suning.netalliance.commoditydetail.query 推广商品详情信息接口",
		APIORGNAME: "suning.netalliance.commoditydetail.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.commoditydetail.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.selectrecommendcommodity.query 商品精选接口",
		APIORGNAME: "suning.netalliance.selectrecommendcommodity.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.selectrecommendcommodity.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.inverstmentcommodity.query 高佣专区商品查询接口",
		APIORGNAME: "suning.netalliance.inverstmentcommodity.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.inverstmentcommodity.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.searchcommoditynew.query 关键字商品查询接口(新)",
		APIORGNAME: "suning.netalliance.searchcommoditynew.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.searchcommoditynew.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.unioninfomation.get 单个查询联盟商品信息",
		APIORGNAME: "suning.netalliance.unioninfomation.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.unioninfomation.get",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.commodityimages.query 商品图文详情查询",
		APIORGNAME: "suning.netalliance.commodityimages.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.commodityimages.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.toolseller.add 生成工具商PID接口",
		APIORGNAME: "suning.netalliance.toolseller.add",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.toolseller.add",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.couponinfo.query 查询券领用情况",
		APIORGNAME: "suning.netalliance.couponinfo.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.couponinfo.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.recommendcommodity.query 小编推荐接口",
		APIORGNAME: "suning.netalliance.recommendcommodity.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.recommendcommodity.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.hoistinglink.query 获取吊起链接接口",
		APIORGNAME: "suning.netalliance.hoistinglink.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.hoistinglink.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.order.query 网盟订单信息批量查询",
		APIORGNAME: "suning.netalliance.order.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.order.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.order.get 网盟订单信息单独查询",
		APIORGNAME: "suning.netalliance.order.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.order.get",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.ordersettle.get 网盟订单结算信息查询",
		APIORGNAME: "suning.netalliance.ordersettle.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.ordersettle.get",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.ordersettlenew.get 订单结算信息接口",
		APIORGNAME: "suning.netalliance.ordersettlenew.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.ordersettle.get",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.riskmanagementorder.query 风控订单查询",
		APIORGNAME: "suning.netalliance.riskmanagementorder.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.riskmanagementorder.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.orderinfo.query 订单批量查询接口",
		APIORGNAME: "suning.netalliance.orderinfo.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.riskmanagementorder.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.morerecommend.get 关联商品推荐接口",
		APIORGNAME: "suning.netalliance.morerecommend.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.morerecommend.get",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.searchcommodity.query 关键字商品查询接口",
		APIORGNAME: "suning.netalliance.searchcommodity.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.searchcommodity.query",
	}

	request = ApiInfo{
		APIDESC:    "suning.netalliance.custompromotionurl.query 获取自定义推广链接接口",
		APIORGNAME: "suning.netalliance.custompromotionurl.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.custompromotionurl.query",
	}
	request = ApiInfo{
		APIDESC:    "suning.netalliance.storepromotionurl.query 获取商品或店铺推广链接接口",
		APIORGNAME: "suning.netalliance.storepromotionurl.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.storepromotionurl.query",
	}
	request = ApiInfo{
		APIDESC:    "suning.netalliance.bacthcustomlink.query 批量转链",
		APIORGNAME: "suning.netalliance.bacthcustomlink.query",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.bacthcustomlink.query",
	}
	request = ApiInfo{
		APIDESC:    "suning.netalliance.appletextensionlink.get 商品和券二合一（小程序）",
		APIORGNAME: "suning.netalliance.appletextensionlink.get",
		APIVERSION: "v1.2",
		APIURL:     "https://open.suning.com/ospos/apipage/toApiMethodDetailMenuNew.do?interCode=suning.netalliance.appletextensionlink.get",
	}
	createAPI(request)
}

func main() {
	//
	createAPIS()
	//createReadmeApiList()
}
