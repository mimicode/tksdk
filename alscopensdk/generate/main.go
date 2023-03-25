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
	sdkPath := filepath.Join(dir, "alscopensdk")
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
	SubFields []CheckField
}

func createMustCheck(checkFields []CheckField) string {

	var joinFunc = func(v CheckField) string {
		var checkStr = ""
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
		return checkStr
	}

	checkStr := ""
	for _, v := range checkFields {
		checkStr += joinFunc(v)
		checkStr += subJoin(v, "")
	}
	return checkStr
}

func subJoin(v CheckField, pkey string) string {

	var joinFunc2 = func(field CheckField, key string) string {
		var checkStr = ""
		//非空检测
		if field.FieldNotNullCheck {
			checkStr += `utils.CheckNotNull(tk.getMapVal(` + key + `,"` + field.FieldName + `"), "` + field.FieldName + `")
`
		}
		//最小值检测
		if field.FieldMinCheck {
			checkStr += `utils.CheckMinValue(tk.getMapVal(` + key + `,"` + field.FieldName + `"),` + field.FieldMin + `, "` + field.FieldName + `")
`
		}
		//最大值检测
		if field.FieldMaxCheck {
			checkStr += `utils.CheckMaxValue(tk.getMapVal(` + key + `,"` + field.FieldName + `"),` + field.FieldMax + `, "` + field.FieldName + `")
`
		}
		//数值检测
		if field.FieldNumberCheck {
			checkStr += `utils.CheckNumber(tk.getMapVal(` + key + `,"` + field.FieldName + `"), "` + field.FieldName + `")
`
		}
		//值长度检测
		if field.FieldLenCheck {
			checkStr += `utils.CheckMaxLength(tk.getMapVal(` + key + `,"` + field.FieldName + `"),` + field.FieldLen + `, "` + field.FieldName + `")
`
		}
		//集合长度检测
		if field.FieldListSizeCheck {
			checkStr += `utils.CheckMaxListSize(tk.getMapVal(` + key + `,"` + field.FieldName + `"),` + field.FieldSize + `, "` + field.FieldName + `")
`
		}
		return checkStr
	}
	checkStr := ""
	if len(v.SubFields) > 0 {

		key := fmt.Sprintf("%s_%s", pkey, v.FieldName)
		if pkey == "" {
			key = v.FieldName
		}
		key = utils.StrFirstToUpper(key, "_")
		if pkey == "" {
			checkStr += `var ` + key + ` = tk.parseSubParameters(tk.Parameters.Get("` + v.FieldName + `"))`
		} else {
			checkStr += `var ` + key + ` = tk.parseSubParameters(tk.getMapVal(` + pkey + `,"` + v.FieldName + `"))`
		}
		checkStr += "\n"

		for _, subField := range v.SubFields {
			checkStr += joinFunc2(subField, key)
			if len(subField.SubFields) > 0 {
				checkStr += subJoin(subField, key)
				checkStr += "\n"
			}
		}

	}
	//fmt.Println("pk====", pkey, "====", checkStr)
	return checkStr
}

