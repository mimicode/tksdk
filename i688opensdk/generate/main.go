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
	sdkPath := filepath.Join(dir, "..")
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
		file.WriteString("## 1688开放平台SDK\n\n")
		file.WriteString("这是1688开放平台的Go语言SDK，用于对接1688开放平台API。\n\n")
		file.WriteString("### 签名规则\n\n")
		file.WriteString("1688开放平台的签名规则如下：\n\n")
		file.WriteString("1. 构造签名因子：urlPath。url中的一部分，从协议（param2）开始截取，到\"?\"为止，urlPath=param2/1/system/currentTime/1000000\n\n")
		file.WriteString("2. 构造签名因子：拼装的参数。参数按照首字母排序，然后将key和value拼在一起，最后按顺序拼在一起\n\n")
		file.WriteString("3. 合并两个签名因子。把前两步的字符串拼起来\n\n")
		file.WriteString("4. 对合并后的签名因子执行hmac_sha1算法。Signature=uppercase(hex(hmac_sha1(s, secretKey)))\n\n")
		file.WriteString("## API列表\n")
		file.WriteString("--\n")
		file.WriteString(strings.Join(apiList, "\n"))
	}
}

type ApiInfo struct {
	APIDESC        string
	APIORGNAME     string
	APIURL         string
	BUSINESSMODULE string
	CHECKFIELDS    []CheckField
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

func createRequestFile(apiInfo ApiInfo) {
	dir, _ := os.Getwd()
	sdkPath := filepath.Join(dir, "..")
	fileName := strings.ToLower(apiInfo.APIORGNAME)
	fileName = strings.Replace(fileName, ".", "", -1)
	fileName = strings.Replace(fileName, "_", "", -1)
	fileName = strings.Replace(fileName, "-", "", -1)
	fileName = strings.Replace(fileName, "1688", "i688", -1)
	fileName = strings.Replace(fileName, "alibaba", "i688", -1)

	var tplRequest string
	tplRequestBytes, err := ioutil.ReadFile(filepath.Join(dir, "tplrequest.tpl"))
	if err != nil {
		fmt.Println(err)
		return
	}
	tplRequest = string(tplRequestBytes)

	//替换模板的变量
	tplRequest = strings.Replace(tplRequest, "--APIDESC--", apiInfo.APIDESC, -1)
	tplRequest = strings.Replace(tplRequest, "--APIURL--", apiInfo.APIURL, -1)
	tplRequest = strings.Replace(tplRequest, "--APINAME--", utils.StrFirstToUpper(fileName, ""), -1)
	tplRequest = strings.Replace(tplRequest, "--APIORGNAME--", apiInfo.APIORGNAME, -1)
	tplRequest = strings.Replace(tplRequest, "--BUSINESSMODULE--", apiInfo.BUSINESSMODULE, -1)
	tplRequest = strings.Replace(tplRequest, "--APIPARAMCHECK--", createMustCheck(apiInfo.CHECKFIELDS), -1)

	//判断是否需要引入工具包
	if len(apiInfo.CHECKFIELDS) > 0 {
		tplRequest = strings.Replace(tplRequest, "--UTIL--", `"github.com/mimicode/tksdk/utils"`, -1)
	} else {
		tplRequest = strings.Replace(tplRequest, "--UTIL--", "", -1)
	}

	//写入文件
	outFileName := filepath.Join(sdkPath, "request", fileName+".go")
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	outFile.WriteString(tplRequest)
	fmt.Println("创建请求文件：" + outFileName)
}

func createResponseDir(apiInfo ApiInfo) {
	dir, _ := os.Getwd()
	sdkPath := filepath.Join(dir, "..")
	fileName := strings.ToLower(apiInfo.APIORGNAME)
	fileName = strings.Replace(fileName, ".", "", -1)
	fileName = strings.Replace(fileName, "_", "", -1)
	fileName = strings.Replace(fileName, "-", "", -1)
	fileName = strings.Replace(fileName, "1688", "i688", -1)
	fileName = strings.Replace(fileName, "alibaba", "i688", -1)

	//创建目录
	respDir := filepath.Join(sdkPath, "response", fileName)
	os.MkdirAll(respDir, 0755)

	//创建响应文件
	var tplResponse string
	tplResponseBytes, err := ioutil.ReadFile(filepath.Join(dir, "tplresponse.tpl"))
	if err != nil {
		fmt.Println(err)
		return
	}
	tplResponse = string(tplResponseBytes)

	//替换模板的变量
	tplResponse = strings.Replace(tplResponse, "--APIDESC--", apiInfo.APIDESC, -1)
	tplResponse = strings.Replace(tplResponse, "--APINAME--", fileName, -1)

	//写入文件
	outFileName := filepath.Join(respDir, fileName+".go")
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	outFile.WriteString(tplResponse)
	fmt.Println("创建响应文件：" + outFileName)
}

func main() {
	//创建系统当前时间API
	// apiInfo := ApiInfo{
	// 	APIDESC:    "alibaba.system.currentTime( 获取系统当前时间 )",
	// 	APIORGNAME: "system/currentTime",
	// 	APIURL:     "https://open.1688.com/api/system/currentTime.html",
	// }
	
	// 我们已经手动创建了这些文件，所以这里不需要再创建
	// createRequestFile(apiInfo)
	// createResponseDir(apiInfo)

	//更新README.MD文件
	createReadmeApiList()
}