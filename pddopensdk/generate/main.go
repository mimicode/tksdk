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

	dir = strings.Replace(dir, "\\", "/", -1) + "/pddopensdk"

	tplrequest := filepath.Join(dir, "generate/tplrequest.tpl")

	bytes, err := ioutil.ReadFile(tplrequest)
	if err != nil {
		panic(err)
	}
	//"taobaoopensdk/utils"
	APIPARAMCHECK := createMustCheck(apiRequest.CHECKFIELDS)
	UTIL := `"github.com/mimicode/tksdk/utils"`
	DIRNAME := strings.ToLower(strings.Replace(apiRequest.APIORGNAME, ".", "", -1))
	FILENAME := DIRNAME + ".go"

	APINAME := strFirstToUpper(apiRequest.APIORGNAME, ".")
	//请求文件
	fileContent := string(bytes)
	fileContent = strings.Replace(fileContent, "--APIDESC--", apiRequest.APIDESC, -1)
	fileContent = strings.Replace(fileContent, "--APIORGNAME--", apiRequest.APIORGNAME, -1)
	fileContent = strings.Replace(fileContent, "--APIURL--", apiRequest.APIURL, -1)
	fileContent = strings.Replace(fileContent, "--APINAME--", APINAME, -1)
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
	response = strings.Replace(response, "--APINAME--", DIRNAME, -1)
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
		APIDESC:     "pdd.ddk.goods.detail（多多进宝商品详情查询）",
		APIORGNAME:  "pdd.ddk.goods.detail",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.detail",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:     "pdd.ddk.goods.search（多多进宝商品查询）",
		APIORGNAME:  "pdd.ddk.goods.search",
		APIURL:      "https://open.pinduoduo.com/#/apidocument/port?id=27",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.goods.prom.url.generate（生成多多进宝推广链接）",
		APIORGNAME: "pdd.ddk.oauth.goods.prom.url.generate",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?id=27",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "p_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "goods_id_list",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.goods.promotion.url.generate（生成多多进宝推广链接非授权）",
		APIORGNAME: "pdd.ddk.goods.promotion.url.generate",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?id=27",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "p_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "goods_id_list",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.goods.pid.generate（创建多多进宝推广位)",
		APIORGNAME: "pdd.ddk.goods.pid.generate",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.goods.pid.generate",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "number",
				FieldNotNullCheck: true,
			},
			{
				FieldName:        "number",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.goods.pid.generate（多多进宝推广位创建接口）",
		APIORGNAME: "pdd.ddk.oauth.goods.pid.generate",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.oauth.goods.pid.generate",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "number",
				FieldNotNullCheck: true,
			},
			{
				FieldName:        "number",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.resource.url.gen（生成多多进宝频道推广）",
		APIORGNAME: "pdd.ddk.resource.url.gen",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.resource.url.gen",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.theme.prom.url.generate（多多进宝主题推广链接生成",
		APIORGNAME: "pdd.ddk.theme.prom.url.generate",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?id=pdd.ddk.resource.url.gen",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "theme_id_list",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "pdd.ddk.top.goods.list.query（多多客获取爆款排行商品接口）",
		APIORGNAME: "pdd.ddk.top.goods.list.query",
		APIURL:     "https://open.pinduoduo.com/#/apidocument/port?portId=pdd.ddk.top.goods.list.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）",
		APIORGNAME: "pdd.ddk.order.list.increment.get",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.order.list.increment.get&permissionId=2",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "start_update_time",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "end_update_time",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.lottery.url.gen多多客生成转盘抽免单url",
		APIORGNAME: "pdd.ddk.lottery.url.gen",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.lottery.url.gen",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid_list",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.rp.prom.url.generate生成营销工具推广链接",
		APIORGNAME: "pdd.ddk.rp.prom.url.generate",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.rp.prom.url.generate",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "p_id_list",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.cms.prom.url.generate生成商城-频道推广链接",
		APIORGNAME: "pdd.ddk.cms.prom.url.generate",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.cms.prom.url.generate",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "p_id_list",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.member.authority.query 查询是否绑定备案",
		APIORGNAME: "pdd.ddk.member.authority.query",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.member.authority.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			}, {
				FieldName:         "custom_parameters",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "pdd.ddk.goods.zs.unit.url.gen多多进宝转链接口",
		APIORGNAME: "pdd.ddk.goods.zs.unit.url.gen",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.zs.unit.url.gen",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			}, {
				FieldName:         "custom_parameters",
				FieldNotNullCheck: true,
			}, {
				FieldName:         "source_url",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:     "pdd.ddk.goods.recommend.get多多进宝商品推荐API",
		APIORGNAME:  "pdd.ddk.goods.recommend.get",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.recommend.get",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.all.order.list.increment.get查询所有授权的多多客订单",
		APIORGNAME: "pdd.ddk.all.order.list.increment.get",
		APIURL:     "https://open.pinduoduo.com/application/document/api?permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "end_update_time",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "start_update_time",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.cashgift.create创建多多礼金",
		APIORGNAME: "pdd.ddk.oauth.cashgift.create",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cashgift.create&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "acquire_end_time",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "acquire_start_time",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "coupon_amount",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "quantity",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "validity_time_type",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "source_url",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.cashgift.status.update多多礼金状态更新",
		APIORGNAME: "pdd.ddk.oauth.cashgift.status.update",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cashgift.status.update&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "cash_gift_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "update_type",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.cms.prom.url.generate生成商城推广链接接口",
		APIORGNAME: "pdd.ddk.oauth.cms.prom.url.generate",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.cms.prom.url.generate&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "p_id_list",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:     "pdd.ddk.oauth.goods.detail多多进宝商品详情查询",
		APIORGNAME:  "pdd.ddk.oauth.goods.detail",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.detail&permissionId=7",
		CHECKFIELDS: []CheckField{},
	}
	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.goods.pid.generate多多进宝推广位创建接口",
		APIORGNAME: "pdd.ddk.oauth.goods.pid.generate",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.pid.generate&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "number",
				FieldMinCheck:    true,
				FieldMaxCheck:    true,
				FieldNumberCheck: true,
				FieldMin:         "1",
				FieldMax:         "100",
			},
		},
	}
	request = ApiInfo{
		APIDESC:     "pdd.ddk.oauth.goods.pid.query多多客已生成推广位信息查询",
		APIORGNAME:  "pdd.ddk.oauth.goods.pid.query",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.pid.query&permissionId=7",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:     "pdd.ddk.oauth.goods.prom.url.generate生成多多进宝推广链接",
		APIORGNAME:  "pdd.ddk.oauth.goods.prom.url.generate",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.prom.url.generate&permissionId=7",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:     "pdd.ddk.oauth.goods.recommend.get运营频道商品查询API",
		APIORGNAME:  "pdd.ddk.oauth.goods.recommend.get",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.recommend.get&permissionId=7",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:     "pdd.ddk.oauth.goods.search多多进宝商品查询",
		APIORGNAME:  "pdd.ddk.oauth.goods.search",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.search&permissionId=7",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.goods.zs.unit.url.gen多多进宝转链接口",
		APIORGNAME: "pdd.ddk.oauth.goods.zs.unit.url.gen",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.goods.zs.unit.url.gen&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "source_url",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:     "pdd.ddk.oauth.member.authority.query查询是否绑定备案",
		APIORGNAME:  "pdd.ddk.oauth.member.authority.query",
		APIURL:      "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.member.authority.query&permissionId=7",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.order.detail.get获取订单详情",
		APIORGNAME: "pdd.ddk.oauth.order.detail.get",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.order.detail.get&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "order_sn",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.pid.mediaid.bind批量绑定推广位的媒体id",
		APIORGNAME: "pdd.ddk.oauth.pid.mediaid.bind",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.pid.mediaid.bind&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "media_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "pid_list",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.resource.url.gen拼多多主站频道推广接口",
		APIORGNAME: "pdd.ddk.oauth.resource.url.gen",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.resource.url.gen&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "pdd.ddk.oauth.rp.prom.url.generate生成营销工具推广链接",
		APIORGNAME: "pdd.ddk.oauth.rp.prom.url.generate",
		APIURL:     "https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.rp.prom.url.generate&permissionId=7",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "p_id_list",
				FieldNotNullCheck: true,
			},
		},
	}

	createAPI(request)
}

func main() {
	//
	createAPIS()
	//createReadmeApiList()
}
