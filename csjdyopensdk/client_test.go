package csjdyopensdk

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	appKey, appSecret  string
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Jd struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"` 
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
	client.Init(appKey, appSecret, sessionKey)
	return client
}


func TestProductSearchRequest(t *testing.T) {
	client := GetClient()
	request := &request.ProductSearchRequest{}
	request.AddParameter("title", "iphone")
	request.AddParameter("page", "1")
	request.AddParameter("page_size", "10")
	response, err := client.Execute(request)
	if err != nil {
		t.Errorf("ProductSearchRequest failed: %v", err)
	}
	t.Logf("ProductSearchRequest response: %v", response)
}
