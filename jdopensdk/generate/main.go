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
	dir = strings.Replace(dir, "\\", "/", -1) + "/jdopensdk"
	tplrequest := filepath.Join(dir, "generate/tplrequest.tpl")
	bytes, err := ioutil.ReadFile(tplrequest)
	if err != nil {
		panic(err)
	}
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
	responseJsonName := strings.ReplaceAll(apiRequest.APIORGNAME,".","_")
	response = strings.Replace(response, "--APIDESC--", apiRequest.APIDESC, -1)
	response = strings.Replace(response, "--PACKAGENAME--", DIRNAME, -1)
	response = strings.Replace(response, "--APINAME--", APINAME, -1)
	response = strings.Replace(response, "--RESPONSEJSONNAME--", responseJsonName , -1)
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
		APIDESC:    "jd.union.open.goods.jingfen.query 京粉精选商品查询接口",
		APIORGNAME: "jd.union.open.goods.jingfen.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.jingfen.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "jd.union.open.goods.query 接口描述：查询商品及优惠券信息",
		APIORGNAME: "jd.union.open.goods.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "jd.union.open.goods.material.query 猜你喜欢商品推荐",
		APIORGNAME: "jd.union.open.goods.material.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.material.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "jd.union.open.goods.promotiongoodsinfo.query 通过SKUID查询推广商品",
		APIORGNAME: "jd.union.open.goods.promotiongoodsinfo.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.promotiongoodsinfo.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "jd.union.open.category.goods.get 根据商品的父类目id查询子类目id信息",
		APIORGNAME: "jd.union.open.category.goods.get",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}


	request = ApiInfo{
		APIDESC:    "jd.union.open.goods.bigfield.query 商品详情查询接口,大字段信息",
		APIORGNAME: "jd.union.open.goods.bigfield.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.bigfield.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}


	request = ApiInfo{
		APIDESC:    "jd.union.open.activity.query 提供联盟官方活动查询",
		APIORGNAME: "jd.union.open.activity.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.activity.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}



	request = ApiInfo{
		APIDESC:    "jd.union.open.activity.recommend.query 提供联盟官方活动查询",
		APIORGNAME: "jd.union.open.activity.recommend.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.activity.recommend.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}




	request = ApiInfo{
		APIDESC:    "jd.union.open.promotion.common.get 网站/APP来获取的推广链接",
		APIORGNAME: "jd.union.open.promotion.common.get",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.common.get",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}





	request = ApiInfo{
		APIDESC:    "jd.union.open.promotion.bysubunionid.get 通过商品链接、领券链接、活动链接获取普通推广链接或优惠券二合一推广链接",
		APIORGNAME: "jd.union.open.promotion.bysubunionid.get",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.bysubunionid.get",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}




	request = ApiInfo{
		APIDESC:    "jd.union.open.promotion.byunionid.get 工具商媒体帮助子站长获取普通推广链接和优惠券二合一推广链接",
		APIORGNAME: "jd.union.open.promotion.byunionid.get",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.byunionid.get",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}



	request = ApiInfo{
		APIDESC:    "jd.union.open.order.row.query 查询推广订单及佣金信息",
		APIORGNAME: "jd.union.open.order.row.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.row.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "jd.union.open.order.bonus.query 奖励订单查询接口",
		APIORGNAME: "jd.union.open.order.bonus.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.bonus.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}
	request = ApiInfo{
		APIDESC:    "jd.union.open.statistics.redpacket.query 京享红包效果数据",
		APIORGNAME: "jd.union.open.statistics.redpacket.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.statistics.redpacket.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}


	request = ApiInfo{
		APIDESC:    "jd.union.open.order.agent.query 工具商订单行查询接口",
		APIORGNAME: "jd.union.open.order.agent.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.agent.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
				FieldNotNullCheck: true,
			},
		},
	}

	request = ApiInfo{
		APIDESC:    "jd.union.open.statistics.redpacket.agent.query 工具商查询帮助其他推客转的链接的红包发放数据",
		APIORGNAME: "jd.union.open.statistics.redpacket.agent.query",
		APIURL:     "https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.statistics.redpacket.agent.query",
		CHECKFIELDS: []CheckField{
			{
				FieldName:         "360buy_param_json",
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
