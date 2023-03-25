package main

import (
	"fmt"
	"github.com/mimicode/tksdk/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func createReadmeApiList() {
	dir, _ := os.Getwd()
	sdkPath := filepath.Join(dir, "tbopensdk")
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
	dir = strings.Replace(dir, "\\", "/", -1) + "/tbopensdk"

	tplrequest := filepath.Join(dir, "generate/tplrequest.tpl")

	bytes, err := ioutil.ReadFile(tplrequest)
	if err != nil {
		panic(err)
	}
	//"taobaoopensdk/utils"
	APIPARAMCHECK := createMustCheck(apiRequest.CHECKFIELDS)
	UTIL := `"github.com/mimicode/tksdk/utils"`
	APINAME := strings.Replace(apiRequest.APIORGNAME, "taobao.", "", -1)

	DIRNAME := strings.Replace(apiRequest.APIORGNAME, ".", "", -1)
	FILENAME := DIRNAME + ".go"

	APINAME = utils.StrFirstToUpper(APINAME, ".")
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
		APIDESC:    "taobao.tbk.shop.get( 淘宝客店铺查询 )",
		APIORGNAME: "taobao.tbk.shop.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=24521&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "fields",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "q",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.shop.recommend.get( 淘宝客店铺关联推荐查询 )",
		APIORGNAME: "taobao.tbk.shop.recommend.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=24522&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "fields",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "user_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.uatm.favorites.item.get( 获取淘宝联盟选品库的宝贝信息 )",
		APIORGNAME: "taobao.tbk.uatm.favorites.item.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=26619&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "fields",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
			{
				FieldName:         "favorites_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.uatm.favorites.get( 获取淘宝联盟选品库列表 )",
		APIORGNAME: "taobao.tbk.uatm.favorites.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=26620&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "fields",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.ju.tqg.get( 淘抢购api )",
		APIORGNAME: "taobao.tbk.ju.tqg.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=27543&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "fields",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "start_time",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "end_time",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:     "taobao.ju.items.search( 聚划算商品搜索接口 )",
		APIORGNAME:  "taobao.ju.items.search",
		APIURL:      "http://open.taobao.com/api.htm?docId=28762&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.dg.item.coupon.get( 好券清单API【导购】 )",
		APIORGNAME: "taobao.tbk.dg.item.coupon.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=29821&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:     "taobao.tbk.coupon.get( 阿里妈妈推广券信息查询 )",
		APIORGNAME:  "taobao.tbk.coupon.get",
		APIURL:      "http://open.taobao.com/api.htm?docId=31106&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.tpwd.create( 淘宝客淘口令 )",
		APIORGNAME: "taobao.tbk.tpwd.create",
		APIURL:     "http://open.taobao.com/api.htm?docId=31127&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "text",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "url",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.dg.newuser.order.get( 淘宝客新用户订单API--导购 )",
		APIORGNAME: "taobao.tbk.dg.newuser.order.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=33892&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "activity_id",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.newuser.order.get( 淘宝客新用户订单API--社交 )",
		APIORGNAME: "taobao.tbk.sc.newuser.order.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=33897&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "activity_id",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.dg.optimus.material( 淘宝客物料下行-导购 )",
		APIORGNAME: "taobao.tbk.dg.optimus.material",
		APIURL:     "http://open.taobao.com/api.htm?docId=33947&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.dg.material.optional( 通用物料搜索API（导购） )",
		APIORGNAME: "taobao.tbk.dg.material.optional",
		APIURL:     "http://open.taobao.com/api.htm?docId=35896&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.dg.newuser.order.sum( 拉新活动汇总API--导购 )",
		APIORGNAME: "taobao.tbk.dg.newuser.order.sum",
		APIURL:     "http://open.taobao.com/api.htm?docId=36836&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "activity_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "page_no",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
			{
				FieldName:         "page_size",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.newuser.order.sum( 拉新活动汇总API--社交 )",
		APIORGNAME: "taobao.tbk.sc.newuser.order.sum",
		APIURL:     "http://open.taobao.com/api.htm?docId=36837&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "activity_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "page_no",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
			{
				FieldName:         "page_size",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.optimus.material( 淘宝客擎天柱通用物料API - 社交 )",
		APIORGNAME: "taobao.tbk.sc.optimus.material",
		APIURL:     "http://open.taobao.com/api.htm?docId=37884&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
			{
				FieldName:         "site_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.activitylink.get( 淘宝联盟官方活动推广API-媒体 )",
		APIORGNAME: "taobao.tbk.activitylink.get",
		APIURL:     "http://open.taobao.com/api.htm?docId=41918&docType=2&scopeId=11655",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
			{
				FieldName:         "promotion_scene_id",
				FieldNotNullCheck: true,
				FieldNumberCheck:  true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.order.details.get( 淘宝客【服务商】所有订单查询 )",
		APIORGNAME: "taobao.tbk.sc.order.details.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=43755&docType=2&scopeId=16322",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "end_time",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "start_time",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.activity.info.get( 淘宝客-服务商-官方活动信息获取 新接口 )",
		APIORGNAME: "taobao.tbk.sc.activity.info.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=48417&docType=2&scopeId=18353",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "activity_material_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "adzone_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "site_id",
				FieldNotNullCheck: true,
			},
		},
	}
	/*
	  per_face	String	true	10	单个淘礼金面额，支持两位小数，单位元
	  send_start_time*/
	request = ApiInfo{
		APIDESC:    "taobao.tbk.dg.vegas.tlj.create( 淘宝客-推广者-淘礼金创建 )",
		APIORGNAME: "taobao.tbk.dg.vegas.tlj.create",
		APIURL:     "https://open.taobao.com/api.htm?docId=40173&docType=2",
		CHECKFIELDS: []CheckField{

			{
				FieldName:        "adzone_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "item_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "total_num",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "user_total_win_num_limit",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "security_switch",
				FieldNotNullCheck: true,
			},
			{
				FieldName:     "name",
				FieldLenCheck: true,
				FieldLen:      "10",
			},
			{
				FieldName:     "per_face",
				FieldMinCheck: true,
				FieldMin:      "1",
			},
			{
				FieldName:         "send_start_time",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.shop.convert( 淘宝客-服务商-店铺链接转换 )",
		APIORGNAME: "taobao.tbk.sc.shop.convert",
		APIURL:     "https://open.taobao.com/api.htm?docId=43878&docType=2&scopeId=16403",
		CHECKFIELDS: []CheckField{

			{
				FieldName:        "site_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "fields",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "user_ids",
				FieldNotNullCheck: true,
			},
			{
				FieldName:        "adzone_id",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.tpwd.convert( 淘宝客-服务商-淘口令解析&转链 )",
		APIORGNAME: "taobao.tbk.sc.tpwd.convert",
		APIURL:     "https://open.taobao.com/api.htm?docId=43873&docType=2&scopeId=16401",
		CHECKFIELDS: []CheckField{

			{
				FieldName:        "site_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "adzone_id",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "password_content",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.item.click.extract( 淘宝客-公用-链接解析出商品id )",
		APIORGNAME: "taobao.tbk.item.click.extract",
		APIURL:     "https://open.taobao.com/api.htm?docId=43873&docType=2&scopeId=16401",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "click_url",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "taobao.tbk.itemid.private.transform( 淘宝客-推广者-商品id转化（二方）（专有） )",
		APIORGNAME: "taobao.tbk.itemid.private.transform",
		APIURL:     "https://open.taobao.com/api.htm?docId=64171&docType=2&scopeId=27337",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "item_ids",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "taobao.tbk.sc.adzone.create( 淘宝客-服务商-创建推广者位 )",
		APIORGNAME: "taobao.tbk.sc.adzone.create",
		APIURL:     "https://open.taobao.com/api.htm?docId=34751&docType=2&scopeId=13878",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "site_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "adzone_name",
				FieldNotNullCheck: true,
			},
		},
	}
	//request = ApiInfo{
	//	APIDESC:    "taobao.tbk.sc.relation.refund( 淘宝客-服务商-维权退款订单查询 )",
	//	APIORGNAME: "taobao.tbk.sc.relation.refund",
	//	APIURL:     "https://open.taobao.com/api.htm?docId=43874&docType=2&scopeId=16322",
	//	CHECKFIELDS: []CheckField{
	//
	//		{
	//			FieldName:        "page_size",
	//			FieldNumberCheck: true,
	//		},
	//		{
	//			FieldName:        "search_type",
	//			FieldNumberCheck: true,
	//		},
	//		{
	//			FieldName:        "refund_type",
	//			FieldNumberCheck: true,
	//		},
	//
	//		{
	//			FieldName:        "page_no",
	//			FieldNumberCheck: true,
	//		},
	//		{
	//			FieldName:        "biz_type",
	//			FieldNumberCheck: true,
	//		},
	//		{
	//			FieldName:         "start_time",
	//			FieldNotNullCheck: true,
	//		},
	//	},
	//}

	createAPI(request)
}

func main() {
	//
	createAPIS()
	//createReadmeApiList()
}