func createAPI(apiRequest ApiInfo) {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1) + "/alscopensdk"

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

	APINAME := utils.StrFirstToUpper(apiRequest.APIORGNAME, ".")
	RESPONSENAME := utils.StrFirstToUpper(apiRequest.APIORGNAME, ".") + "Response"
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
	response = strings.Replace(response, "--RESPONSJSON--", strings.Replace(apiRequest.APIORGNAME, ".", "_", -1), -1)
	response = strings.Replace(response, "--RESPONSENAME--", RESPONSENAME, -1)
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
		APIDESC:    "alibaba.alsc.union.kb.item.promotion( 本地生活媒体平台口碑选品 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.promotion",
		APIURL:     "https://open.taobao.com/api.htm?docId=58766&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "page_number",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_size",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
			{
				FieldName:        "settle_type",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.media.zone.add( 本地生活媒体推广位创建 )",
		APIORGNAME: "alibaba.alsc.union.media.zone.add",
		APIURL:     "https://open.taobao.com/api.htm?docId=58768&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "zone_name",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.item.promotion.filter.list( 本地生活媒体平台口碑选品筛选项集合 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.promotion.filter.list",
		APIURL:     "https://open.taobao.com/api.htm?docId=59273&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "filter_type",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.media.zone.get( 本地生活媒体推广位查询 )",
		APIORGNAME: "alibaba.alsc.union.media.zone.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=59276&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "page",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "limit",
				FieldNumberCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kbcpx.positive.order.get( 本地生活媒体推广订单明细报表查询 )",
		APIORGNAME: "alibaba.alsc.union.kbcpx.positive.order.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=59449&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "date_type",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "biz_unit",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_size",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_number",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "start_date",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kbcpx.refund.order.get( 本地生活媒体推广用户维权订单数据报表 )",
		APIORGNAME: "alibaba.alsc.union.kbcpx.refund.order.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=59450&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "date_type",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "biz_unit",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_size",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_number",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "start_date",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kbcpx.punish.order.get( 本地生活媒体推广用户反作弊订单数据报表 )",
		APIORGNAME: "alibaba.alsc.union.kbcpx.punish.order.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=59451&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "date_type",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "biz_unit",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_size",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_number",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "start_date",
				FieldNumberCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.item.promotion.share.create( 本地生活媒体创建商品推广链接 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.promotion.share.create",
		APIURL:     "https://open.taobao.com/api.htm?docId=59620&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "item_id",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.eleme.promotion.storepromotion.query( 本地联盟饿了么单店推广店铺列表 )",
		APIORGNAME: "alibaba.alsc.union.eleme.promotion.storepromotion.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=60447&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "pid",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "longitude",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "latitude",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.eleme.promotion.storepromotion.get( 本地联盟饿了么单店推广单店铺查询 )",
		APIORGNAME: "alibaba.alsc.union.eleme.promotion.storepromotion.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=60448&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "pid",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "shop_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.eleme.promotion.officialactivity.get( 本地联盟饿了么推广官方活动查询 )",
		APIORGNAME: "alibaba.alsc.union.eleme.promotion.officialactivity.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=60449&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "pid",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "activity_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.item.query( 本地联盟口碑商品列表 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=63314&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{
			{
				FieldName:        "page_number",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_size",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "biz_type",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.store.item.query( 本地联盟口碑门店商品列表 )",
		APIORGNAME: "alibaba.alsc.union.kb.store.item.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=63268&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "store_id",
				FieldNotNullCheck: true,
			},
			{
				FieldName:         "biz_type",
				FieldNotNullCheck: true,
			}, {
				FieldName:         "pid",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.store.query( 本地联盟口碑商户列表 )",
		APIORGNAME: "alibaba.alsc.union.kb.store.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=63267&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{

			{
				FieldName:        "page_number",
				FieldNumberCheck: true,
			},
			{
				FieldName:        "page_size",
				FieldNumberCheck: true,
			},
			{
				FieldName:         "biz_type",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.item.store.detail.get( 本地联盟口碑门店详情 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.store.detail.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=62009&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "store_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.item.store.relation.query( 本地联盟口碑商品门店关系 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.store.relation.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=62008&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "item_id",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "biz_type",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.item.detail.get( 本地联盟口碑商品详情 )",
		APIORGNAME: "alibaba.alsc.union.kb.item.detail.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=62007&docType=2&scopeId=24408",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "item_id",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "biz_type",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.order.create( openapi订单创建 )",
		APIORGNAME: "alibaba.alsc.union.kb.order.create",
		APIURL:     "https://open.taobao.com/api.htm?docId=59883&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "order_dto",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "outer_order_id",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "item_id",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "title",
						FieldNotNullCheck: true,
					},
					{
						FieldName:        "quantity",
						FieldNumberCheck: true,
					},
					{
						FieldName:        "pay_order_fee",
						FieldNumberCheck: true,
					},
					{
						FieldName:        "order_fee",
						FieldNumberCheck: true,
					},
					{
						FieldName:        "sell_price",
						FieldNumberCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.order.query( openapi订单查询 )",
		APIORGNAME: "alibaba.alsc.union.kb.order.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=59884&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "order_query_dto",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "biz_order_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.order.refund( openapi订单售中退 )",
		APIORGNAME: "alibaba.alsc.union.kb.order.refund",
		APIURL:     "https://open.taobao.com/api.htm?docId=59885&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{

			{
				FieldName:         "order_refund_dto",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "voucher_list",
						FieldNotNullCheck: true,
						SubFields: []CheckField{
							{
								FieldName:         "item_id",
								FieldNotNullCheck: true,
							},
							{
								FieldName:         "voucher_id",
								FieldNotNullCheck: true,
							},
						},
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.bbt.item.query( 本地联盟爆爆团商品列表 )",
		APIORGNAME: "alibaba.alsc.union.kb.bbt.item.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=59941&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.bbt.item.detail.get( 本地联盟爆爆团商品详情 )",
		APIORGNAME: "alibaba.alsc.union.kb.bbt.item.detail.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=59975&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "item_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.bbt.item.store.relation.query( 本地联盟爆爆团商品门店关系 )",
		APIORGNAME: "alibaba.alsc.union.kb.bbt.item.store.relation.query",
		APIURL:     "https://open.taobao.com/api.htm?docId=59976&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "item_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.bbt.item.store.detail.get( 本地联盟爆爆团门店详情 )",
		APIORGNAME: "alibaba.alsc.union.kb.bbt.item.store.detail.get",
		APIURL:     "https://open.taobao.com/api.htm?docId=59977&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "query_request",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "store_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.order.pay( openapi预下单订单支付 )",
		APIORGNAME: "alibaba.alsc.union.kb.order.pay",
		APIURL:     "https://open.taobao.com/api.htm?docId=61523&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "order_pay_dto",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "outer_order_id",
						FieldNotNullCheck: true,
					},
					{
						FieldName:         "biz_order_id",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.common.encrypt( 加密方法 )",
		APIORGNAME: "alibaba.alsc.union.kb.common.encrypt",
		APIURL:     "https://open.taobao.com/api.htm?docId=61584&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "encrypt_model",
				FieldNotNullCheck: true,
				SubFields: []CheckField{
					{
						FieldName:         "text",
						FieldNotNullCheck: true,
					},
				},
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "alibaba.alsc.union.kb.bbt.item.promotion.filter.list( 本地生活爆爆团选品筛选集合 )",
		APIORGNAME: "alibaba.alsc.union.kb.bbt.item.promotion.filter.list",
		APIURL:     "https://open.taobao.com/api.htm?docId=61957&docType=2&scopeId=24987",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "filter_type",
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
