package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SortParamters 参数排序返回排序的KEY 通过KEY循环取出
func SortParamters(params url.Values) []string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
func SortMapParamters(params map[string]interface{}) []string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}

func Md5(data string) string {
	md5obj := md5.New()
	if _, err := md5obj.Write([]byte(data)); err != nil {
		return ""
	}
	md5Data := md5obj.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func Hmac(key, data string) string {
	hmac1 := hmac.New(md5.New, []byte(key))
	hmac1.Write([]byte(data))
	return hex.EncodeToString(hmac1.Sum([]byte("")))
}

// CheckMaxLength 检验字段fieldName的值value 的长度
func CheckMaxLength(value string, maxLength int, fieldName string) {
	if len([]rune(value)) > maxLength {
		panic(fmt.Sprintf("Invalid Arguments:the length of %s can not be larger than %d.", fieldName, maxLength))
	}
}

// CheckMaxListSize 检验字段fieldName的值value的最大列表长度
func CheckMaxListSize(value string, maxSize int, fieldName string) {

	if !CheckEmpty(value) && len(strings.Split(value, ",")) > maxSize {
		panic(fmt.Sprintf("Invalid Arguments:the listsize(the string split by \",\") of %s must be less than %d .", fieldName, maxSize))
	}
}

// CheckMaxValue 检测最大值
func CheckMaxValue(value string, maxValue int, fieldName string) {
	if CheckEmpty(value) {
		return
	}
	if i, e := strconv.Atoi(value); e != nil {
		panic("fieldName cid Value Is Not Number")
	} else {
		if i > maxValue {
			panic(fmt.Sprintf("Invalid Arguments:the value of %s can not be larger than %d .", fieldName, maxValue))
		}
	}
}

// CheckMinValue 检测最小值
func CheckMinValue(value string, minVal int, fieldName string) {

	if CheckEmpty(value) {
		return
	}
	if i, e := strconv.Atoi(value); e != nil {
		panic("fieldName cid Value Is Not Number")
	} else {
		if i < minVal {
			panic(fmt.Sprintf("Invalid Arguments:the value of %s can not be less than %d .", fieldName, minVal))
		}
	}

}

// CheckNotNull 检测非空
func CheckNotNull(value, fieldName string) {
	if CheckEmpty(value) {
		panic(fmt.Sprintf("Missing Required Arguments: %s", fieldName))
	}
}

// CheckEmpty 检测是否为空 空则返回真
func CheckEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// 检测是否为数字
func CheckNumber(value, fieldName string) {
	compile := regexp.MustCompile(`\d+`)
	if !compile.MatchString(value) {
		panic(fmt.Sprintf("Invalid Arguments:the value of %s is not number : %s.", fieldName, value))
	}
}

func NowTime() *time.Time {
	var cstSh = time.FixedZone("CST", 8*3600)
	in := time.Now().In(cstSh)
	return &in
}

func GetUUID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 16)
}

// 检测最小值
func CheckMinFloatValue(value string, minVal float64, fieldName string) {
	if CheckEmpty(value) {
		return
	}
	if i, e := strconv.ParseFloat(value, 64); e != nil {
		panic("fieldName cid Value Is Not Number")
	} else {
		if i < minVal {
			panic(fmt.Sprintf("Invalid Arguments:the value of %s can not be less than %.2f .", fieldName, minVal))
		}
	}

}

func StrFirstToUpper(str string, step string) string {
	temp := strings.Split(str, step)
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 'a' && vv[i] <= 'z' {
					vv[i] -= 32
				}
			}
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
