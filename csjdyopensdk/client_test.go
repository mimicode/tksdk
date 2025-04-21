package csjdyopensdk

import (
	"encoding/json"
	"fmt"
	"github.com/mimicode/tksdk/csjdyopensdk/request"
	"github.com/mimicode/tksdk/csjdyopensdk/response/productsearch"
	"io/ioutil"
	"os"
	"testing"
)

var (
	appKey, appSecret string
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Jd struct {
					AppKey    string `json:"app_key"`
					AppSecret string `json:"app_secret"`
				} `json:"csjdy"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Jd.AppKey
				appSecret = data.Jd.AppSecret
			}
		}
	}
}
func GetClient() *TopClient {
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")
	return client
}

func TestProductSearchRequest(t *testing.T) {
	client := GetClient()
	req := &request.ProductSearchRequest{}
	req.AddParameter("title", "iphone")
	req.AddParameter("page", 1)
	req.AddParameter("page_size", 10)

	var getResponse DefaultResponse = &productsearch.Response{}

	err := client.Exec(req, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*productsearch.Response)

		if result.IsError() {
			fmt.Println(result.Body)
		} else {

		}

	}
}
